// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xfeed_proto/event_points.proto

package xfeed_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PointsGroup_PointType int32

const (
	PointsGroup_unknown      PointsGroup_PointType = 0
	PointsGroup_SCORE        PointsGroup_PointType = 1
	PointsGroup_RED_CARDS    PointsGroup_PointType = 2
	PointsGroup_YELLOW_CARDS PointsGroup_PointType = 3
	PointsGroup_PENALTIES    PointsGroup_PointType = 4
	PointsGroup_CORNERS      PointsGroup_PointType = 5
)

var PointsGroup_PointType_name = map[int32]string{
	0: "unknown",
	1: "SCORE",
	2: "RED_CARDS",
	3: "YELLOW_CARDS",
	4: "PENALTIES",
	5: "CORNERS",
}
var PointsGroup_PointType_value = map[string]int32{
	"unknown":      0,
	"SCORE":        1,
	"RED_CARDS":    2,
	"YELLOW_CARDS": 3,
	"PENALTIES":    4,
	"CORNERS":      5,
}

func (x PointsGroup_PointType) String() string {
	return proto.EnumName(PointsGroup_PointType_name, int32(x))
}
func (PointsGroup_PointType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_event_points_09dcb4ce6c1c64ff, []int{3, 0}
}

type GroupParams struct {
	Period               int32    `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupParams) Reset()         { *m = GroupParams{} }
func (m *GroupParams) String() string { return proto.CompactTextString(m) }
func (*GroupParams) ProtoMessage()    {}
func (*GroupParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_points_09dcb4ce6c1c64ff, []int{0}
}
func (m *GroupParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupParams.Unmarshal(m, b)
}
func (m *GroupParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupParams.Marshal(b, m, deterministic)
}
func (dst *GroupParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupParams.Merge(dst, src)
}
func (m *GroupParams) XXX_Size() int {
	return xxx_messageInfo_GroupParams.Size(m)
}
func (m *GroupParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupParams.DiscardUnknown(m)
}

var xxx_messageInfo_GroupParams proto.InternalMessageInfo

func (m *GroupParams) GetPeriod() int32 {
	if m != nil {
		return m.Period
	}
	return 0
}

type StateParams struct {
	Participant          int32    `protobuf:"varint,1,opt,name=participant,proto3" json:"participant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateParams) Reset()         { *m = StateParams{} }
func (m *StateParams) String() string { return proto.CompactTextString(m) }
func (*StateParams) ProtoMessage()    {}
func (*StateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_points_09dcb4ce6c1c64ff, []int{1}
}
func (m *StateParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateParams.Unmarshal(m, b)
}
func (m *StateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateParams.Marshal(b, m, deterministic)
}
func (dst *StateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateParams.Merge(dst, src)
}
func (m *StateParams) XXX_Size() int {
	return xxx_messageInfo_StateParams.Size(m)
}
func (m *StateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_StateParams.DiscardUnknown(m)
}

var xxx_messageInfo_StateParams proto.InternalMessageInfo

func (m *StateParams) GetParticipant() int32 {
	if m != nil {
		return m.Participant
	}
	return 0
}

type State struct {
	StateParams          *StateParams `protobuf:"bytes,1,opt,name=state_params,json=stateParams,proto3" json:"state_params,omitempty"`
	Value                int32        `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_points_09dcb4ce6c1c64ff, []int{2}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetStateParams() *StateParams {
	if m != nil {
		return m.StateParams
	}
	return nil
}

func (m *State) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type PointsGroup struct {
	PointType            PointsGroup_PointType `protobuf:"varint,1,opt,name=point_type,json=pointType,proto3,enum=xfeed_proto.PointsGroup_PointType" json:"point_type,omitempty"`
	GroupParams          *GroupParams          `protobuf:"bytes,2,opt,name=group_params,json=groupParams,proto3" json:"group_params,omitempty"`
	State                []*State              `protobuf:"bytes,3,rep,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PointsGroup) Reset()         { *m = PointsGroup{} }
func (m *PointsGroup) String() string { return proto.CompactTextString(m) }
func (*PointsGroup) ProtoMessage()    {}
func (*PointsGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_points_09dcb4ce6c1c64ff, []int{3}
}
func (m *PointsGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointsGroup.Unmarshal(m, b)
}
func (m *PointsGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointsGroup.Marshal(b, m, deterministic)
}
func (dst *PointsGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointsGroup.Merge(dst, src)
}
func (m *PointsGroup) XXX_Size() int {
	return xxx_messageInfo_PointsGroup.Size(m)
}
func (m *PointsGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PointsGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PointsGroup proto.InternalMessageInfo

func (m *PointsGroup) GetPointType() PointsGroup_PointType {
	if m != nil {
		return m.PointType
	}
	return PointsGroup_unknown
}

func (m *PointsGroup) GetGroupParams() *GroupParams {
	if m != nil {
		return m.GroupParams
	}
	return nil
}

func (m *PointsGroup) GetState() []*State {
	if m != nil {
		return m.State
	}
	return nil
}

type EventPoints struct {
	PointGroups          []*PointsGroup `protobuf:"bytes,1,rep,name=point_groups,json=pointGroups,proto3" json:"point_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EventPoints) Reset()         { *m = EventPoints{} }
func (m *EventPoints) String() string { return proto.CompactTextString(m) }
func (*EventPoints) ProtoMessage()    {}
func (*EventPoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_points_09dcb4ce6c1c64ff, []int{4}
}
func (m *EventPoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventPoints.Unmarshal(m, b)
}
func (m *EventPoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventPoints.Marshal(b, m, deterministic)
}
func (dst *EventPoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPoints.Merge(dst, src)
}
func (m *EventPoints) XXX_Size() int {
	return xxx_messageInfo_EventPoints.Size(m)
}
func (m *EventPoints) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPoints.DiscardUnknown(m)
}

var xxx_messageInfo_EventPoints proto.InternalMessageInfo

func (m *EventPoints) GetPointGroups() []*PointsGroup {
	if m != nil {
		return m.PointGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupParams)(nil), "xfeed_proto.GroupParams")
	proto.RegisterType((*StateParams)(nil), "xfeed_proto.StateParams")
	proto.RegisterType((*State)(nil), "xfeed_proto.State")
	proto.RegisterType((*PointsGroup)(nil), "xfeed_proto.PointsGroup")
	proto.RegisterType((*EventPoints)(nil), "xfeed_proto.EventPoints")
	proto.RegisterEnum("xfeed_proto.PointsGroup_PointType", PointsGroup_PointType_name, PointsGroup_PointType_value)
}

func init() {
	proto.RegisterFile("xfeed_proto/event_points.proto", fileDescriptor_event_points_09dcb4ce6c1c64ff)
}

var fileDescriptor_event_points_09dcb4ce6c1c64ff = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x5f, 0x4b, 0xfb, 0x30,
	0x14, 0xfd, 0xad, 0xfd, 0x75, 0xd2, 0x9b, 0x29, 0x25, 0x88, 0xf4, 0x49, 0x46, 0x40, 0xd8, 0xd3,
	0x06, 0xf3, 0xd1, 0xa7, 0xb1, 0x05, 0x51, 0xc6, 0x36, 0xd2, 0x81, 0xe8, 0x4b, 0xad, 0x2e, 0x8e,
	0xa2, 0x36, 0xa1, 0xcd, 0xa6, 0xfb, 0x28, 0x7e, 0x5b, 0xc9, 0x6d, 0x27, 0x91, 0xf9, 0x76, 0x72,
	0xef, 0x49, 0xce, 0x9f, 0xc0, 0xf9, 0xe7, 0x8b, 0x94, 0xab, 0x54, 0x97, 0xca, 0xa8, 0x81, 0xdc,
	0xca, 0xc2, 0xa4, 0x5a, 0xe5, 0x85, 0xa9, 0xfa, 0x38, 0xa2, 0xc4, 0xd9, 0xb3, 0x0b, 0x20, 0xd7,
	0xa5, 0xda, 0xe8, 0x45, 0x56, 0x66, 0xef, 0x15, 0x3d, 0x83, 0xb6, 0x96, 0x65, 0xae, 0x56, 0x71,
	0xab, 0xdb, 0xea, 0x05, 0xa2, 0x39, 0xb1, 0x01, 0x90, 0xc4, 0x64, 0x46, 0x36, 0xb4, 0x2e, 0x10,
	0x9d, 0x95, 0x26, 0x7f, 0xce, 0x75, 0x56, 0x98, 0x86, 0xeb, 0x8e, 0xd8, 0x03, 0x04, 0x78, 0x81,
	0x5e, 0x41, 0xa7, 0xb2, 0x20, 0xd5, 0x78, 0x15, 0xb9, 0x64, 0x18, 0xf7, 0x1d, 0x13, 0x7d, 0xe7,
	0x69, 0x41, 0x2a, 0x47, 0xe7, 0x14, 0x82, 0x6d, 0xf6, 0xb6, 0x91, 0xb1, 0x87, 0x0a, 0xf5, 0x81,
	0x7d, 0x79, 0x40, 0x16, 0x98, 0x08, 0xad, 0xd3, 0x11, 0x00, 0x06, 0x4c, 0xcd, 0x4e, 0x4b, 0x14,
	0x38, 0x19, 0xb2, 0x5f, 0x02, 0x0e, 0xbb, 0xc6, 0xcb, 0x9d, 0x96, 0x22, 0xd4, 0x7b, 0x68, 0x5d,
	0xae, 0xed, 0x76, 0xef, 0xd2, 0xfb, 0xc3, 0xa5, 0xd3, 0x93, 0x20, 0x6b, 0xa7, 0xb4, 0x1e, 0x04,
	0x68, 0x3a, 0xf6, 0xbb, 0x7e, 0x8f, 0x0c, 0xe9, 0x61, 0x36, 0x51, 0x13, 0xd8, 0x23, 0x84, 0x3f,
	0xf2, 0x94, 0xc0, 0xd1, 0xa6, 0x78, 0x2d, 0xd4, 0x47, 0x11, 0xfd, 0xa3, 0x21, 0x04, 0xc9, 0x78,
	0x2e, 0x78, 0xd4, 0xa2, 0xc7, 0x10, 0x0a, 0x3e, 0x49, 0xc7, 0x23, 0x31, 0x49, 0x22, 0x8f, 0x46,
	0xd0, 0xb9, 0xe7, 0xd3, 0xe9, 0xfc, 0xae, 0x99, 0xf8, 0x96, 0xb0, 0xe0, 0xb3, 0xd1, 0x74, 0x79,
	0xc3, 0x93, 0xe8, 0xbf, 0x7d, 0x67, 0x3c, 0x17, 0x33, 0x2e, 0x92, 0x28, 0x60, 0xb7, 0x40, 0xb8,
	0xfd, 0xf2, 0x3a, 0xb1, 0xcd, 0x55, 0x57, 0x83, 0x7e, 0x6d, 0xfb, 0xfe, 0x41, 0x2e, 0xa7, 0x1c,
	0x41, 0x90, 0x8d, 0xb8, 0x7a, 0x6a, 0xe3, 0xfe, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x0b,
	0x38, 0x44, 0x51, 0x02, 0x00, 0x00,
}
