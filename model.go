package xa

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	pb "github.com/x-feed/x-feed-sdk-golang/pkg/feed"
	"time"
)

// sport description DTOs
type (
	SportDescription struct {
		SportID     int32
		SportName   string
		Periods     []*Period
		MarketTypes []*MarketType
	}

	Period struct {
		PeriodID   int32
		PeriodName string
	}

	MarketType struct {
		MarketTypeID       int32
		MarketNameTemplate string
		OutcomeTypes       []*OutcomeType
	}

	OutcomeType struct {
		OutcomeTypeID       int32
		OutcomeNameTemplate string
	}
)

// Event DTOs
type (
	Event struct {
		EventId      string
		SportId      int32
		Category     string
		League       string
		Status       EventStatus
		Start        *time.Time
		Participants []string
		Timer        *EventTimer
	}
	EventStatus int32 // just consts

	EventTimer struct {
		Changed *time.Time
		Time    *time.Duration
		State   TimerState
	}
	TimerState int32 // constants
)

const (
	EventStatusUnknown  EventStatus = 0
	EventStatusPrematch EventStatus = 1
	EventStatusLive     EventStatus = 2
)

const (
	TimerStateUnknown  TimerState = 0
	TimerStateForward  TimerState = 1
	TimerStateBackward TimerState = 2
	TimerStatePause    TimerState = 3
)

// market DTOs
type (
	Market struct {
		MarketId     string
		MarketType   int32
		MarketParams []*MarketParam
		Outcomes     []*Outcome
	}
	MarketParam struct {
		Type  MarketParamType
		Value string
	}
	MarketParamType int32 // constants
	Outcome         struct {
		OutcomeId   string
		OutcomeType int32
		Value       string
		Suspended   bool
	}
)

// settlement DTO
type (
	EventSettlement struct {
		EventId   string
		Resulting *Resulting
		Outcomes  map[string]*OutcomeSettlement
	}
	Resulting struct {
		ResultGroups []*ResultGroup
	}
	ResultGroup struct {
		ResultGroupId int32
		Params        *ResultGroupPeriod
		Results       []*ResultGroupResult
	}
	ResultGroupPeriod int32
	ResultGroupResult struct {
		Params *ResultGroupResultTeam
		Value  int32
	}
	ResultGroupResultTeam int32

	OutcomeSettlement       OutcomeSettlementStatus
	OutcomeSettlementStatus int32 // constants
)

func NewSportDescription(sportDescription pb.SportDescription) *SportDescription {
	result := &SportDescription{
		SportID:   sportDescription.GetSportId(),
		SportName: sportDescription.GetSportName(),
	}

	result.Periods = make([]*Period, 0, len(sportDescription.GetPeriods()))
	for _, period := range sportDescription.GetPeriods() {
		result.Periods = append(result.Periods, &Period{
			PeriodID:   period.GetPeriodId(),
			PeriodName: period.GetPeriodName(),
		})
	}

	result.MarketTypes = make([]*MarketType, 0, len(sportDescription.GetMarketTypes()))
	for _, marketType := range sportDescription.GetMarketTypes() {
		mt := &MarketType{
			MarketTypeID:       marketType.GetMarketTypeId(),
			MarketNameTemplate: marketType.GetMarketNameTemplate(),
		}
		mt.OutcomeTypes = make([]*OutcomeType, 0, len(marketType.GetOutcomeTypes()))
		for _, outcomeType := range marketType.GetOutcomeTypes() {
			mt.OutcomeTypes = append(mt.OutcomeTypes, &OutcomeType{
				OutcomeTypeID:       outcomeType.GetOutcomeTypeId(),
				OutcomeNameTemplate: outcomeType.GetOutcomeNameTemplate(),
			})
		}
	}

	return result
}

func NewEvent(feedEvent pb.FeedEvent) (*Event, error) {
	startTs, err := ptypes.Timestamp(feedEvent.GetStartTs())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse Event StartTs")
	}

	timer, err := NewEventTimer(feedEvent.GetTimer())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse Event Timer")
	}

	return &Event{
		EventId:      feedEvent.GetEventId(),
		SportId:      feedEvent.GetSportId(),
		Category:     feedEvent.GetCategory(),
		League:       feedEvent.GetLeague(),
		Status:       NewEventStatus(feedEvent.GetStatus()),
		Start:        &startTs,
		Participants: feedEvent.GetParticipants(),
		Timer:        timer,
	}, nil
}

func NewEventStatus(eventStatus pb.FeedEvent_EventStatus) EventStatus {
	switch eventStatus {
	case pb.FeedEvent_LIVE:
		return EventStatusLive
	case pb.FeedEvent_PREMATCH:
		return EventStatusPrematch
	case pb.FeedEvent_unknown:
		fallthrough
	default:
		return EventStatusUnknown
	}
}

func NewEventTimer(eventTimer *pb.EventTimer) (*EventTimer, error) {
	changedTs, err := ptypes.Timestamp(eventTimer.GetChangedTs())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse EventTimer ChangedTs")
	}

	eventTime, err := ptypes.Duration(eventTimer.GetTime())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse EventTimer Time")
	}

	state := NewTimerState(eventTimer.GetState())

	return &EventTimer{
		Changed: &changedTs,
		Time:    &eventTime,
		State:   state,
	}, nil
}

func NewTimerState(timerState pb.EventTimer_TimerState) TimerState {
	switch timerState {
	case pb.EventTimer_FORWARD:
		return TimerStateForward
	case pb.EventTimer_BACKWARD:
		return TimerStateBackward
	case pb.EventTimer_PAUSE:
		return TimerStatePause
	case pb.EventTimer_unknown:
		fallthrough
	default:
		return TimerStateUnknown
	}
}
