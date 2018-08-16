package xa

import pb "github.com/x-feed/x-feed-sdk-golang/pkg/feed"

type (
	SportDescription pb.SportDescription
	Period           pb.Period
	MarketType       pb.MarketType
	OutcomeType      pb.OutcomeType

	Event       pb.FeedEvent
	EventStatus pb.FeedEvent_EventStatus // just const

	EventTimer      pb.EventTimer
	EventTimerState pb.EventTimer_TimerState // constants

	Market          pb.FeedMarket
	MarketParam     pb.FeedMarketParam
	MarketParamType pb.FeedMarketParam_MarketParamType // constants
	Outcome         pb.FeedOutcome

	EventSettlement       pb.EventSettlement
	Resulting             pb.FeedResulting
	ResultGroup           pb.FeedResulting_ResultGroup
	ResultGroupPeriod     pb.FeedResulting_ResultGroup_ResultGroupParams
	ResultGroupResult     pb.FeedResulting_ResultGroup_Result
	ResultGroupResultTeam pb.FeedResulting_ResultGroup_Result_ResultParams

	OutcomeSettlement       pb.OutcomeSettlement
	OutcomeSettlementStatus pb.OutcomeSettlement_SettlementType // constants
)
