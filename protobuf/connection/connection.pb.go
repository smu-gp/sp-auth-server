// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection.proto

package connection

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AuthResponse_ResultMessage int32

const (
	AuthResponse_MESSAGE_FAILED  AuthResponse_ResultMessage = 0
	AuthResponse_MESSAGE_SUCCESS AuthResponse_ResultMessage = 1
)

var AuthResponse_ResultMessage_name = map[int32]string{
	0: "MESSAGE_FAILED",
	1: "MESSAGE_SUCCESS",
}

var AuthResponse_ResultMessage_value = map[string]int32{
	"MESSAGE_FAILED":  0,
	"MESSAGE_SUCCESS": 1,
}

func (x AuthResponse_ResultMessage) String() string {
	return proto.EnumName(AuthResponse_ResultMessage_name, int32(x))
}

func (AuthResponse_ResultMessage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{3, 0}
}

type AuthResponse_FailedReason int32

const (
	AuthResponse_NONE             AuthResponse_FailedReason = 0
	AuthResponse_AUTH_FAILED      AuthResponse_FailedReason = 1
	AuthResponse_INTERNAL_ERR     AuthResponse_FailedReason = 2
	AuthResponse_REJECT_HOST      AuthResponse_FailedReason = 3
	AuthResponse_NO_HOST_WAITED   AuthResponse_FailedReason = 4
	AuthResponse_RESPONSE_TIMEOUT AuthResponse_FailedReason = 5
)

var AuthResponse_FailedReason_name = map[int32]string{
	0: "NONE",
	1: "AUTH_FAILED",
	2: "INTERNAL_ERR",
	3: "REJECT_HOST",
	4: "NO_HOST_WAITED",
	5: "RESPONSE_TIMEOUT",
}

var AuthResponse_FailedReason_value = map[string]int32{
	"NONE":             0,
	"AUTH_FAILED":      1,
	"INTERNAL_ERR":     2,
	"REJECT_HOST":      3,
	"NO_HOST_WAITED":   4,
	"RESPONSE_TIMEOUT": 5,
}

func (x AuthResponse_FailedReason) String() string {
	return proto.EnumName(AuthResponse_FailedReason_name, int32(x))
}

func (AuthResponse_FailedReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{3, 1}
}

type AuthDeviceInfo_DeviceType int32

const (
	AuthDeviceInfo_DEVICE_ANDROID AuthDeviceInfo_DeviceType = 0
	AuthDeviceInfo_DEVICE_WEB     AuthDeviceInfo_DeviceType = 1
)

var AuthDeviceInfo_DeviceType_name = map[int32]string{
	0: "DEVICE_ANDROID",
	1: "DEVICE_WEB",
}

var AuthDeviceInfo_DeviceType_value = map[string]int32{
	"DEVICE_ANDROID": 0,
	"DEVICE_WEB":     1,
}

func (x AuthDeviceInfo_DeviceType) String() string {
	return proto.EnumName(AuthDeviceInfo_DeviceType_name, int32(x))
}

func (AuthDeviceInfo_DeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{6, 0}
}

type ConnectionRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionRequest) Reset()         { *m = ConnectionRequest{} }
func (m *ConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectionRequest) ProtoMessage()    {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

func (m *ConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionRequest.Unmarshal(m, b)
}
func (m *ConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionRequest.Marshal(b, m, deterministic)
}
func (m *ConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionRequest.Merge(m, src)
}
func (m *ConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectionRequest.Size(m)
}
func (m *ConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionRequest proto.InternalMessageInfo

func (m *ConnectionRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ConnectionResponse struct {
	ConnectionCode       string   `protobuf:"bytes,1,opt,name=connectionCode,proto3" json:"connectionCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionResponse) Reset()         { *m = ConnectionResponse{} }
func (m *ConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectionResponse) ProtoMessage()    {}
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

func (m *ConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionResponse.Unmarshal(m, b)
}
func (m *ConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionResponse.Marshal(b, m, deterministic)
}
func (m *ConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionResponse.Merge(m, src)
}
func (m *ConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectionResponse.Size(m)
}
func (m *ConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionResponse proto.InternalMessageInfo

func (m *ConnectionResponse) GetConnectionCode() string {
	if m != nil {
		return m.ConnectionCode
	}
	return ""
}

type AuthRequest struct {
	ConnectionCode       string          `protobuf:"bytes,1,opt,name=connectionCode,proto3" json:"connectionCode,omitempty"`
	DeviceInfo           *AuthDeviceInfo `protobuf:"bytes,2,opt,name=deviceInfo,proto3" json:"deviceInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{2}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetConnectionCode() string {
	if m != nil {
		return m.ConnectionCode
	}
	return ""
}

func (m *AuthRequest) GetDeviceInfo() *AuthDeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

type AuthResponse struct {
	Message              AuthResponse_ResultMessage `protobuf:"varint,1,opt,name=message,proto3,enum=connection_grpc.AuthResponse_ResultMessage" json:"message,omitempty"`
	UserId               string                     `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	FailedReason         AuthResponse_FailedReason  `protobuf:"varint,3,opt,name=failedReason,proto3,enum=connection_grpc.AuthResponse_FailedReason" json:"failedReason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{3}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetMessage() AuthResponse_ResultMessage {
	if m != nil {
		return m.Message
	}
	return AuthResponse_MESSAGE_FAILED
}

func (m *AuthResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthResponse) GetFailedReason() AuthResponse_FailedReason {
	if m != nil {
		return m.FailedReason
	}
	return AuthResponse_NONE
}

type WaitAuthRequest struct {
	UserId               string          `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AuthDevice           *AuthDeviceInfo `protobuf:"bytes,2,opt,name=authDevice,proto3" json:"authDevice,omitempty"`
	AcceptDevice         bool            `protobuf:"varint,3,opt,name=acceptDevice,proto3" json:"acceptDevice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WaitAuthRequest) Reset()         { *m = WaitAuthRequest{} }
func (m *WaitAuthRequest) String() string { return proto.CompactTextString(m) }
func (*WaitAuthRequest) ProtoMessage()    {}
func (*WaitAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{4}
}

func (m *WaitAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitAuthRequest.Unmarshal(m, b)
}
func (m *WaitAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitAuthRequest.Marshal(b, m, deterministic)
}
func (m *WaitAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitAuthRequest.Merge(m, src)
}
func (m *WaitAuthRequest) XXX_Size() int {
	return xxx_messageInfo_WaitAuthRequest.Size(m)
}
func (m *WaitAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WaitAuthRequest proto.InternalMessageInfo

func (m *WaitAuthRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *WaitAuthRequest) GetAuthDevice() *AuthDeviceInfo {
	if m != nil {
		return m.AuthDevice
	}
	return nil
}

func (m *WaitAuthRequest) GetAcceptDevice() bool {
	if m != nil {
		return m.AcceptDevice
	}
	return false
}

type WaitAuthResponse struct {
	AuthDevice           *AuthDeviceInfo `protobuf:"bytes,1,opt,name=authDevice,proto3" json:"authDevice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WaitAuthResponse) Reset()         { *m = WaitAuthResponse{} }
func (m *WaitAuthResponse) String() string { return proto.CompactTextString(m) }
func (*WaitAuthResponse) ProtoMessage()    {}
func (*WaitAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{5}
}

func (m *WaitAuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitAuthResponse.Unmarshal(m, b)
}
func (m *WaitAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitAuthResponse.Marshal(b, m, deterministic)
}
func (m *WaitAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitAuthResponse.Merge(m, src)
}
func (m *WaitAuthResponse) XXX_Size() int {
	return xxx_messageInfo_WaitAuthResponse.Size(m)
}
func (m *WaitAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WaitAuthResponse proto.InternalMessageInfo

func (m *WaitAuthResponse) GetAuthDevice() *AuthDeviceInfo {
	if m != nil {
		return m.AuthDevice
	}
	return nil
}

type AuthDeviceInfo struct {
	DeviceType           AuthDeviceInfo_DeviceType `protobuf:"varint,1,opt,name=deviceType,proto3,enum=connection_grpc.AuthDeviceInfo_DeviceType" json:"deviceType,omitempty"`
	DeviceName           string                    `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AuthDeviceInfo) Reset()         { *m = AuthDeviceInfo{} }
func (m *AuthDeviceInfo) String() string { return proto.CompactTextString(m) }
func (*AuthDeviceInfo) ProtoMessage()    {}
func (*AuthDeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{6}
}

func (m *AuthDeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthDeviceInfo.Unmarshal(m, b)
}
func (m *AuthDeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthDeviceInfo.Marshal(b, m, deterministic)
}
func (m *AuthDeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthDeviceInfo.Merge(m, src)
}
func (m *AuthDeviceInfo) XXX_Size() int {
	return xxx_messageInfo_AuthDeviceInfo.Size(m)
}
func (m *AuthDeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthDeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AuthDeviceInfo proto.InternalMessageInfo

func (m *AuthDeviceInfo) GetDeviceType() AuthDeviceInfo_DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return AuthDeviceInfo_DEVICE_ANDROID
}

func (m *AuthDeviceInfo) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{7}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("connection_grpc.AuthResponse_ResultMessage", AuthResponse_ResultMessage_name, AuthResponse_ResultMessage_value)
	proto.RegisterEnum("connection_grpc.AuthResponse_FailedReason", AuthResponse_FailedReason_name, AuthResponse_FailedReason_value)
	proto.RegisterEnum("connection_grpc.AuthDeviceInfo_DeviceType", AuthDeviceInfo_DeviceType_name, AuthDeviceInfo_DeviceType_value)
	proto.RegisterType((*ConnectionRequest)(nil), "connection_grpc.ConnectionRequest")
	proto.RegisterType((*ConnectionResponse)(nil), "connection_grpc.ConnectionResponse")
	proto.RegisterType((*AuthRequest)(nil), "connection_grpc.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "connection_grpc.AuthResponse")
	proto.RegisterType((*WaitAuthRequest)(nil), "connection_grpc.WaitAuthRequest")
	proto.RegisterType((*WaitAuthResponse)(nil), "connection_grpc.WaitAuthResponse")
	proto.RegisterType((*AuthDeviceInfo)(nil), "connection_grpc.AuthDeviceInfo")
	proto.RegisterType((*Empty)(nil), "connection_grpc.Empty")
}

func init() { proto.RegisterFile("connection.proto", fileDescriptor_51baa40a1cc6b48b) }

var fileDescriptor_51baa40a1cc6b48b = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x9d, 0xbb, 0xaf, 0x72, 0x5b, 0x5a, 0x63, 0x10, 0xaa, 0x26, 0x3e, 0x8a, 0x91, 0x50, 0xb5,
	0x49, 0xd5, 0x54, 0x5e, 0x78, 0x40, 0x42, 0x21, 0xf1, 0x58, 0xa6, 0x35, 0x41, 0x76, 0x4a, 0x1f,
	0xa3, 0x90, 0x7a, 0xa3, 0xd2, 0x9a, 0x84, 0x26, 0xad, 0xb4, 0x77, 0x9e, 0xf9, 0x29, 0xfc, 0x35,
	0xfe, 0x02, 0x4a, 0x9b, 0x34, 0x4e, 0xcb, 0xca, 0x78, 0xcc, 0xc9, 0x3d, 0xe7, 0x5c, 0xdf, 0x7b,
	0x6c, 0xc0, 0x7e, 0x18, 0x04, 0xd2, 0x4f, 0xc6, 0x61, 0xd0, 0x8d, 0xa6, 0x61, 0x12, 0x92, 0x66,
	0x81, 0xb8, 0xd7, 0xd3, 0xc8, 0xa7, 0x27, 0xf0, 0x48, 0x5f, 0x41, 0x5c, 0x7e, 0x9f, 0xc9, 0x38,
	0x21, 0x4f, 0xe1, 0x60, 0x16, 0xcb, 0xa9, 0x39, 0x6a, 0xa1, 0x36, 0xea, 0x3c, 0xe0, 0xd9, 0x17,
	0x7d, 0x0f, 0x44, 0x2d, 0x8e, 0xa3, 0x30, 0x88, 0x25, 0x79, 0x03, 0x8d, 0x42, 0x55, 0x0f, 0x47,
	0x32, 0x63, 0xad, 0xa1, 0x74, 0x0e, 0x35, 0x6d, 0x96, 0x7c, 0xcb, 0x4d, 0xee, 0x49, 0x23, 0x1f,
	0x00, 0x46, 0x72, 0x3e, 0xf6, 0xa5, 0x19, 0x5c, 0x85, 0xad, 0x4a, 0x1b, 0x75, 0x6a, 0xbd, 0x97,
	0xdd, 0xb5, 0x73, 0x74, 0x53, 0x65, 0x63, 0x55, 0xc6, 0x15, 0x0a, 0xfd, 0x5d, 0x81, 0xfa, 0xd2,
	0x38, 0x6b, 0x98, 0xc1, 0xe1, 0x44, 0xc6, 0xb1, 0x77, 0xbd, 0xb4, 0x6c, 0xf4, 0x4e, 0xfe, 0x2a,
	0x97, 0xd7, 0x77, 0xb9, 0x8c, 0x67, 0x37, 0x49, 0x7f, 0x49, 0xe1, 0x39, 0x57, 0x99, 0x52, 0x45,
	0x9d, 0x12, 0xb1, 0xa0, 0x7e, 0xe5, 0x8d, 0x6f, 0xe4, 0x88, 0x4b, 0x2f, 0x0e, 0x83, 0xd6, 0xee,
	0xc2, 0xe3, 0x78, 0xbb, 0xc7, 0x99, 0xc2, 0xe0, 0x25, 0x3e, 0x7d, 0x07, 0x0f, 0x4b, 0x1d, 0x10,
	0x02, 0x8d, 0x3e, 0x13, 0x42, 0xfb, 0xc4, 0xdc, 0x33, 0xcd, 0xbc, 0x64, 0x06, 0xde, 0x21, 0x8f,
	0xa1, 0x99, 0x63, 0x62, 0xa0, 0xeb, 0x4c, 0x08, 0x8c, 0xe8, 0x1c, 0xea, 0xaa, 0x2e, 0xa9, 0xc2,
	0x9e, 0x65, 0x5b, 0x0c, 0xef, 0x90, 0x26, 0xd4, 0xb4, 0x81, 0x73, 0x9e, 0xf3, 0x11, 0xc1, 0x50,
	0x37, 0x2d, 0x87, 0x71, 0x4b, 0xbb, 0x74, 0x19, 0xe7, 0xb8, 0x92, 0x96, 0x70, 0x76, 0xc1, 0x74,
	0xc7, 0x3d, 0xb7, 0x85, 0x83, 0x77, 0x53, 0x5b, 0xcb, 0x5e, 0x7c, 0xb8, 0x43, 0xcd, 0x74, 0x98,
	0x81, 0xf7, 0xc8, 0x13, 0xc0, 0x9c, 0x89, 0xcf, 0xb6, 0x25, 0x98, 0xeb, 0x98, 0x7d, 0x66, 0x0f,
	0x1c, 0xbc, 0x4f, 0x7f, 0x22, 0x68, 0x0e, 0xbd, 0x71, 0xa2, 0xae, 0xfb, 0x8e, 0x4c, 0xa5, 0xeb,
	0xf5, 0x56, 0xbb, 0xbb, 0xf7, 0x7a, 0x0b, 0x0a, 0xa1, 0x50, 0xf7, 0x7c, 0x5f, 0x46, 0x49, 0x26,
	0x91, 0x8e, 0xbb, 0xca, 0x4b, 0x18, 0x15, 0x80, 0x8b, 0x7e, 0xb2, 0x14, 0x94, 0x8d, 0xd1, 0x7f,
	0x1b, 0xd3, 0x5f, 0x08, 0x1a, 0xe5, 0xdf, 0xe4, 0x22, 0xcf, 0xaa, 0x73, 0x1b, 0xe5, 0xe1, 0x3a,
	0xfe, 0x87, 0x66, 0xd7, 0x58, 0x31, 0xb8, 0xc2, 0x26, 0x2f, 0x72, 0x2d, 0xcb, 0x9b, 0xc8, 0x2c,
	0x62, 0x0a, 0x42, 0x4f, 0x01, 0x0a, 0x66, 0xba, 0x1c, 0x83, 0x7d, 0x31, 0x75, 0xe6, 0x6a, 0x96,
	0xc1, 0x6d, 0x33, 0xcd, 0x44, 0x03, 0x20, 0xc3, 0x86, 0xec, 0x23, 0x46, 0xf4, 0x10, 0xf6, 0xd9,
	0x24, 0x4a, 0x6e, 0x7b, 0x3f, 0x2a, 0xea, 0xad, 0x17, 0x72, 0xba, 0x18, 0xe4, 0x00, 0xa0, 0x00,
	0x09, 0xdd, 0x68, 0x7b, 0xe3, 0x9d, 0x38, 0x7a, 0xbd, 0xb5, 0x26, 0x9b, 0xb3, 0x0e, 0x7b, 0xe9,
	0x81, 0xc9, 0xb3, 0x3b, 0x2e, 0xc0, 0x52, 0xea, 0xf9, 0xd6, 0xeb, 0x41, 0x04, 0x54, 0xf3, 0x05,
	0x92, 0xf6, 0x46, 0xe9, 0x5a, 0xd6, 0x8e, 0x5e, 0x6d, 0xa9, 0x58, 0x0a, 0x76, 0xd0, 0x29, 0xfa,
	0x7a, 0xb0, 0x78, 0x13, 0xdf, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x37, 0x35, 0xf8, 0x27,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	Connection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	WaitAuth(ctx context.Context, opts ...grpc.CallOption) (ConnectionService_WaitAuthClient, error)
}

type connectionServiceClient struct {
	cc *grpc.ClientConn
}

func NewConnectionServiceClient(cc *grpc.ClientConn) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) Connection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection_grpc.ConnectionService/Connection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/connection_grpc.ConnectionService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) WaitAuth(ctx context.Context, opts ...grpc.CallOption) (ConnectionService_WaitAuthClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConnectionService_serviceDesc.Streams[0], "/connection_grpc.ConnectionService/WaitAuth", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionServiceWaitAuthClient{stream}
	return x, nil
}

type ConnectionService_WaitAuthClient interface {
	Send(*WaitAuthRequest) error
	Recv() (*WaitAuthResponse, error)
	grpc.ClientStream
}

type connectionServiceWaitAuthClient struct {
	grpc.ClientStream
}

func (x *connectionServiceWaitAuthClient) Send(m *WaitAuthRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionServiceWaitAuthClient) Recv() (*WaitAuthResponse, error) {
	m := new(WaitAuthResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
type ConnectionServiceServer interface {
	Connection(context.Context, *ConnectionRequest) (*ConnectionResponse, error)
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	WaitAuth(ConnectionService_WaitAuthServer) error
}

func RegisterConnectionServiceServer(s *grpc.Server, srv ConnectionServiceServer) {
	s.RegisterService(&_ConnectionService_serviceDesc, srv)
}

func _ConnectionService_Connection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).Connection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_grpc.ConnectionService/Connection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).Connection(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_grpc.ConnectionService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_WaitAuth_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionServiceServer).WaitAuth(&connectionServiceWaitAuthServer{stream})
}

type ConnectionService_WaitAuthServer interface {
	Send(*WaitAuthResponse) error
	Recv() (*WaitAuthRequest, error)
	grpc.ServerStream
}

type connectionServiceWaitAuthServer struct {
	grpc.ServerStream
}

func (x *connectionServiceWaitAuthServer) Send(m *WaitAuthResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionServiceWaitAuthServer) Recv() (*WaitAuthRequest, error) {
	m := new(WaitAuthRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ConnectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connection_grpc.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connection",
			Handler:    _ConnectionService_Connection_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _ConnectionService_Auth_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WaitAuth",
			Handler:       _ConnectionService_WaitAuth_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "connection.proto",
}
