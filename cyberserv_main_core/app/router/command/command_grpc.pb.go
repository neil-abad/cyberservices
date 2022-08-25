// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package command

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RoutingServiceClient is the client API for RoutingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutingServiceClient interface {
	SubscribeRoutingStats(ctx context.Context, in *SubscribeRoutingStatsRequest, opts ...grpc.CallOption) (RoutingService_SubscribeRoutingStatsClient, error)
	TestRoute(ctx context.Context, in *TestRouteRequest, opts ...grpc.CallOption) (*RoutingContext, error)
}

type routingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutingServiceClient(cc grpc.ClientConnInterface) RoutingServiceClient {
	return &routingServiceClient{cc}
}

func (c *routingServiceClient) SubscribeRoutingStats(ctx context.Context, in *SubscribeRoutingStatsRequest, opts ...grpc.CallOption) (RoutingService_SubscribeRoutingStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoutingService_serviceDesc.Streams[0], "/cyberservices.core.app.router.command.RoutingService/SubscribeRoutingStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingServiceSubscribeRoutingStatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutingService_SubscribeRoutingStatsClient interface {
	Recv() (*RoutingContext, error)
	grpc.ClientStream
}

type routingServiceSubscribeRoutingStatsClient struct {
	grpc.ClientStream
}

func (x *routingServiceSubscribeRoutingStatsClient) Recv() (*RoutingContext, error) {
	m := new(RoutingContext)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routingServiceClient) TestRoute(ctx context.Context, in *TestRouteRequest, opts ...grpc.CallOption) (*RoutingContext, error) {
	out := new(RoutingContext)
	err := c.cc.Invoke(ctx, "/cyberservices.core.app.router.command.RoutingService/TestRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutingServiceServer is the server API for RoutingService service.
// All implementations must embed UnimplementedRoutingServiceServer
// for forward compatibility
type RoutingServiceServer interface {
	SubscribeRoutingStats(*SubscribeRoutingStatsRequest, RoutingService_SubscribeRoutingStatsServer) error
	TestRoute(context.Context, *TestRouteRequest) (*RoutingContext, error)
	mustEmbedUnimplementedRoutingServiceServer()
}

// UnimplementedRoutingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutingServiceServer struct {
}

func (UnimplementedRoutingServiceServer) SubscribeRoutingStats(*SubscribeRoutingStatsRequest, RoutingService_SubscribeRoutingStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeRoutingStats not implemented")
}
func (UnimplementedRoutingServiceServer) TestRoute(context.Context, *TestRouteRequest) (*RoutingContext, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRoute not implemented")
}
func (UnimplementedRoutingServiceServer) mustEmbedUnimplementedRoutingServiceServer() {}

// UnsafeRoutingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutingServiceServer will
// result in compilation errors.
type UnsafeRoutingServiceServer interface {
	mustEmbedUnimplementedRoutingServiceServer()
}

func RegisterRoutingServiceServer(s *grpc.Server, srv RoutingServiceServer) {
	s.RegisterService(&_RoutingService_serviceDesc, srv)
}

func _RoutingService_SubscribeRoutingStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRoutingStatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutingServiceServer).SubscribeRoutingStats(m, &routingServiceSubscribeRoutingStatsServer{stream})
}

type RoutingService_SubscribeRoutingStatsServer interface {
	Send(*RoutingContext) error
	grpc.ServerStream
}

type routingServiceSubscribeRoutingStatsServer struct {
	grpc.ServerStream
}

func (x *routingServiceSubscribeRoutingStatsServer) Send(m *RoutingContext) error {
	return x.ServerStream.SendMsg(m)
}

func _RoutingService_TestRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingServiceServer).TestRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyberservices.core.app.router.command.RoutingService/TestRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingServiceServer).TestRoute(ctx, req.(*TestRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoutingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cyberservices.core.app.router.command.RoutingService",
	HandlerType: (*RoutingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestRoute",
			Handler:    _RoutingService_TestRoute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeRoutingStats",
			Handler:       _RoutingService_SubscribeRoutingStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "app/router/command/command.proto",
}
