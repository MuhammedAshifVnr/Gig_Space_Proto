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

type CommonGigRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  int64                 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error   string                `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Data    map[string]*AnyValues `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CommonGigRes) Reset() {
	*x = CommonGigRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonGigRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonGigRes) ProtoMessage() {}

func (x *CommonGigRes) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CommonGigRes.ProtoReflect.Descriptor instead.
func (*CommonGigRes) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{0}
}

func (x *CommonGigRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommonGigRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CommonGigRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CommonGigRes) GetData() map[string]*AnyValues {
	if x != nil {
		return x.Data
	}
	return nil
}

type AnyValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*AnyValues_StringValue
	//	*AnyValues_IntValue
	//	*AnyValues_BoolValue
	//	*AnyValues_DoubleValue
	Value isAnyValues_Value `protobuf_oneof:"value"`
}

func (x *AnyValues) Reset() {
	*x = AnyValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyValues) ProtoMessage() {}

func (x *AnyValues) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AnyValues.ProtoReflect.Descriptor instead.
func (*AnyValues) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{1}
}

func (m *AnyValues) GetValue() isAnyValues_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *AnyValues) GetStringValue() string {
	if x, ok := x.GetValue().(*AnyValues_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *AnyValues) GetIntValue() int32 {
	if x, ok := x.GetValue().(*AnyValues_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *AnyValues) GetBoolValue() bool {
	if x, ok := x.GetValue().(*AnyValues_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *AnyValues) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*AnyValues_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

type isAnyValues_Value interface {
	isAnyValues_Value()
}

type AnyValues_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type AnyValues_IntValue struct {
	IntValue int32 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type AnyValues_BoolValue struct {
	BoolValue bool `protobuf:"varint,3,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type AnyValues_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

func (*AnyValues_StringValue) isAnyValues_Value() {}

func (*AnyValues_IntValue) isAnyValues_Value() {}

func (*AnyValues_BoolValue) isAnyValues_Value() {}

func (*AnyValues_DoubleValue) isAnyValues_Value() {}

type UpdateGigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64   `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	Title             string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description       string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Category          string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Type              string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Price             float64  `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	DeliveryDays      int64    `protobuf:"varint,6,opt,name=delivery_days,json=deliveryDays,proto3" json:"delivery_days,omitempty"`
	NumberOfRevisions int64    `protobuf:"varint,7,opt,name=number_of_revisions,json=numberOfRevisions,proto3" json:"number_of_revisions,omitempty"`
	UserId            uint32   `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Images            [][]byte `protobuf:"bytes,9,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *UpdateGigRequest) Reset() {
	*x = UpdateGigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGigRequest) ProtoMessage() {}

func (x *UpdateGigRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateGigRequest.ProtoReflect.Descriptor instead.
func (*UpdateGigRequest) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateGigRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGigRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateGigRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateGigRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateGigRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateGigRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateGigRequest) GetDeliveryDays() int64 {
	if x != nil {
		return x.DeliveryDays
	}
	return 0
}

func (x *UpdateGigRequest) GetNumberOfRevisions() int64 {
	if x != nil {
		return x.NumberOfRevisions
	}
	return 0
}

func (x *UpdateGigRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateGigRequest) GetImages() [][]byte {
	if x != nil {
		return x.Images
	}
	return nil
}

type GetGigsByFreelancerIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FreelancerId uint64 `protobuf:"varint,1,opt,name=freelancer_id,json=freelancerId,proto3" json:"freelancer_id,omitempty"`
}

func (x *GetGigsByFreelancerIDRequest) Reset() {
	*x = GetGigsByFreelancerIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGigsByFreelancerIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGigsByFreelancerIDRequest) ProtoMessage() {}

func (x *GetGigsByFreelancerIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetGigsByFreelancerIDRequest.ProtoReflect.Descriptor instead.
func (*GetGigsByFreelancerIDRequest) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{3}
}

func (x *GetGigsByFreelancerIDRequest) GetFreelancerId() uint64 {
	if x != nil {
		return x.FreelancerId
	}
	return 0
}

type GetGigsByFreelancerIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gigs []*Gig `protobuf:"bytes,1,rep,name=gigs,proto3" json:"gigs,omitempty"`
}

func (x *GetGigsByFreelancerIDResponse) Reset() {
	*x = GetGigsByFreelancerIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGigsByFreelancerIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGigsByFreelancerIDResponse) ProtoMessage() {}

func (x *GetGigsByFreelancerIDResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetGigsByFreelancerIDResponse.ProtoReflect.Descriptor instead.
func (*GetGigsByFreelancerIDResponse) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{4}
}

func (x *GetGigsByFreelancerIDResponse) GetGigs() []*Gig {
	if x != nil {
		return x.Gigs
	}
	return nil
}

type Gig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category     uint32   `protobuf:"varint,4,opt,name=category,proto3" json:"category,omitempty"`
	FreelancerId uint64   `protobuf:"varint,5,opt,name=freelancer_id,json=freelancerId,proto3" json:"freelancer_id,omitempty"`
	Price        float32  `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	DeliveryDays int32    `protobuf:"varint,7,opt,name=delivery_days,json=deliveryDays,proto3" json:"delivery_days,omitempty"`
	Revisions    int32    `protobuf:"varint,8,opt,name=revisions,proto3" json:"revisions,omitempty"`
	Image        []*Image `protobuf:"bytes,9,rep,name=image,proto3" json:"image,omitempty"`
}

func (x *Gig) Reset() {
	*x = Gig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gig) ProtoMessage() {}

func (x *Gig) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[5]
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
	return file_gig_proto_rawDescGZIP(), []int{5}
}

func (x *Gig) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Gig) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Gig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Gig) GetCategory() uint32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *Gig) GetFreelancerId() uint64 {
	if x != nil {
		return x.FreelancerId
	}
	return 0
}

func (x *Gig) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Gig) GetDeliveryDays() int32 {
	if x != nil {
		return x.DeliveryDays
	}
	return 0
}

func (x *Gig) GetRevisions() int32 {
	if x != nil {
		return x.Revisions
	}
	return 0
}

func (x *Gig) GetImage() []*Image {
	if x != nil {
		return x.Image
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
	Images            [][]byte `protobuf:"bytes,9,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *CreateGigReq) Reset() {
	*x = CreateGigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGigReq) ProtoMessage() {}

func (x *CreateGigReq) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[6]
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
	return file_gig_proto_rawDescGZIP(), []int{6}
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

func (x *CreateGigReq) GetImages() [][]byte {
	if x != nil {
		return x.Images
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_gig_proto_rawDescGZIP(), []int{7}
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gig_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gig_proto_msgTypes[8]
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
	return file_gig_proto_rawDescGZIP(), []int{8}
}

var File_gig_proto protoreflect.FileDescriptor

var file_gig_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x69, 0x67,
	0x22, 0xd0, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x69, 0x67, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x69, 0x67, 0x52, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x47, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x69, 0x67, 0x2e, 0x41,
	0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x9e, 0x01, 0x0a, 0x09, 0x41, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xa6, 0x02, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
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
	0x09, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x43, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x47, 0x69, 0x67, 0x73, 0x42, 0x79, 0x46, 0x72, 0x65, 0x65, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x66, 0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x3d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x47, 0x69, 0x67, 0x73, 0x42, 0x79, 0x46,
	0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x67, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x67, 0x69, 0x67, 0x2e, 0x47, 0x69, 0x67, 0x52, 0x04, 0x67, 0x69, 0x67,
	0x73, 0x22, 0x89, 0x02, 0x0a, 0x03, 0x47, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x66, 0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x69, 0x67,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x92, 0x02,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x79,
	0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x19, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdb,
	0x01, 0x0a, 0x0a, 0x47, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x12, 0x11, 0x2e, 0x67, 0x69, 0x67,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x67, 0x69, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x69, 0x67, 0x73, 0x42, 0x79, 0x46, 0x72,
	0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x67, 0x69, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x69, 0x67, 0x73, 0x42, 0x79, 0x46, 0x72, 0x65, 0x65, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x67, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x69, 0x67, 0x73, 0x42, 0x79, 0x46, 0x72, 0x65,
	0x65, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x69, 0x67, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x15, 0x2e, 0x67, 0x69, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x69, 0x67, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x69, 0x67, 0x52, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_gig_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gig_proto_goTypes = []any{
	(*CommonGigRes)(nil),                  // 0: gig.CommonGigRes
	(*AnyValues)(nil),                     // 1: gig.AnyValues
	(*UpdateGigRequest)(nil),              // 2: gig.UpdateGigRequest
	(*GetGigsByFreelancerIDRequest)(nil),  // 3: gig.GetGigsByFreelancerIDRequest
	(*GetGigsByFreelancerIDResponse)(nil), // 4: gig.GetGigsByFreelancerIDResponse
	(*Gig)(nil),                           // 5: gig.Gig
	(*CreateGigReq)(nil),                  // 6: gig.CreateGigReq
	(*Image)(nil),                         // 7: gig.Image
	(*EmptyResponse)(nil),                 // 8: gig.EmptyResponse
	nil,                                   // 9: gig.CommonGigRes.DataEntry
}
var file_gig_proto_depIdxs = []int32{
	9, // 0: gig.CommonGigRes.data:type_name -> gig.CommonGigRes.DataEntry
	5, // 1: gig.GetGigsByFreelancerIDResponse.gigs:type_name -> gig.Gig
	7, // 2: gig.Gig.image:type_name -> gig.Image
	1, // 3: gig.CommonGigRes.DataEntry.value:type_name -> gig.AnyValues
	6, // 4: gig.GigService.CreateGig:input_type -> gig.CreateGigReq
	3, // 5: gig.GigService.GetGigsByFreelancerID:input_type -> gig.GetGigsByFreelancerIDRequest
	2, // 6: gig.GigService.UpdateGigByID:input_type -> gig.UpdateGigRequest
	8, // 7: gig.GigService.CreateGig:output_type -> gig.EmptyResponse
	4, // 8: gig.GigService.GetGigsByFreelancerID:output_type -> gig.GetGigsByFreelancerIDResponse
	0, // 9: gig.GigService.UpdateGigByID:output_type -> gig.CommonGigRes
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gig_proto_init() }
func file_gig_proto_init() {
	if File_gig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gig_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonGigRes); i {
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
			switch v := v.(*AnyValues); i {
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
			switch v := v.(*UpdateGigRequest); i {
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
			switch v := v.(*GetGigsByFreelancerIDRequest); i {
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
			switch v := v.(*GetGigsByFreelancerIDResponse); i {
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
		file_gig_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
		file_gig_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
		file_gig_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Image); i {
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
		file_gig_proto_msgTypes[8].Exporter = func(v any, i int) any {
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
	file_gig_proto_msgTypes[1].OneofWrappers = []any{
		(*AnyValues_StringValue)(nil),
		(*AnyValues_IntValue)(nil),
		(*AnyValues_BoolValue)(nil),
		(*AnyValues_DoubleValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
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
