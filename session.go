package xfeed

import (
	"context"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	pb "github.com/x-feed/x-feed-sdk-golang/pkg/feed"
	"github.com/x-feed/x-feed-sdk-golang/pkg/logger"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

type Session struct {
	Lg             logger.LogEntry
	requestTimeout time.Duration

	limiter    *rate.Limiter
	clientConn *grpc.ClientConn

	eventsFeedMutex sync.Mutex
	eventsStream    chan *EventEnvelope
	marketsStream   chan *MarketEnvelope

	eventSettlementsMutex sync.Mutex
	eventSettlements      chan *EventSettlementEnvelope

	entitiesMutex sync.Mutex
}

func (s *Session) EventsFeed(clientName string) (chan *EventEnvelope, chan *MarketEnvelope, error) {
	s.eventsFeedMutex.Lock()
	s.eventsFeedMutex.Unlock()

	if s.eventsStream != nil && s.marketsStream != nil {
		return s.eventsStream, s.marketsStream, nil
	}

	eventRequest := &pb.StreamEventsRequest{
		ClientName: clientName,
	}

	err := s.limiter.Wait(context.Background())
	if err != nil {

		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.requestTimeout)
	eventResponseStream, err := pb.NewFeedClient(s.clientConn).StreamEvents(ctx, eventRequest)
	if err != nil {
		cancel()

		return nil, nil, err
	}

	s.eventsStream = make(chan *EventEnvelope)
	s.marketsStream = make(chan *MarketEnvelope)

	go func(cancelFunc context.CancelFunc) {
		defer cancelFunc()
		for {
			eventsResponse, err := eventResponseStream.Recv()
			if err != nil {
				s.Lg.Errorf("Can't get EventsResponse %v", err)
				close(s.eventsStream)
				close(s.marketsStream)
				s.eventsStream = nil
				s.marketsStream = nil

				return
			}

			s.publish(eventsResponse)
		}
	}(cancel)

	return s.eventsStream, s.marketsStream, nil
}

func (s *Session) SettlementsFeed(clientName string, lastConsumed time.Time) (chan *EventSettlementEnvelope, error) {
	s.eventSettlementsMutex.Lock()
	s.eventSettlementsMutex.Unlock()

	if s.eventSettlements != nil {
		return s.eventSettlements, nil
	}

	lConsumed, err := ptypes.TimestampProto(lastConsumed)
	if err != nil {

		return nil, errors.Wrapf(err, "timestamp %v is invalid time", lastConsumed)
	}

	settlementRequest := &pb.StreamSettlementsRequest{
		ClientName:            clientName,
		LastConsumedTimestamp: lConsumed,
	}

	err = s.limiter.Wait(context.Background())
	if err != nil {

		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.requestTimeout)
	settlementResponseStream, err := pb.NewFeedClient(s.clientConn).StreamSettlements(ctx, settlementRequest)
	if err != nil {
		cancel()

		return nil, err
	}

	s.eventSettlements = make(chan *EventSettlementEnvelope)

	go func(cancelFunc context.CancelFunc) {
		defer cancelFunc()
		for {
			settlementResponse, err := settlementResponseStream.Recv()
			if err != nil {
				s.Lg.Errorf("Can't get settlementResponse %v", err)
				close(s.eventSettlements)
				s.eventSettlements = nil

				return
			}

			generatedTs, err := parseTimestamp(settlementResponse.GetDiffTimestamp())
			if err != nil {
				generatedTs = time.Now()
			}

			if eventSettlements := settlementResponse.GetMultipleEventsSettlement(); eventSettlements == nil {

				continue
			}
			for _, eventSettlement := range settlementResponse.GetMultipleEventsSettlement().GetEventSettlement() {
				s.eventSettlements <- &EventSettlementEnvelope{
					EventSettlement: NewEventSettlement(eventSettlement),
					GeneratedAt:     &generatedTs,
				}
			}
		}
	}(cancel)

	return s.eventSettlements, nil
}

func (s *Session) Entities(language string) ([]*SportDescription, error) {
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
		result = append(result, NewSportDescription(sportDescription))
	}

	return result, nil
}

func (s *Session) publish(eventsResponse *pb.StreamEventsResponse) {
	if diff := eventsResponse.GetDiffsMessage(); diff != nil {
		generatedTs, err := parseTimestamp(eventsResponse.GetGeneratedTs())
		if err != nil {
			generatedTs = time.Now()
		}

		// Looks like it is better to publish Events and Markets diffs for single event in parallel
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func(eventDiffs []*pb.EventDiff) {
			for _, eventDiff := range eventDiffs {
				if event := eventDiff.GetEvent(); event == nil {

					continue
				}
				e, err := NewEvent(eventDiff.GetEvent())
				if err != nil {
					s.Lg.Errorf("can't parse FeedEvent: %v", err)

					continue
				}

				s.eventsStream <- &EventEnvelope{
					EventDiff:   e,
					GeneratedAt: &generatedTs,
					Action:      NewFeedAction(eventDiff.GetDiffType()),
				}
			}
			wg.Done()
		}(diff.GetEventDiffs())

		go func(marketDiffs []*pb.MarketsDiff) {
			for _, marketsDiffs := range marketDiffs {
				for _, marketDiff := range marketsDiffs.GetEventMarketsDiffs() {
					if marketDiff == nil || marketDiff.GetMarket() == nil {

						continue
					}
					market := NewMarket(marketDiff.GetMarket())

					s.marketsStream <- &MarketEnvelope{
						EventID:     marketsDiffs.GetEventId(),
						MarketDiff:  market,
						GeneratedAt: &generatedTs,
						Action:      NewFeedAction(marketDiff.GetDiffType()),
					}
				}
			}
			wg.Done()
		}(diff.GetMarketDiffs())

		wg.Wait()
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
