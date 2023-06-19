// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: advent_product_service.proto

package storage_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdventProductServiceClient is the client API for AdventProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdventProductServiceClient interface {
	Create(ctx context.Context, in *CreateAdventProductRequest, opts ...grpc.CallOption) (*AdventProduct, error)
	Get(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*AdventProduct, error)
	GetList(ctx context.Context, in *GetListAdventProductRequest, opts ...grpc.CallOption) (*GetAdventProductsListResponse, error)
	Update(ctx context.Context, in *UpdateAdventProductRequest, opts ...grpc.CallOption) (*AdventProduct, error)
	PatchUpdate(ctx context.Context, in *PatchUpdateRequest, opts ...grpc.CallOption) (*AdventProduct, error)
	Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type adventProductServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdventProductServiceClient(cc grpc.ClientConnInterface) AdventProductServiceClient {
	return &adventProductServiceClient{cc}
}

func (c *adventProductServiceClient) Create(ctx context.Context, in *CreateAdventProductRequest, opts ...grpc.CallOption) (*AdventProduct, error) {
	out := new(AdventProduct)
	err := c.cc.Invoke(ctx, "/storage_service.AdventProductService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventProductServiceClient) Get(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*AdventProduct, error) {
	out := new(AdventProduct)
	err := c.cc.Invoke(ctx, "/storage_service.AdventProductService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventProductServiceClient) GetList(ctx context.Context, in *GetListAdventProductRequest, opts ...grpc.CallOption) (*GetAdventProductsListResponse, error) {
	out := new(GetAdventProductsListResponse)
	err := c.cc.Invoke(ctx, "/storage_service.AdventProductService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventProductServiceClient) Update(ctx context.Context, in *UpdateAdventProductRequest, opts ...grpc.CallOption) (*AdventProduct, error) {
	out := new(AdventProduct)
	err := c.cc.Invoke(ctx, "/storage_service.AdventProductService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventProductServiceClient) PatchUpdate(ctx context.Context, in *PatchUpdateRequest, opts ...grpc.CallOption) (*AdventProduct, error) {
	out := new(AdventProduct)
	err := c.cc.Invoke(ctx, "/storage_service.AdventProductService/PatchUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adventProductServiceClient) Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/storage_service.AdventProductService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdventProductServiceServer is the server API for AdventProductService service.
// All implementations must embed UnimplementedAdventProductServiceServer
// for forward compatibility
type AdventProductServiceServer interface {
	Create(context.Context, *CreateAdventProductRequest) (*AdventProduct, error)
	Get(context.Context, *PrimaryKey) (*AdventProduct, error)
	GetList(context.Context, *GetListAdventProductRequest) (*GetAdventProductsListResponse, error)
	Update(context.Context, *UpdateAdventProductRequest) (*AdventProduct, error)
	PatchUpdate(context.Context, *PatchUpdateRequest) (*AdventProduct, error)
	Delete(context.Context, *PrimaryKey) (*empty.Empty, error)
	mustEmbedUnimplementedAdventProductServiceServer()
}

// UnimplementedAdventProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdventProductServiceServer struct {
}

func (UnimplementedAdventProductServiceServer) Create(context.Context, *CreateAdventProductRequest) (*AdventProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdventProductServiceServer) Get(context.Context, *PrimaryKey) (*AdventProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAdventProductServiceServer) GetList(context.Context, *GetListAdventProductRequest) (*GetAdventProductsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdventProductServiceServer) Update(context.Context, *UpdateAdventProductRequest) (*AdventProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdventProductServiceServer) PatchUpdate(context.Context, *PatchUpdateRequest) (*AdventProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchUpdate not implemented")
}
func (UnimplementedAdventProductServiceServer) Delete(context.Context, *PrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdventProductServiceServer) mustEmbedUnimplementedAdventProductServiceServer() {}

// UnsafeAdventProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdventProductServiceServer will
// result in compilation errors.
type UnsafeAdventProductServiceServer interface {
	mustEmbedUnimplementedAdventProductServiceServer()
}

func RegisterAdventProductServiceServer(s grpc.ServiceRegistrar, srv AdventProductServiceServer) {
	s.RegisterService(&AdventProductService_ServiceDesc, srv)
}

func _AdventProductService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdventProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventProductServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage_service.AdventProductService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventProductServiceServer).Create(ctx, req.(*CreateAdventProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventProductService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventProductServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage_service.AdventProductService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventProductServiceServer).Get(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventProductService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAdventProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventProductServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage_service.AdventProductService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventProductServiceServer).GetList(ctx, req.(*GetListAdventProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventProductService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdventProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventProductServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage_service.AdventProductService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventProductServiceServer).Update(ctx, req.(*UpdateAdventProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventProductService_PatchUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventProductServiceServer).PatchUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage_service.AdventProductService/PatchUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventProductServiceServer).PatchUpdate(ctx, req.(*PatchUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdventProductService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventProductServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage_service.AdventProductService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventProductServiceServer).Delete(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// AdventProductService_ServiceDesc is the grpc.ServiceDesc for AdventProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdventProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage_service.AdventProductService",
	HandlerType: (*AdventProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AdventProductService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AdventProductService_Get_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _AdventProductService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdventProductService_Update_Handler,
		},
		{
			MethodName: "PatchUpdate",
			Handler:    _AdventProductService_PatchUpdate_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdventProductService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advent_product_service.proto",
}