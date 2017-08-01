// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alerts.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	alerts.proto
	base.proto
	demo.proto

It has these top-level messages:
	AlertRule
	AlertsListRequest
	AlertsListResponse
	DemoVersionRequest
	DemoVersionResponse
	DemoPingRequest
	DemoPingResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AlertRule struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Disabled bool   `protobuf:"varint,3,opt,name=disabled" json:"disabled,omitempty"`
}

func (m *AlertRule) Reset()                    { *m = AlertRule{} }
func (m *AlertRule) String() string            { return proto.CompactTextString(m) }
func (*AlertRule) ProtoMessage()               {}
func (*AlertRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AlertRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlertRule) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *AlertRule) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type AlertsListRequest struct {
}

func (m *AlertsListRequest) Reset()                    { *m = AlertsListRequest{} }
func (m *AlertsListRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertsListRequest) ProtoMessage()               {}
func (*AlertsListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AlertsListResponse struct {
	Error      Error        `protobuf:"varint,1,opt,name=error,enum=api.Error" json:"error,omitempty"`
	AlertRules []*AlertRule `protobuf:"bytes,2,rep,name=alert_rules,json=alertRules" json:"alert_rules,omitempty"`
}

func (m *AlertsListResponse) Reset()                    { *m = AlertsListResponse{} }
func (m *AlertsListResponse) String() string            { return proto.CompactTextString(m) }
func (*AlertsListResponse) ProtoMessage()               {}
func (*AlertsListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AlertsListResponse) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_NO_ERROR
}

func (m *AlertsListResponse) GetAlertRules() []*AlertRule {
	if m != nil {
		return m.AlertRules
	}
	return nil
}

func init() {
	proto.RegisterType((*AlertRule)(nil), "api.AlertRule")
	proto.RegisterType((*AlertsListRequest)(nil), "api.AlertsListRequest")
	proto.RegisterType((*AlertsListResponse)(nil), "api.AlertsListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Alerts service

type AlertsClient interface {
	List(ctx context.Context, in *AlertsListRequest, opts ...grpc.CallOption) (*AlertsListResponse, error)
}

type alertsClient struct {
	cc *grpc.ClientConn
}

func NewAlertsClient(cc *grpc.ClientConn) AlertsClient {
	return &alertsClient{cc}
}

func (c *alertsClient) List(ctx context.Context, in *AlertsListRequest, opts ...grpc.CallOption) (*AlertsListResponse, error) {
	out := new(AlertsListResponse)
	err := grpc.Invoke(ctx, "/api.Alerts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Alerts service

type AlertsServer interface {
	List(context.Context, *AlertsListRequest) (*AlertsListResponse, error)
}

func RegisterAlertsServer(s *grpc.Server, srv AlertsServer) {
	s.RegisterService(&_Alerts_serviceDesc, srv)
}

func _Alerts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Alerts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).List(ctx, req.(*AlertsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alerts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Alerts",
	HandlerType: (*AlertsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Alerts_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alerts.proto",
}

func init() { proto.RegisterFile("alerts.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0xbb, 0x2e, 0xbb, 0xb3, 0xcb, 0x82, 0x23, 0x68, 0x29, 0x1e, 0x4a, 0x4f, 0x3d,
	0xb5, 0x52, 0x9f, 0xc0, 0x83, 0x27, 0x05, 0x21, 0xe0, 0x59, 0xa6, 0xec, 0x50, 0x02, 0xb5, 0xa9,
	0x99, 0x54, 0x3c, 0xfb, 0x0a, 0x3e, 0x9a, 0xaf, 0xe0, 0x83, 0x48, 0x13, 0x5d, 0x17, 0xbc, 0xfd,
	0xf9, 0xf2, 0x27, 0xff, 0x3f, 0x03, 0x5b, 0xea, 0xd9, 0x3a, 0xa9, 0x46, 0x6b, 0x9c, 0xc1, 0x84,
	0x46, 0x9d, 0x5d, 0x76, 0xc6, 0x74, 0x3d, 0xd7, 0x34, 0xea, 0x9a, 0x86, 0xc1, 0x38, 0x72, 0xda,
	0x0c, 0x3f, 0x96, 0x0c, 0x5a, 0x12, 0x0e, 0xba, 0x78, 0x80, 0xf5, 0xcd, 0xfc, 0x5c, 0x4d, 0x3d,
	0x23, 0xc2, 0x62, 0xa0, 0x67, 0x4e, 0xa3, 0x3c, 0x2a, 0xd7, 0xca, 0xeb, 0x99, 0x39, 0x7e, 0x73,
	0x69, 0x1c, 0xd8, 0xac, 0x31, 0x83, 0xd5, 0x5e, 0x0b, 0xb5, 0x3d, 0xef, 0xd3, 0x24, 0x8f, 0xca,
	0x95, 0x3a, 0x9c, 0x8b, 0x33, 0x38, 0xf5, 0x1f, 0xca, 0xbd, 0x16, 0xa7, 0xf8, 0x65, 0x62, 0x71,
	0x45, 0x07, 0x78, 0x0c, 0x65, 0x34, 0x83, 0x30, 0xe6, 0x70, 0xc2, 0xd6, 0x1a, 0xeb, 0xf3, 0x76,
	0x0d, 0x54, 0x34, 0xea, 0xea, 0x76, 0x26, 0x2a, 0x5c, 0x60, 0x0d, 0x1b, 0x3f, 0xdc, 0x93, 0x9d,
	0x7a, 0x96, 0x34, 0xce, 0x93, 0x72, 0xd3, 0xec, 0xbc, 0xef, 0xd0, 0x5a, 0x01, 0xfd, 0x4a, 0x69,
	0x1e, 0x61, 0x19, 0x82, 0xf0, 0x0e, 0x16, 0x73, 0x18, 0x9e, 0xff, 0xb9, 0x8f, 0x2b, 0x65, 0x17,
	0xff, 0x78, 0x68, 0x55, 0xe0, 0xfb, 0xe7, 0xd7, 0x47, 0xbc, 0x45, 0xa8, 0x5f, 0xaf, 0xea, 0xb0,
	0xda, 0x76, 0xe9, 0x97, 0x75, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x65, 0x4b, 0x4b, 0x6b,
	0x01, 0x00, 0x00,
}