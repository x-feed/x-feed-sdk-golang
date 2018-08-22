package xfeed

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	pb "github.com/x-feed/x-feed-sdk-golang/pkg/feed"
)

type (
	// SportDescription is sent for each sport which x-feed supports.
	// It contains list of Periods and MarketTypes which are used for specific sport
	SportDescription struct {
		ID          int32
		Name        string
		Periods     []*Period
		MarketTypes []*MarketType
	}

	// Period represents Sport specific timespan of the game, e. g. full time, first half, set, etc.
	Period struct {
		ID   int32
		Name string
	}

	// MarketType contains template for the Market name
	// Variables:
	// "{%participant}" - get participant name by number in market_param "team" (1, 2)
	// "{$participantN}" - participant by predefined number.
	// "{+$handicap}", "{-$handicap}" - market_param "handicap"
	// "{$total}" - market_param "total"
	MarketType struct {
		ID int32
		// event.participants = ["Dinamo", "Shakhtar"]
		// market_params.team = 1
		// Ex.: "{%participant} Total" -> "Dinamo Total"
		NameTemplate string
		OutcomeTypes []*OutcomeType
	}

	// OutcomeType contains template for the Outcome name
	OutcomeType struct {
		ID int32
		// event.participants = ["Dinamo", "Shakhtar"]
		// market_params.handicap = 1.5
		// Ex.: "{$participant2} ({-$handicap})" -> "Shakhtar (-1.5)"
		NameTemplate string
	}

	// EventPoints represents live game resulting or final game statistics
	EventPoints struct {
		PointGroups []*PointsGroup
	}

	// PointsGroup represents statistics unit of specific type for specific game period
	PointsGroup struct {
		PointType     PointType
		GroupPeriodID int32
		State         []*State
	}

	// PointType represents type of statistics unit: Score, Red cards, Corners, etc.
	PointType int32

	// State represents statistics entry for specific participant
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

type (
	// FeedAction defines operation which shall be done on entity: Insert, Update, Delete
	FeedAction int32

	// EventEnvelope represents state update of specific Event
	EventEnvelope struct {
		EventDiff   *Event
		GeneratedAt *time.Time
		Action      FeedAction
	}

	// Event represents Sport event (game) within specific sport/category/league
	Event struct {
		ID           string
		SportID      int32
		Category     string
		League       string
		Status       EventStatus
		Start        *time.Time
		Participants []string
		Timer        *EventTimer
		Statistics   *EventPoints
	}

	// EventStatus represents state of event, prematch or live
	EventStatus int32

	// EventTimer represents game timer. For some Sports time direction is counterclockwise (for example for basketball)
	EventTimer struct {
		Changed *time.Time
		Time    *time.Duration
		State   TimerState
	}

	// TimerState represents state of the timer: paused, forward, Backward
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
	// MarketEnvelope represents state update of specific Market
	MarketEnvelope struct {
		EventID     string
		MarketDiff  *Market
		GeneratedAt *time.Time
		Action      FeedAction
	}

	// Market represents market instance
	Market struct {
		ID           string
		MarketTypeID int32
		MarketParams []*MarketParam
		Outcomes     []*Outcome
	}

	// MarketParam is key value for specific market parameter (for parametrized markets)
	MarketParam struct {
		Type  MarketParamType
		Value string
	}

	// MarketParamType represents type of parameter of parametrised market
	MarketParamType int32

	// Outcome represents specific instance of outcome
	Outcome struct {
		ID        string
		Type      int32
		Value     string
		Suspended bool
	}
)

const (
	MarketParamTypeUnknown  MarketParamType = 0
	MarketParamTypePeriod   MarketParamType = 1
	MarketParamTypeTotal    MarketParamType = 2
	MarketParamTypeHandicap MarketParamType = 3
	MarketParamTypeTeam     MarketParamType = 4
)

type (
	// EventSettlementEnvelope represents state update of specific EventSettlement
	EventSettlementEnvelope struct {
		EventSettlement *EventSettlement
		GeneratedAt     *time.Time
	}

	// EventSettlement contains settelments for Outcomes for specific event
	EventSettlement struct {
		EventID   string
		Resulting *EventPoints
		Outcomes  map[string]OutcomeSettlementStatus
	}

	// OutcomeSettlementStatus represents result status of outcome: unsettled, win, lose, return
	OutcomeSettlementStatus int32
)

const (
	OutcomeSettlementUnknown   OutcomeSettlementStatus = 0
	OutcomeSettlementUnsettled OutcomeSettlementStatus = 1
	OutcomeSettlementWin       OutcomeSettlementStatus = 2
	OutcomeSettlementLose      OutcomeSettlementStatus = 3
	OutcomeSettlementReturn    OutcomeSettlementStatus = 4
)

func newSportDescription(sportDescription *pb.SportDescription) *SportDescription {
	result := &SportDescription{
		ID:   sportDescription.GetSportId(),
		Name: sportDescription.GetSportName(),
	}

	result.Periods = make([]*Period, 0, len(sportDescription.GetPeriods()))
	for _, period := range sportDescription.GetPeriods() {
		if period == nil {
			continue
		}
		result.Periods = append(result.Periods, &Period{
			ID:   period.GetPeriodId(),
			Name: period.GetPeriodName(),
		})
	}

	result.MarketTypes = make([]*MarketType, 0, len(sportDescription.GetMarketTypes()))
	for _, marketType := range sportDescription.GetMarketTypes() {
		if marketType == nil {
			continue
		}
		mt := &MarketType{
			ID:           marketType.GetMarketTypeId(),
			NameTemplate: marketType.GetMarketNameTemplate(),
		}
		mt.OutcomeTypes = make([]*OutcomeType, 0, len(marketType.GetOutcomeTypes()))
		for _, outcomeType := range marketType.GetOutcomeTypes() {
			if outcomeType == nil {
				continue
			}
			mt.OutcomeTypes = append(mt.OutcomeTypes, &OutcomeType{
				ID:           outcomeType.GetOutcomeTypeId(),
				NameTemplate: outcomeType.GetOutcomeNameTemplate(),
			})
		}
	}

	return result
}

func newEvent(feedEvent *pb.FeedEvent) (*Event, error) {
	startTs, err := ptypes.Timestamp(feedEvent.GetStartTs())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse Event StartTs")
	}

	timer, err := newEventTimer(feedEvent.GetTimer())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse Event Timer")
	}

	return &Event{
		ID:           feedEvent.GetEventId(),
		SportID:      feedEvent.GetSportId(),
		Category:     feedEvent.GetCategory(),
		League:       feedEvent.GetLeague(),
		Status:       newEventStatus(feedEvent.GetStatus()),
		Start:        &startTs,
		Participants: feedEvent.GetParticipants(),
		Timer:        timer,
	}, nil
}

func newEventStatus(eventStatus pb.FeedEvent_EventStatus) EventStatus {
	switch eventStatus {
	case pb.FeedEvent_LIVE:
		return EventStatusLive
	case pb.FeedEvent_PREMATCH:
		return EventStatusPrematch
	default:
		return EventStatusUnknown
	}
}

func newEventTimer(eventTimer *pb.EventTimer) (*EventTimer, error) {
	changedTs, err := ptypes.Timestamp(eventTimer.GetChangedTs())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse EventTimer ChangedTs")
	}

	eventTime, err := ptypes.Duration(eventTimer.GetTime())
	if err != nil {
		return nil, errors.Wrap(err, "can't parse EventTimer Time")
	}

	state := newTimerState(eventTimer.GetState())

	return &EventTimer{
		Changed: &changedTs,
		Time:    &eventTime,
		State:   state,
	}, nil
}

func newTimerState(timerState pb.EventTimer_TimerState) TimerState {
	switch timerState {
	case pb.EventTimer_FORWARD:
		return TimerStateForward
	case pb.EventTimer_BACKWARD:
		return TimerStateBackward
	case pb.EventTimer_PAUSE:
		return TimerStatePause
	default:
		return TimerStateUnknown
	}
}

func newMarket(feedMarket *pb.FeedMarket) *Market {
	market := &Market{
		ID:           feedMarket.GetMarketId(),
		MarketTypeID: feedMarket.GetMarketType(),
	}

	market.MarketParams = make([]*MarketParam, 0, len(feedMarket.GetMarketParams()))
	for _, marketParam := range feedMarket.GetMarketParams() {
		if marketParam == nil {
			continue
		}
		market.MarketParams = append(market.MarketParams, &MarketParam{
			Type:  newMarketParamType(marketParam.GetType()),
			Value: marketParam.GetValue(),
		})
	}

	market.Outcomes = make([]*Outcome, 0, len(feedMarket.GetOutcomes()))
	for _, outcome := range feedMarket.GetOutcomes() {
		if outcome == nil {
			continue
		}
		market.Outcomes = append(market.Outcomes, &Outcome{
			ID:        outcome.GetOutcomeId(),
			Type:      outcome.GetOutcomeType(),
			Value:     outcome.GetValue(),
			Suspended: outcome.GetSuspended(),
		})
	}

	return market
}

func newMarketParamType(marketParamType pb.FeedMarketParam_MarketParamType) MarketParamType {
	switch marketParamType {
	case pb.FeedMarketParam_TEAM:
		return MarketParamTypeTeam
	case pb.FeedMarketParam_HANDICAP:
		return MarketParamTypeHandicap
	case pb.FeedMarketParam_TOTAL:
		return MarketParamTypeTotal
	case pb.FeedMarketParam_PERIOD:
		return MarketParamTypePeriod
	default:
		return MarketParamTypeUnknown
	}
}

func newEventSettlement(settlement *pb.EventSettlement) *EventSettlement {
	eventSettlement := &EventSettlement{
		EventID:   settlement.GetEventId(),
		Resulting: newEventPoints(settlement.GetResulting()),
		Outcomes:  make(map[string]OutcomeSettlementStatus),
	}
	for outcomeID, settlementStatus := range settlement.GetOutcomes() {
		if settlementStatus == nil {
			eventSettlement.Outcomes[outcomeID] = OutcomeSettlementUnknown
			continue
		}
		status := settlementStatus.GetSettlement()
		eventSettlement.Outcomes[outcomeID] = newOutcomeSettlementStatus(status)
	}

	return eventSettlement
}

func newEventPoints(eventPoints *pb.EventPoints) *EventPoints {
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
			PointType:     newPointType(pointGroup.GetPointType()),
			GroupPeriodID: periodID,
			State:         states,
		})
	}

	return points
}

func newPointType(pointType pb.PointsGroup_PointType) PointType {
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
	default:
		return PointTypeUnknown
	}
}

func newOutcomeSettlementStatus(status pb.OutcomeSettlement_SettlementType) OutcomeSettlementStatus {
	switch status {
	case pb.OutcomeSettlement_RETURN:
		return OutcomeSettlementReturn
	case pb.OutcomeSettlement_LOSE:
		return OutcomeSettlementLose
	case pb.OutcomeSettlement_WIN:
		return OutcomeSettlementWin
	case pb.OutcomeSettlement_UNSETTLED:
		return OutcomeSettlementUnsettled
	default:
		return OutcomeSettlementUnknown
	}
}

func newFeedAction(action pb.DiffType) FeedAction {
	switch action {
	case pb.DiffType_UPDATE:
		return Update
	case pb.DiffType_DELETE:
		return Delete
	case pb.DiffType_INSERT:
		return Insert
	default:
		return Unknown
	}
}
