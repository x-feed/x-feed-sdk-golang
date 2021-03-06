package xfeed

import (
	"context"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	"github.com/x-feed/x-feed-sdk-golang/pkg/logging"
	pb "github.com/x-feed/x-feed-sdk-golang/pkg/xfeed_proto"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

// Session represents started and working session of x-feed.
type Session struct {
	logger         logging.Logger
	requestTimeout time.Duration
	clientID       string

	limiter    *rate.Limiter
	clientConn *grpc.ClientConn
	// eventsFeedMutex, eventSettlementsMutex, entitiesMutex mutexes prevent concurrent invocation of specific API method.
	eventsFeedMutex sync.Mutex
	eventsStream    chan *EventEnvelope
	marketsStream   chan *MarketEnvelope

	eventSettlementsMutex sync.Mutex
	eventSettlements      chan *EventSettlementEnvelope

	entitiesMutex sync.Mutex

	statusMutex         sync.Mutex
	statusChangeHandler func(ConnectionStatus)
	currentStatus       ConnectionStatus
}

// EventsFeed returns channels of state updates for Events and Markets.
// when there are communication errors with X-feed servers it closes the channels
func (s *Session) EventsFeed() (<-chan *EventEnvelope, <-chan *MarketEnvelope, error) {
	if s == nil {
		s.setStatus(StatusRed, s.currentStatus.SettlementsStream)
		return nil, nil, errors.New("session is not initialised")
	}
	s.eventsFeedMutex.Lock()
	defer s.eventsFeedMutex.Unlock()

	if s.eventsStream != nil && s.marketsStream != nil {
		return s.eventsStream, s.marketsStream, nil
	}

	eventRequest := &pb.StreamEventsRequest{
		ClientName: s.clientID,
	}

	err := s.limiter.Wait(context.Background())
	if err != nil {

		s.setStatus(StatusRed, s.currentStatus.SettlementsStream)
		return nil, nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	eventResponseStream, err := pb.NewFeedClient(s.clientConn).StreamEvents(ctx, eventRequest)
	if err != nil {
		cancel()

		s.setStatus(StatusRed, s.currentStatus.SettlementsStream)
		return nil, nil, err
	}

	s.eventsStream = make(chan *EventEnvelope)
	s.marketsStream = make(chan *MarketEnvelope)

	s.setStatus(StatusYellow, s.currentStatus.SettlementsStream)
	go func(cancelFunc context.CancelFunc) {
		defer cancelFunc()
		for {
			eventsResponse, err := eventResponseStream.Recv()
			if err != nil {
				s.logger.Errorf("Can't get EventsResponse %v", err)
				close(s.eventsStream)
				close(s.marketsStream)
				s.eventsStream = nil
				s.marketsStream = nil
				s.setStatus(StatusRed, s.currentStatus.SettlementsStream)

				return
			}

			s.publish(eventsResponse)
		}
	}(cancel)

	return s.eventsStream, s.marketsStream, nil
}

// SettlementsFeed returns channel of state updates for event settlements from specific point of time.
// when there is communication errors with X-feed servers it closes the channels
func (s *Session) SettlementsFeed(lastConsumed time.Time) (<-chan *EventSettlementEnvelope, error) {
	if s == nil {
		s.setStatus(s.currentStatus.EventsStreamStatus, StatusRed)
		return nil, errors.New("session is not initialised")
	}
	s.eventSettlementsMutex.Lock()
	defer s.eventSettlementsMutex.Unlock()

	if s.eventSettlements != nil {
		return s.eventSettlements, nil
	}

	lConsumed, err := ptypes.TimestampProto(lastConsumed)
	if err != nil {

		s.setStatus(s.currentStatus.EventsStreamStatus, StatusRed)
		return nil, errors.Wrapf(err, "timestamp %v is invalid time", lastConsumed)
	}

	settlementRequest := &pb.StreamSettlementsRequest{
		ClientName:            s.clientID,
		LastConsumedTimestamp: lConsumed,
	}

	err = s.limiter.Wait(context.Background())
	if err != nil {

		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	settlementResponseStream, err := pb.NewFeedClient(s.clientConn).StreamSettlements(ctx, settlementRequest)
	if err != nil {
		cancel()

		return nil, err
	}

	s.eventSettlements = make(chan *EventSettlementEnvelope)

	s.setStatus(s.currentStatus.EventsStreamStatus, StatusYellow)
	go func(cancelFunc context.CancelFunc) {
		defer cancelFunc()
		for {
			settlementResponse, err := settlementResponseStream.Recv()
			if err != nil {
				s.logger.Errorf("can't get settlementResponse %v", err)
				close(s.eventSettlements)
				s.eventSettlements = nil
				s.setStatus(s.currentStatus.EventsStreamStatus, StatusRed)

				return
			}

			generatedTs, err := parseTimestamp(settlementResponse.GetDiffTimestamp())
			if err != nil {
				generatedTs = time.Now()
			}

			if recoveryComplete := settlementResponse.GetRecoveryComplete(); recoveryComplete == nil {
				s.setStatus(s.currentStatus.EventsStreamStatus, StatusGreen)
			}

			if eventSettlements := settlementResponse.GetMultipleEventsSettlement(); eventSettlements == nil {
				s.logger.Infof("eventSettlements is empty %v", err)

				continue
			}
			for _, eventSettlement := range settlementResponse.GetMultipleEventsSettlement().GetEventSettlement() {
				s.eventSettlements <- &EventSettlementEnvelope{
					EventSettlement: newEventSettlement(eventSettlement),
					GeneratedAt:     &generatedTs,
				}
			}
		}
	}(cancel)

	return s.eventSettlements, nil
}

// Entities returns current snapshot of SportDescriptions.
// when there is communication errors with X-feed servers error is returned
func (s *Session) Entities(language string) ([]*SportDescription, error) {
	if s == nil {
		return nil, errors.New("session is not initialised")
	}
	s.entitiesMutex.Lock()
	defer s.entitiesMutex.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), s.requestTimeout)
	defer cancel()

	in := &pb.SportDescriptionsRequest{
		Lang: language,
	}

	err := s.limiter.Wait(context.Background())
	if err != nil {

		return nil, err
	}

	entities, err := pb.NewFeedClient(s.clientConn).GetSportDescriptions(ctx, in)
	if err != nil {

		return nil, errors.Wrap(err, "can't get SportEntities")
	}

	result := make([]*SportDescription, 0, len(entities.GetSportDescriptions()))
	for _, sportDescription := range entities.GetSportDescriptions() {
		result = append(result, newSportDescription(sportDescription, language))
	}

	return result, nil
}

func (s *Session) publish(eventsResponse *pb.StreamEventsResponse) {
	if recoveryComplete := eventsResponse.GetRecoveryComplete(); recoveryComplete != nil {
		s.setStatus(StatusGreen, s.currentStatus.SettlementsStream)
	}

	if diff := eventsResponse.GetDiffsMessage(); diff != nil {
		generatedTs, err := parseTimestamp(eventsResponse.GetGeneratedTs())
		if err != nil {
			generatedTs = time.Now()
		}

		func(eventDiffs []*pb.EventDiff) {
			for _, eventDiff := range eventDiffs {
				if event := eventDiff.GetEvent(); event == nil {

					continue
				}
				e, err := newEvent(eventDiff.GetEvent())
				if err != nil {
					s.logger.Debugf("can't parse FeedEvent error: %v, message: %+v", err, eventDiff.GetEvent())
					// TODO: fix this after model stabilizing
				}

				s.eventsStream <- &EventEnvelope{
					EventDiff:   e,
					GeneratedAt: &generatedTs,
					Action:      newFeedAction(eventDiff.GetDiffType()),
				}
			}
		}(diff.GetEventDiffs())

		func(marketDiffs []*pb.MarketsDiff) {
			for _, marketsDiffs := range marketDiffs {
				for _, marketDiff := range marketsDiffs.GetEventMarketsDiffs() {
					if marketDiff == nil || marketDiff.GetMarket() == nil {

						continue
					}
					market := newMarket(marketDiff.GetMarket())

					s.marketsStream <- &MarketEnvelope{
						EventID:     marketsDiffs.GetEventId(),
						MarketDiff:  market,
						GeneratedAt: &generatedTs,
						Action:      newFeedAction(marketDiff.GetDiffType()),
					}
				}
			}
		}(diff.GetMarketDiffs())
	}
}

func parseTimestamp(genTs *timestamp.Timestamp) (time.Time, error) {
	var generatedTs time.Time
	var err error
	if genTs != nil {
		generatedTs, err = ptypes.Timestamp(genTs)
		if err != nil {

			return generatedTs, errors.Errorf("can't parse x-feed GeneratedTs timestamp: %v, err: %v", genTs, err)
		}
	}

	if generatedTs.IsZero() {
		return generatedTs, errors.Errorf("generatedTs timestamp: %v is zero", generatedTs)
	}

	return generatedTs, nil
}

func (s *Session) setStatus(eventsStreamStatus, settlementsStream EntitiesStreamStatus) {
	s.statusMutex.Lock()
	s.currentStatus.EventsStreamStatus = eventsStreamStatus
	s.currentStatus.SettlementsStream = settlementsStream

	s.statusChangeHandler(s.currentStatus)
	s.statusMutex.Unlock()
}
