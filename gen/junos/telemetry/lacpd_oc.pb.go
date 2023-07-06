// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lacpd_oc.proto

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

type LacpLacp struct {
	State                *LacpLacpStateType      `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	Interfaces           *LacpLacpInterfacesType `protobuf:"bytes,152,opt,name=interfaces" json:"interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LacpLacp) Reset()         { *m = LacpLacp{} }
func (m *LacpLacp) String() string { return proto.CompactTextString(m) }
func (*LacpLacp) ProtoMessage()    {}
func (*LacpLacp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0}
}
func (m *LacpLacp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacp.Unmarshal(m, b)
}
func (m *LacpLacp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacp.Marshal(b, m, deterministic)
}
func (m *LacpLacp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacp.Merge(m, src)
}
func (m *LacpLacp) XXX_Size() int {
	return xxx_messageInfo_LacpLacp.Size(m)
}
func (m *LacpLacp) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacp.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacp proto.InternalMessageInfo

func (m *LacpLacp) GetState() *LacpLacpStateType {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *LacpLacp) GetInterfaces() *LacpLacpInterfacesType {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type LacpLacpStateType struct {
	SystemPriority       *uint32  `protobuf:"varint,51,opt,name=system_priority,json=systemPriority" json:"system_priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LacpLacpStateType) Reset()         { *m = LacpLacpStateType{} }
func (m *LacpLacpStateType) String() string { return proto.CompactTextString(m) }
func (*LacpLacpStateType) ProtoMessage()    {}
func (*LacpLacpStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0, 0}
}
func (m *LacpLacpStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacpStateType.Unmarshal(m, b)
}
func (m *LacpLacpStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacpStateType.Marshal(b, m, deterministic)
}
func (m *LacpLacpStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacpStateType.Merge(m, src)
}
func (m *LacpLacpStateType) XXX_Size() int {
	return xxx_messageInfo_LacpLacpStateType.Size(m)
}
func (m *LacpLacpStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacpStateType.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacpStateType proto.InternalMessageInfo

func (m *LacpLacpStateType) GetSystemPriority() uint32 {
	if m != nil && m.SystemPriority != nil {
		return *m.SystemPriority
	}
	return 0
}

type LacpLacpInterfacesType struct {
	Interface            []*LacpLacpInterfacesTypeInterfaceList `protobuf:"bytes,151,rep,name=interface" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *LacpLacpInterfacesType) Reset()         { *m = LacpLacpInterfacesType{} }
func (m *LacpLacpInterfacesType) String() string { return proto.CompactTextString(m) }
func (*LacpLacpInterfacesType) ProtoMessage()    {}
func (*LacpLacpInterfacesType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0, 1}
}
func (m *LacpLacpInterfacesType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacpInterfacesType.Unmarshal(m, b)
}
func (m *LacpLacpInterfacesType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacpInterfacesType.Marshal(b, m, deterministic)
}
func (m *LacpLacpInterfacesType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacpInterfacesType.Merge(m, src)
}
func (m *LacpLacpInterfacesType) XXX_Size() int {
	return xxx_messageInfo_LacpLacpInterfacesType.Size(m)
}
func (m *LacpLacpInterfacesType) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacpInterfacesType.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacpInterfacesType proto.InternalMessageInfo

func (m *LacpLacpInterfacesType) GetInterface() []*LacpLacpInterfacesTypeInterfaceList {
	if m != nil {
		return m.Interface
	}
	return nil
}

type LacpLacpInterfacesTypeInterfaceList struct {
	Name                 *string                                         `protobuf:"bytes,51,opt,name=name" json:"name,omitempty"`
	State                *LacpLacpInterfacesTypeInterfaceListStateType   `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	Members              *LacpLacpInterfacesTypeInterfaceListMembersType `protobuf:"bytes,152,opt,name=members" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *LacpLacpInterfacesTypeInterfaceList) Reset()         { *m = LacpLacpInterfacesTypeInterfaceList{} }
func (m *LacpLacpInterfacesTypeInterfaceList) String() string { return proto.CompactTextString(m) }
func (*LacpLacpInterfacesTypeInterfaceList) ProtoMessage()    {}
func (*LacpLacpInterfacesTypeInterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0, 1, 0}
}
func (m *LacpLacpInterfacesTypeInterfaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceList.Unmarshal(m, b)
}
func (m *LacpLacpInterfacesTypeInterfaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceList.Marshal(b, m, deterministic)
}
func (m *LacpLacpInterfacesTypeInterfaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceList.Merge(m, src)
}
func (m *LacpLacpInterfacesTypeInterfaceList) XXX_Size() int {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceList.Size(m)
}
func (m *LacpLacpInterfacesTypeInterfaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceList.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacpInterfacesTypeInterfaceList proto.InternalMessageInfo

func (m *LacpLacpInterfacesTypeInterfaceList) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceList) GetState() *LacpLacpInterfacesTypeInterfaceListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *LacpLacpInterfacesTypeInterfaceList) GetMembers() *LacpLacpInterfacesTypeInterfaceListMembersType {
	if m != nil {
		return m.Members
	}
	return nil
}

type LacpLacpInterfacesTypeInterfaceListStateType struct {
	Name                 *string  `protobuf:"bytes,51,opt,name=name" json:"name,omitempty"`
	Interval             *string  `protobuf:"bytes,52,opt,name=interval" json:"interval,omitempty"`
	LacpMode             *string  `protobuf:"bytes,53,opt,name=lacp_mode,json=lacpMode" json:"lacp_mode,omitempty"`
	SystemIdMac          *string  `protobuf:"bytes,54,opt,name=system_id_mac,json=systemIdMac" json:"system_id_mac,omitempty"`
	SystemPriority       *uint32  `protobuf:"varint,55,opt,name=system_priority,json=systemPriority" json:"system_priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LacpLacpInterfacesTypeInterfaceListStateType) Reset() {
	*m = LacpLacpInterfacesTypeInterfaceListStateType{}
}
func (m *LacpLacpInterfacesTypeInterfaceListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*LacpLacpInterfacesTypeInterfaceListStateType) ProtoMessage() {}
func (*LacpLacpInterfacesTypeInterfaceListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0, 1, 0, 0}
}
func (m *LacpLacpInterfacesTypeInterfaceListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListStateType.Unmarshal(m, b)
}
func (m *LacpLacpInterfacesTypeInterfaceListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListStateType.Marshal(b, m, deterministic)
}
func (m *LacpLacpInterfacesTypeInterfaceListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListStateType.Merge(m, src)
}
func (m *LacpLacpInterfacesTypeInterfaceListStateType) XXX_Size() int {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListStateType.Size(m)
}
func (m *LacpLacpInterfacesTypeInterfaceListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListStateType proto.InternalMessageInfo

func (m *LacpLacpInterfacesTypeInterfaceListStateType) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListStateType) GetInterval() string {
	if m != nil && m.Interval != nil {
		return *m.Interval
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListStateType) GetLacpMode() string {
	if m != nil && m.LacpMode != nil {
		return *m.LacpMode
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListStateType) GetSystemIdMac() string {
	if m != nil && m.SystemIdMac != nil {
		return *m.SystemIdMac
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListStateType) GetSystemPriority() uint32 {
	if m != nil && m.SystemPriority != nil {
		return *m.SystemPriority
	}
	return 0
}

type LacpLacpInterfacesTypeInterfaceListMembersType struct {
	Member               []*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList `protobuf:"bytes,151,rep,name=member" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersType) Reset() {
	*m = LacpLacpInterfacesTypeInterfaceListMembersType{}
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersType) String() string {
	return proto.CompactTextString(m)
}
func (*LacpLacpInterfacesTypeInterfaceListMembersType) ProtoMessage() {}
func (*LacpLacpInterfacesTypeInterfaceListMembersType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0, 1, 0, 1}
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersType.Unmarshal(m, b)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersType.Marshal(b, m, deterministic)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersType.Merge(m, src)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersType) XXX_Size() int {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersType.Size(m)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersType) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersType.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersType proto.InternalMessageInfo

func (m *LacpLacpInterfacesTypeInterfaceListMembersType) GetMember() []*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList {
	if m != nil {
		return m.Member
	}
	return nil
}

type LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList struct {
	Interface            *string                                                            `protobuf:"bytes,51,opt,name=interface" json:"interface,omitempty"`
	State                *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                           `json:"-"`
	XXX_unrecognized     []byte                                                             `json:"-"`
	XXX_sizecache        int32                                                              `json:"-"`
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) Reset() {
	*m = LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList{}
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) String() string {
	return proto.CompactTextString(m)
}
func (*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) ProtoMessage() {}
func (*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0, 1, 0, 1, 0}
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList.Unmarshal(m, b)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList.Marshal(b, m, deterministic)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList.Merge(m, src)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) XXX_Size() int {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList.Size(m)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList proto.InternalMessageInfo

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) GetInterface() string {
	if m != nil && m.Interface != nil {
		return *m.Interface
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList) GetState() *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType struct {
	Interface              *string                                                                        `protobuf:"bytes,51,opt,name=interface" json:"interface,omitempty"`
	Activity               *string                                                                        `protobuf:"bytes,52,opt,name=activity" json:"activity,omitempty"`
	Timeout                *string                                                                        `protobuf:"bytes,53,opt,name=timeout" json:"timeout,omitempty"`
	Synchronization        *string                                                                        `protobuf:"bytes,54,opt,name=synchronization" json:"synchronization,omitempty"`
	Aggregatable           *bool                                                                          `protobuf:"varint,55,opt,name=aggregatable" json:"aggregatable,omitempty"`
	Collecting             *bool                                                                          `protobuf:"varint,56,opt,name=collecting" json:"collecting,omitempty"`
	Distributing           *bool                                                                          `protobuf:"varint,57,opt,name=distributing" json:"distributing,omitempty"`
	SystemId               *string                                                                        `protobuf:"bytes,58,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
	OperKey                *uint32                                                                        `protobuf:"varint,59,opt,name=oper_key,json=operKey" json:"oper_key,omitempty"`
	PartnerId              *string                                                                        `protobuf:"bytes,61,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
	PartnerKey             *uint32                                                                        `protobuf:"varint,62,opt,name=partner_key,json=partnerKey" json:"partner_key,omitempty"`
	PortNum                *uint32                                                                        `protobuf:"varint,60,opt,name=port_num,json=portNum" json:"port_num,omitempty"`
	PartnerPortNum         *uint32                                                                        `protobuf:"varint,63,opt,name=partner_port_num,json=partnerPortNum" json:"partner_port_num,omitempty"`
	LastChange             *uint64                                                                        `protobuf:"varint,70,opt,name=last_change,json=lastChange" json:"last_change,omitempty"`
	Counters               *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType `protobuf:"bytes,151,opt,name=counters" json:"counters,omitempty"`
	MuxState               *string                                                                        `protobuf:"bytes,64,opt,name=mux_state,json=muxState" json:"mux_state,omitempty"`
	PartnerTimeout         *string                                                                        `protobuf:"bytes,65,opt,name=partner_timeout,json=partnerTimeout" json:"partner_timeout,omitempty"`
	PartnerSynchronization *string                                                                        `protobuf:"bytes,66,opt,name=partner_synchronization,json=partnerSynchronization" json:"partner_synchronization,omitempty"`
	PartnerAggregatable    *bool                                                                          `protobuf:"varint,67,opt,name=partner_aggregatable,json=partnerAggregatable" json:"partner_aggregatable,omitempty"`
	PartnerCollecting      *bool                                                                          `protobuf:"varint,68,opt,name=partner_collecting,json=partnerCollecting" json:"partner_collecting,omitempty"`
	PartnerDistributing    *bool                                                                          `protobuf:"varint,69,opt,name=partner_distributing,json=partnerDistributing" json:"partner_distributing,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                                                                       `json:"-"`
	XXX_unrecognized       []byte                                                                         `json:"-"`
	XXX_sizecache          int32                                                                          `json:"-"`
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) Reset() {
	*m = LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType{}
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) ProtoMessage() {}
func (*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0, 1, 0, 1, 0, 0}
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType.Unmarshal(m, b)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType.Marshal(b, m, deterministic)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType.Merge(m, src)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) XXX_Size() int {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType.Size(m)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType proto.InternalMessageInfo

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetInterface() string {
	if m != nil && m.Interface != nil {
		return *m.Interface
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetActivity() string {
	if m != nil && m.Activity != nil {
		return *m.Activity
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetTimeout() string {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetSynchronization() string {
	if m != nil && m.Synchronization != nil {
		return *m.Synchronization
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetAggregatable() bool {
	if m != nil && m.Aggregatable != nil {
		return *m.Aggregatable
	}
	return false
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetCollecting() bool {
	if m != nil && m.Collecting != nil {
		return *m.Collecting
	}
	return false
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetDistributing() bool {
	if m != nil && m.Distributing != nil {
		return *m.Distributing
	}
	return false
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetSystemId() string {
	if m != nil && m.SystemId != nil {
		return *m.SystemId
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetOperKey() uint32 {
	if m != nil && m.OperKey != nil {
		return *m.OperKey
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPartnerId() string {
	if m != nil && m.PartnerId != nil {
		return *m.PartnerId
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPartnerKey() uint32 {
	if m != nil && m.PartnerKey != nil {
		return *m.PartnerKey
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPortNum() uint32 {
	if m != nil && m.PortNum != nil {
		return *m.PortNum
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPartnerPortNum() uint32 {
	if m != nil && m.PartnerPortNum != nil {
		return *m.PartnerPortNum
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetLastChange() uint64 {
	if m != nil && m.LastChange != nil {
		return *m.LastChange
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetCounters() *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetMuxState() string {
	if m != nil && m.MuxState != nil {
		return *m.MuxState
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPartnerTimeout() string {
	if m != nil && m.PartnerTimeout != nil {
		return *m.PartnerTimeout
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPartnerSynchronization() string {
	if m != nil && m.PartnerSynchronization != nil {
		return *m.PartnerSynchronization
	}
	return ""
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPartnerAggregatable() bool {
	if m != nil && m.PartnerAggregatable != nil {
		return *m.PartnerAggregatable
	}
	return false
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPartnerCollecting() bool {
	if m != nil && m.PartnerCollecting != nil {
		return *m.PartnerCollecting
	}
	return false
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType) GetPartnerDistributing() bool {
	if m != nil && m.PartnerDistributing != nil {
		return *m.PartnerDistributing
	}
	return false
}

type LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType struct {
	LacpInPkts             *uint64  `protobuf:"varint,51,opt,name=lacp_in_pkts,json=lacpInPkts" json:"lacp_in_pkts,omitempty"`
	LacpOutPkts            *uint64  `protobuf:"varint,52,opt,name=lacp_out_pkts,json=lacpOutPkts" json:"lacp_out_pkts,omitempty"`
	LacpRxErrors           *uint64  `protobuf:"varint,53,opt,name=lacp_rx_errors,json=lacpRxErrors" json:"lacp_rx_errors,omitempty"`
	LacpTxErrors           *uint64  `protobuf:"varint,54,opt,name=lacp_tx_errors,json=lacpTxErrors" json:"lacp_tx_errors,omitempty"`
	LacpUnknownErrors      *uint64  `protobuf:"varint,55,opt,name=lacp_unknown_errors,json=lacpUnknownErrors" json:"lacp_unknown_errors,omitempty"`
	LacpErrors             *uint64  `protobuf:"varint,56,opt,name=lacp_errors,json=lacpErrors" json:"lacp_errors,omitempty"`
	LacpTimeoutTransitions *uint64  `protobuf:"varint,58,opt,name=lacp_timeout_transitions,json=lacpTimeoutTransitions" json:"lacp_timeout_transitions,omitempty"`
	CollectionTime         *string  `protobuf:"bytes,57,opt,name=collection_time,json=collectionTime" json:"collection_time,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) Reset() {
	*m = LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType{}
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) String() string {
	return proto.CompactTextString(m)
}
func (*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) ProtoMessage() {
}
func (*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bfdbbc56d07876e, []int{0, 1, 0, 1, 0, 0, 0}
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType.Unmarshal(m, b)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType.Marshal(b, m, deterministic)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType.Merge(m, src)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) XXX_Size() int {
	return xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType.Size(m)
}
func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType proto.InternalMessageInfo

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) GetLacpInPkts() uint64 {
	if m != nil && m.LacpInPkts != nil {
		return *m.LacpInPkts
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) GetLacpOutPkts() uint64 {
	if m != nil && m.LacpOutPkts != nil {
		return *m.LacpOutPkts
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) GetLacpRxErrors() uint64 {
	if m != nil && m.LacpRxErrors != nil {
		return *m.LacpRxErrors
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) GetLacpTxErrors() uint64 {
	if m != nil && m.LacpTxErrors != nil {
		return *m.LacpTxErrors
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) GetLacpUnknownErrors() uint64 {
	if m != nil && m.LacpUnknownErrors != nil {
		return *m.LacpUnknownErrors
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) GetLacpErrors() uint64 {
	if m != nil && m.LacpErrors != nil {
		return *m.LacpErrors
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) GetLacpTimeoutTransitions() uint64 {
	if m != nil && m.LacpTimeoutTransitions != nil {
		return *m.LacpTimeoutTransitions
	}
	return 0
}

func (m *LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType) GetCollectionTime() string {
	if m != nil && m.CollectionTime != nil {
		return *m.CollectionTime
	}
	return ""
}

var E_JnprLacpLacpExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*LacpLacp)(nil),
	Field:         52,
	Name:          "jnpr_lacp_lacp_ext",
	Tag:           "bytes,52,opt,name=jnpr_lacp_lacp_ext",
	Filename:      "lacpd_oc.proto",
}

func init() {
	proto.RegisterType((*LacpLacp)(nil), "lacp_lacp")
	proto.RegisterType((*LacpLacpStateType)(nil), "lacp_lacp.state_type")
	proto.RegisterType((*LacpLacpInterfacesType)(nil), "lacp_lacp.interfaces_type")
	proto.RegisterType((*LacpLacpInterfacesTypeInterfaceList)(nil), "lacp_lacp.interfaces_type.interface_list")
	proto.RegisterType((*LacpLacpInterfacesTypeInterfaceListStateType)(nil), "lacp_lacp.interfaces_type.interface_list.state_type")
	proto.RegisterType((*LacpLacpInterfacesTypeInterfaceListMembersType)(nil), "lacp_lacp.interfaces_type.interface_list.members_type")
	proto.RegisterType((*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberList)(nil), "lacp_lacp.interfaces_type.interface_list.members_type.member_list")
	proto.RegisterType((*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateType)(nil), "lacp_lacp.interfaces_type.interface_list.members_type.member_list.state_type")
	proto.RegisterType((*LacpLacpInterfacesTypeInterfaceListMembersTypeMemberListStateTypeCountersType)(nil), "lacp_lacp.interfaces_type.interface_list.members_type.member_list.state_type.counters_type")
	proto.RegisterExtension(E_JnprLacpLacpExt)
}

func init() { proto.RegisterFile("lacpd_oc.proto", fileDescriptor_8bfdbbc56d07876e) }

var fileDescriptor_8bfdbbc56d07876e = []byte{
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0x4b, 0x4a, 0xec, 0xe3, 0xc4, 0x6e, 0x27, 0xd0, 0x6e, 0x96, 0x9f, 0x46, 0x01, 0xa9,
	0x46, 0x02, 0x4b, 0x94, 0x34, 0x09, 0x2d, 0x3f, 0x6d, 0xd2, 0x20, 0x42, 0x49, 0x31, 0x9b, 0x70,
	0xd5, 0x8b, 0xd1, 0x64, 0x3d, 0xb8, 0x8b, 0xbd, 0x33, 0xab, 0xd9, 0xd9, 0xd6, 0xe6, 0x92, 0x67,
	0x40, 0x82, 0x27, 0xe0, 0x11, 0x7a, 0xcb, 0x0d, 0x2f, 0xc0, 0x05, 0xef, 0x83, 0xe6, 0xcc, 0xcc,
	0xfe, 0xb8, 0x41, 0xaa, 0x80, 0x8b, 0x44, 0x9a, 0xef, 0x7c, 0xdf, 0x99, 0x99, 0xf3, 0x7d, 0xb3,
	0x86, 0xde, 0x8c, 0xc5, 0xd9, 0x98, 0xca, 0x78, 0x98, 0x29, 0xa9, 0x65, 0xb8, 0xa1, 0xf9, 0x8c,
	0xa7, 0x5c, 0xab, 0x05, 0xd5, 0x32, 0xb3, 0xe0, 0xf6, 0x1f, 0x57, 0xa0, 0x63, 0x78, 0xd4, 0xfc,
	0x23, 0xef, 0xc3, 0xe5, 0x5c, 0x33, 0xcd, 0x83, 0x5f, 0x5a, 0x5b, 0xad, 0x41, 0xf7, 0xd6, 0xeb,
	0xc3, 0xb2, 0x36, 0xc4, 0x02, 0xd5, 0x8b, 0x8c, 0x47, 0x96, 0x44, 0xee, 0x02, 0x24, 0x42, 0x73,
	0xf5, 0x3d, 0x8b, 0x79, 0x1e, 0xfc, 0x6a, 0x25, 0x61, 0x4d, 0x52, 0x55, 0xad, 0xae, 0x46, 0x0f,
	0x6f, 0x03, 0x54, 0x1d, 0xc9, 0x4d, 0xe8, 0xe7, 0x8b, 0x5c, 0xf3, 0x94, 0x66, 0x2a, 0x91, 0x2a,
	0xd1, 0x8b, 0xe0, 0xa3, 0xad, 0xd6, 0x60, 0x3d, 0xea, 0x59, 0x78, 0xe4, 0xd0, 0xf0, 0xe7, 0x3e,
	0xf4, 0x97, 0xda, 0x92, 0x2f, 0xa1, 0x53, 0x42, 0xe6, 0xe4, 0xaf, 0x0c, 0xba, 0xb7, 0xde, 0xfb,
	0xe7, 0x63, 0x54, 0x6b, 0x3a, 0x4b, 0x72, 0x1d, 0x55, 0xe2, 0xf0, 0x79, 0x0f, 0x7a, 0xcd, 0x2a,
	0xd9, 0x84, 0x15, 0xc1, 0x52, 0x8e, 0xc7, 0xe9, 0x1c, 0x5c, 0xfe, 0xe9, 0xde, 0xa5, 0x76, 0x2b,
	0x42, 0x88, 0x3c, 0x5c, 0x9a, 0xd6, 0xce, 0x4b, 0xef, 0x79, 0xc1, 0x30, 0xbf, 0x85, 0xd5, 0x94,
	0xa7, 0xe7, 0x5c, 0x95, 0x93, 0xdc, 0x7d, 0xf9, 0x76, 0x4e, 0x69, 0x1b, 0xfa, 0x3e, 0xe1, 0x6f,
	0xad, 0xc6, 0x8c, 0x49, 0xfd, 0x26, 0xee, 0x0a, 0x21, 0xb4, 0xb1, 0xd5, 0x53, 0x36, 0x0b, 0x76,
	0x10, 0x2f, 0xd7, 0xe4, 0x0d, 0x97, 0x8c, 0x54, 0x8e, 0x79, 0x70, 0xdb, 0x16, 0x0d, 0x70, 0x22,
	0xc7, 0x9c, 0x6c, 0xc3, 0xba, 0x33, 0x2c, 0x19, 0xd3, 0x94, 0xc5, 0xc1, 0x2e, 0x12, 0xba, 0x16,
	0x3c, 0x1e, 0x9f, 0xb0, 0xf8, 0x22, 0x53, 0xf7, 0x2e, 0x34, 0xf5, 0xf7, 0x2e, 0xac, 0xd5, 0xaf,
	0x40, 0x1e, 0xc3, 0xab, 0x76, 0xed, 0xed, 0x3c, 0xf8, 0x77, 0xb3, 0x70, 0x0b, 0xeb, 0xb3, 0x6b,
	0x19, 0xfe, 0x09, 0xd0, 0xad, 0xe1, 0xe4, 0x9d, 0x7a, 0x7c, 0x1a, 0x36, 0x57, 0x38, 0x19, 0x2f,
	0x79, 0x7d, 0xf2, 0xdf, 0x0f, 0xf4, 0x62, 0x08, 0xc2, 0xe7, 0x9d, 0x86, 0x63, 0x6f, 0xbe, 0x70,
	0xb2, 0xfa, 0x91, 0x42, 0x68, 0xb3, 0x58, 0x27, 0x4f, 0xcd, 0x5c, 0x9d, 0x77, 0x7e, 0x4d, 0x02,
	0x58, 0xd5, 0x49, 0xca, 0x65, 0xa1, 0x9d, 0x73, 0x7e, 0x49, 0x06, 0xc6, 0x14, 0x11, 0x3f, 0x51,
	0x52, 0x24, 0x3f, 0x32, 0x9d, 0x48, 0xe1, 0xac, 0x5b, 0x86, 0xc9, 0x36, 0xac, 0xb1, 0xc9, 0x44,
	0xf1, 0x09, 0xd3, 0xec, 0x7c, 0xc6, 0xd1, 0xbb, 0x76, 0xd4, 0xc0, 0xc8, 0xdb, 0x00, 0xb1, 0x9c,
	0xcd, 0x78, 0xac, 0x13, 0x31, 0x09, 0xf6, 0x91, 0x51, 0x43, 0x4c, 0x8f, 0x71, 0x92, 0x6b, 0x95,
	0x9c, 0x17, 0xc8, 0xf8, 0xd8, 0xf6, 0xa8, 0x63, 0x26, 0x67, 0x65, 0x94, 0x82, 0x3b, 0xf6, 0x22,
	0x3e, 0x46, 0x64, 0x13, 0xda, 0x32, 0xe3, 0x8a, 0x4e, 0xf9, 0x22, 0xb8, 0x8b, 0xe1, 0x59, 0x35,
	0xeb, 0x87, 0x7c, 0x41, 0xde, 0x02, 0xc8, 0x98, 0xd2, 0x82, 0x2b, 0x23, 0xfc, 0xd4, 0x8e, 0xc7,
	0x21, 0xc7, 0x63, 0x72, 0x03, 0xba, 0xbe, 0x6c, 0xc4, 0x9f, 0xa1, 0xd8, 0x2b, 0x8c, 0x7e, 0x13,
	0xda, 0x99, 0x54, 0x9a, 0x8a, 0x22, 0x0d, 0x3e, 0xb1, 0xad, 0xcd, 0xfa, 0x51, 0x91, 0x92, 0x01,
	0x5c, 0xf1, 0xda, 0x92, 0xf2, 0xb9, 0x8d, 0xae, 0xc3, 0x47, 0x8e, 0x79, 0x03, 0xba, 0x33, 0x96,
	0x6b, 0x1a, 0x3f, 0x61, 0x62, 0xc2, 0x83, 0x2f, 0xb6, 0x5a, 0x83, 0x95, 0x08, 0x0c, 0x74, 0x88,
	0x08, 0x99, 0x43, 0x3b, 0x96, 0x85, 0x31, 0x2d, 0xf7, 0xd9, 0x79, 0xfc, 0xbf, 0x66, 0x67, 0xe8,
	0xdb, 0xdb, 0x24, 0x95, 0xbb, 0x99, 0xb9, 0xa6, 0xc5, 0x9c, 0xda, 0xd8, 0xde, 0xb3, 0x73, 0x4d,
	0x8b, 0xf9, 0x29, 0x7e, 0x6e, 0x6e, 0x42, 0xdf, 0xdf, 0xd0, 0x07, 0xe5, 0x3e, 0x52, 0xfc, 0x05,
	0xcf, 0x5c, 0x5e, 0xf6, 0xe0, 0xba, 0x27, 0x2e, 0xe7, 0xe6, 0x00, 0x05, 0xd7, 0x5c, 0xf9, 0x74,
	0x29, 0x3e, 0x1f, 0xc2, 0x6b, 0x5e, 0xd8, 0x88, 0xd1, 0x21, 0x46, 0x60, 0xc3, 0xd5, 0xee, 0xd7,
	0xd3, 0xf4, 0x01, 0x10, 0x2f, 0xa9, 0xa5, 0xea, 0x01, 0x0a, 0xae, 0xba, 0xca, 0x61, 0x15, 0xae,
	0xda, 0x0e, 0x8d, 0x90, 0x1d, 0x35, 0x76, 0x78, 0x50, 0x2b, 0x85, 0x7f, 0x5d, 0x82, 0xf5, 0xc6,
	0xbc, 0xc8, 0x16, 0xac, 0xa1, 0x1b, 0x89, 0xa0, 0xd9, 0x54, 0xe7, 0xf8, 0xcc, 0xd0, 0xc1, 0x38,
	0x3b, 0x16, 0xa3, 0xa9, 0xce, 0xcd, 0xa7, 0x0e, 0x19, 0xb2, 0xd0, 0x96, 0xb2, 0x83, 0x94, 0xae,
	0x01, 0xbf, 0x29, 0x34, 0x72, 0xde, 0xb5, 0xbf, 0xb6, 0x54, 0xcd, 0x29, 0x57, 0x4a, 0xaa, 0x1c,
	0x9f, 0xdd, 0x4a, 0x84, 0xbd, 0xa3, 0xf9, 0x11, 0x62, 0x25, 0x4b, 0x97, 0xac, 0xdd, 0x8a, 0x75,
	0xe6, 0x59, 0x43, 0xd8, 0x40, 0x56, 0x21, 0xa6, 0x42, 0x3e, 0x13, 0x9e, 0xba, 0x87, 0xd4, 0xab,
	0xa6, 0xf4, 0x9d, 0xad, 0x38, 0x3e, 0x46, 0x30, 0xce, 0x3c, 0x6f, 0xbf, 0xba, 0x80, 0x23, 0xec,
	0x43, 0x60, 0xb7, 0xb5, 0x96, 0x52, 0xad, 0x98, 0xc8, 0x13, 0x63, 0x52, 0x8e, 0xef, 0x6d, 0x25,
	0xba, 0x86, 0x07, 0xb0, 0xe5, 0xb3, 0xaa, 0x6a, 0x52, 0xe2, 0x8d, 0x90, 0x02, 0xf5, 0xf8, 0x82,
	0x3b, 0x51, 0xaf, 0x82, 0x8d, 0xec, 0xce, 0x08, 0xc8, 0x0f, 0x22, 0x53, 0xb4, 0x0c, 0x36, 0xe5,
	0x73, 0x4d, 0xae, 0x0f, 0xbf, 0x2a, 0x44, 0x92, 0x71, 0xf5, 0x88, 0xeb, 0x67, 0x52, 0x4d, 0xf3,
	0x53, 0x2e, 0x72, 0x73, 0xc6, 0x1d, 0x7c, 0x07, 0x50, 0xbd, 0x83, 0xa8, 0x6f, 0xe4, 0x5f, 0xb3,
	0x38, 0x33, 0x7f, 0x47, 0x73, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xe9, 0x8c, 0x4c,
	0xbe, 0x08, 0x00, 0x00,
}
