// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: kubeappsapis/plugins/fluxv2/packages/v1alpha1/fluxv2.proto

package v1alpha1

import (
	context "context"
	v1alpha1 "github.com/kubeapps/kubeapps/cmd/kubeapps-apis/gen/core/packages/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FluxV2PackagesServiceClient is the client API for FluxV2PackagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FluxV2PackagesServiceClient interface {
	// GetAvailablePackageSummaries returns the available packages managed by the 'fluxv2' plugin
	GetAvailablePackageSummaries(ctx context.Context, in *v1alpha1.GetAvailablePackageSummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageSummariesResponse, error)
	// GetAvailablePackageDetail returns the package metadata managed by the 'fluxv2' plugin
	GetAvailablePackageDetail(ctx context.Context, in *v1alpha1.GetAvailablePackageDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageDetailResponse, error)
	// GetAvailablePackageVersions returns the package versions managed by the 'fluxv2' plugin
	GetAvailablePackageVersions(ctx context.Context, in *v1alpha1.GetAvailablePackageVersionsRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageVersionsResponse, error)
	// GetInstalledPackageSummaries returns the installed packages managed by the 'fluxv2' plugin
	GetInstalledPackageSummaries(ctx context.Context, in *v1alpha1.GetInstalledPackageSummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageSummariesResponse, error)
	// GetInstalledPackageDetail returns the requested installed package managed by the 'fluxv2' plugin
	GetInstalledPackageDetail(ctx context.Context, in *v1alpha1.GetInstalledPackageDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageDetailResponse, error)
	// CreateInstalledPackage creates an installed package based on the request.
	CreateInstalledPackage(ctx context.Context, in *v1alpha1.CreateInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.CreateInstalledPackageResponse, error)
	// UpdateInstalledPackage updates an installed package based on the request.
	UpdateInstalledPackage(ctx context.Context, in *v1alpha1.UpdateInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.UpdateInstalledPackageResponse, error)
	// DeleteInstalledPackage deletes an installed package based on the request.
	DeleteInstalledPackage(ctx context.Context, in *v1alpha1.DeleteInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.DeleteInstalledPackageResponse, error)
	// GetInstalledPackageResourceRefs returns the references for the Kubernetes
	// resources created by an installed package.
	GetInstalledPackageResourceRefs(ctx context.Context, in *v1alpha1.GetInstalledPackageResourceRefsRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageResourceRefsResponse, error)
}

type fluxV2PackagesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFluxV2PackagesServiceClient(cc grpc.ClientConnInterface) FluxV2PackagesServiceClient {
	return &fluxV2PackagesServiceClient{cc}
}

func (c *fluxV2PackagesServiceClient) GetAvailablePackageSummaries(ctx context.Context, in *v1alpha1.GetAvailablePackageSummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageSummariesResponse, error) {
	out := new(v1alpha1.GetAvailablePackageSummariesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetAvailablePackageSummaries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) GetAvailablePackageDetail(ctx context.Context, in *v1alpha1.GetAvailablePackageDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageDetailResponse, error) {
	out := new(v1alpha1.GetAvailablePackageDetailResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetAvailablePackageDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) GetAvailablePackageVersions(ctx context.Context, in *v1alpha1.GetAvailablePackageVersionsRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackageVersionsResponse, error) {
	out := new(v1alpha1.GetAvailablePackageVersionsResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetAvailablePackageVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) GetInstalledPackageSummaries(ctx context.Context, in *v1alpha1.GetInstalledPackageSummariesRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageSummariesResponse, error) {
	out := new(v1alpha1.GetInstalledPackageSummariesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetInstalledPackageSummaries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) GetInstalledPackageDetail(ctx context.Context, in *v1alpha1.GetInstalledPackageDetailRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageDetailResponse, error) {
	out := new(v1alpha1.GetInstalledPackageDetailResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetInstalledPackageDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) CreateInstalledPackage(ctx context.Context, in *v1alpha1.CreateInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.CreateInstalledPackageResponse, error) {
	out := new(v1alpha1.CreateInstalledPackageResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/CreateInstalledPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) UpdateInstalledPackage(ctx context.Context, in *v1alpha1.UpdateInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.UpdateInstalledPackageResponse, error) {
	out := new(v1alpha1.UpdateInstalledPackageResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/UpdateInstalledPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) DeleteInstalledPackage(ctx context.Context, in *v1alpha1.DeleteInstalledPackageRequest, opts ...grpc.CallOption) (*v1alpha1.DeleteInstalledPackageResponse, error) {
	out := new(v1alpha1.DeleteInstalledPackageResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/DeleteInstalledPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) GetInstalledPackageResourceRefs(ctx context.Context, in *v1alpha1.GetInstalledPackageResourceRefsRequest, opts ...grpc.CallOption) (*v1alpha1.GetInstalledPackageResourceRefsResponse, error) {
	out := new(v1alpha1.GetInstalledPackageResourceRefsResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetInstalledPackageResourceRefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FluxV2PackagesServiceServer is the server API for FluxV2PackagesService service.
// All implementations should embed UnimplementedFluxV2PackagesServiceServer
// for forward compatibility
type FluxV2PackagesServiceServer interface {
	// GetAvailablePackageSummaries returns the available packages managed by the 'fluxv2' plugin
	GetAvailablePackageSummaries(context.Context, *v1alpha1.GetAvailablePackageSummariesRequest) (*v1alpha1.GetAvailablePackageSummariesResponse, error)
	// GetAvailablePackageDetail returns the package metadata managed by the 'fluxv2' plugin
	GetAvailablePackageDetail(context.Context, *v1alpha1.GetAvailablePackageDetailRequest) (*v1alpha1.GetAvailablePackageDetailResponse, error)
	// GetAvailablePackageVersions returns the package versions managed by the 'fluxv2' plugin
	GetAvailablePackageVersions(context.Context, *v1alpha1.GetAvailablePackageVersionsRequest) (*v1alpha1.GetAvailablePackageVersionsResponse, error)
	// GetInstalledPackageSummaries returns the installed packages managed by the 'fluxv2' plugin
	GetInstalledPackageSummaries(context.Context, *v1alpha1.GetInstalledPackageSummariesRequest) (*v1alpha1.GetInstalledPackageSummariesResponse, error)
	// GetInstalledPackageDetail returns the requested installed package managed by the 'fluxv2' plugin
	GetInstalledPackageDetail(context.Context, *v1alpha1.GetInstalledPackageDetailRequest) (*v1alpha1.GetInstalledPackageDetailResponse, error)
	// CreateInstalledPackage creates an installed package based on the request.
	CreateInstalledPackage(context.Context, *v1alpha1.CreateInstalledPackageRequest) (*v1alpha1.CreateInstalledPackageResponse, error)
	// UpdateInstalledPackage updates an installed package based on the request.
	UpdateInstalledPackage(context.Context, *v1alpha1.UpdateInstalledPackageRequest) (*v1alpha1.UpdateInstalledPackageResponse, error)
	// DeleteInstalledPackage deletes an installed package based on the request.
	DeleteInstalledPackage(context.Context, *v1alpha1.DeleteInstalledPackageRequest) (*v1alpha1.DeleteInstalledPackageResponse, error)
	// GetInstalledPackageResourceRefs returns the references for the Kubernetes
	// resources created by an installed package.
	GetInstalledPackageResourceRefs(context.Context, *v1alpha1.GetInstalledPackageResourceRefsRequest) (*v1alpha1.GetInstalledPackageResourceRefsResponse, error)
}

// UnimplementedFluxV2PackagesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFluxV2PackagesServiceServer struct {
}

func (UnimplementedFluxV2PackagesServiceServer) GetAvailablePackageSummaries(context.Context, *v1alpha1.GetAvailablePackageSummariesRequest) (*v1alpha1.GetAvailablePackageSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageSummaries not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) GetAvailablePackageDetail(context.Context, *v1alpha1.GetAvailablePackageDetailRequest) (*v1alpha1.GetAvailablePackageDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageDetail not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) GetAvailablePackageVersions(context.Context, *v1alpha1.GetAvailablePackageVersionsRequest) (*v1alpha1.GetAvailablePackageVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageVersions not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) GetInstalledPackageSummaries(context.Context, *v1alpha1.GetInstalledPackageSummariesRequest) (*v1alpha1.GetInstalledPackageSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageSummaries not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) GetInstalledPackageDetail(context.Context, *v1alpha1.GetInstalledPackageDetailRequest) (*v1alpha1.GetInstalledPackageDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageDetail not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) CreateInstalledPackage(context.Context, *v1alpha1.CreateInstalledPackageRequest) (*v1alpha1.CreateInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstalledPackage not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) UpdateInstalledPackage(context.Context, *v1alpha1.UpdateInstalledPackageRequest) (*v1alpha1.UpdateInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstalledPackage not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) DeleteInstalledPackage(context.Context, *v1alpha1.DeleteInstalledPackageRequest) (*v1alpha1.DeleteInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstalledPackage not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) GetInstalledPackageResourceRefs(context.Context, *v1alpha1.GetInstalledPackageResourceRefsRequest) (*v1alpha1.GetInstalledPackageResourceRefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageResourceRefs not implemented")
}

// UnsafeFluxV2PackagesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FluxV2PackagesServiceServer will
// result in compilation errors.
type UnsafeFluxV2PackagesServiceServer interface {
	mustEmbedUnimplementedFluxV2PackagesServiceServer()
}

func RegisterFluxV2PackagesServiceServer(s grpc.ServiceRegistrar, srv FluxV2PackagesServiceServer) {
	s.RegisterService(&FluxV2PackagesService_ServiceDesc, srv)
}

func _FluxV2PackagesService_GetAvailablePackageSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetAvailablePackageSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetAvailablePackageSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetAvailablePackageSummaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetAvailablePackageSummaries(ctx, req.(*v1alpha1.GetAvailablePackageSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_GetAvailablePackageDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetAvailablePackageDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetAvailablePackageDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetAvailablePackageDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetAvailablePackageDetail(ctx, req.(*v1alpha1.GetAvailablePackageDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_GetAvailablePackageVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetAvailablePackageVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetAvailablePackageVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetAvailablePackageVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetAvailablePackageVersions(ctx, req.(*v1alpha1.GetAvailablePackageVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_GetInstalledPackageSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetInstalledPackageSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetInstalledPackageSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetInstalledPackageSummaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetInstalledPackageSummaries(ctx, req.(*v1alpha1.GetInstalledPackageSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_GetInstalledPackageDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetInstalledPackageDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetInstalledPackageDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetInstalledPackageDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetInstalledPackageDetail(ctx, req.(*v1alpha1.GetInstalledPackageDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_CreateInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.CreateInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).CreateInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/CreateInstalledPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).CreateInstalledPackage(ctx, req.(*v1alpha1.CreateInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_UpdateInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.UpdateInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).UpdateInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/UpdateInstalledPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).UpdateInstalledPackage(ctx, req.(*v1alpha1.UpdateInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_DeleteInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.DeleteInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).DeleteInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/DeleteInstalledPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).DeleteInstalledPackage(ctx, req.(*v1alpha1.DeleteInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_GetInstalledPackageResourceRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetInstalledPackageResourceRefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetInstalledPackageResourceRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetInstalledPackageResourceRefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetInstalledPackageResourceRefs(ctx, req.(*v1alpha1.GetInstalledPackageResourceRefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FluxV2PackagesService_ServiceDesc is the grpc.ServiceDesc for FluxV2PackagesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FluxV2PackagesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService",
	HandlerType: (*FluxV2PackagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailablePackageSummaries",
			Handler:    _FluxV2PackagesService_GetAvailablePackageSummaries_Handler,
		},
		{
			MethodName: "GetAvailablePackageDetail",
			Handler:    _FluxV2PackagesService_GetAvailablePackageDetail_Handler,
		},
		{
			MethodName: "GetAvailablePackageVersions",
			Handler:    _FluxV2PackagesService_GetAvailablePackageVersions_Handler,
		},
		{
			MethodName: "GetInstalledPackageSummaries",
			Handler:    _FluxV2PackagesService_GetInstalledPackageSummaries_Handler,
		},
		{
			MethodName: "GetInstalledPackageDetail",
			Handler:    _FluxV2PackagesService_GetInstalledPackageDetail_Handler,
		},
		{
			MethodName: "CreateInstalledPackage",
			Handler:    _FluxV2PackagesService_CreateInstalledPackage_Handler,
		},
		{
			MethodName: "UpdateInstalledPackage",
			Handler:    _FluxV2PackagesService_UpdateInstalledPackage_Handler,
		},
		{
			MethodName: "DeleteInstalledPackage",
			Handler:    _FluxV2PackagesService_DeleteInstalledPackage_Handler,
		},
		{
			MethodName: "GetInstalledPackageResourceRefs",
			Handler:    _FluxV2PackagesService_GetInstalledPackageResourceRefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeappsapis/plugins/fluxv2/packages/v1alpha1/fluxv2.proto",
}

// FluxV2RepositoriesServiceClient is the client API for FluxV2RepositoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FluxV2RepositoriesServiceClient interface {
	// AddPackageRepository add an existing package repository to the set of ones already managed by the
	// 'fluxv2' plugin
	AddPackageRepository(ctx context.Context, in *v1alpha1.AddPackageRepositoryRequest, opts ...grpc.CallOption) (*v1alpha1.AddPackageRepositoryResponse, error)
}

type fluxV2RepositoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFluxV2RepositoriesServiceClient(cc grpc.ClientConnInterface) FluxV2RepositoriesServiceClient {
	return &fluxV2RepositoriesServiceClient{cc}
}

func (c *fluxV2RepositoriesServiceClient) AddPackageRepository(ctx context.Context, in *v1alpha1.AddPackageRepositoryRequest, opts ...grpc.CallOption) (*v1alpha1.AddPackageRepositoryResponse, error) {
	out := new(v1alpha1.AddPackageRepositoryResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2RepositoriesService/AddPackageRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FluxV2RepositoriesServiceServer is the server API for FluxV2RepositoriesService service.
// All implementations should embed UnimplementedFluxV2RepositoriesServiceServer
// for forward compatibility
type FluxV2RepositoriesServiceServer interface {
	// AddPackageRepository add an existing package repository to the set of ones already managed by the
	// 'fluxv2' plugin
	AddPackageRepository(context.Context, *v1alpha1.AddPackageRepositoryRequest) (*v1alpha1.AddPackageRepositoryResponse, error)
}

// UnimplementedFluxV2RepositoriesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFluxV2RepositoriesServiceServer struct {
}

func (UnimplementedFluxV2RepositoriesServiceServer) AddPackageRepository(context.Context, *v1alpha1.AddPackageRepositoryRequest) (*v1alpha1.AddPackageRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPackageRepository not implemented")
}

// UnsafeFluxV2RepositoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FluxV2RepositoriesServiceServer will
// result in compilation errors.
type UnsafeFluxV2RepositoriesServiceServer interface {
	mustEmbedUnimplementedFluxV2RepositoriesServiceServer()
}

func RegisterFluxV2RepositoriesServiceServer(s grpc.ServiceRegistrar, srv FluxV2RepositoriesServiceServer) {
	s.RegisterService(&FluxV2RepositoriesService_ServiceDesc, srv)
}

func _FluxV2RepositoriesService_AddPackageRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.AddPackageRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2RepositoriesServiceServer).AddPackageRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2RepositoriesService/AddPackageRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2RepositoriesServiceServer).AddPackageRepository(ctx, req.(*v1alpha1.AddPackageRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FluxV2RepositoriesService_ServiceDesc is the grpc.ServiceDesc for FluxV2RepositoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FluxV2RepositoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2RepositoriesService",
	HandlerType: (*FluxV2RepositoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPackageRepository",
			Handler:    _FluxV2RepositoriesService_AddPackageRepository_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeappsapis/plugins/fluxv2/packages/v1alpha1/fluxv2.proto",
}
