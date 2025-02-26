// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: packager/v1/packager.proto

package packagerv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProbeService_GeneratePlaylist_FullMethodName = "/packager.v1.ProbeService/GeneratePlaylist"
)

// ProbeServiceClient is the client API for ProbeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProbeServiceClient interface {
	GeneratePlaylist(ctx context.Context, in *GeneratePlaylistRequest, opts ...grpc.CallOption) (*GeneratePlaylistResponse, error)
}

type probeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProbeServiceClient(cc grpc.ClientConnInterface) ProbeServiceClient {
	return &probeServiceClient{cc}
}

func (c *probeServiceClient) GeneratePlaylist(ctx context.Context, in *GeneratePlaylistRequest, opts ...grpc.CallOption) (*GeneratePlaylistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneratePlaylistResponse)
	err := c.cc.Invoke(ctx, ProbeService_GeneratePlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProbeServiceServer is the server API for ProbeService service.
// All implementations should embed UnimplementedProbeServiceServer
// for forward compatibility.
type ProbeServiceServer interface {
	GeneratePlaylist(context.Context, *GeneratePlaylistRequest) (*GeneratePlaylistResponse, error)
}

// UnimplementedProbeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProbeServiceServer struct{}

func (UnimplementedProbeServiceServer) GeneratePlaylist(context.Context, *GeneratePlaylistRequest) (*GeneratePlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePlaylist not implemented")
}
func (UnimplementedProbeServiceServer) testEmbeddedByValue() {}

// UnsafeProbeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProbeServiceServer will
// result in compilation errors.
type UnsafeProbeServiceServer interface {
	mustEmbedUnimplementedProbeServiceServer()
}

func RegisterProbeServiceServer(s grpc.ServiceRegistrar, srv ProbeServiceServer) {
	// If the following call pancis, it indicates UnimplementedProbeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProbeService_ServiceDesc, srv)
}

func _ProbeService_GeneratePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProbeServiceServer).GeneratePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProbeService_GeneratePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProbeServiceServer).GeneratePlaylist(ctx, req.(*GeneratePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProbeService_ServiceDesc is the grpc.ServiceDesc for ProbeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProbeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "packager.v1.ProbeService",
	HandlerType: (*ProbeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratePlaylist",
			Handler:    _ProbeService_GeneratePlaylist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packager/v1/packager.proto",
}
