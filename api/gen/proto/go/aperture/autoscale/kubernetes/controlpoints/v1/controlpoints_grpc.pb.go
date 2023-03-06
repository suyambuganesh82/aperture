// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package controlpointsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AutoscaleKubernetesControlPointsServiceClient is the client API for AutoscaleKubernetesControlPointsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutoscaleKubernetesControlPointsServiceClient interface {
	GetControlPoints(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AutoscaleKubernetesControlPoints, error)
}

type autoscaleKubernetesControlPointsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoscaleKubernetesControlPointsServiceClient(cc grpc.ClientConnInterface) AutoscaleKubernetesControlPointsServiceClient {
	return &autoscaleKubernetesControlPointsServiceClient{cc}
}

func (c *autoscaleKubernetesControlPointsServiceClient) GetControlPoints(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AutoscaleKubernetesControlPoints, error) {
	out := new(AutoscaleKubernetesControlPoints)
	err := c.cc.Invoke(ctx, "/aperture.autoscale.kubernetes.controlpoints.v1.AutoscaleKubernetesControlPointsService/GetControlPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoscaleKubernetesControlPointsServiceServer is the server API for AutoscaleKubernetesControlPointsService service.
// All implementations should embed UnimplementedAutoscaleKubernetesControlPointsServiceServer
// for forward compatibility
type AutoscaleKubernetesControlPointsServiceServer interface {
	GetControlPoints(context.Context, *emptypb.Empty) (*AutoscaleKubernetesControlPoints, error)
}

// UnimplementedAutoscaleKubernetesControlPointsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAutoscaleKubernetesControlPointsServiceServer struct {
}

func (UnimplementedAutoscaleKubernetesControlPointsServiceServer) GetControlPoints(context.Context, *emptypb.Empty) (*AutoscaleKubernetesControlPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetControlPoints not implemented")
}

// UnsafeAutoscaleKubernetesControlPointsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutoscaleKubernetesControlPointsServiceServer will
// result in compilation errors.
type UnsafeAutoscaleKubernetesControlPointsServiceServer interface {
	mustEmbedUnimplementedAutoscaleKubernetesControlPointsServiceServer()
}

func RegisterAutoscaleKubernetesControlPointsServiceServer(s grpc.ServiceRegistrar, srv AutoscaleKubernetesControlPointsServiceServer) {
	s.RegisterService(&AutoscaleKubernetesControlPointsService_ServiceDesc, srv)
}

func _AutoscaleKubernetesControlPointsService_GetControlPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscaleKubernetesControlPointsServiceServer).GetControlPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aperture.autoscale.kubernetes.controlpoints.v1.AutoscaleKubernetesControlPointsService/GetControlPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscaleKubernetesControlPointsServiceServer).GetControlPoints(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AutoscaleKubernetesControlPointsService_ServiceDesc is the grpc.ServiceDesc for AutoscaleKubernetesControlPointsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutoscaleKubernetesControlPointsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aperture.autoscale.kubernetes.controlpoints.v1.AutoscaleKubernetesControlPointsService",
	HandlerType: (*AutoscaleKubernetesControlPointsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetControlPoints",
			Handler:    _AutoscaleKubernetesControlPointsService_GetControlPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aperture/autoscale/kubernetes/controlpoints/v1/controlpoints.proto",
}
