// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: advent_product.proto

package storage_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdventProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Photo      string  `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo,omitempty"`
	CategoryId string  `protobuf:"bytes,4,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	BranchId   string  `protobuf:"bytes,5,opt,name=branchId,proto3" json:"branchId,omitempty"`
	Barcode    string  `protobuf:"bytes,6,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Count      int64   `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	Price      float32 `protobuf:"fixed32,8,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AdventProduct) Reset() {
	*x = AdventProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdventProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdventProduct) ProtoMessage() {}

func (x *AdventProduct) ProtoReflect() protoreflect.Message {
	mi := &file_advent_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdventProduct.ProtoReflect.Descriptor instead.
func (*AdventProduct) Descriptor() ([]byte, []int) {
	return file_advent_product_proto_rawDescGZIP(), []int{0}
}

func (x *AdventProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdventProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdventProduct) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *AdventProduct) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *AdventProduct) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *AdventProduct) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *AdventProduct) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AdventProduct) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateAdventProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Photo      string  `protobuf:"bytes,2,opt,name=photo,proto3" json:"photo,omitempty"`
	CategoryId string  `protobuf:"bytes,3,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	BrachId    string  `protobuf:"bytes,4,opt,name=brachId,proto3" json:"brachId,omitempty"`
	Barcode    string  `protobuf:"bytes,5,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Count      int64   `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	Price      float32 `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateAdventProductRequest) Reset() {
	*x = CreateAdventProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdventProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdventProductRequest) ProtoMessage() {}

func (x *CreateAdventProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdventProductRequest.ProtoReflect.Descriptor instead.
func (*CreateAdventProductRequest) Descriptor() ([]byte, []int) {
	return file_advent_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAdventProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAdventProductRequest) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CreateAdventProductRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *CreateAdventProductRequest) GetBrachId() string {
	if x != nil {
		return x.BrachId
	}
	return ""
}

func (x *CreateAdventProductRequest) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *CreateAdventProductRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CreateAdventProductRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetListAdventProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barcode  string `protobuf:"bytes,1,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetListAdventProductRequest) Reset() {
	*x = GetListAdventProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAdventProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAdventProductRequest) ProtoMessage() {}

func (x *GetListAdventProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAdventProductRequest.ProtoReflect.Descriptor instead.
func (*GetListAdventProductRequest) Descriptor() ([]byte, []int) {
	return file_advent_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetListAdventProductRequest) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *GetListAdventProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type GetAdventProductsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdventProducts []*AdventProduct `protobuf:"bytes,1,rep,name=adventProducts,proto3" json:"adventProducts,omitempty"`
	Count          int32            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAdventProductsListResponse) Reset() {
	*x = GetAdventProductsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdventProductsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdventProductsListResponse) ProtoMessage() {}

func (x *GetAdventProductsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advent_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdventProductsListResponse.ProtoReflect.Descriptor instead.
func (*GetAdventProductsListResponse) Descriptor() ([]byte, []int) {
	return file_advent_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetAdventProductsListResponse) GetAdventProducts() []*AdventProduct {
	if x != nil {
		return x.AdventProducts
	}
	return nil
}

func (x *GetAdventProductsListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateAdventProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdventProduct *AdventProduct `protobuf:"bytes,1,opt,name=adventProduct,proto3" json:"adventProduct,omitempty"`
}

func (x *UpdateAdventProductRequest) Reset() {
	*x = UpdateAdventProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdventProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdventProductRequest) ProtoMessage() {}

func (x *UpdateAdventProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdventProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdventProductRequest) Descriptor() ([]byte, []int) {
	return file_advent_product_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAdventProductRequest) GetAdventProduct() *AdventProduct {
	if x != nil {
		return x.AdventProduct
	}
	return nil
}

var File_advent_product_proto protoreflect.FileDescriptor

var file_advent_product_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x72, 0x61, 0x63, 0x68, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x72, 0x61, 0x63, 0x68, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x53,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x22, 0x7d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x64, 0x76, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x64, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x61, 0x64,
	0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x62, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x44, 0x0a, 0x0d, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0d, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_advent_product_proto_rawDescOnce sync.Once
	file_advent_product_proto_rawDescData = file_advent_product_proto_rawDesc
)

func file_advent_product_proto_rawDescGZIP() []byte {
	file_advent_product_proto_rawDescOnce.Do(func() {
		file_advent_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_advent_product_proto_rawDescData)
	})
	return file_advent_product_proto_rawDescData
}

var file_advent_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_advent_product_proto_goTypes = []interface{}{
	(*AdventProduct)(nil),                 // 0: storage_service.AdventProduct
	(*CreateAdventProductRequest)(nil),    // 1: storage_service.CreateAdventProductRequest
	(*GetListAdventProductRequest)(nil),   // 2: storage_service.GetListAdventProductRequest
	(*GetAdventProductsListResponse)(nil), // 3: storage_service.GetAdventProductsListResponse
	(*UpdateAdventProductRequest)(nil),    // 4: storage_service.UpdateAdventProductRequest
}
var file_advent_product_proto_depIdxs = []int32{
	0, // 0: storage_service.GetAdventProductsListResponse.adventProducts:type_name -> storage_service.AdventProduct
	0, // 1: storage_service.UpdateAdventProductRequest.adventProduct:type_name -> storage_service.AdventProduct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_advent_product_proto_init() }
func file_advent_product_proto_init() {
	if File_advent_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_advent_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdventProduct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_advent_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdventProductRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_advent_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAdventProductRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_advent_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdventProductsListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_advent_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdventProductRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_advent_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_advent_product_proto_goTypes,
		DependencyIndexes: file_advent_product_proto_depIdxs,
		MessageInfos:      file_advent_product_proto_msgTypes,
	}.Build()
	File_advent_product_proto = out.File
	file_advent_product_proto_rawDesc = nil
	file_advent_product_proto_goTypes = nil
	file_advent_product_proto_depIdxs = nil
}
