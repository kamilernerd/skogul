// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pfe_iflset_qstats.proto

package telemetry

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type JunosPfeIflset_234 struct {
	System               *JunosPfeIflset_234SystemType `protobuf:"bytes,151,opt,name=system" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *JunosPfeIflset_234) Reset()         { *m = JunosPfeIflset_234{} }
func (m *JunosPfeIflset_234) String() string { return proto.CompactTextString(m) }
func (*JunosPfeIflset_234) ProtoMessage()    {}
func (*JunosPfeIflset_234) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b66e3e3153fc1b, []int{0}
}
func (m *JunosPfeIflset_234) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeIflset_234.Unmarshal(m, b)
}
func (m *JunosPfeIflset_234) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeIflset_234.Marshal(b, m, deterministic)
}
func (m *JunosPfeIflset_234) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeIflset_234.Merge(m, src)
}
func (m *JunosPfeIflset_234) XXX_Size() int {
	return xxx_messageInfo_JunosPfeIflset_234.Size(m)
}
func (m *JunosPfeIflset_234) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeIflset_234.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeIflset_234 proto.InternalMessageInfo

func (m *JunosPfeIflset_234) GetSystem() *JunosPfeIflset_234SystemType {
	if m != nil {
		return m.System
	}
	return nil
}

type JunosPfeIflset_234SystemType struct {
	Linecard             *JunosPfeIflset_234SystemTypeLinecardType `protobuf:"bytes,151,opt,name=linecard" json:"linecard,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *JunosPfeIflset_234SystemType) Reset()         { *m = JunosPfeIflset_234SystemType{} }
func (m *JunosPfeIflset_234SystemType) String() string { return proto.CompactTextString(m) }
func (*JunosPfeIflset_234SystemType) ProtoMessage()    {}
func (*JunosPfeIflset_234SystemType) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b66e3e3153fc1b, []int{0, 0}
}
func (m *JunosPfeIflset_234SystemType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeIflset_234SystemType.Unmarshal(m, b)
}
func (m *JunosPfeIflset_234SystemType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeIflset_234SystemType.Marshal(b, m, deterministic)
}
func (m *JunosPfeIflset_234SystemType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeIflset_234SystemType.Merge(m, src)
}
func (m *JunosPfeIflset_234SystemType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeIflset_234SystemType.Size(m)
}
func (m *JunosPfeIflset_234SystemType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeIflset_234SystemType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeIflset_234SystemType proto.InternalMessageInfo

func (m *JunosPfeIflset_234SystemType) GetLinecard() *JunosPfeIflset_234SystemTypeLinecardType {
	if m != nil {
		return m.Linecard
	}
	return nil
}

type JunosPfeIflset_234SystemTypeLinecardType struct {
	Cos                  *JunosPfeIflset_234SystemTypeLinecardTypeCosType `protobuf:"bytes,151,opt,name=cos" json:"cos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *JunosPfeIflset_234SystemTypeLinecardType) Reset() {
	*m = JunosPfeIflset_234SystemTypeLinecardType{}
}
func (m *JunosPfeIflset_234SystemTypeLinecardType) String() string { return proto.CompactTextString(m) }
func (*JunosPfeIflset_234SystemTypeLinecardType) ProtoMessage()    {}
func (*JunosPfeIflset_234SystemTypeLinecardType) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b66e3e3153fc1b, []int{0, 0, 0}
}
func (m *JunosPfeIflset_234SystemTypeLinecardType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardType.Unmarshal(m, b)
}
func (m *JunosPfeIflset_234SystemTypeLinecardType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardType.Marshal(b, m, deterministic)
}
func (m *JunosPfeIflset_234SystemTypeLinecardType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardType.Merge(m, src)
}
func (m *JunosPfeIflset_234SystemTypeLinecardType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardType.Size(m)
}
func (m *JunosPfeIflset_234SystemTypeLinecardType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardType proto.InternalMessageInfo

func (m *JunosPfeIflset_234SystemTypeLinecardType) GetCos() *JunosPfeIflset_234SystemTypeLinecardTypeCosType {
	if m != nil {
		return m.Cos
	}
	return nil
}

type JunosPfeIflset_234SystemTypeLinecardTypeCosType struct {
	Interface            []*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList `protobuf:"bytes,151,rep,name=interface" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosType) Reset() {
	*m = JunosPfeIflset_234SystemTypeLinecardTypeCosType{}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosType) ProtoMessage() {}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosType) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b66e3e3153fc1b, []int{0, 0, 0, 0}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosType.Unmarshal(m, b)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosType.Marshal(b, m, deterministic)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosType.Merge(m, src)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosType.Size(m)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosType proto.InternalMessageInfo

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosType) GetInterface() []*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList {
	if m != nil {
		return m.Interface
	}
	return nil
}

type JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList struct {
	Name                 *string                                                                         `protobuf:"bytes,51,opt,name=name" json:"name,omitempty"`
	InterfaceSet         []*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList `protobuf:"bytes,151,rep,name=interface_set,json=interfaceSet" json:"interface_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                        `json:"-"`
	XXX_unrecognized     []byte                                                                          `json:"-"`
	XXX_sizecache        int32                                                                           `json:"-"`
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) Reset() {
	*m = JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList{}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) ProtoMessage() {}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b66e3e3153fc1b, []int{0, 0, 0, 0, 0}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList.Unmarshal(m, b)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList.Marshal(b, m, deterministic)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList.Merge(m, src)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) XXX_Size() int {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList.Size(m)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList proto.InternalMessageInfo

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList) GetInterfaceSet() []*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList {
	if m != nil {
		return m.InterfaceSet
	}
	return nil
}

type JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList struct {
	Iflsetindex          *uint64                                                                                 `protobuf:"varint,51,opt,name=iflsetindex" json:"iflsetindex,omitempty"`
	Output               *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType `protobuf:"bytes,151,opt,name=output" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                `json:"-"`
	XXX_unrecognized     []byte                                                                                  `json:"-"`
	XXX_sizecache        int32                                                                                   `json:"-"`
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) Reset() {
	*m = JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList{}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) ProtoMessage() {}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b66e3e3153fc1b, []int{0, 0, 0, 0, 0, 0}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList.Unmarshal(m, b)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList.Marshal(b, m, deterministic)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList.Merge(m, src)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) XXX_Size() int {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList.Size(m)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList proto.InternalMessageInfo

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) GetIflsetindex() uint64 {
	if m != nil && m.Iflsetindex != nil {
		return *m.Iflsetindex
	}
	return 0
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList) GetOutput() *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType {
	if m != nil {
		return m.Output
	}
	return nil
}

type JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType struct {
	Queue                []*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList `protobuf:"bytes,151,rep,name=queue" json:"queue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                           `json:"-"`
	XXX_unrecognized     []byte                                                                                             `json:"-"`
	XXX_sizecache        int32                                                                                              `json:"-"`
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) Reset() {
	*m = JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType{}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) ProtoMessage() {
}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b66e3e3153fc1b, []int{0, 0, 0, 0, 0, 0, 0}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType.Unmarshal(m, b)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType.Marshal(b, m, deterministic)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType.Merge(m, src)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType.Size(m)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType proto.InternalMessageInfo

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType) GetQueue() []*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList {
	if m != nil {
		return m.Queue
	}
	return nil
}

type JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList struct {
	QueueNumber          *uint64  `protobuf:"varint,51,opt,name=queue_number,json=queueNumber" json:"queue_number,omitempty"`
	QueuedPkts           *uint64  `protobuf:"varint,52,opt,name=queued_pkts,json=queuedPkts" json:"queued_pkts,omitempty"`
	QueuedOctets         *uint64  `protobuf:"varint,53,opt,name=queued_octets,json=queuedOctets" json:"queued_octets,omitempty"`
	TransmitPkts         *uint64  `protobuf:"varint,54,opt,name=transmit_pkts,json=transmitPkts" json:"transmit_pkts,omitempty"`
	TransmitOctets       *uint64  `protobuf:"varint,55,opt,name=transmit_octets,json=transmitOctets" json:"transmit_octets,omitempty"`
	TailDropPkts         *uint64  `protobuf:"varint,56,opt,name=tail_drop_pkts,json=tailDropPkts" json:"tail_drop_pkts,omitempty"`
	RedDropPkts          *uint64  `protobuf:"varint,57,opt,name=red_drop_pkts,json=redDropPkts" json:"red_drop_pkts,omitempty"`
	RedDropOctets        *uint64  `protobuf:"varint,58,opt,name=red_drop_octets,json=redDropOctets" json:"red_drop_octets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) Reset() {
	*m = JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList{}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) ProtoMessage() {
}
func (*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b66e3e3153fc1b, []int{0, 0, 0, 0, 0, 0, 0, 0}
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList.Unmarshal(m, b)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList.Marshal(b, m, deterministic)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList.Merge(m, src)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) XXX_Size() int {
	return xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList.Size(m)
}
func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList proto.InternalMessageInfo

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) GetQueueNumber() uint64 {
	if m != nil && m.QueueNumber != nil {
		return *m.QueueNumber
	}
	return 0
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) GetQueuedPkts() uint64 {
	if m != nil && m.QueuedPkts != nil {
		return *m.QueuedPkts
	}
	return 0
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) GetQueuedOctets() uint64 {
	if m != nil && m.QueuedOctets != nil {
		return *m.QueuedOctets
	}
	return 0
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) GetTransmitPkts() uint64 {
	if m != nil && m.TransmitPkts != nil {
		return *m.TransmitPkts
	}
	return 0
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) GetTransmitOctets() uint64 {
	if m != nil && m.TransmitOctets != nil {
		return *m.TransmitOctets
	}
	return 0
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) GetTailDropPkts() uint64 {
	if m != nil && m.TailDropPkts != nil {
		return *m.TailDropPkts
	}
	return 0
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) GetRedDropPkts() uint64 {
	if m != nil && m.RedDropPkts != nil {
		return *m.RedDropPkts
	}
	return 0
}

func (m *JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList) GetRedDropOctets() uint64 {
	if m != nil && m.RedDropOctets != nil {
		return *m.RedDropOctets
	}
	return 0
}

var E_JnprJunosPfeIflset_234Ext = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosPfeIflset_234)(nil),
	Field:         234,
	Name:          "jnpr_junos_pfe_iflset_234_ext",
	Tag:           "bytes,234,opt,name=jnpr_junos_pfe_iflset_234_ext",
	Filename:      "pfe_iflset_qstats.proto",
}

func init() {
	proto.RegisterType((*JunosPfeIflset_234)(nil), "junos_pfe_iflset_234")
	proto.RegisterType((*JunosPfeIflset_234SystemType)(nil), "junos_pfe_iflset_234.system_type")
	proto.RegisterType((*JunosPfeIflset_234SystemTypeLinecardType)(nil), "junos_pfe_iflset_234.system_type.linecard_type")
	proto.RegisterType((*JunosPfeIflset_234SystemTypeLinecardTypeCosType)(nil), "junos_pfe_iflset_234.system_type.linecard_type.cos_type")
	proto.RegisterType((*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceList)(nil), "junos_pfe_iflset_234.system_type.linecard_type.cos_type.interface_list")
	proto.RegisterType((*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetList)(nil), "junos_pfe_iflset_234.system_type.linecard_type.cos_type.interface_list.interface_set_list")
	proto.RegisterType((*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputType)(nil), "junos_pfe_iflset_234.system_type.linecard_type.cos_type.interface_list.interface_set_list.output_type")
	proto.RegisterType((*JunosPfeIflset_234SystemTypeLinecardTypeCosTypeInterfaceListInterfaceSetListOutputTypeQueueList)(nil), "junos_pfe_iflset_234.system_type.linecard_type.cos_type.interface_list.interface_set_list.output_type.queue_list")
	proto.RegisterExtension(E_JnprJunosPfeIflset_234Ext)
}

func init() { proto.RegisterFile("pfe_iflset_qstats.proto", fileDescriptor_e1b66e3e3153fc1b) }

var fileDescriptor_e1b66e3e3153fc1b = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x86, 0xe5, 0xb6, 0xc9, 0x4d, 0x8e, 0x9b, 0x54, 0x9a, 0x0b, 0xaa, 0xb1, 0x84, 0x28, 0x05,
	0xd1, 0xac, 0x8c, 0x94, 0x04, 0x28, 0x59, 0x21, 0x04, 0x42, 0x14, 0x54, 0xaa, 0x74, 0xc7, 0xc6,
	0x32, 0xf1, 0xb1, 0xe4, 0xc6, 0xf1, 0x4c, 0x66, 0x8e, 0x45, 0xb2, 0x02, 0xb1, 0xe2, 0x0d, 0x60,
	0xc3, 0x9b, 0x20, 0xf6, 0xec, 0xd8, 0xf1, 0x0e, 0xbc, 0x02, 0x1b, 0xe4, 0x19, 0xdb, 0x49, 0x20,
	0x12, 0x02, 0x01, 0xcb, 0xf9, 0xcf, 0x37, 0xdf, 0xf9, 0xe5, 0x31, 0xec, 0x8a, 0x08, 0xfd, 0x38,
	0x4a, 0x14, 0x92, 0x3f, 0x55, 0x14, 0x90, 0xf2, 0x84, 0xe4, 0xc4, 0xdd, 0xff, 0x09, 0x13, 0x9c,
	0x20, 0xc9, 0xb9, 0x4f, 0x5c, 0x98, 0x70, 0xff, 0x53, 0x13, 0xce, 0x9d, 0x65, 0x29, 0x57, 0xfe,
	0xd2, 0xb5, 0x6e, 0xaf, 0xcf, 0x06, 0x50, 0x57, 0x73, 0x45, 0x38, 0x71, 0xde, 0x58, 0x7b, 0x56,
	0xc7, 0xee, 0x5e, 0xf6, 0xd6, 0x71, 0x9e, 0x81, 0x7c, 0x9a, 0x0b, 0x1c, 0x16, 0x37, 0xdc, 0xaf,
	0x0d, 0xb0, 0x97, 0x72, 0xf6, 0x18, 0x1a, 0x49, 0x9c, 0xe2, 0x28, 0x90, 0x61, 0x69, 0xbb, 0xfe,
	0x53, 0x9b, 0x57, 0x5e, 0x31, 0xee, 0xca, 0xe0, 0xbe, 0x6b, 0x40, 0x6b, 0x65, 0xc6, 0x1e, 0xc1,
	0xe6, 0x88, 0xab, 0x52, 0x7d, 0xf8, 0x8b, 0x6a, 0x6f, 0xc4, 0x95, 0xd9, 0x91, 0x5b, 0xdc, 0x0f,
	0xff, 0x41, 0xa3, 0x4c, 0x58, 0x04, 0xcd, 0x38, 0x25, 0x94, 0x51, 0x30, 0xc2, 0xdc, 0xbf, 0xd9,
	0xb1, 0xbb, 0x0f, 0x7e, 0xd7, 0xef, 0x55, 0x2a, 0x3f, 0x89, 0x15, 0x0d, 0x17, 0x6a, 0xf7, 0x63,
	0x1d, 0xda, 0xab, 0x53, 0x76, 0x01, 0xb6, 0xd2, 0x60, 0x82, 0x4e, 0x6f, 0xcf, 0xea, 0x34, 0xef,
	0xd6, 0x5e, 0xdd, 0xd9, 0x68, 0x58, 0x43, 0x1d, 0xb1, 0x97, 0x16, 0xb4, 0x16, 0xb4, 0x42, 0x2a,
	0xab, 0x3d, 0xfd, 0x43, 0xd5, 0xbc, 0x15, 0xbb, 0x69, 0xbb, 0x5d, 0x65, 0xa7, 0x48, 0xee, 0xdb,
	0x1a, 0xb0, 0x1f, 0x21, 0x76, 0x00, 0xb6, 0xd9, 0x1b, 0xa7, 0x21, 0xce, 0x74, 0xf7, 0xad, 0xb2,
	0xfb, 0xf2, 0x84, 0xbd, 0x80, 0x3a, 0xcf, 0x48, 0x64, 0x54, 0xbe, 0x5a, 0xf4, 0xf7, 0xaa, 0x7b,
	0x66, 0x53, 0xf1, 0x8f, 0x9a, 0x83, 0xfb, 0x79, 0x13, 0xec, 0xa5, 0x9c, 0xbd, 0xb6, 0xa0, 0x36,
	0xcd, 0x30, 0xab, 0x9e, 0x79, 0xfa, 0x6f, 0x0a, 0x79, 0x7a, 0xa9, 0xf9, 0xc4, 0xa6, 0x80, 0xfb,
	0x7e, 0x03, 0x60, 0x91, 0xb2, 0x0e, 0x6c, 0x9b, 0x53, 0x9a, 0x4d, 0x9e, 0xa1, 0xfc, 0xee, 0xa3,
	0xea, 0xd1, 0xb1, 0x9e, 0xb0, 0x4b, 0x60, 0x8e, 0xa1, 0x2f, 0xc6, 0xa4, 0x9c, 0x7e, 0x0e, 0x0e,
	0x8d, 0x2a, 0x3c, 0x19, 0x93, 0x62, 0x57, 0xa0, 0x55, 0x00, 0x7c, 0x44, 0x48, 0xca, 0xb9, 0xa1,
	0x11, 0xe3, 0x0f, 0x9f, 0xe8, 0x2c, 0x87, 0x48, 0x06, 0xa9, 0x9a, 0xc4, 0x64, 0x3c, 0x37, 0x0d,
	0x54, 0x86, 0xda, 0x74, 0x00, 0x3b, 0x15, 0x54, 0xb8, 0x6e, 0x69, 0xac, 0x5d, 0xc6, 0x85, 0xed,
	0x2a, 0xb4, 0x29, 0x88, 0x13, 0x3f, 0x94, 0x5c, 0x18, 0xdd, 0x61, 0xa1, 0x0b, 0xe2, 0xe4, 0x9e,
	0xe4, 0x42, 0xeb, 0xf6, 0xa1, 0x25, 0x31, 0x5c, 0x82, 0x6e, 0x6b, 0xc8, 0x96, 0x18, 0x56, 0xcc,
	0x35, 0xd8, 0xa9, 0x98, 0x62, 0xe5, 0x40, 0x53, 0xad, 0x82, 0x32, 0x1b, 0x07, 0x02, 0x2e, 0x9e,
	0xa5, 0x42, 0xfa, 0xeb, 0x9e, 0xcf, 0xc7, 0x19, 0xb1, 0x5d, 0xef, 0x28, 0x4b, 0x63, 0x81, 0xf2,
	0x18, 0xe9, 0x39, 0x97, 0x63, 0x75, 0x8a, 0xa9, 0xe2, 0x52, 0x39, 0x5f, 0xcc, 0xaf, 0x78, 0x7e,
	0xed, 0xcb, 0x0f, 0x9d, 0xdc, 0x7a, 0x94, 0x4f, 0x4e, 0x22, 0x7c, 0xa8, 0xf3, 0x6e, 0xaf, 0x7f,
	0x7f, 0x46, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x11, 0x43, 0x5a, 0x73, 0x05, 0x00, 0x00,
}
