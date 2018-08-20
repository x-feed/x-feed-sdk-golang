package xfeed

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	pb "github.com/x-feed/x-feed-sdk-golang/pkg/feed"
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

type (
	EventPoints struct {
		PointGroups []*PointsGroup
	}

	PointsGroup struct {
		PointType     PointType
		GroupPeriodID int32
		State         []*State
	}

	PointType int32

	State struct {
		Participant int32
		Value       int32
	}
)

const (
	PointTypeUnknown     PointType = 0
	PointTypeScore       PointType = 1
	PointTypeRedСards    PointType = 2
	PointTypeYellowСards PointType = 3
	PointTypePenalties   PointType = 4
	PointTypeCorners     PointType = 5
)

// Event DTOs
type (
	FeedAction int32

	EventEnvelope struct {
		EventDiff   *Event
		GeneratedAt *time.Time
		Action      FeedAction
	}

	Event struct {
		EventID      string
		SportID      int32
		Category     string
		League       string
		Status       EventStatus
		Start        *time.Time
		Participants []string
		Timer        *EventTimer
		Statistics   *EventPoints
	}

	EventStatus int32

	EventTimer struct {
		Changed *time.Time
		Time    *time.Duration
		State   TimerState
	}

	TimerState int32
)

const (
	Unknown FeedAction = 0
	Insert  FeedAction = 1
	Delete  FeedAction = 2
	Update  FeedAction = 3
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
	MarketEnvelope struct {
		EventID     string
		MarketDiff  *Market
		GeneratedAt *time.Time
		Action      FeedAction
	}

	Market struct {
		MarketID     string
		MarketType   int32
		MarketParams []*MarketParam
		Outcomes     []*Outcome
	}

	MarketParam struct {
		Type  MarketParamType
		Value string
	}

	MarketParamType int32

	Outcome struct {
		OutcomeID   string
		OutcomeType int32
		Value       string
		Suspended   bool
	}
)

const (
	MarketParamTypeUnknown  MarketParamType = 0
	MarketParamTypePeriod   MarketParamType = 1
	MarketParamTypeTotal    MarketParamType = 2
	MarketParamTypeHandicap MarketParamType = 3
	MarketParamTypeTeam     MarketParamType = 4
)

// settlement DTO
type (
	EventSettlementEnvelope struct {
		EventSettlement *EventSettlement
		GeneratedAt     *time.Time
	}

	EventSettlement struct {
		EventID   string
		Resulting *EventPoints
		Outcomes  map[string]OutcomeSettlementStatus
	}

	OutcomeSettlementStatus int32
)

const (
	OutcomeSettlementUnknown   OutcomeSettlementStatus = 0
	OutcomeSettlementUnsettled OutcomeSettlementStatus = 1
	OutcomeSettlementWin       OutcomeSettlementStatus = 2
	OutcomeSettlementLose      OutcomeSettlementStatus = 3
	OutcomeSettlementReturn    OutcomeSettlementStatus = 4
)

func NewSportDescription(sportDescription *pb.SportDescription) *SportDescription {
	result := &SportDescription{
		SportID:   sportDescription.GetSportId(),
		SportName: sportDescription.GetSportName(),
	}

	result.Periods = make([]*Period, 0, len(sportDescription.GetPeriods()))
	for _, period := range sportDescription.GetPeriods() {
		if period == nil {
			continue
		}
		result.Periods = append(result.Periods, &Period{
			PeriodID:   period.GetPeriodId(),
			PeriodName: period.GetPeriodName(),
		})
	}

	result.MarketTypes = make([]*MarketType, 0, len(sportDescription.GetMarketTypes()))
	for _, marketType := range sportDescription.GetMarketTypes() {
		if marketType == nil {
			continue
		}
		mt := &MarketType{
			MarketTypeID:       marketType.GetMarketTypeId(),
			MarketNameTemplate: marketType.GetMarketNameTemplate(),
		}
		mt.OutcomeTypes = make([]*OutcomeType, 0, len(marketType.GetOutcomeTypes()))
		for _, outcomeType := range marketType.GetOutcomeTypes() {
			if outcomeType == nil {
				continue
			}
			mt.OutcomeTypes = append(mt.OutcomeTypes, &OutcomeType{
				OutcomeTypeID:       outcomeType.GetOutcomeTypeId(),
				OutcomeNameTemplate: outcomeType.GetOutcomeNameTemplate(),
			})
		}
	}

	return result
}

func NewEvent(feedEvent *pb.FeedEvent) (*Event, error) {
	startTs, err := ptypes.Timestamp(feedEvent.GetStartTs())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse Event StartTs")
	}

	timer, err := NewEventTimer(feedEvent.GetTimer())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse Event Timer")
	}

	return &Event{
		EventID:      feedEvent.GetEventId(),
		SportID:      feedEvent.GetSportId(),
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

func NewMarket(feedMarket *pb.FeedMarket) *Market {
	market := &Market{
		MarketID:   feedMarket.GetMarketId(),
		MarketType: feedMarket.GetMarketType(),
	}

	market.MarketParams = make([]*MarketParam, 0, len(feedMarket.GetMarketParams()))
	for _, marketParam := range feedMarket.GetMarketParams() {
		if marketParam == nil {
			continue
		}
		market.MarketParams = append(market.MarketParams, &MarketParam{
			Type:  NewMarketParamType(marketParam.GetType()),
			Value: marketParam.GetValue(),
		})
	}

	market.Outcomes = make([]*Outcome, 0, len(feedMarket.GetOutcomes()))
	for _, outcome := range feedMarket.GetOutcomes() {
		if outcome == nil {
			continue
		}
		market.Outcomes = append(market.Outcomes, &Outcome{
			OutcomeID:   outcome.GetOutcomeId(),
			OutcomeType: outcome.GetOutcomeType(),
			Value:       outcome.GetValue(),
			Suspended:   outcome.GetSuspended(),
		})
	}

	return market
}

func NewMarketParamType(marketParamType pb.FeedMarketParam_MarketParamType) MarketParamType {
	switch marketParamType {
	case pb.FeedMarketParam_TEAM:
		return MarketParamTypeTeam
	case pb.FeedMarketParam_HANDICAP:
		return MarketParamTypeHandicap
	case pb.FeedMarketParam_TOTAL:
		return MarketParamTypeTotal
	case pb.FeedMarketParam_PERIOD:
		return MarketParamTypePeriod
	case pb.FeedMarketParam_unknown:
		fallthrough
	default:
		return MarketParamTypeUnknown
	}
}

func NewEventSettlement(settlement *pb.EventSettlement) *EventSettlement {
	eventSettlement := &EventSettlement{
		EventID:   settlement.GetEventId(),
		Resulting: NewEventPoints(settlement.GetResulting()),
		Outcomes:  make(map[string]OutcomeSettlementStatus),
	}
	for outcomeID, settlementStatus := range settlement.GetOutcomes() {
		if settlementStatus == nil {
			eventSettlement.Outcomes[outcomeID] = OutcomeSettlementUnknown
			continue
		}
		status := settlementStatus.GetSettlement()
		eventSettlement.Outcomes[outcomeID] = NewOutcomeSettlementStatus(status)
	}

	return eventSettlement
}

func NewEventPoints(eventPoints *pb.EventPoints) *EventPoints {
	if eventPoints == nil {
		return nil
	}

	points := &EventPoints{
		PointGroups: make([]*PointsGroup, 0, len(eventPoints.GetPointGroups())),
	}
	for _, pointGroup := range eventPoints.GetPointGroups() {
		var periodID int32
		var states []*State
		if pointGroup.GetGroupParams() != nil {
			periodID = pointGroup.GetGroupParams().Period
		}

		for _, state := range pointGroup.GetState() {
			if state == nil {
				continue
			}
			s := &State{
				Value: state.GetValue(),
			}
			if state.GetStateParams() != nil {
				s.Participant = state.GetStateParams().GetParticipant()
			}
			states = append(states, s)
		}

		points.PointGroups = append(points.PointGroups, &PointsGroup{
			PointType:     NewPointType(pointGroup.GetPointType()),
			GroupPeriodID: periodID,
			State:         states,
		})
	}

	return points
}

func NewPointType(pointType pb.PointsGroup_PointType) PointType {
	switch pointType {
	case pb.PointsGroup_SCORE:
		return PointTypeScore
	case pb.PointsGroup_RED_CARDS:
		return PointTypeRedСards
	case pb.PointsGroup_YELLOW_CARDS:
		return PointTypeYellowСards
	case pb.PointsGroup_PENALTIES:
		return PointTypePenalties
	case pb.PointsGroup_CORNERS:
		return PointTypeCorners
	case pb.PointsGroup_unknown:
		fallthrough
	default:
		return PointTypeUnknown
	}
}

func NewOutcomeSettlementStatus(status pb.OutcomeSettlement_SettlementType) OutcomeSettlementStatus {
	switch status {
	case pb.OutcomeSettlement_RETURN:
		return OutcomeSettlementReturn
	case pb.OutcomeSettlement_LOSE:
		return OutcomeSettlementLose
	case pb.OutcomeSettlement_WIN:
		return OutcomeSettlementWin
	case pb.OutcomeSettlement_UNSETTLED:
		return OutcomeSettlementUnsettled
	case pb.OutcomeSettlement_unknown:
		fallthrough
	default:
		return OutcomeSettlementUnknown
	}
}

func NewFeedAction(action pb.DiffType) FeedAction {
	switch action {
	case pb.DiffType_UPDATE:
		return Update
	case pb.DiffType_DELETE:
		return Delete
	case pb.DiffType_INSERT:
		return Insert
	case pb.DiffType_unknown:
		fallthrough
	default:
		return Unknown
	}
}
