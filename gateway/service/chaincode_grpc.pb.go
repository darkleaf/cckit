// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChaincodeClient is the client API for Chaincode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaincodeClient interface {
	// Exec: Query or Invoke
	Exec(ctx context.Context, in *ChaincodeExec, opts ...grpc.CallOption) (*peer.ProposalResponse, error)
	// Query chaincode on home peer. Do NOT send to orderer.
	Query(ctx context.Context, in *ChaincodeInput, opts ...grpc.CallOption) (*peer.ProposalResponse, error)
	// Invoke chaincode on peers, according to endorsement policy and the SEND to orderer
	Invoke(ctx context.Context, in *ChaincodeInput, opts ...grpc.CallOption) (*peer.ProposalResponse, error)
	// Chaincode events stream
	Events(ctx context.Context, in *ChaincodeLocator, opts ...grpc.CallOption) (Chaincode_EventsClient, error)
}

type chaincodeClient struct {
	cc grpc.ClientConnInterface
}

func NewChaincodeClient(cc grpc.ClientConnInterface) ChaincodeClient {
	return &chaincodeClient{cc}
}

func (c *chaincodeClient) Exec(ctx context.Context, in *ChaincodeExec, opts ...grpc.CallOption) (*peer.ProposalResponse, error) {
	out := new(peer.ProposalResponse)
	err := c.cc.Invoke(ctx, "/s7techlab.gateway.service.Chaincode/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) Query(ctx context.Context, in *ChaincodeInput, opts ...grpc.CallOption) (*peer.ProposalResponse, error) {
	out := new(peer.ProposalResponse)
	err := c.cc.Invoke(ctx, "/s7techlab.gateway.service.Chaincode/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) Invoke(ctx context.Context, in *ChaincodeInput, opts ...grpc.CallOption) (*peer.ProposalResponse, error) {
	out := new(peer.ProposalResponse)
	err := c.cc.Invoke(ctx, "/s7techlab.gateway.service.Chaincode/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) Events(ctx context.Context, in *ChaincodeLocator, opts ...grpc.CallOption) (Chaincode_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chaincode_ServiceDesc.Streams[0], "/s7techlab.gateway.service.Chaincode/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chaincode_EventsClient interface {
	Recv() (*peer.ChaincodeEvent, error)
	grpc.ClientStream
}

type chaincodeEventsClient struct {
	grpc.ClientStream
}

func (x *chaincodeEventsClient) Recv() (*peer.ChaincodeEvent, error) {
	m := new(peer.ChaincodeEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChaincodeServer is the server API for Chaincode service.
// All implementations must embed UnimplementedChaincodeServer
// for forward compatibility
type ChaincodeServer interface {
	// Exec: Query or Invoke
	Exec(context.Context, *ChaincodeExec) (*peer.ProposalResponse, error)
	// Query chaincode on home peer. Do NOT send to orderer.
	Query(context.Context, *ChaincodeInput) (*peer.ProposalResponse, error)
	// Invoke chaincode on peers, according to endorsement policy and the SEND to orderer
	Invoke(context.Context, *ChaincodeInput) (*peer.ProposalResponse, error)
	// Chaincode events stream
	Events(*ChaincodeLocator, Chaincode_EventsServer) error
	mustEmbedUnimplementedChaincodeServer()
}

// UnimplementedChaincodeServer must be embedded to have forward compatible implementations.
type UnimplementedChaincodeServer struct {
}

func (UnimplementedChaincodeServer) Exec(context.Context, *ChaincodeExec) (*peer.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedChaincodeServer) Query(context.Context, *ChaincodeInput) (*peer.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedChaincodeServer) Invoke(context.Context, *ChaincodeInput) (*peer.ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (UnimplementedChaincodeServer) Events(*ChaincodeLocator, Chaincode_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}
func (UnimplementedChaincodeServer) mustEmbedUnimplementedChaincodeServer() {}

// UnsafeChaincodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaincodeServer will
// result in compilation errors.
type UnsafeChaincodeServer interface {
	mustEmbedUnimplementedChaincodeServer()
}

func RegisterChaincodeServer(s grpc.ServiceRegistrar, srv ChaincodeServer) {
	s.RegisterService(&Chaincode_ServiceDesc, srv)
}

func _Chaincode_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeExec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/s7techlab.gateway.service.Chaincode/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).Exec(ctx, req.(*ChaincodeExec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/s7techlab.gateway.service.Chaincode/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).Query(ctx, req.(*ChaincodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/s7techlab.gateway.service.Chaincode/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).Invoke(ctx, req.(*ChaincodeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChaincodeLocator)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChaincodeServer).Events(m, &chaincodeEventsServer{stream})
}

type Chaincode_EventsServer interface {
	Send(*peer.ChaincodeEvent) error
	grpc.ServerStream
}

type chaincodeEventsServer struct {
	grpc.ServerStream
}

func (x *chaincodeEventsServer) Send(m *peer.ChaincodeEvent) error {
	return x.ServerStream.SendMsg(m)
}

// Chaincode_ServiceDesc is the grpc.ServiceDesc for Chaincode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chaincode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "s7techlab.gateway.service.Chaincode",
	HandlerType: (*ChaincodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _Chaincode_Exec_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Chaincode_Query_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _Chaincode_Invoke_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _Chaincode_Events_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/chaincode.proto",
}