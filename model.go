package xa

import (
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
		StartTs      *time.Time
		Participants []string
		Timer        *EventTimer
	}
	EventStatus int32 // just const

	EventTimer struct {
		ChangedTs *time.Time
		Time      *time.Duration
		State     EventTimerState
	}
	EventTimerState int32 // constants
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
