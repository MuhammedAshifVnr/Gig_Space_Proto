// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: gig.proto

package proto

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

type GetAllGigsByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAllGigsByIDReq) Reset() {
	*x = GetAllGigsByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGigsByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGigsByIDReq) ProtoMessage() {}

func (x *GetAllGigsByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGigsByIDReq.ProtoReflect.Descriptor instead.
func (*GetAllGigsByIDReq) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllGigsByIDReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAllGigsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gig []*Gig `protobuf:"bytes,1,rep,name=gig,proto3" json:"gig,omitempty"`
}

func (x *GetAllGigsResp) Reset() {
	*x = GetAllGigsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGigsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGigsResp) ProtoMessage() {}

func (x *GetAllGigsResp) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGigsResp.ProtoReflect.Descriptor instead.
func (*GetAllGigsResp) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllGigsResp) GetGig() []*Gig {
	if x != nil {
		return x.Gig
	}
	return nil
}

type Gig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Titel        string   `protobuf:"bytes,2,opt,name=Titel,proto3" json:"Titel,omitempty"`
	Description  string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	FrelancerID  uint32   `protobuf:"varint,4,opt,name=FrelancerID,proto3" json:"FrelancerID,omitempty"`
	Price        float32  `protobuf:"fixed32,5,opt,name=Price,proto3" json:"Price,omitempty"`
	Category     string   `protobuf:"bytes,6,opt,name=Category,proto3" json:"Category,omitempty"`
	DeliveryDays int64    `protobuf:"varint,7,opt,name=DeliveryDays,proto3" json:"DeliveryDays,omitempty"`
	Revisions    int64    `protobuf:"varint,8,opt,name=Revisions,proto3" json:"Revisions,omitempty"`
	Image_Urls   []string `protobuf:"bytes,9,rep,name=Image_Urls,json=ImageUrls,proto3" json:"Image_Urls,omitempty"`
}

func (x *Gig) Reset() {
	*x = Gig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gig) ProtoMessage() {}

func (x *Gig) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gig.ProtoReflect.Descriptor instead.
func (*Gig) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{2}
}

func (x *Gig) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Gig) GetTitel() string {
	if x != nil {
		return x.Titel
	}
	return ""
}

func (x *Gig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Gig) GetFrelancerID() uint32 {
	if x != nil {
		return x.FrelancerID
	}
	return 0
}

func (x *Gig) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Gig) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Gig) GetDeliveryDays() int64 {
	if x != nil {
		return x.DeliveryDays
	}
	return 0
}

func (x *Gig) GetRevisions() int64 {
	if x != nil {
		return x.Revisions
	}
	return 0
}

func (x *Gig) GetImage_Urls() []string {
	if x != nil {
		return x.Image_Urls
	}
	return nil
}

type CreateGigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title             string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description       string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Category          string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Type              string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Price             float64  `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	DeliveryDays      int64    `protobuf:"varint,6,opt,name=delivery_days,json=deliveryDays,proto3" json:"delivery_days,omitempty"`
	NumberOfRevisions int64    `protobuf:"varint,7,opt,name=number_of_revisions,json=numberOfRevisions,proto3" json:"number_of_revisions,omitempty"`
	UserId            uint32   `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Images            []string `protobuf:"bytes,9,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *CreateGigReq) Reset() {
	*x = CreateGigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGigReq) ProtoMessage() {}

func (x *CreateGigReq) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGigReq.ProtoReflect.Descriptor instead.
func (*CreateGigReq) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGigReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateGigReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateGigReq) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateGigReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateGigReq) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateGigReq) GetDeliveryDays() int64 {
	if x != nil {
		return x.DeliveryDays
	}
	return 0
}

func (x *CreateGigReq) GetNumberOfRevisions() int64 {
	if x != nil {
		return x.NumberOfRevisions
	}
	return 0
}

func (x *CreateGigReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateGigReq) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{4}
}

var File_gig_proto protoreflect.FileDescriptor

var file_gig_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x69, 0x67,
	0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x69, 0x67, 0x73, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47,
	0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x03, 0x67, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x67, 0x69, 0x67, 0x2e, 0x47, 0x69, 0x67, 0x52, 0x03,
	0x67, 0x69, 0x67, 0x22, 0x82, 0x02, 0x0a, 0x03, 0x47, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x65,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x72, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x72, 0x65, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x44, 0x61, 0x79, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x92, 0x02, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2e, 0x0a, 0x13,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x7e,
	0x0a, 0x0a, 0x47, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x12, 0x11, 0x2e, 0x67, 0x69, 0x67, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x67,
	0x69, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x69, 0x67, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x16, 0x2e, 0x67, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x69,
	0x67, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x69, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gig_proto_rawDescOnce sync.Once
	file_gig_proto_rawDescData = file_gig_proto_rawDesc
)

func file_gig_proto_rawDescGZIP() []byte {
	file_gig_proto_rawDescOnce.Do(func() {
		file_gig_proto_rawDescData = protoimpl.X.CompressGZIP(file_gig_proto_rawDescData)
	})
	return file_gig_proto_rawDescData
}

var file_gig_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gig_proto_goTypes = []any{
	(*GetAllGigsByIDReq)(nil), // 0: gig.GetAllGigsByIDReq
	(*GetAllGigsResp)(nil),    // 1: gig.GetAllGigsResp
	(*Gig)(nil),               // 2: gig.Gig
	(*CreateGigReq)(nil),      // 3: gig.CreateGigReq
	(*EmptyResponse)(nil),     // 4: gig.EmptyResponse
}
var file_gig_proto_depIdxs = []int32{
	2, // 0: gig.GetAllGigsResp.gig:type_name -> gig.Gig
	3, // 1: gig.GigService.CreateGig:input_type -> gig.CreateGigReq
	0, // 2: gig.GigService.GetAllGigByID:input_type -> gig.GetAllGigsByIDReq
	4, // 3: gig.GigService.CreateGig:output_type -> gig.EmptyResponse
	1, // 4: gig.GigService.GetAllGigByID:output_type -> gig.GetAllGigsResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gig_proto_init() }
func file_gig_proto_init() {
	if File_gig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gig_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllGigsByIDReq); i {
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
		file_gig_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllGigsResp); i {
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
		file_gig_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Gig); i {
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
		file_gig_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateGigReq); i {
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
		file_gig_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_gig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gig_proto_goTypes,
		DependencyIndexes: file_gig_proto_depIdxs,
		MessageInfos:      file_gig_proto_msgTypes,
	}.Build()
	File_gig_proto = out.File
	file_gig_proto_rawDesc = nil
	file_gig_proto_goTypes = nil
	file_gig_proto_depIdxs = nil
}
