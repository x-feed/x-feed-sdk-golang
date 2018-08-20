package xfeed

import (
	"context"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	pb "github.com/x-feed/x-feed-sdk-golang/pkg/feed"
	"github.com/x-feed/x-feed-sdk-golang/pkg/logger"
	"golang.org/x/time/rate"
)

type Session struct {
	lg logger.LogEntry
	requestTimeout time.Duration

	m sync.Mutex
	limiter        *rate.Limiter
	client         pb.FeedClient
	eventsStream chan *EventEnvelope
	marketsStream chan *MarketEnvelope
}

func (s *Session) EventsFeed(clientName string) (chan *EventEnvelope, chan *MarketEnvelope,  error) {
	s.m.Lock()
	s.m.Unlock()

	if s.eventsStream != nil && s.marketsStream != nil {
		return s.eventsStream, s.marketsStream, nil
	}

	eventRequest := &pb.StreamEventsRequest{
		ClientName: clientName,
	}
	s.eventsStream = make(chan *EventEnvelope)
	s.marketsStream = make(chan *MarketEnvelope)

	err := s.limiter.Wait(context.Background())
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.requestTimeout)
	eventResponseStream, err := s.client.StreamEvents(ctx, eventRequest)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	go func(cancelFunc context.CancelFunc) {
		defer cancelFunc()
		for {
			eventsResponse, err := eventResponseStream.Recv()
			if err != nil {
				s.lg.Errorf("Can't get EventsResponse %v", err)
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

func (s *Session) SettlementsFeed(clientName string, lastConsumed time.Time) (chan *pb.SettlementResponse, error) {
	lConsumed, err := ptypes.TimestampProto(lastConsumed)
	if err != nil {
		return nil, errors.Wrapf(err, "timestamp %v is invalid time", lastConsumed)
	}

	settlementRequest := &pb.SettlementRequest{
		ClientName:            clientName,
		LastConsumedTimestamp: lConsumed,
	}
	result := make(chan *pb.SettlementResponse)

	err = s.limiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.requestTimeout)
	settlementResponseStream, err := s.client.SettlementFeed(ctx, settlementRequest)
	if err != nil {
		cancel()
		return nil, err
	}

	go func(cancelFunc context.CancelFunc) {
		defer cancelFunc()
		for {
			settlementResponse, err := settlementResponseStream.Recv()
			if err != nil {
				s.log.Errorf("Can't get settlementResponse %v", err)
				close(result)
				return
			}
			result <- settlementResponse
		}
	}(cancel)

	return result, nil
}

func (s *Session) Entities(language string) ([]*pb.SportEntities, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.requestTimeout)
	defer cancel()

	in := &pb.EntitiesRequest{
		Lang: language,
	}

	err := s.limiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}

	entities, err := s.client.Entities(ctx, in)
	if err != nil {
		return nil, errors.Wrap(err, "can't get SportEntities")
	}

	if entities != nil {
		return entities.Sports, nil
	}

	return nil, nil
}

func (s *Session) publish(eventsResponse *pb.StreamEventsResponse) {
	if diff := eventsResponse.GetDiffsMessage(); diff != nil {
		var generatedTs time.Time
		var err error
		if genTs := eventsResponse.GetGeneratedTs(); genTs != nil {
			generatedTs, err = ptypes.Timestamp(genTs)
			if err != nil {
				s.lg.Errorf("can't parse x-feed GeneratedTs timestamp: %v, err: %v", genTs, err)
			}
		}
		if generatedTs.IsZero() {
			s.lg.Errorf("generatedTs timestamp: %v is zero", generatedTs)
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
					s.lg.Errorf("can't parse FeedEvent: %v", err)
					continue
				}

				s.eventsStream <- &EventEnvelope{
					EventDiff: e,
					GeneratedAt: &generatedTs,
					Action: NewFeedAction(eventDiff.GetDiffType()),
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
						EventID: marketsDiffs.GetEventId(),
						MarketDiff: market,
						GeneratedAt: &generatedTs,
						Action: NewFeedAction(marketDiff.GetDiffType()),
					}
				}
			}
			wg.Done()
		}(diff.GetMarketDiffs())

		wg.Wait()
	}
}