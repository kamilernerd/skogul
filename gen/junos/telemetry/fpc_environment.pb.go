// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fpc_environment.proto

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

// Temperature sensor information
type GTempInfo struct {
	// Temperature sensor name
	TempSensorName *string `protobuf:"bytes,1,req,name=temp_sensor_name,json=tempSensorName" json:"temp_sensor_name,omitempty"`
	// Gauge: the current temperature read by this sensor, in degrees Celcius
	TempValue            *int32   `protobuf:"varint,2,req,name=temp_value,json=tempValue" json:"temp_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GTempInfo) Reset()         { *m = GTempInfo{} }
func (m *GTempInfo) String() string { return proto.CompactTextString(m) }
func (*GTempInfo) ProtoMessage()    {}
func (*GTempInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_80caabd2d4d5ba4d, []int{0}
}
func (m *GTempInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GTempInfo.Unmarshal(m, b)
}
func (m *GTempInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GTempInfo.Marshal(b, m, deterministic)
}
func (m *GTempInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GTempInfo.Merge(m, src)
}
func (m *GTempInfo) XXX_Size() int {
	return xxx_messageInfo_GTempInfo.Size(m)
}
func (m *GTempInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GTempInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GTempInfo proto.InternalMessageInfo

func (m *GTempInfo) GetTempSensorName() string {
	if m != nil && m.TempSensorName != nil {
		return *m.TempSensorName
	}
	return ""
}

func (m *GTempInfo) GetTempValue() int32 {
	if m != nil && m.TempValue != nil {
		return *m.TempValue
	}
	return 0
}

// Voltage sensor information
type GVoltageInfo struct {
	// Voltage sensor name
	VoltageSensorName *string `protobuf:"bytes,1,req,name=voltage_sensor_name,json=voltageSensorName" json:"voltage_sensor_name,omitempty"`
	// Gauge: The current voltage (in milli volts) read by this sensor
	VoltageValue         *int32   `protobuf:"varint,2,req,name=voltage_value,json=voltageValue" json:"voltage_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GVoltageInfo) Reset()         { *m = GVoltageInfo{} }
func (m *GVoltageInfo) String() string { return proto.CompactTextString(m) }
func (*GVoltageInfo) ProtoMessage()    {}
func (*GVoltageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_80caabd2d4d5ba4d, []int{1}
}
func (m *GVoltageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GVoltageInfo.Unmarshal(m, b)
}
func (m *GVoltageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GVoltageInfo.Marshal(b, m, deterministic)
}
func (m *GVoltageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GVoltageInfo.Merge(m, src)
}
func (m *GVoltageInfo) XXX_Size() int {
	return xxx_messageInfo_GVoltageInfo.Size(m)
}
func (m *GVoltageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GVoltageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GVoltageInfo proto.InternalMessageInfo

func (m *GVoltageInfo) GetVoltageSensorName() string {
	if m != nil && m.VoltageSensorName != nil {
		return *m.VoltageSensorName
	}
	return ""
}

func (m *GVoltageInfo) GetVoltageValue() int32 {
	if m != nil && m.VoltageValue != nil {
		return *m.VoltageValue
	}
	return 0
}

// Power consumed by FPC, ASICs & Max consumption
type GPowerInfo struct {
	// This is maximum power, in Watts, that this fpc is allowed to consume
	MaxFpcPower *int32 `protobuf:"varint,1,opt,name=max_fpc_power,json=maxFpcPower" json:"max_fpc_power,omitempty"`
	// This is the amount of power, in Watts, that this fpc currently is consuming
	FpcPower             *int32   `protobuf:"varint,2,opt,name=fpc_power,json=fpcPower" json:"fpc_power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GPowerInfo) Reset()         { *m = GPowerInfo{} }
func (m *GPowerInfo) String() string { return proto.CompactTextString(m) }
func (*GPowerInfo) ProtoMessage()    {}
func (*GPowerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_80caabd2d4d5ba4d, []int{2}
}
func (m *GPowerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPowerInfo.Unmarshal(m, b)
}
func (m *GPowerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPowerInfo.Marshal(b, m, deterministic)
}
func (m *GPowerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPowerInfo.Merge(m, src)
}
func (m *GPowerInfo) XXX_Size() int {
	return xxx_messageInfo_GPowerInfo.Size(m)
}
func (m *GPowerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GPowerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GPowerInfo proto.InternalMessageInfo

func (m *GPowerInfo) GetMaxFpcPower() int32 {
	if m != nil && m.MaxFpcPower != nil {
		return *m.MaxFpcPower
	}
	return 0
}

func (m *GPowerInfo) GetFpcPower() int32 {
	if m != nil && m.FpcPower != nil {
		return *m.FpcPower
	}
	return 0
}

// Main FPC Environment message, which will be shipped out
type GFpcEnvironment struct {
	// FPC power info
	PowerRecord *GPowerInfo `protobuf:"bytes,1,opt,name=power_record,json=powerRecord" json:"power_record,omitempty"`
	// Various temp sensor info
	TempRecord []*GTempInfo `protobuf:"bytes,2,rep,name=temp_record,json=tempRecord" json:"temp_record,omitempty"`
	// Various voltage sensor info
	VoltageRecord        []*GVoltageInfo                `protobuf:"bytes,3,rep,name=voltage_record,json=voltageRecord" json:"voltage_record,omitempty"`
	Linecard             []*GFpcEnvironmentLinecardList `protobuf:"bytes,4,rep,name=linecard" json:"linecard,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GFpcEnvironment) Reset()         { *m = GFpcEnvironment{} }
func (m *GFpcEnvironment) String() string { return proto.CompactTextString(m) }
func (*GFpcEnvironment) ProtoMessage()    {}
func (*GFpcEnvironment) Descriptor() ([]byte, []int) {
	return fileDescriptor_80caabd2d4d5ba4d, []int{3}
}
func (m *GFpcEnvironment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GFpcEnvironment.Unmarshal(m, b)
}
func (m *GFpcEnvironment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GFpcEnvironment.Marshal(b, m, deterministic)
}
func (m *GFpcEnvironment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GFpcEnvironment.Merge(m, src)
}
func (m *GFpcEnvironment) XXX_Size() int {
	return xxx_messageInfo_GFpcEnvironment.Size(m)
}
func (m *GFpcEnvironment) XXX_DiscardUnknown() {
	xxx_messageInfo_GFpcEnvironment.DiscardUnknown(m)
}

var xxx_messageInfo_GFpcEnvironment proto.InternalMessageInfo

func (m *GFpcEnvironment) GetPowerRecord() *GPowerInfo {
	if m != nil {
		return m.PowerRecord
	}
	return nil
}

func (m *GFpcEnvironment) GetTempRecord() []*GTempInfo {
	if m != nil {
		return m.TempRecord
	}
	return nil
}

func (m *GFpcEnvironment) GetVoltageRecord() []*GVoltageInfo {
	if m != nil {
		return m.VoltageRecord
	}
	return nil
}

func (m *GFpcEnvironment) GetLinecard() []*GFpcEnvironmentLinecardList {
	if m != nil {
		return m.Linecard
	}
	return nil
}

type GFpcEnvironmentLinecardList struct {
	Name                 *string         `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	PowerRecord          *GPowerInfo     `protobuf:"bytes,1,opt,name=power_record,json=powerRecord" json:"power_record,omitempty"`
	TempRecord           []*GTempInfo    `protobuf:"bytes,2,rep,name=temp_record,json=tempRecord" json:"temp_record,omitempty"`
	VoltageRecord        []*GVoltageInfo `protobuf:"bytes,3,rep,name=voltage_record,json=voltageRecord" json:"voltage_record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GFpcEnvironmentLinecardList) Reset()         { *m = GFpcEnvironmentLinecardList{} }
func (m *GFpcEnvironmentLinecardList) String() string { return proto.CompactTextString(m) }
func (*GFpcEnvironmentLinecardList) ProtoMessage()    {}
func (*GFpcEnvironmentLinecardList) Descriptor() ([]byte, []int) {
	return fileDescriptor_80caabd2d4d5ba4d, []int{3, 0}
}
func (m *GFpcEnvironmentLinecardList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GFpcEnvironmentLinecardList.Unmarshal(m, b)
}
func (m *GFpcEnvironmentLinecardList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GFpcEnvironmentLinecardList.Marshal(b, m, deterministic)
}
func (m *GFpcEnvironmentLinecardList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GFpcEnvironmentLinecardList.Merge(m, src)
}
func (m *GFpcEnvironmentLinecardList) XXX_Size() int {
	return xxx_messageInfo_GFpcEnvironmentLinecardList.Size(m)
}
func (m *GFpcEnvironmentLinecardList) XXX_DiscardUnknown() {
	xxx_messageInfo_GFpcEnvironmentLinecardList.DiscardUnknown(m)
}

var xxx_messageInfo_GFpcEnvironmentLinecardList proto.InternalMessageInfo

func (m *GFpcEnvironmentLinecardList) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GFpcEnvironmentLinecardList) GetPowerRecord() *GPowerInfo {
	if m != nil {
		return m.PowerRecord
	}
	return nil
}

func (m *GFpcEnvironmentLinecardList) GetTempRecord() []*GTempInfo {
	if m != nil {
		return m.TempRecord
	}
	return nil
}

func (m *GFpcEnvironmentLinecardList) GetVoltageRecord() []*GVoltageInfo {
	if m != nil {
		return m.VoltageRecord
	}
	return nil
}

var E_FpcEnvironmentExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*GFpcEnvironment)(nil),
	Field:         4,
	Name:          "fpcEnvironmentExt",
	Tag:           "bytes,4,opt,name=fpcEnvironmentExt",
	Filename:      "fpc_environment.proto",
}

func init() {
	proto.RegisterType((*GTempInfo)(nil), "g_temp_info")
	proto.RegisterType((*GVoltageInfo)(nil), "g_voltage_info")
	proto.RegisterType((*GPowerInfo)(nil), "g_power_info")
	proto.RegisterType((*GFpcEnvironment)(nil), "g_fpc_environment")
	proto.RegisterType((*GFpcEnvironmentLinecardList)(nil), "g_fpc_environment.linecard_list")
	proto.RegisterExtension(E_FpcEnvironmentExt)
}

func init() { proto.RegisterFile("fpc_environment.proto", fileDescriptor_80caabd2d4d5ba4d) }

var fileDescriptor_80caabd2d4d5ba4d = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0x26, 0x69, 0x0b, 0xed, 0x49, 0x52, 0xed, 0x5c, 0xc4, 0x78, 0xdd, 0x84, 0xe0, 0x22, 0x08,
	0x46, 0x29, 0xe8, 0x42, 0x5c, 0x88, 0x70, 0x5d, 0xb8, 0xb8, 0xca, 0x08, 0x82, 0xab, 0x21, 0xa4,
	0x27, 0x21, 0x98, 0xcc, 0x0c, 0x93, 0xe9, 0x8f, 0x5b, 0x9f, 0xc8, 0x77, 0xf0, 0xc5, 0x24, 0x27,
	0x89, 0x6d, 0xad, 0xee, 0xef, 0x2e, 0xf3, 0xfd, 0xcc, 0x97, 0xf3, 0x9d, 0x81, 0x07, 0x85, 0xce,
	0x05, 0xca, 0x5d, 0x65, 0x94, 0x6c, 0x50, 0xda, 0x54, 0x1b, 0x65, 0xd5, 0xf5, 0x95, 0xc5, 0x1a,
	0x1b, 0xb4, 0xe6, 0xbb, 0xb0, 0x4a, 0xf7, 0x60, 0xbc, 0x01, 0xaf, 0x14, 0x16, 0x1b, 0x2d, 0x2a,
	0x59, 0x28, 0xf6, 0x1c, 0xee, 0xd3, 0xa1, 0x45, 0xd9, 0x2a, 0x23, 0x64, 0xd6, 0x60, 0xe8, 0x44,
	0x6e, 0xb2, 0x78, 0x37, 0xfb, 0xf1, 0xd6, 0x9d, 0x3b, 0x7c, 0xd9, 0xd1, 0x9f, 0x89, 0xbd, 0xcd,
	0x1a, 0x64, 0x4f, 0x00, 0xc8, 0xb0, 0xcb, 0xea, 0x2d, 0x86, 0x6e, 0xe4, 0x26, 0x33, 0x92, 0x46,
	0x0e, 0x5f, 0x74, 0xc4, 0x97, 0x0e, 0x8f, 0x5b, 0x58, 0x96, 0x62, 0xa7, 0x6a, 0x9b, 0x95, 0xd8,
	0x07, 0xbd, 0x84, 0xab, 0xf1, 0xfc, 0xdf, 0xac, 0xd5, 0xa0, 0x38, 0x89, 0x7b, 0x0a, 0xc1, 0x68,
	0xfb, 0x47, 0xa2, 0x3f, 0x70, 0x7d, 0xe8, 0x47, 0xf0, 0x4b, 0xa1, 0xd5, 0x1e, 0x4d, 0x1f, 0x19,
	0x43, 0xd0, 0x64, 0x07, 0xd1, 0x95, 0x43, 0x68, 0xe8, 0x44, 0x4e, 0x32, 0xe3, 0x5e, 0x93, 0x1d,
	0xde, 0xeb, 0xfc, 0x53, 0x07, 0xb1, 0xc7, 0xb0, 0x38, 0xf2, 0x2e, 0xf1, 0xf3, 0x62, 0x20, 0xe3,
	0x9f, 0x13, 0x58, 0x95, 0xe2, 0xaf, 0x72, 0xd9, 0x0b, 0xf0, 0xfb, 0x10, 0x83, 0xb9, 0x32, 0x1b,
	0xba, 0xd5, 0x5b, 0x07, 0xe9, 0x69, 0x36, 0xf7, 0xe8, 0x9b, 0x93, 0x82, 0x3d, 0x03, 0x8f, 0x3a,
	0x1b, 0x0c, 0x6e, 0x34, 0x49, 0xbc, 0xb5, 0x9f, 0x9e, 0xec, 0x81, 0x53, 0xa9, 0x83, 0xfc, 0x15,
	0x2c, 0xc7, 0x99, 0x07, 0xc7, 0x84, 0x1c, 0xf7, 0xd2, 0xf3, 0x4e, 0xf9, 0x58, 0xcd, 0xe0, 0x7b,
	0x03, 0xf3, 0xba, 0x92, 0x98, 0x67, 0x66, 0x13, 0x4e, 0xc9, 0x11, 0xa5, 0x17, 0xbf, 0x9f, 0x8e,
	0x12, 0x51, 0x57, 0xad, 0xe5, 0x7f, 0x1c, 0xd7, 0xbf, 0x1c, 0x08, 0xce, 0x38, 0xf6, 0x08, 0xa6,
	0xb4, 0xa3, 0x69, 0xe4, 0x1c, 0x77, 0x44, 0xd0, 0x9d, 0xed, 0xe0, 0xf5, 0x57, 0x58, 0x15, 0x3a,
	0xbf, 0x39, 0xce, 0x7b, 0x73, 0xb0, 0xec, 0x61, 0xfa, 0x61, 0x2b, 0x2b, 0x8d, 0xe6, 0x16, 0xed,
	0x5e, 0x99, 0x6f, 0x6d, 0xff, 0xc0, 0x5a, 0x9a, 0xcc, 0x5b, 0xb3, 0xcb, 0x96, 0xf8, 0xe5, 0x2d,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x55, 0x85, 0x1a, 0x3b, 0x66, 0x03, 0x00, 0x00,
}
