// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.2.0
// source: bangunruang.proto

package proto

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
	BangunRuang_HitungVolumeKubus_FullMethodName = "/proto.BangunRuang/HitungVolumeKubus"
)

// BangunRuangClient is the client API for BangunRuang service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BangunRuangClient interface {
	HitungVolumeKubus(ctx context.Context, in *RequestKubus, opts ...grpc.CallOption) (*ResponseVolumeKubus, error)
}

type bangunRuangClient struct {
	cc grpc.ClientConnInterface
}

func NewBangunRuangClient(cc grpc.ClientConnInterface) BangunRuangClient {
	return &bangunRuangClient{cc}
}

func (c *bangunRuangClient) HitungVolumeKubus(ctx context.Context, in *RequestKubus, opts ...grpc.CallOption) (*ResponseVolumeKubus, error) {
	out := new(ResponseVolumeKubus)
	err := c.cc.Invoke(ctx, BangunRuang_HitungVolumeKubus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BangunRuangServer is the server API for BangunRuang service.
// All implementations should embed UnimplementedBangunRuangServer
// for forward compatibility
type BangunRuangServer interface {
	HitungVolumeKubus(context.Context, *RequestKubus) (*ResponseVolumeKubus, error)
}

// UnimplementedBangunRuangServer should be embedded to have forward compatible implementations.
type UnimplementedBangunRuangServer struct {
}

func (UnimplementedBangunRuangServer) HitungVolumeKubus(context.Context, *RequestKubus) (*ResponseVolumeKubus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HitungVolumeKubus not implemented")
}

// UnsafeBangunRuangServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BangunRuangServer will
// result in compilation errors.
type UnsafeBangunRuangServer interface {
	mustEmbedUnimplementedBangunRuangServer()
}

func RegisterBangunRuangServer(s grpc.ServiceRegistrar, srv BangunRuangServer) {
	s.RegisterService(&BangunRuang_ServiceDesc, srv)
}

func _BangunRuang_HitungVolumeKubus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKubus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BangunRuangServer).HitungVolumeKubus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BangunRuang_HitungVolumeKubus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BangunRuangServer).HitungVolumeKubus(ctx, req.(*RequestKubus))
	}
	return interceptor(ctx, in, info, handler)
}

// BangunRuang_ServiceDesc is the grpc.ServiceDesc for BangunRuang service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BangunRuang_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BangunRuang",
	HandlerType: (*BangunRuangServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HitungVolumeKubus",
			Handler:    _BangunRuang_HitungVolumeKubus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bangunruang.proto",
}
