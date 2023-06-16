// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: supplier_service.proto

package organization_service

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

// SupplierServiceClient is the client API for SupplierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupplierServiceClient interface {
	Create(ctx context.Context, in *CreateSupplierRequest, opts ...grpc.CallOption) (*Supplier, error)
	Get(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Supplier, error)
	GetList(ctx context.Context, in *GetListSupplierRequest, opts ...grpc.CallOption) (*GetSuppliersListResponse, error)
	Update(ctx context.Context, in *UpdateSupplierRequest, opts ...grpc.CallOption) (*Supplier, error)
	PatchUpdate(ctx context.Context, in *PatchUpdateRequest, opts ...grpc.CallOption) (*Supplier, error)
	Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type supplierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupplierServiceClient(cc grpc.ClientConnInterface) SupplierServiceClient {
	return &supplierServiceClient{cc}
}

func (c *supplierServiceClient) Create(ctx context.Context, in *CreateSupplierRequest, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/organization_service.SupplierService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) Get(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/organization_service.SupplierService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) GetList(ctx context.Context, in *GetListSupplierRequest, opts ...grpc.CallOption) (*GetSuppliersListResponse, error) {
	out := new(GetSuppliersListResponse)
	err := c.cc.Invoke(ctx, "/organization_service.SupplierService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) Update(ctx context.Context, in *UpdateSupplierRequest, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/organization_service.SupplierService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) PatchUpdate(ctx context.Context, in *PatchUpdateRequest, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/organization_service.SupplierService/PatchUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/organization_service.SupplierService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupplierServiceServer is the server API for SupplierService service.
// All implementations must embed UnimplementedSupplierServiceServer
// for forward compatibility
type SupplierServiceServer interface {
	Create(context.Context, *CreateSupplierRequest) (*Supplier, error)
	Get(context.Context, *PrimaryKey) (*Supplier, error)
	GetList(context.Context, *GetListSupplierRequest) (*GetSuppliersListResponse, error)
	Update(context.Context, *UpdateSupplierRequest) (*Supplier, error)
	PatchUpdate(context.Context, *PatchUpdateRequest) (*Supplier, error)
	Delete(context.Context, *PrimaryKey) (*empty.Empty, error)
	mustEmbedUnimplementedSupplierServiceServer()
}

// UnimplementedSupplierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupplierServiceServer struct {
}

func (UnimplementedSupplierServiceServer) Create(context.Context, *CreateSupplierRequest) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSupplierServiceServer) Get(context.Context, *PrimaryKey) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSupplierServiceServer) GetList(context.Context, *GetListSupplierRequest) (*GetSuppliersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSupplierServiceServer) Update(context.Context, *UpdateSupplierRequest) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSupplierServiceServer) PatchUpdate(context.Context, *PatchUpdateRequest) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchUpdate not implemented")
}
func (UnimplementedSupplierServiceServer) Delete(context.Context, *PrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSupplierServiceServer) mustEmbedUnimplementedSupplierServiceServer() {}

// UnsafeSupplierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupplierServiceServer will
// result in compilation errors.
type UnsafeSupplierServiceServer interface {
	mustEmbedUnimplementedSupplierServiceServer()
}

func RegisterSupplierServiceServer(s grpc.ServiceRegistrar, srv SupplierServiceServer) {
	s.RegisterService(&SupplierService_ServiceDesc, srv)
}

func _SupplierService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.SupplierService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Create(ctx, req.(*CreateSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.SupplierService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Get(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.SupplierService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).GetList(ctx, req.(*GetListSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.SupplierService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Update(ctx, req.(*UpdateSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_PatchUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).PatchUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.SupplierService/PatchUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).PatchUpdate(ctx, req.(*PatchUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.SupplierService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Delete(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// SupplierService_ServiceDesc is the grpc.ServiceDesc for SupplierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupplierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organization_service.SupplierService",
	HandlerType: (*SupplierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SupplierService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SupplierService_Get_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SupplierService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SupplierService_Update_Handler,
		},
		{
			MethodName: "PatchUpdate",
			Handler:    _SupplierService_PatchUpdate_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SupplierService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supplier_service.proto",
}