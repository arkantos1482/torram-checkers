// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: alice/checkers/v1/torram.proto

package checkersv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CheckersTorram_CheckersCreateGm_FullMethodName = "/alice.checkers.v1.CheckersTorram/CheckersCreateGm"
)

// CheckersTorramClient is the client API for CheckersTorram service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckersTorramClient interface {
	CheckersCreateGm(ctx context.Context, in *ReqCheckersTorram, opts ...grpc.CallOption) (*ResCheckersTorram, error)
}

type checkersTorramClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckersTorramClient(cc grpc.ClientConnInterface) CheckersTorramClient {
	return &checkersTorramClient{cc}
}

func (c *checkersTorramClient) CheckersCreateGm(ctx context.Context, in *ReqCheckersTorram, opts ...grpc.CallOption) (*ResCheckersTorram, error) {
	out := new(ResCheckersTorram)
	err := c.cc.Invoke(ctx, CheckersTorram_CheckersCreateGm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckersTorramServer is the server API for CheckersTorram service.
// All implementations must embed UnimplementedCheckersTorramServer
// for forward compatibility
type CheckersTorramServer interface {
	CheckersCreateGm(context.Context, *ReqCheckersTorram) (*ResCheckersTorram, error)
	mustEmbedUnimplementedCheckersTorramServer()
}

// UnimplementedCheckersTorramServer must be embedded to have forward compatible implementations.
type UnimplementedCheckersTorramServer struct {
}

func (UnimplementedCheckersTorramServer) CheckersCreateGm(context.Context, *ReqCheckersTorram) (*ResCheckersTorram, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckersCreateGm not implemented")
}
func (UnimplementedCheckersTorramServer) mustEmbedUnimplementedCheckersTorramServer() {}

// UnsafeCheckersTorramServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckersTorramServer will
// result in compilation errors.
type UnsafeCheckersTorramServer interface {
	mustEmbedUnimplementedCheckersTorramServer()
}

func RegisterCheckersTorramServer(s grpc.ServiceRegistrar, srv CheckersTorramServer) {
	s.RegisterService(&CheckersTorram_ServiceDesc, srv)
}

func _CheckersTorram_CheckersCreateGm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCheckersTorram)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckersTorramServer).CheckersCreateGm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckersTorram_CheckersCreateGm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckersTorramServer).CheckersCreateGm(ctx, req.(*ReqCheckersTorram))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckersTorram_ServiceDesc is the grpc.ServiceDesc for CheckersTorram service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckersTorram_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alice.checkers.v1.CheckersTorram",
	HandlerType: (*CheckersTorramServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckersCreateGm",
			Handler:    _CheckersTorram_CheckersCreateGm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alice/checkers/v1/torram.proto",
}