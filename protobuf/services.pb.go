// Code generated by protoc-gen-go.
// source: services.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Protocol service

type ProtocolClient interface {
	Run(ctx context.Context, opts ...grpc.CallOption) (Protocol_RunClient, error)
}

type protocolClient struct {
	cc *grpc.ClientConn
}

func NewProtocolClient(cc *grpc.ClientConn) ProtocolClient {
	return &protocolClient{cc}
}

func (c *protocolClient) Run(ctx context.Context, opts ...grpc.CallOption) (Protocol_RunClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Protocol_serviceDesc.Streams[0], c.cc, "/protobuf.Protocol/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &protocolRunClient{stream}
	return x, nil
}

type Protocol_RunClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type protocolRunClient struct {
	grpc.ClientStream
}

func (x *protocolRunClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *protocolRunClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Protocol service

type ProtocolServer interface {
	Run(Protocol_RunServer) error
}

func RegisterProtocolServer(s *grpc.Server, srv ProtocolServer) {
	s.RegisterService(&_Protocol_serviceDesc, srv)
}

func _Protocol_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProtocolServer).Run(&protocolRunServer{stream})
}

type Protocol_RunServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type protocolRunServer struct {
	grpc.ServerStream
}

func (x *protocolRunServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *protocolRunServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Protocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Protocol",
	HandlerType: (*ProtocolServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _Protocol_Run_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

// Client API for PseudonymSystemCA service

type PseudonymSystemCAClient interface {
	GenerateCertificate(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystemCA_GenerateCertificateClient, error)
	GenerateCertificate_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystemCA_GenerateCertificate_ECClient, error)
}

type pseudonymSystemCAClient struct {
	cc *grpc.ClientConn
}

func NewPseudonymSystemCAClient(cc *grpc.ClientConn) PseudonymSystemCAClient {
	return &pseudonymSystemCAClient{cc}
}

func (c *pseudonymSystemCAClient) GenerateCertificate(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystemCA_GenerateCertificateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystemCA_serviceDesc.Streams[0], c.cc, "/protobuf.PseudonymSystemCA/GenerateCertificate", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemCAGenerateCertificateClient{stream}
	return x, nil
}

type PseudonymSystemCA_GenerateCertificateClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemCAGenerateCertificateClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemCAGenerateCertificateClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemCAGenerateCertificateClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemCAClient) GenerateCertificate_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystemCA_GenerateCertificate_ECClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystemCA_serviceDesc.Streams[1], c.cc, "/protobuf.PseudonymSystemCA/GenerateCertificate_EC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemCAGenerateCertificate_ECClient{stream}
	return x, nil
}

type PseudonymSystemCA_GenerateCertificate_ECClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemCAGenerateCertificate_ECClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemCAGenerateCertificate_ECClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemCAGenerateCertificate_ECClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PseudonymSystemCA service

type PseudonymSystemCAServer interface {
	GenerateCertificate(PseudonymSystemCA_GenerateCertificateServer) error
	GenerateCertificate_EC(PseudonymSystemCA_GenerateCertificate_ECServer) error
}

func RegisterPseudonymSystemCAServer(s *grpc.Server, srv PseudonymSystemCAServer) {
	s.RegisterService(&_PseudonymSystemCA_serviceDesc, srv)
}

func _PseudonymSystemCA_GenerateCertificate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemCAServer).GenerateCertificate(&pseudonymSystemCAGenerateCertificateServer{stream})
}

type PseudonymSystemCA_GenerateCertificateServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemCAGenerateCertificateServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemCAGenerateCertificateServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemCAGenerateCertificateServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystemCA_GenerateCertificate_EC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemCAServer).GenerateCertificate_EC(&pseudonymSystemCAGenerateCertificate_ECServer{stream})
}

type PseudonymSystemCA_GenerateCertificate_ECServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemCAGenerateCertificate_ECServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemCAGenerateCertificate_ECServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemCAGenerateCertificate_ECServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PseudonymSystemCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.PseudonymSystemCA",
	HandlerType: (*PseudonymSystemCAServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateCertificate",
			Handler:       _PseudonymSystemCA_GenerateCertificate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GenerateCertificate_EC",
			Handler:       _PseudonymSystemCA_GenerateCertificate_EC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

// Client API for PseudonymSystem service

type PseudonymSystemClient interface {
	GenerateNym(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_GenerateNymClient, error)
	GenerateNym_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_GenerateNym_ECClient, error)
	ObtainCredential(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_ObtainCredentialClient, error)
	ObtainCredential_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_ObtainCredential_ECClient, error)
	TransferCredential(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_TransferCredentialClient, error)
	TransferCredential_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_TransferCredential_ECClient, error)
}

type pseudonymSystemClient struct {
	cc *grpc.ClientConn
}

func NewPseudonymSystemClient(cc *grpc.ClientConn) PseudonymSystemClient {
	return &pseudonymSystemClient{cc}
}

func (c *pseudonymSystemClient) GenerateNym(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_GenerateNymClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[0], c.cc, "/protobuf.PseudonymSystem/GenerateNym", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemGenerateNymClient{stream}
	return x, nil
}

type PseudonymSystem_GenerateNymClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemGenerateNymClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemGenerateNymClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemGenerateNymClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) GenerateNym_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_GenerateNym_ECClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[1], c.cc, "/protobuf.PseudonymSystem/GenerateNym_EC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemGenerateNym_ECClient{stream}
	return x, nil
}

type PseudonymSystem_GenerateNym_ECClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemGenerateNym_ECClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemGenerateNym_ECClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemGenerateNym_ECClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) ObtainCredential(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_ObtainCredentialClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[2], c.cc, "/protobuf.PseudonymSystem/ObtainCredential", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemObtainCredentialClient{stream}
	return x, nil
}

type PseudonymSystem_ObtainCredentialClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemObtainCredentialClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemObtainCredentialClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemObtainCredentialClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) ObtainCredential_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_ObtainCredential_ECClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[3], c.cc, "/protobuf.PseudonymSystem/ObtainCredential_EC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemObtainCredential_ECClient{stream}
	return x, nil
}

type PseudonymSystem_ObtainCredential_ECClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemObtainCredential_ECClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemObtainCredential_ECClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemObtainCredential_ECClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) TransferCredential(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_TransferCredentialClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[4], c.cc, "/protobuf.PseudonymSystem/TransferCredential", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemTransferCredentialClient{stream}
	return x, nil
}

type PseudonymSystem_TransferCredentialClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemTransferCredentialClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemTransferCredentialClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemTransferCredentialClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pseudonymSystemClient) TransferCredential_EC(ctx context.Context, opts ...grpc.CallOption) (PseudonymSystem_TransferCredential_ECClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PseudonymSystem_serviceDesc.Streams[5], c.cc, "/protobuf.PseudonymSystem/TransferCredential_EC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pseudonymSystemTransferCredential_ECClient{stream}
	return x, nil
}

type PseudonymSystem_TransferCredential_ECClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type pseudonymSystemTransferCredential_ECClient struct {
	grpc.ClientStream
}

func (x *pseudonymSystemTransferCredential_ECClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pseudonymSystemTransferCredential_ECClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PseudonymSystem service

type PseudonymSystemServer interface {
	GenerateNym(PseudonymSystem_GenerateNymServer) error
	GenerateNym_EC(PseudonymSystem_GenerateNym_ECServer) error
	ObtainCredential(PseudonymSystem_ObtainCredentialServer) error
	ObtainCredential_EC(PseudonymSystem_ObtainCredential_ECServer) error
	TransferCredential(PseudonymSystem_TransferCredentialServer) error
	TransferCredential_EC(PseudonymSystem_TransferCredential_ECServer) error
}

func RegisterPseudonymSystemServer(s *grpc.Server, srv PseudonymSystemServer) {
	s.RegisterService(&_PseudonymSystem_serviceDesc, srv)
}

func _PseudonymSystem_GenerateNym_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).GenerateNym(&pseudonymSystemGenerateNymServer{stream})
}

type PseudonymSystem_GenerateNymServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemGenerateNymServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemGenerateNymServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemGenerateNymServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_GenerateNym_EC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).GenerateNym_EC(&pseudonymSystemGenerateNym_ECServer{stream})
}

type PseudonymSystem_GenerateNym_ECServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemGenerateNym_ECServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemGenerateNym_ECServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemGenerateNym_ECServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_ObtainCredential_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).ObtainCredential(&pseudonymSystemObtainCredentialServer{stream})
}

type PseudonymSystem_ObtainCredentialServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemObtainCredentialServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemObtainCredentialServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemObtainCredentialServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_ObtainCredential_EC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).ObtainCredential_EC(&pseudonymSystemObtainCredential_ECServer{stream})
}

type PseudonymSystem_ObtainCredential_ECServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemObtainCredential_ECServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemObtainCredential_ECServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemObtainCredential_ECServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_TransferCredential_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).TransferCredential(&pseudonymSystemTransferCredentialServer{stream})
}

type PseudonymSystem_TransferCredentialServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemTransferCredentialServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemTransferCredentialServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemTransferCredentialServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PseudonymSystem_TransferCredential_EC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PseudonymSystemServer).TransferCredential_EC(&pseudonymSystemTransferCredential_ECServer{stream})
}

type PseudonymSystem_TransferCredential_ECServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pseudonymSystemTransferCredential_ECServer struct {
	grpc.ServerStream
}

func (x *pseudonymSystemTransferCredential_ECServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pseudonymSystemTransferCredential_ECServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PseudonymSystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.PseudonymSystem",
	HandlerType: (*PseudonymSystemServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateNym",
			Handler:       _PseudonymSystem_GenerateNym_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GenerateNym_EC",
			Handler:       _PseudonymSystem_GenerateNym_EC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ObtainCredential",
			Handler:       _PseudonymSystem_ObtainCredential_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ObtainCredential_EC",
			Handler:       _PseudonymSystem_ObtainCredential_EC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TransferCredential",
			Handler:       _PseudonymSystem_TransferCredential_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TransferCredential_EC",
			Handler:       _PseudonymSystem_TransferCredential_EC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

// Client API for Info service

type InfoClient interface {
	GetServiceInfo(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*ServiceInfo, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) GetServiceInfo(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*ServiceInfo, error) {
	out := new(ServiceInfo)
	err := grpc.Invoke(ctx, "/protobuf.Info/GetServiceInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Info service

type InfoServer interface {
	GetServiceInfo(context.Context, *EmptyMsg) (*ServiceInfo, error)
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Info/GetServiceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).GetServiceInfo(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceInfo",
			Handler:    _Info_GetServiceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

func init() { proto.RegisterFile("services.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x52, 0x7c, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x19, 0x23, 0x5b, 0x2e, 0x8e, 0x00,
	0x10, 0x23, 0x39, 0x3f, 0x47, 0xc8, 0x90, 0x8b, 0x39, 0xa8, 0x34, 0x4f, 0x48, 0x50, 0x0f, 0xa6,
	0x5a, 0xcf, 0x17, 0xa2, 0x58, 0x0a, 0x53, 0x48, 0x89, 0x41, 0x83, 0xd1, 0x80, 0xd1, 0x68, 0x0e,
	0x23, 0x97, 0x60, 0x40, 0x71, 0x6a, 0x69, 0x4a, 0x7e, 0x5e, 0x65, 0x6e, 0x70, 0x65, 0x71, 0x49,
	0x6a, 0xae, 0xb3, 0xa3, 0x90, 0x23, 0x97, 0xb0, 0x7b, 0x6a, 0x5e, 0x6a, 0x51, 0x62, 0x49, 0xaa,
	0x73, 0x6a, 0x51, 0x49, 0x66, 0x5a, 0x66, 0x32, 0x90, 0x49, 0x8a, 0xc1, 0x42, 0x2e, 0x5c, 0x62,
	0x58, 0x8c, 0x88, 0x77, 0x75, 0x26, 0xc9, 0x79, 0x53, 0x98, 0xb9, 0xf8, 0xd1, 0x9c, 0x27, 0x64,
	0xc9, 0xc5, 0x0d, 0x33, 0xd9, 0xaf, 0x32, 0x97, 0x24, 0x47, 0xd9, 0x70, 0xf1, 0x21, 0x69, 0x25,
	0xd1, 0x31, 0x42, 0x76, 0x5c, 0x02, 0xfe, 0x49, 0x25, 0x89, 0x99, 0x79, 0xce, 0x45, 0xa9, 0x29,
	0xa9, 0x79, 0x25, 0x99, 0x89, 0x39, 0x24, 0xe9, 0x07, 0x86, 0x2a, 0xba, 0x7e, 0x52, 0x9d, 0xe0,
	0xc0, 0x25, 0x14, 0x52, 0x94, 0x98, 0x57, 0x9c, 0x96, 0x5a, 0x44, 0xa6, 0x23, 0x9c, 0xb9, 0x44,
	0x31, 0x4d, 0x20, 0x35, 0x5a, 0x5c, 0xb9, 0x58, 0x3c, 0xf3, 0xd2, 0xf2, 0x85, 0x6c, 0x41, 0xe1,
	0x59, 0x12, 0x0c, 0x49, 0xab, 0x60, 0x11, 0x21, 0x84, 0x16, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0xdf,
	0xe2, 0x74, 0x29, 0x51, 0x84, 0x18, 0x92, 0x52, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xb8, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xa1, 0xa5, 0x5b, 0x08, 0xee, 0x02, 0x00, 0x00,
}
