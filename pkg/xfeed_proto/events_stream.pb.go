// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xfeed_proto/events_stream.proto

package xfeed_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DiffType int32

const (
	DiffType_unknown DiffType = 0
	DiffType_INSERT  DiffType = 1
	DiffType_DELETE  DiffType = 2
	DiffType_UPDATE  DiffType = 3
)

var DiffType_name = map[int32]string{
	0: "unknown",
	1: "INSERT",
	2: "DELETE",
	3: "UPDATE",
}
var DiffType_value = map[string]int32{
	"unknown": 0,
	"INSERT":  1,
	"DELETE":  2,
	"UPDATE":  3,
}

func (x DiffType) String() string {
	return proto.EnumName(DiffType_name, int32(x))
}
func (DiffType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{0}
}

type EventTimer_TimerState int32

const (
	EventTimer_unknown  EventTimer_TimerState = 0
	EventTimer_FORWARD  EventTimer_TimerState = 1
	EventTimer_BACKWARD EventTimer_TimerState = 2
	EventTimer_PAUSE    EventTimer_TimerState = 3
)

var EventTimer_TimerState_name = map[int32]string{
	0: "unknown",
	1: "FORWARD",
	2: "BACKWARD",
	3: "PAUSE",
}
var EventTimer_TimerState_value = map[string]int32{
	"unknown":  0,
	"FORWARD":  1,
	"BACKWARD": 2,
	"PAUSE":    3,
}

func (x EventTimer_TimerState) String() string {
	return proto.EnumName(EventTimer_TimerState_name, int32(x))
}
func (EventTimer_TimerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{5, 0}
}

type FeedEvent_EventStatus int32

const (
	FeedEvent_unknown  FeedEvent_EventStatus = 0
	FeedEvent_PREMATCH FeedEvent_EventStatus = 1
	FeedEvent_LIVE     FeedEvent_EventStatus = 2
)

var FeedEvent_EventStatus_name = map[int32]string{
	0: "unknown",
	1: "PREMATCH",
	2: "LIVE",
}
var FeedEvent_EventStatus_value = map[string]int32{
	"unknown":  0,
	"PREMATCH": 1,
	"LIVE":     2,
}

func (x FeedEvent_EventStatus) String() string {
	return proto.EnumName(FeedEvent_EventStatus_name, int32(x))
}
func (FeedEvent_EventStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{6, 0}
}

type FeedMarketParam_MarketParamType int32

const (
	FeedMarketParam_unknown  FeedMarketParam_MarketParamType = 0
	FeedMarketParam_PERIOD   FeedMarketParam_MarketParamType = 1
	FeedMarketParam_TOTAL    FeedMarketParam_MarketParamType = 2
	FeedMarketParam_HANDICAP FeedMarketParam_MarketParamType = 3
	FeedMarketParam_TEAM     FeedMarketParam_MarketParamType = 4
)

var FeedMarketParam_MarketParamType_name = map[int32]string{
	0: "unknown",
	1: "PERIOD",
	2: "TOTAL",
	3: "HANDICAP",
	4: "TEAM",
}
var FeedMarketParam_MarketParamType_value = map[string]int32{
	"unknown":  0,
	"PERIOD":   1,
	"TOTAL":    2,
	"HANDICAP": 3,
	"TEAM":     4,
}

func (x FeedMarketParam_MarketParamType) String() string {
	return proto.EnumName(FeedMarketParam_MarketParamType_name, int32(x))
}
func (FeedMarketParam_MarketParamType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{11, 0}
}

type StreamEventsRequest struct {
	ClientName           string   `protobuf:"bytes,1,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEventsRequest) Reset()         { *m = StreamEventsRequest{} }
func (m *StreamEventsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamEventsRequest) ProtoMessage()    {}
func (*StreamEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{0}
}
func (m *StreamEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventsRequest.Unmarshal(m, b)
}
func (m *StreamEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventsRequest.Marshal(b, m, deterministic)
}
func (dst *StreamEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventsRequest.Merge(dst, src)
}
func (m *StreamEventsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamEventsRequest.Size(m)
}
func (m *StreamEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventsRequest proto.InternalMessageInfo

func (m *StreamEventsRequest) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type StreamEventsResponse struct {
	// Types that are valid to be assigned to Data:
	//	*StreamEventsResponse_DiffsMessage
	//	*StreamEventsResponse_RecoveryComplete
	Data                 isStreamEventsResponse_Data `protobuf_oneof:"data"`
	GeneratedTs          *timestamp.Timestamp        `protobuf:"bytes,3,opt,name=generated_ts,json=generatedTs,proto3" json:"generated_ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StreamEventsResponse) Reset()         { *m = StreamEventsResponse{} }
func (m *StreamEventsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamEventsResponse) ProtoMessage()    {}
func (*StreamEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{1}
}
func (m *StreamEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventsResponse.Unmarshal(m, b)
}
func (m *StreamEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventsResponse.Marshal(b, m, deterministic)
}
func (dst *StreamEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventsResponse.Merge(dst, src)
}
func (m *StreamEventsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamEventsResponse.Size(m)
}
func (m *StreamEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventsResponse proto.InternalMessageInfo

type isStreamEventsResponse_Data interface {
	isStreamEventsResponse_Data()
}

type StreamEventsResponse_DiffsMessage struct {
	DiffsMessage *DiffsMessage `protobuf:"bytes,1,opt,name=diffs_message,json=diffsMessage,proto3,oneof"`
}

type StreamEventsResponse_RecoveryComplete struct {
	RecoveryComplete *EventsRecoveryComplete `protobuf:"bytes,2,opt,name=recovery_complete,json=recoveryComplete,proto3,oneof"`
}

func (*StreamEventsResponse_DiffsMessage) isStreamEventsResponse_Data() {}

func (*StreamEventsResponse_RecoveryComplete) isStreamEventsResponse_Data() {}

func (m *StreamEventsResponse) GetData() isStreamEventsResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *StreamEventsResponse) GetDiffsMessage() *DiffsMessage {
	if x, ok := m.GetData().(*StreamEventsResponse_DiffsMessage); ok {
		return x.DiffsMessage
	}
	return nil
}

func (m *StreamEventsResponse) GetRecoveryComplete() *EventsRecoveryComplete {
	if x, ok := m.GetData().(*StreamEventsResponse_RecoveryComplete); ok {
		return x.RecoveryComplete
	}
	return nil
}

func (m *StreamEventsResponse) GetGeneratedTs() *timestamp.Timestamp {
	if m != nil {
		return m.GeneratedTs
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamEventsResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamEventsResponse_OneofMarshaler, _StreamEventsResponse_OneofUnmarshaler, _StreamEventsResponse_OneofSizer, []interface{}{
		(*StreamEventsResponse_DiffsMessage)(nil),
		(*StreamEventsResponse_RecoveryComplete)(nil),
	}
}

func _StreamEventsResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamEventsResponse)
	// data
	switch x := m.Data.(type) {
	case *StreamEventsResponse_DiffsMessage:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DiffsMessage); err != nil {
			return err
		}
	case *StreamEventsResponse_RecoveryComplete:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RecoveryComplete); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamEventsResponse.Data has unexpected type %T", x)
	}
	return nil
}

func _StreamEventsResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamEventsResponse)
	switch tag {
	case 1: // data.diffs_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DiffsMessage)
		err := b.DecodeMessage(msg)
		m.Data = &StreamEventsResponse_DiffsMessage{msg}
		return true, err
	case 2: // data.recovery_complete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventsRecoveryComplete)
		err := b.DecodeMessage(msg)
		m.Data = &StreamEventsResponse_RecoveryComplete{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamEventsResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamEventsResponse)
	// data
	switch x := m.Data.(type) {
	case *StreamEventsResponse_DiffsMessage:
		s := proto.Size(x.DiffsMessage)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamEventsResponse_RecoveryComplete:
		s := proto.Size(x.RecoveryComplete)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DiffsMessage struct {
	EventDiffs           []*EventDiff   `protobuf:"bytes,1,rep,name=event_diffs,json=eventDiffs,proto3" json:"event_diffs,omitempty"`
	MarketDiffs          []*MarketsDiff `protobuf:"bytes,2,rep,name=market_diffs,json=marketDiffs,proto3" json:"market_diffs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiffsMessage) Reset()         { *m = DiffsMessage{} }
func (m *DiffsMessage) String() string { return proto.CompactTextString(m) }
func (*DiffsMessage) ProtoMessage()    {}
func (*DiffsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{2}
}
func (m *DiffsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiffsMessage.Unmarshal(m, b)
}
func (m *DiffsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiffsMessage.Marshal(b, m, deterministic)
}
func (dst *DiffsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiffsMessage.Merge(dst, src)
}
func (m *DiffsMessage) XXX_Size() int {
	return xxx_messageInfo_DiffsMessage.Size(m)
}
func (m *DiffsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DiffsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DiffsMessage proto.InternalMessageInfo

func (m *DiffsMessage) GetEventDiffs() []*EventDiff {
	if m != nil {
		return m.EventDiffs
	}
	return nil
}

func (m *DiffsMessage) GetMarketDiffs() []*MarketsDiff {
	if m != nil {
		return m.MarketDiffs
	}
	return nil
}

type EventsRecoveryComplete struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventsRecoveryComplete) Reset()         { *m = EventsRecoveryComplete{} }
func (m *EventsRecoveryComplete) String() string { return proto.CompactTextString(m) }
func (*EventsRecoveryComplete) ProtoMessage()    {}
func (*EventsRecoveryComplete) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{3}
}
func (m *EventsRecoveryComplete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsRecoveryComplete.Unmarshal(m, b)
}
func (m *EventsRecoveryComplete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsRecoveryComplete.Marshal(b, m, deterministic)
}
func (dst *EventsRecoveryComplete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsRecoveryComplete.Merge(dst, src)
}
func (m *EventsRecoveryComplete) XXX_Size() int {
	return xxx_messageInfo_EventsRecoveryComplete.Size(m)
}
func (m *EventsRecoveryComplete) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsRecoveryComplete.DiscardUnknown(m)
}

var xxx_messageInfo_EventsRecoveryComplete proto.InternalMessageInfo

type EventDiff struct {
	DiffType             DiffType   `protobuf:"varint,1,opt,name=diff_type,json=diffType,proto3,enum=xfeed_proto.DiffType" json:"diff_type,omitempty"`
	Event                *FeedEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EventDiff) Reset()         { *m = EventDiff{} }
func (m *EventDiff) String() string { return proto.CompactTextString(m) }
func (*EventDiff) ProtoMessage()    {}
func (*EventDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{4}
}
func (m *EventDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventDiff.Unmarshal(m, b)
}
func (m *EventDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventDiff.Marshal(b, m, deterministic)
}
func (dst *EventDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDiff.Merge(dst, src)
}
func (m *EventDiff) XXX_Size() int {
	return xxx_messageInfo_EventDiff.Size(m)
}
func (m *EventDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDiff.DiscardUnknown(m)
}

var xxx_messageInfo_EventDiff proto.InternalMessageInfo

func (m *EventDiff) GetDiffType() DiffType {
	if m != nil {
		return m.DiffType
	}
	return DiffType_unknown
}

func (m *EventDiff) GetEvent() *FeedEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type EventTimer struct {
	ChangedTs            *timestamp.Timestamp  `protobuf:"bytes,1,opt,name=changed_ts,json=changedTs,proto3" json:"changed_ts,omitempty"`
	Time                 *duration.Duration    `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	State                EventTimer_TimerState `protobuf:"varint,3,opt,name=state,proto3,enum=xfeed_proto.EventTimer_TimerState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EventTimer) Reset()         { *m = EventTimer{} }
func (m *EventTimer) String() string { return proto.CompactTextString(m) }
func (*EventTimer) ProtoMessage()    {}
func (*EventTimer) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{5}
}
func (m *EventTimer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventTimer.Unmarshal(m, b)
}
func (m *EventTimer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventTimer.Marshal(b, m, deterministic)
}
func (dst *EventTimer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTimer.Merge(dst, src)
}
func (m *EventTimer) XXX_Size() int {
	return xxx_messageInfo_EventTimer.Size(m)
}
func (m *EventTimer) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTimer.DiscardUnknown(m)
}

var xxx_messageInfo_EventTimer proto.InternalMessageInfo

func (m *EventTimer) GetChangedTs() *timestamp.Timestamp {
	if m != nil {
		return m.ChangedTs
	}
	return nil
}

func (m *EventTimer) GetTime() *duration.Duration {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *EventTimer) GetState() EventTimer_TimerState {
	if m != nil {
		return m.State
	}
	return EventTimer_unknown
}

type FeedEvent struct {
	EventId      string                `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	SportId      int32                 `protobuf:"varint,2,opt,name=sport_id,json=sportId,proto3" json:"sport_id,omitempty"`
	Category     string                `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	League       string                `protobuf:"bytes,4,opt,name=league,proto3" json:"league,omitempty"`
	Status       FeedEvent_EventStatus `protobuf:"varint,5,opt,name=status,proto3,enum=xfeed_proto.FeedEvent_EventStatus" json:"status,omitempty"`
	StartTs      *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	Participants []string              `protobuf:"bytes,7,rep,name=participants,proto3" json:"participants,omitempty"`
	// only live
	Timer                *EventTimer  `protobuf:"bytes,8,opt,name=timer,proto3" json:"timer,omitempty"`
	Statistics           *EventPoints `protobuf:"bytes,9,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FeedEvent) Reset()         { *m = FeedEvent{} }
func (m *FeedEvent) String() string { return proto.CompactTextString(m) }
func (*FeedEvent) ProtoMessage()    {}
func (*FeedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{6}
}
func (m *FeedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedEvent.Unmarshal(m, b)
}
func (m *FeedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedEvent.Marshal(b, m, deterministic)
}
func (dst *FeedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedEvent.Merge(dst, src)
}
func (m *FeedEvent) XXX_Size() int {
	return xxx_messageInfo_FeedEvent.Size(m)
}
func (m *FeedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_FeedEvent proto.InternalMessageInfo

func (m *FeedEvent) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *FeedEvent) GetSportId() int32 {
	if m != nil {
		return m.SportId
	}
	return 0
}

func (m *FeedEvent) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *FeedEvent) GetLeague() string {
	if m != nil {
		return m.League
	}
	return ""
}

func (m *FeedEvent) GetStatus() FeedEvent_EventStatus {
	if m != nil {
		return m.Status
	}
	return FeedEvent_unknown
}

func (m *FeedEvent) GetStartTs() *timestamp.Timestamp {
	if m != nil {
		return m.StartTs
	}
	return nil
}

func (m *FeedEvent) GetParticipants() []string {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *FeedEvent) GetTimer() *EventTimer {
	if m != nil {
		return m.Timer
	}
	return nil
}

func (m *FeedEvent) GetStatistics() *EventPoints {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type FeedOutcome struct {
	OutcomeId            string   `protobuf:"bytes,1,opt,name=outcome_id,json=outcomeId,proto3" json:"outcome_id,omitempty"`
	OutcomeType          int32    `protobuf:"varint,4,opt,name=outcome_type,json=outcomeType,proto3" json:"outcome_type,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Suspended            bool     `protobuf:"varint,3,opt,name=suspended,proto3" json:"suspended,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedOutcome) Reset()         { *m = FeedOutcome{} }
func (m *FeedOutcome) String() string { return proto.CompactTextString(m) }
func (*FeedOutcome) ProtoMessage()    {}
func (*FeedOutcome) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{7}
}
func (m *FeedOutcome) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedOutcome.Unmarshal(m, b)
}
func (m *FeedOutcome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedOutcome.Marshal(b, m, deterministic)
}
func (dst *FeedOutcome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedOutcome.Merge(dst, src)
}
func (m *FeedOutcome) XXX_Size() int {
	return xxx_messageInfo_FeedOutcome.Size(m)
}
func (m *FeedOutcome) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedOutcome.DiscardUnknown(m)
}

var xxx_messageInfo_FeedOutcome proto.InternalMessageInfo

func (m *FeedOutcome) GetOutcomeId() string {
	if m != nil {
		return m.OutcomeId
	}
	return ""
}

func (m *FeedOutcome) GetOutcomeType() int32 {
	if m != nil {
		return m.OutcomeType
	}
	return 0
}

func (m *FeedOutcome) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FeedOutcome) GetSuspended() bool {
	if m != nil {
		return m.Suspended
	}
	return false
}

type MarketsDiff struct {
	EventId              string              `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventMarketsDiffs    []*SingleMarketDiff `protobuf:"bytes,2,rep,name=event_markets_diffs,json=eventMarketsDiffs,proto3" json:"event_markets_diffs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MarketsDiff) Reset()         { *m = MarketsDiff{} }
func (m *MarketsDiff) String() string { return proto.CompactTextString(m) }
func (*MarketsDiff) ProtoMessage()    {}
func (*MarketsDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{8}
}
func (m *MarketsDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketsDiff.Unmarshal(m, b)
}
func (m *MarketsDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketsDiff.Marshal(b, m, deterministic)
}
func (dst *MarketsDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketsDiff.Merge(dst, src)
}
func (m *MarketsDiff) XXX_Size() int {
	return xxx_messageInfo_MarketsDiff.Size(m)
}
func (m *MarketsDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketsDiff.DiscardUnknown(m)
}

var xxx_messageInfo_MarketsDiff proto.InternalMessageInfo

func (m *MarketsDiff) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *MarketsDiff) GetEventMarketsDiffs() []*SingleMarketDiff {
	if m != nil {
		return m.EventMarketsDiffs
	}
	return nil
}

type SingleMarketDiff struct {
	DiffType             DiffType    `protobuf:"varint,1,opt,name=diff_type,json=diffType,proto3,enum=xfeed_proto.DiffType" json:"diff_type,omitempty"`
	Market               *FeedMarket `protobuf:"bytes,2,opt,name=market,proto3" json:"market,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SingleMarketDiff) Reset()         { *m = SingleMarketDiff{} }
func (m *SingleMarketDiff) String() string { return proto.CompactTextString(m) }
func (*SingleMarketDiff) ProtoMessage()    {}
func (*SingleMarketDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{9}
}
func (m *SingleMarketDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleMarketDiff.Unmarshal(m, b)
}
func (m *SingleMarketDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleMarketDiff.Marshal(b, m, deterministic)
}
func (dst *SingleMarketDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleMarketDiff.Merge(dst, src)
}
func (m *SingleMarketDiff) XXX_Size() int {
	return xxx_messageInfo_SingleMarketDiff.Size(m)
}
func (m *SingleMarketDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleMarketDiff.DiscardUnknown(m)
}

var xxx_messageInfo_SingleMarketDiff proto.InternalMessageInfo

func (m *SingleMarketDiff) GetDiffType() DiffType {
	if m != nil {
		return m.DiffType
	}
	return DiffType_unknown
}

func (m *SingleMarketDiff) GetMarket() *FeedMarket {
	if m != nil {
		return m.Market
	}
	return nil
}

type FeedMarket struct {
	MarketId             string             `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	MarketType           int32              `protobuf:"varint,2,opt,name=market_type,json=marketType,proto3" json:"market_type,omitempty"`
	MarketParams         []*FeedMarketParam `protobuf:"bytes,3,rep,name=market_params,json=marketParams,proto3" json:"market_params,omitempty"`
	Outcomes             []*FeedOutcome     `protobuf:"bytes,4,rep,name=outcomes,proto3" json:"outcomes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FeedMarket) Reset()         { *m = FeedMarket{} }
func (m *FeedMarket) String() string { return proto.CompactTextString(m) }
func (*FeedMarket) ProtoMessage()    {}
func (*FeedMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{10}
}
func (m *FeedMarket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMarket.Unmarshal(m, b)
}
func (m *FeedMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMarket.Marshal(b, m, deterministic)
}
func (dst *FeedMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMarket.Merge(dst, src)
}
func (m *FeedMarket) XXX_Size() int {
	return xxx_messageInfo_FeedMarket.Size(m)
}
func (m *FeedMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMarket.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMarket proto.InternalMessageInfo

func (m *FeedMarket) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *FeedMarket) GetMarketType() int32 {
	if m != nil {
		return m.MarketType
	}
	return 0
}

func (m *FeedMarket) GetMarketParams() []*FeedMarketParam {
	if m != nil {
		return m.MarketParams
	}
	return nil
}

func (m *FeedMarket) GetOutcomes() []*FeedOutcome {
	if m != nil {
		return m.Outcomes
	}
	return nil
}

type FeedMarketParam struct {
	Type                 FeedMarketParam_MarketParamType `protobuf:"varint,1,opt,name=type,proto3,enum=xfeed_proto.FeedMarketParam_MarketParamType" json:"type,omitempty"`
	Value                string                          `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FeedMarketParam) Reset()         { *m = FeedMarketParam{} }
func (m *FeedMarketParam) String() string { return proto.CompactTextString(m) }
func (*FeedMarketParam) ProtoMessage()    {}
func (*FeedMarketParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_stream_e0d6208e9ad5a493, []int{11}
}
func (m *FeedMarketParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMarketParam.Unmarshal(m, b)
}
func (m *FeedMarketParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMarketParam.Marshal(b, m, deterministic)
}
func (dst *FeedMarketParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMarketParam.Merge(dst, src)
}
func (m *FeedMarketParam) XXX_Size() int {
	return xxx_messageInfo_FeedMarketParam.Size(m)
}
func (m *FeedMarketParam) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMarketParam.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMarketParam proto.InternalMessageInfo

func (m *FeedMarketParam) GetType() FeedMarketParam_MarketParamType {
	if m != nil {
		return m.Type
	}
	return FeedMarketParam_unknown
}

func (m *FeedMarketParam) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamEventsRequest)(nil), "xfeed_proto.StreamEventsRequest")
	proto.RegisterType((*StreamEventsResponse)(nil), "xfeed_proto.StreamEventsResponse")
	proto.RegisterType((*DiffsMessage)(nil), "xfeed_proto.DiffsMessage")
	proto.RegisterType((*EventsRecoveryComplete)(nil), "xfeed_proto.EventsRecoveryComplete")
	proto.RegisterType((*EventDiff)(nil), "xfeed_proto.EventDiff")
	proto.RegisterType((*EventTimer)(nil), "xfeed_proto.EventTimer")
	proto.RegisterType((*FeedEvent)(nil), "xfeed_proto.FeedEvent")
	proto.RegisterType((*FeedOutcome)(nil), "xfeed_proto.FeedOutcome")
	proto.RegisterType((*MarketsDiff)(nil), "xfeed_proto.MarketsDiff")
	proto.RegisterType((*SingleMarketDiff)(nil), "xfeed_proto.SingleMarketDiff")
	proto.RegisterType((*FeedMarket)(nil), "xfeed_proto.FeedMarket")
	proto.RegisterType((*FeedMarketParam)(nil), "xfeed_proto.FeedMarketParam")
	proto.RegisterEnum("xfeed_proto.DiffType", DiffType_name, DiffType_value)
	proto.RegisterEnum("xfeed_proto.EventTimer_TimerState", EventTimer_TimerState_name, EventTimer_TimerState_value)
	proto.RegisterEnum("xfeed_proto.FeedEvent_EventStatus", FeedEvent_EventStatus_name, FeedEvent_EventStatus_value)
	proto.RegisterEnum("xfeed_proto.FeedMarketParam_MarketParamType", FeedMarketParam_MarketParamType_name, FeedMarketParam_MarketParamType_value)
}

func init() {
	proto.RegisterFile("xfeed_proto/events_stream.proto", fileDescriptor_events_stream_e0d6208e9ad5a493)
}

var fileDescriptor_events_stream_e0d6208e9ad5a493 = []byte{
	// 979 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x0d, 0x75, 0x33, 0x39, 0x54, 0x12, 0x66, 0x93, 0x3a, 0xb4, 0x9b, 0xd4, 0x2e, 0xfb, 0x62,
	0x14, 0x89, 0x0c, 0xa8, 0xb7, 0xb4, 0x41, 0xd1, 0x30, 0x16, 0x03, 0x0b, 0xb5, 0x6c, 0x61, 0xc5,
	0xb4, 0x8f, 0xc2, 0x46, 0x5c, 0xab, 0x44, 0xc4, 0x4b, 0xb9, 0x2b, 0xbb, 0x7e, 0x6e, 0x7f, 0xa6,
	0x7f, 0xd0, 0xf7, 0x7e, 0x51, 0xd1, 0x1f, 0x28, 0x76, 0x96, 0x92, 0x29, 0xc9, 0x49, 0x80, 0xbe,
	0x71, 0x66, 0xce, 0xdc, 0xce, 0xcc, 0x0e, 0x61, 0xef, 0xb7, 0x73, 0xce, 0xa3, 0x71, 0x5e, 0x64,
	0x32, 0x3b, 0xe4, 0x17, 0x3c, 0x95, 0x62, 0x2c, 0x64, 0xc1, 0x59, 0xd2, 0x41, 0x1d, 0xb1, 0x2b,
	0x80, 0xdd, 0xbd, 0x69, 0x96, 0x4d, 0x67, 0xfc, 0x10, 0xa5, 0x37, 0xf3, 0xf3, 0x43, 0x19, 0x27,
	0x5c, 0x48, 0x96, 0xe4, 0x1a, 0xbd, 0xfb, 0xc9, 0x3a, 0x20, 0x9a, 0x17, 0x4c, 0xc6, 0x59, 0xba,
	0xb0, 0x6f, 0xa4, 0x1b, 0xe7, 0x59, 0x9c, 0x4a, 0xa1, 0xed, 0xde, 0xd7, 0x70, 0x7f, 0x84, 0xd9,
	0x03, 0x2c, 0x85, 0xf2, 0x5f, 0xe7, 0x5c, 0x48, 0xb2, 0x07, 0xf6, 0x64, 0x16, 0x2b, 0x74, 0xca,
	0x12, 0xee, 0x1a, 0xfb, 0xc6, 0x81, 0x45, 0x41, 0xab, 0x4e, 0x59, 0xc2, 0xbd, 0x7f, 0x0d, 0x78,
	0xb0, 0xea, 0x28, 0xf2, 0x2c, 0x15, 0x9c, 0xbc, 0x80, 0xdb, 0x51, 0x7c, 0x7e, 0x2e, 0xc6, 0x09,
	0x17, 0x82, 0x4d, 0xb5, 0xaf, 0xdd, 0xdd, 0xe9, 0x54, 0x0a, 0xe9, 0xf4, 0x14, 0x62, 0xa0, 0x01,
	0xc7, 0xb7, 0x68, 0x3b, 0xaa, 0xc8, 0x84, 0xc2, 0xbd, 0x82, 0x4f, 0xb2, 0x0b, 0x5e, 0x5c, 0x8d,
	0x27, 0x59, 0x92, 0xcf, 0xb8, 0xe4, 0x6e, 0x0d, 0xa3, 0x7c, 0xb6, 0x12, 0x65, 0x91, 0x59, 0x63,
	0x8f, 0x4a, 0xe8, 0xf1, 0x2d, 0xea, 0x14, 0x6b, 0x3a, 0xf2, 0x3d, 0xb4, 0xa7, 0x3c, 0xe5, 0x05,
	0x93, 0x3c, 0x1a, 0x4b, 0xe1, 0xd6, 0x31, 0xdc, 0x6e, 0x47, 0xb3, 0xd7, 0x59, 0xb0, 0xd7, 0x09,
	0x17, 0xf4, 0x52, 0x7b, 0x89, 0x0f, 0xc5, 0xcb, 0x16, 0x34, 0x22, 0x26, 0x99, 0xf7, 0x87, 0x01,
	0xed, 0x6a, 0xed, 0xe4, 0x1b, 0xb0, 0x35, 0xa9, 0xd8, 0x81, 0x6b, 0xec, 0xd7, 0x0f, 0xec, 0xee,
	0xf6, 0x66, 0x95, 0xca, 0x89, 0x02, 0x5f, 0x7c, 0x0a, 0xf2, 0x1c, 0xda, 0x09, 0x2b, 0xde, 0xf2,
	0x85, 0x67, 0x0d, 0x3d, 0xdd, 0x15, 0xcf, 0x01, 0x02, 0x04, 0xfa, 0xda, 0x1a, 0x8d, 0xce, 0x9e,
	0x0b, 0xdb, 0x37, 0xf7, 0xee, 0x25, 0x60, 0x2d, 0xf3, 0x91, 0x2e, 0x58, 0x2a, 0xf8, 0x58, 0x5e,
	0xe5, 0x7a, 0x0c, 0x77, 0xba, 0x1f, 0x6d, 0x8c, 0x21, 0xbc, 0xca, 0x39, 0x35, 0xa3, 0xf2, 0x8b,
	0x3c, 0x81, 0x26, 0x56, 0x59, 0x12, 0xbe, 0xda, 0xca, 0x2b, 0xce, 0x23, 0x0c, 0x4f, 0x35, 0xc8,
	0xfb, 0xc7, 0x00, 0x40, 0x85, 0xe2, 0xad, 0x20, 0xdf, 0x02, 0x4c, 0x7e, 0x61, 0xe9, 0x54, 0x73,
	0x6c, 0x7c, 0x90, 0x63, 0xab, 0x44, 0x87, 0x82, 0x3c, 0x85, 0x86, 0x5a, 0xed, 0x32, 0xed, 0xce,
	0x86, 0x53, 0xaf, 0x5c, 0x6b, 0x8a, 0x30, 0xf2, 0x0c, 0x9a, 0x42, 0x32, 0xc9, 0x71, 0x90, 0x77,
	0xba, 0xde, 0x26, 0xe3, 0x58, 0x11, 0xe6, 0x2a, 0x46, 0x0a, 0x49, 0xb5, 0x83, 0xf7, 0x03, 0xc0,
	0xb5, 0x92, 0xd8, 0xb0, 0x35, 0x4f, 0xdf, 0xa6, 0xd9, 0x65, 0xea, 0xdc, 0x52, 0xc2, 0xab, 0x33,
	0xfa, 0xb3, 0x4f, 0x7b, 0x8e, 0x41, 0xda, 0x60, 0xbe, 0xf4, 0x8f, 0x7e, 0x44, 0xa9, 0x46, 0x2c,
	0x68, 0x0e, 0xfd, 0xd7, 0xa3, 0xc0, 0xa9, 0x7b, 0x7f, 0xd6, 0xc1, 0x5a, 0x12, 0x41, 0x76, 0xc0,
	0xd4, 0x0b, 0x10, 0x47, 0xe5, 0x2b, 0xd9, 0x42, 0xb9, 0x1f, 0x29, 0x93, 0xc8, 0xb3, 0x02, 0x4d,
	0xaa, 0xad, 0x26, 0xdd, 0x42, 0xb9, 0x1f, 0x91, 0x5d, 0x30, 0x27, 0x4c, 0xf2, 0x69, 0x56, 0x5c,
	0x61, 0x07, 0x16, 0x5d, 0xca, 0x64, 0x1b, 0x5a, 0x33, 0xce, 0xa6, 0x73, 0xee, 0x36, 0xd0, 0x52,
	0x4a, 0xe4, 0x3b, 0x68, 0xa9, 0x0e, 0xe6, 0xc2, 0x6d, 0xde, 0xd0, 0xf3, 0xb2, 0x22, 0xdd, 0xfd,
	0x08, 0x91, 0xb4, 0xf4, 0x20, 0x5f, 0x81, 0x29, 0x24, 0x2b, 0xa4, 0x1a, 0x4b, 0xeb, 0x83, 0x63,
	0xd9, 0x42, 0x6c, 0x28, 0x88, 0x07, 0xed, 0x9c, 0x15, 0x32, 0x9e, 0xc4, 0x39, 0x4b, 0xa5, 0x70,
	0xb7, 0xf6, 0xeb, 0x07, 0x16, 0x5d, 0xd1, 0x91, 0xa7, 0xd0, 0x54, 0x13, 0x29, 0x5c, 0x13, 0xe3,
	0x3e, 0x7c, 0xc7, 0x24, 0xa8, 0x46, 0x91, 0x67, 0x00, 0xaa, 0xa6, 0x58, 0xc8, 0x78, 0x22, 0x5c,
	0x0b, 0x7d, 0xdc, 0x4d, 0x9f, 0x21, 0xde, 0x28, 0x5a, 0xc1, 0x7a, 0x5d, 0xb0, 0x2b, 0xad, 0xad,
	0x4e, 0xae, 0x0d, 0xe6, 0x90, 0x06, 0x03, 0x3f, 0x3c, 0x3a, 0x76, 0x0c, 0x62, 0x42, 0xe3, 0xa4,
	0xff, 0x53, 0xe0, 0xd4, 0xbc, 0xdf, 0x0d, 0xb0, 0x15, 0x33, 0x67, 0x73, 0x39, 0xc9, 0x12, 0x4e,
	0x1e, 0x03, 0x64, 0xfa, 0xf3, 0x7a, 0x5e, 0x56, 0xa9, 0xe9, 0x47, 0xe4, 0x53, 0x68, 0x2f, 0xcc,
	0xf8, 0x66, 0x1a, 0x38, 0x35, 0xbb, 0xd4, 0xe1, 0xfb, 0x78, 0x00, 0xcd, 0x0b, 0x36, 0x9b, 0xeb,
	0x45, 0xb5, 0xa8, 0x16, 0xc8, 0x23, 0xb0, 0xc4, 0x5c, 0xe4, 0x3c, 0x8d, 0x78, 0x84, 0x03, 0x35,
	0xe9, 0xb5, 0xc2, 0xbb, 0x04, 0xbb, 0xf2, 0x94, 0xdf, 0xb7, 0x32, 0x03, 0xb8, 0xaf, 0x4d, 0xfa,
	0xb5, 0x8b, 0x95, 0xe3, 0xf0, 0x78, 0x85, 0xa6, 0x51, 0x9c, 0x4e, 0x67, 0x7c, 0xb0, 0xbc, 0x0a,
	0xf4, 0x1e, 0x7a, 0x56, 0x12, 0x09, 0xef, 0x12, 0x9c, 0x75, 0xd8, 0xff, 0x3a, 0x0a, 0x87, 0xd0,
	0xd2, 0x05, 0x95, 0xcf, 0xf3, 0xe1, 0xc6, 0xea, 0xe9, 0x04, 0xb4, 0x84, 0x79, 0x7f, 0x1b, 0x00,
	0xd7, 0x6a, 0xf2, 0x31, 0x58, 0xe5, 0xb1, 0x5b, 0xb6, 0x6c, 0x6a, 0x45, 0x3f, 0x52, 0xbf, 0x9a,
	0xd2, 0x88, 0x25, 0xe9, 0x97, 0x02, 0x5a, 0x85, 0xd9, 0x7d, 0xb8, 0x5d, 0x02, 0x72, 0x56, 0xb0,
	0x44, 0x1d, 0x6f, 0x45, 0xc7, 0xa3, 0x77, 0x14, 0x31, 0x54, 0x20, 0x5a, 0x5e, 0x57, 0x14, 0x04,
	0xf9, 0x12, 0xcc, 0x72, 0x88, 0xc2, 0x6d, 0xdc, 0x70, 0x69, 0x2b, 0x3b, 0x42, 0x97, 0x48, 0xef,
	0x2f, 0x03, 0xee, 0xae, 0xc5, 0x25, 0x2f, 0xa0, 0x51, 0x61, 0xee, 0xc9, 0xfb, 0x6a, 0xe8, 0x54,
	0xbe, 0x91, 0x50, 0xf4, 0xbc, 0x79, 0x83, 0xbc, 0x01, 0xdc, 0x5d, 0x83, 0xaf, 0x6e, 0x38, 0x40,
	0x6b, 0x18, 0xd0, 0xfe, 0x99, 0x3a, 0x4d, 0x16, 0x34, 0xc3, 0xb3, 0xd0, 0x3f, 0x71, 0x6a, 0x6a,
	0xf1, 0x8f, 0xfd, 0xd3, 0x5e, 0xff, 0xc8, 0x1f, 0x3a, 0x75, 0xb5, 0xf8, 0x61, 0xe0, 0x0f, 0x9c,
	0xc6, 0xe7, 0xcf, 0xc1, 0x5c, 0xcc, 0x71, 0x23, 0x4e, 0xff, 0x74, 0x14, 0xd0, 0xd0, 0x31, 0xd4,
	0x77, 0x2f, 0x38, 0x09, 0xc2, 0xc0, 0xa9, 0xa9, 0xef, 0xd7, 0xc3, 0x9e, 0x1f, 0x06, 0x4e, 0xfd,
	0x4d, 0x0b, 0xdb, 0xf9, 0xe2, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x30, 0x85, 0x9a, 0xab,
	0x08, 0x00, 0x00,
}
