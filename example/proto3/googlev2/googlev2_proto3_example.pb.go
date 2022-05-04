// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: googlev2_proto3_example.proto

package googlev2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_EVENT_TYPE_UNDEFINED EventType = 0
	EventType_EVENT_TYPE_ONE       EventType = 1
	EventType_EVENT_TYPE_TWO       EventType = 2
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNDEFINED",
		1: "EVENT_TYPE_ONE",
		2: "EVENT_TYPE_TWO",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNDEFINED": 0,
		"EVENT_TYPE_ONE":       1,
		"EVENT_TYPE_TWO":       2,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_googlev2_proto3_example_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_googlev2_proto3_example_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_googlev2_proto3_example_proto_rawDescGZIP(), []int{0}
}

type TestEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Info      string         `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	IsAwesome bool           `protobuf:"varint,3,opt,name=isAwesome,proto3" json:"isAwesome,omitempty"`
	Labels    []string       `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Embedded  *EmbeddedEvent `protobuf:"bytes,5,opt,name=embedded,proto3" json:"embedded,omitempty"`
	// Types that are assignable to Path:
	//	*TestEvent_Jedi
	//	*TestEvent_Sith
	//	*TestEvent_Other
	Path    isTestEvent_Path       `protobuf_oneof:"path"`
	Nested  *TestEvent_NestedMsg   `protobuf:"bytes,9,opt,name=nested,proto3" json:"nested,omitempty"`
	Ts      *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=ts,proto3" json:"ts,omitempty"`
	NullVal structpb.NullValue     `protobuf:"varint,11,opt,name=nullVal,proto3,enum=google.protobuf.NullValue" json:"nullVal,omitempty"`
}

func (x *TestEvent) Reset() {
	*x = TestEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_googlev2_proto3_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestEvent) ProtoMessage() {}

func (x *TestEvent) ProtoReflect() protoreflect.Message {
	mi := &file_googlev2_proto3_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestEvent.ProtoReflect.Descriptor instead.
func (*TestEvent) Descriptor() ([]byte, []int) {
	return file_googlev2_proto3_example_proto_rawDescGZIP(), []int{0}
}

func (x *TestEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestEvent) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *TestEvent) GetIsAwesome() bool {
	if x != nil {
		return x.IsAwesome
	}
	return false
}

func (x *TestEvent) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TestEvent) GetEmbedded() *EmbeddedEvent {
	if x != nil {
		return x.Embedded
	}
	return nil
}

func (m *TestEvent) GetPath() isTestEvent_Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (x *TestEvent) GetJedi() bool {
	if x, ok := x.GetPath().(*TestEvent_Jedi); ok {
		return x.Jedi
	}
	return false
}

func (x *TestEvent) GetSith() bool {
	if x, ok := x.GetPath().(*TestEvent_Sith); ok {
		return x.Sith
	}
	return false
}

func (x *TestEvent) GetOther() string {
	if x, ok := x.GetPath().(*TestEvent_Other); ok {
		return x.Other
	}
	return ""
}

func (x *TestEvent) GetNested() *TestEvent_NestedMsg {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *TestEvent) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *TestEvent) GetNullVal() structpb.NullValue {
	if x != nil {
		return x.NullVal
	}
	return structpb.NullValue(0)
}

type isTestEvent_Path interface {
	isTestEvent_Path()
}

type TestEvent_Jedi struct {
	Jedi bool `protobuf:"varint,6,opt,name=jedi,proto3,oneof"`
}

type TestEvent_Sith struct {
	Sith bool `protobuf:"varint,7,opt,name=sith,proto3,oneof"`
}

type TestEvent_Other struct {
	Other string `protobuf:"bytes,8,opt,name=other,proto3,oneof"`
}

func (*TestEvent_Jedi) isTestEvent_Path() {}

func (*TestEvent_Sith) isTestEvent_Path() {}

func (*TestEvent_Other) isTestEvent_Path() {}

type EmbeddedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Stuff           string   `protobuf:"bytes,2,opt,name=stuff,proto3" json:"stuff,omitempty"`
	FavoriteNumbers []int32  `protobuf:"varint,3,rep,packed,name=favoriteNumbers,proto3" json:"favoriteNumbers,omitempty"`
	RandomThings    [][]byte `protobuf:"bytes,4,rep,name=randomThings,proto3" json:"randomThings,omitempty"`
}

func (x *EmbeddedEvent) Reset() {
	*x = EmbeddedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_googlev2_proto3_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbeddedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbeddedEvent) ProtoMessage() {}

func (x *EmbeddedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_googlev2_proto3_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbeddedEvent.ProtoReflect.Descriptor instead.
func (*EmbeddedEvent) Descriptor() ([]byte, []int) {
	return file_googlev2_proto3_example_proto_rawDescGZIP(), []int{1}
}

func (x *EmbeddedEvent) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *EmbeddedEvent) GetStuff() string {
	if x != nil {
		return x.Stuff
	}
	return ""
}

func (x *EmbeddedEvent) GetFavoriteNumbers() []int32 {
	if x != nil {
		return x.FavoriteNumbers
	}
	return nil
}

func (x *EmbeddedEvent) GetRandomThings() [][]byte {
	if x != nil {
		return x.RandomThings
	}
	return nil
}

type AllTheThings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int32          `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TheString    string         `protobuf:"bytes,2,opt,name=theString,proto3" json:"theString,omitempty"`
	TheBool      bool           `protobuf:"varint,3,opt,name=theBool,proto3" json:"theBool,omitempty"`
	TheInt32     int32          `protobuf:"varint,4,opt,name=theInt32,proto3" json:"theInt32,omitempty"`
	TheInt64     int64          `protobuf:"varint,5,opt,name=theInt64,proto3" json:"theInt64,omitempty"`
	TheUInt32    uint32         `protobuf:"varint,6,opt,name=theUInt32,proto3" json:"theUInt32,omitempty"`
	TheUInt64    uint64         `protobuf:"varint,7,opt,name=theUInt64,proto3" json:"theUInt64,omitempty"`
	TheSInt32    int32          `protobuf:"zigzag32,8,opt,name=theSInt32,proto3" json:"theSInt32,omitempty"`
	TheSInt64    int64          `protobuf:"zigzag64,9,opt,name=theSInt64,proto3" json:"theSInt64,omitempty"`
	TheFixed32   uint32         `protobuf:"fixed32,10,opt,name=theFixed32,proto3" json:"theFixed32,omitempty"`
	TheFixed64   uint64         `protobuf:"fixed64,11,opt,name=theFixed64,proto3" json:"theFixed64,omitempty"`
	TheSFixed32  int32          `protobuf:"fixed32,12,opt,name=theSFixed32,proto3" json:"theSFixed32,omitempty"`
	TheSFixed64  int64          `protobuf:"fixed64,13,opt,name=theSFixed64,proto3" json:"theSFixed64,omitempty"`
	TheFloat     float32        `protobuf:"fixed32,14,opt,name=theFloat,proto3" json:"theFloat,omitempty"`
	TheDouble    float64        `protobuf:"fixed64,15,opt,name=theDouble,proto3" json:"theDouble,omitempty"`
	TheEventType EventType      `protobuf:"varint,16,opt,name=theEventType,proto3,enum=crowdstrike.csproto.example.proto3.googlev2.EventType" json:"theEventType,omitempty"`
	TheBytes     []byte         `protobuf:"bytes,17,opt,name=theBytes,proto3" json:"theBytes,omitempty"`
	TheMessage   *EmbeddedEvent `protobuf:"bytes,18,opt,name=theMessage,proto3" json:"theMessage,omitempty"`
}

func (x *AllTheThings) Reset() {
	*x = AllTheThings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_googlev2_proto3_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllTheThings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllTheThings) ProtoMessage() {}

func (x *AllTheThings) ProtoReflect() protoreflect.Message {
	mi := &file_googlev2_proto3_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllTheThings.ProtoReflect.Descriptor instead.
func (*AllTheThings) Descriptor() ([]byte, []int) {
	return file_googlev2_proto3_example_proto_rawDescGZIP(), []int{2}
}

func (x *AllTheThings) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AllTheThings) GetTheString() string {
	if x != nil {
		return x.TheString
	}
	return ""
}

func (x *AllTheThings) GetTheBool() bool {
	if x != nil {
		return x.TheBool
	}
	return false
}

func (x *AllTheThings) GetTheInt32() int32 {
	if x != nil {
		return x.TheInt32
	}
	return 0
}

func (x *AllTheThings) GetTheInt64() int64 {
	if x != nil {
		return x.TheInt64
	}
	return 0
}

func (x *AllTheThings) GetTheUInt32() uint32 {
	if x != nil {
		return x.TheUInt32
	}
	return 0
}

func (x *AllTheThings) GetTheUInt64() uint64 {
	if x != nil {
		return x.TheUInt64
	}
	return 0
}

func (x *AllTheThings) GetTheSInt32() int32 {
	if x != nil {
		return x.TheSInt32
	}
	return 0
}

func (x *AllTheThings) GetTheSInt64() int64 {
	if x != nil {
		return x.TheSInt64
	}
	return 0
}

func (x *AllTheThings) GetTheFixed32() uint32 {
	if x != nil {
		return x.TheFixed32
	}
	return 0
}

func (x *AllTheThings) GetTheFixed64() uint64 {
	if x != nil {
		return x.TheFixed64
	}
	return 0
}

func (x *AllTheThings) GetTheSFixed32() int32 {
	if x != nil {
		return x.TheSFixed32
	}
	return 0
}

func (x *AllTheThings) GetTheSFixed64() int64 {
	if x != nil {
		return x.TheSFixed64
	}
	return 0
}

func (x *AllTheThings) GetTheFloat() float32 {
	if x != nil {
		return x.TheFloat
	}
	return 0
}

func (x *AllTheThings) GetTheDouble() float64 {
	if x != nil {
		return x.TheDouble
	}
	return 0
}

func (x *AllTheThings) GetTheEventType() EventType {
	if x != nil {
		return x.TheEventType
	}
	return EventType_EVENT_TYPE_UNDEFINED
}

func (x *AllTheThings) GetTheBytes() []byte {
	if x != nil {
		return x.TheBytes
	}
	return nil
}

func (x *AllTheThings) GetTheMessage() *EmbeddedEvent {
	if x != nil {
		return x.TheMessage
	}
	return nil
}

type RepeatAllTheThings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TheStrings    []string         `protobuf:"bytes,2,rep,name=theStrings,proto3" json:"theStrings,omitempty"`
	TheBools      []bool           `protobuf:"varint,3,rep,packed,name=theBools,proto3" json:"theBools,omitempty"`
	TheInt32S     []int32          `protobuf:"varint,4,rep,packed,name=theInt32s,proto3" json:"theInt32s,omitempty"`
	TheInt64S     []int64          `protobuf:"varint,5,rep,packed,name=theInt64s,proto3" json:"theInt64s,omitempty"`
	TheUInt32S    []uint32         `protobuf:"varint,6,rep,packed,name=theUInt32s,proto3" json:"theUInt32s,omitempty"`
	TheUInt64S    []uint64         `protobuf:"varint,7,rep,packed,name=theUInt64s,proto3" json:"theUInt64s,omitempty"`
	TheSInt32S    []int32          `protobuf:"zigzag32,8,rep,packed,name=theSInt32s,proto3" json:"theSInt32s,omitempty"`
	TheSInt64S    []int64          `protobuf:"zigzag64,9,rep,packed,name=theSInt64s,proto3" json:"theSInt64s,omitempty"`
	TheFixed32S   []uint32         `protobuf:"fixed32,10,rep,packed,name=theFixed32s,proto3" json:"theFixed32s,omitempty"`
	TheFixed64S   []uint64         `protobuf:"fixed64,11,rep,packed,name=theFixed64s,proto3" json:"theFixed64s,omitempty"`
	TheSFixed32S  []int32          `protobuf:"fixed32,12,rep,packed,name=theSFixed32s,proto3" json:"theSFixed32s,omitempty"`
	TheSFixed64S  []int64          `protobuf:"fixed64,13,rep,packed,name=theSFixed64s,proto3" json:"theSFixed64s,omitempty"`
	TheFloats     []float32        `protobuf:"fixed32,14,rep,packed,name=theFloats,proto3" json:"theFloats,omitempty"`
	TheDoubles    []float64        `protobuf:"fixed64,15,rep,packed,name=theDoubles,proto3" json:"theDoubles,omitempty"`
	TheEventTypes []EventType      `protobuf:"varint,16,rep,packed,name=theEventTypes,proto3,enum=crowdstrike.csproto.example.proto3.googlev2.EventType" json:"theEventTypes,omitempty"`
	TheBytes      [][]byte         `protobuf:"bytes,17,rep,name=theBytes,proto3" json:"theBytes,omitempty"`
	TheMessages   []*EmbeddedEvent `protobuf:"bytes,18,rep,name=theMessages,proto3" json:"theMessages,omitempty"`
}

func (x *RepeatAllTheThings) Reset() {
	*x = RepeatAllTheThings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_googlev2_proto3_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatAllTheThings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatAllTheThings) ProtoMessage() {}

func (x *RepeatAllTheThings) ProtoReflect() protoreflect.Message {
	mi := &file_googlev2_proto3_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatAllTheThings.ProtoReflect.Descriptor instead.
func (*RepeatAllTheThings) Descriptor() ([]byte, []int) {
	return file_googlev2_proto3_example_proto_rawDescGZIP(), []int{3}
}

func (x *RepeatAllTheThings) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RepeatAllTheThings) GetTheStrings() []string {
	if x != nil {
		return x.TheStrings
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheBools() []bool {
	if x != nil {
		return x.TheBools
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheInt32S() []int32 {
	if x != nil {
		return x.TheInt32S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheInt64S() []int64 {
	if x != nil {
		return x.TheInt64S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheUInt32S() []uint32 {
	if x != nil {
		return x.TheUInt32S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheUInt64S() []uint64 {
	if x != nil {
		return x.TheUInt64S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheSInt32S() []int32 {
	if x != nil {
		return x.TheSInt32S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheSInt64S() []int64 {
	if x != nil {
		return x.TheSInt64S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheFixed32S() []uint32 {
	if x != nil {
		return x.TheFixed32S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheFixed64S() []uint64 {
	if x != nil {
		return x.TheFixed64S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheSFixed32S() []int32 {
	if x != nil {
		return x.TheSFixed32S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheSFixed64S() []int64 {
	if x != nil {
		return x.TheSFixed64S
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheFloats() []float32 {
	if x != nil {
		return x.TheFloats
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheDoubles() []float64 {
	if x != nil {
		return x.TheDoubles
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheEventTypes() []EventType {
	if x != nil {
		return x.TheEventTypes
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheBytes() [][]byte {
	if x != nil {
		return x.TheBytes
	}
	return nil
}

func (x *RepeatAllTheThings) GetTheMessages() []*EmbeddedEvent {
	if x != nil {
		return x.TheMessages
	}
	return nil
}

type TestEvent_NestedMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details string `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *TestEvent_NestedMsg) Reset() {
	*x = TestEvent_NestedMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_googlev2_proto3_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestEvent_NestedMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestEvent_NestedMsg) ProtoMessage() {}

func (x *TestEvent_NestedMsg) ProtoReflect() protoreflect.Message {
	mi := &file_googlev2_proto3_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestEvent_NestedMsg.ProtoReflect.Descriptor instead.
func (*TestEvent_NestedMsg) Descriptor() ([]byte, []int) {
	return file_googlev2_proto3_example_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TestEvent_NestedMsg) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

var File_googlev2_proto3_example_proto protoreflect.FileDescriptor

var file_googlev2_proto3_example_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x32, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2b, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x2e, 0x63, 0x73, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x03, 0x0a, 0x09,
	0x54, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x41, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x56, 0x0a, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x72, 0x6f, 0x77,
	0x64, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x2e, 0x63, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x32, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x04, 0x6a, 0x65, 0x64, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x04, 0x6a, 0x65, 0x64, 0x69, 0x12, 0x14, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x73, 0x69, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x73, 0x74, 0x72, 0x69, 0x6b,
	0x65, 0x2e, 0x63, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x76,
	0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x6e, 0x75, 0x6c,
	0x6c, 0x56, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x75, 0x6c,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x1a,
	0x25, 0x0a, 0x09, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x83,
	0x01, 0x0a, 0x0d, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x75, 0x66, 0x66, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0x98, 0x05, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x54, 0x68, 0x65, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x68, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x74, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x65,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x68, 0x65,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x68, 0x65, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x68, 0x65, 0x55, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x53, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x74, 0x68, 0x65, 0x53, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x53, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x09, 0x74, 0x68, 0x65, 0x53, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x07, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x06, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x68, 0x65, 0x53, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0f, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x53, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x68, 0x65, 0x53, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x10, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x53, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36,
	0x34, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x74, 0x68, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x68, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x74, 0x68, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x5a, 0x0a, 0x0c, 0x74,
	0x68, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x36, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x2e,
	0x63, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x32, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x74, 0x68, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x68, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x73,
	0x74, 0x72, 0x69, 0x6b, 0x65, 0x2e, 0x63, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x76, 0x32, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xbe, 0x05, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x68, 0x65,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x65, 0x42, 0x6f, 0x6f,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x08, 0x52, 0x08, 0x74, 0x68, 0x65, 0x42, 0x6f, 0x6f,
	0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x74, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x09, 0x74, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x68, 0x65, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x68, 0x65, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x68, 0x65, 0x53, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x11, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x53, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x68, 0x65, 0x53, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x12, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x53, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x07, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x06, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36,
	0x34, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x65, 0x53, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0f, 0x52, 0x0c, 0x74, 0x68, 0x65, 0x53, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x65, 0x53, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x10, 0x52, 0x0c, 0x74, 0x68,
	0x65, 0x53, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68,
	0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x02, 0x52, 0x09, 0x74,
	0x68, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x65, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x68,
	0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x0d, 0x74, 0x68, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x36, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x73, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x2e, 0x63, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x32, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x74, 0x68, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x68, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x5c, 0x0a, 0x0b, 0x74, 0x68, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x73,
	0x74, 0x72, 0x69, 0x6b, 0x65, 0x2e, 0x63, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x76, 0x32, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2a, 0x4d, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72,
	0x6f, 0x77, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6b, 0x65, 0x2f, 0x63, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_googlev2_proto3_example_proto_rawDescOnce sync.Once
	file_googlev2_proto3_example_proto_rawDescData = file_googlev2_proto3_example_proto_rawDesc
)

func file_googlev2_proto3_example_proto_rawDescGZIP() []byte {
	file_googlev2_proto3_example_proto_rawDescOnce.Do(func() {
		file_googlev2_proto3_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_googlev2_proto3_example_proto_rawDescData)
	})
	return file_googlev2_proto3_example_proto_rawDescData
}

var file_googlev2_proto3_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_googlev2_proto3_example_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_googlev2_proto3_example_proto_goTypes = []interface{}{
	(EventType)(0),                // 0: crowdstrike.csproto.example.proto3.googlev2.EventType
	(*TestEvent)(nil),             // 1: crowdstrike.csproto.example.proto3.googlev2.TestEvent
	(*EmbeddedEvent)(nil),         // 2: crowdstrike.csproto.example.proto3.googlev2.EmbeddedEvent
	(*AllTheThings)(nil),          // 3: crowdstrike.csproto.example.proto3.googlev2.AllTheThings
	(*RepeatAllTheThings)(nil),    // 4: crowdstrike.csproto.example.proto3.googlev2.RepeatAllTheThings
	(*TestEvent_NestedMsg)(nil),   // 5: crowdstrike.csproto.example.proto3.googlev2.TestEvent.NestedMsg
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(structpb.NullValue)(0),       // 7: google.protobuf.NullValue
}
var file_googlev2_proto3_example_proto_depIdxs = []int32{
	2, // 0: crowdstrike.csproto.example.proto3.googlev2.TestEvent.embedded:type_name -> crowdstrike.csproto.example.proto3.googlev2.EmbeddedEvent
	5, // 1: crowdstrike.csproto.example.proto3.googlev2.TestEvent.nested:type_name -> crowdstrike.csproto.example.proto3.googlev2.TestEvent.NestedMsg
	6, // 2: crowdstrike.csproto.example.proto3.googlev2.TestEvent.ts:type_name -> google.protobuf.Timestamp
	7, // 3: crowdstrike.csproto.example.proto3.googlev2.TestEvent.nullVal:type_name -> google.protobuf.NullValue
	0, // 4: crowdstrike.csproto.example.proto3.googlev2.AllTheThings.theEventType:type_name -> crowdstrike.csproto.example.proto3.googlev2.EventType
	2, // 5: crowdstrike.csproto.example.proto3.googlev2.AllTheThings.theMessage:type_name -> crowdstrike.csproto.example.proto3.googlev2.EmbeddedEvent
	0, // 6: crowdstrike.csproto.example.proto3.googlev2.RepeatAllTheThings.theEventTypes:type_name -> crowdstrike.csproto.example.proto3.googlev2.EventType
	2, // 7: crowdstrike.csproto.example.proto3.googlev2.RepeatAllTheThings.theMessages:type_name -> crowdstrike.csproto.example.proto3.googlev2.EmbeddedEvent
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_googlev2_proto3_example_proto_init() }
func file_googlev2_proto3_example_proto_init() {
	if File_googlev2_proto3_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_googlev2_proto3_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_googlev2_proto3_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbeddedEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_googlev2_proto3_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllTheThings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_googlev2_proto3_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatAllTheThings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_googlev2_proto3_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestEvent_NestedMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_googlev2_proto3_example_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TestEvent_Jedi)(nil),
		(*TestEvent_Sith)(nil),
		(*TestEvent_Other)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_googlev2_proto3_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_googlev2_proto3_example_proto_goTypes,
		DependencyIndexes: file_googlev2_proto3_example_proto_depIdxs,
		EnumInfos:         file_googlev2_proto3_example_proto_enumTypes,
		MessageInfos:      file_googlev2_proto3_example_proto_msgTypes,
	}.Build()
	File_googlev2_proto3_example_proto = out.File
	file_googlev2_proto3_example_proto_rawDesc = nil
	file_googlev2_proto3_example_proto_goTypes = nil
	file_googlev2_proto3_example_proto_depIdxs = nil
}
