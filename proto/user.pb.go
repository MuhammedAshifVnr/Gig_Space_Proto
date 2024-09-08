// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.4
// source: user.proto

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

type SignupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Country   string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Role      string `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *SignupReq) Reset() {
	*x = SignupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupReq) ProtoMessage() {}

func (x *SignupReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupReq.ProtoReflect.Descriptor instead.
func (*SignupReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *SignupReq) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *SignupReq) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *SignupReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignupReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *SignupReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignupReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type SignupRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Otp     int64  `protobuf:"varint,4,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *SignupRes) Reset() {
	*x = SignupRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRes) ProtoMessage() {}

func (x *SignupRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRes.ProtoReflect.Descriptor instead.
func (*SignupRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *SignupRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SignupRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SignupRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SignupRes) GetOtp() int64 {
	if x != nil {
		return x.Otp
	}
	return 0
}

type VerifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Otp   string `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyReq) Reset() {
	*x = VerifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyReq) ProtoMessage() {}

func (x *VerifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyReq.ProtoReflect.Descriptor instead.
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyReq) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *VerifyReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VerifyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *VerifyRes) Reset() {
	*x = VerifyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRes) ProtoMessage() {}

func (x *VerifyRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRes.ProtoReflect.Descriptor instead.
func (*VerifyRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VerifyRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *VerifyRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *LoginRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LoginRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CommonRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  int64                `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error   string               `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Data    map[string]*AnyValue `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CommonRes) Reset() {
	*x = CommonRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRes) ProtoMessage() {}

func (x *CommonRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRes.ProtoReflect.Descriptor instead.
func (*CommonRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *CommonRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommonRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CommonRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CommonRes) GetData() map[string]*AnyValue {
	if x != nil {
		return x.Data
	}
	return nil
}

type AnyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*AnyValue_StringValue
	//	*AnyValue_IntValue
	//	*AnyValue_BoolValue
	//	*AnyValue_DoubleValue
	Value isAnyValue_Value `protobuf_oneof:"value"`
}

func (x *AnyValue) Reset() {
	*x = AnyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyValue) ProtoMessage() {}

func (x *AnyValue) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyValue.ProtoReflect.Descriptor instead.
func (*AnyValue) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (m *AnyValue) GetValue() isAnyValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *AnyValue) GetStringValue() string {
	if x, ok := x.GetValue().(*AnyValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *AnyValue) GetIntValue() int32 {
	if x, ok := x.GetValue().(*AnyValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *AnyValue) GetBoolValue() bool {
	if x, ok := x.GetValue().(*AnyValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *AnyValue) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*AnyValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

type isAnyValue_Value interface {
	isAnyValue_Value()
}

type AnyValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type AnyValue_IntValue struct {
	IntValue int32 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type AnyValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,3,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type AnyValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"` // Add more types as needed.
}

func (*AnyValue_StringValue) isAnyValue_Value() {}

func (*AnyValue_IntValue) isAnyValue_Value() {}

func (*AnyValue_BoolValue) isAnyValue_Value() {}

func (*AnyValue_DoubleValue) isAnyValue_Value() {}

type AdLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AdLoginReq) Reset() {
	*x = AdLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdLoginReq) ProtoMessage() {}

func (x *AdLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdLoginReq.ProtoReflect.Descriptor instead.
func (*AdLoginReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *AdLoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CategoryReq) Reset() {
	*x = CategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryReq) ProtoMessage() {}

func (x *CategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryReq.ProtoReflect.Descriptor instead.
func (*CategoryReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *CategoryReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddSkillReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillName string `protobuf:"bytes,1,opt,name=skill_name,json=skillName,proto3" json:"skill_name,omitempty"`
}

func (x *AddSkillReq) Reset() {
	*x = AddSkillReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSkillReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSkillReq) ProtoMessage() {}

func (x *AddSkillReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSkillReq.ProtoReflect.Descriptor instead.
func (*AddSkillReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *AddSkillReq) GetSkillName() string {
	if x != nil {
		return x.SkillName
	}
	return ""
}

type EmtpyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmtpyRes) Reset() {
	*x = EmtpyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmtpyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmtpyRes) ProtoMessage() {}

func (x *EmtpyRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmtpyRes.ProtoReflect.Descriptor instead.
func (*EmtpyRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

type SkillInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillName        string `protobuf:"bytes,1,opt,name=skill_name,json=skillName,proto3" json:"skill_name,omitempty"`
	ProficiencyLevel int32  `protobuf:"varint,2,opt,name=proficiency_level,json=proficiencyLevel,proto3" json:"proficiency_level,omitempty"`
}

func (x *SkillInput) Reset() {
	*x = SkillInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillInput) ProtoMessage() {}

func (x *SkillInput) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillInput.ProtoReflect.Descriptor instead.
func (*SkillInput) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *SkillInput) GetSkillName() string {
	if x != nil {
		return x.SkillName
	}
	return ""
}

func (x *SkillInput) GetProficiencyLevel() int32 {
	if x != nil {
		return x.ProficiencyLevel
	}
	return 0
}

type FreeAddSkillsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Skills []*SkillInput `protobuf:"bytes,2,rep,name=skills,proto3" json:"skills,omitempty"`
}

func (x *FreeAddSkillsReq) Reset() {
	*x = FreeAddSkillsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeAddSkillsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeAddSkillsReq) ProtoMessage() {}

func (x *FreeAddSkillsReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeAddSkillsReq.ProtoReflect.Descriptor instead.
func (*FreeAddSkillsReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{13}
}

func (x *FreeAddSkillsReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FreeAddSkillsReq) GetSkills() []*SkillInput {
	if x != nil {
		return x.Skills
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x22, 0x65, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x33, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x53, 0x0a, 0x09,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x3c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x68, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xcb, 0x01, 0x0a, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x47,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9d, 0x01, 0x0a, 0x08, 0x41, 0x6e, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x0a, 0x41, 0x64, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x0b, 0x41, 0x64,
	0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x74, 0x70,
	0x79, 0x52, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x0a, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x55,
	0x0a, 0x10, 0x46, 0x72, 0x65, 0x65, 0x41, 0x64, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x06, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x32, 0xec, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x69,
	0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x11, 0x46, 0x72, 0x65, 0x65, 0x6c, 0x61,
	0x63, 0x65, 0x72, 0x41, 0x64, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x41, 0x64, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_user_proto_goTypes = []any{
	(*SignupReq)(nil),        // 0: user.SignupReq
	(*SignupRes)(nil),        // 1: user.SignupRes
	(*VerifyReq)(nil),        // 2: user.VerifyReq
	(*VerifyRes)(nil),        // 3: user.VerifyRes
	(*LoginReq)(nil),         // 4: user.LoginReq
	(*LoginRes)(nil),         // 5: user.LoginRes
	(*CommonRes)(nil),        // 6: user.CommonRes
	(*AnyValue)(nil),         // 7: user.AnyValue
	(*AdLoginReq)(nil),       // 8: user.AdLoginReq
	(*CategoryReq)(nil),      // 9: user.CategoryReq
	(*AddSkillReq)(nil),      // 10: user.AddSkillReq
	(*EmtpyRes)(nil),         // 11: user.EmtpyRes
	(*SkillInput)(nil),       // 12: user.SkillInput
	(*FreeAddSkillsReq)(nil), // 13: user.FreeAddSkillsReq
	nil,                      // 14: user.CommonRes.DataEntry
}
var file_user_proto_depIdxs = []int32{
	14, // 0: user.CommonRes.data:type_name -> user.CommonRes.DataEntry
	12, // 1: user.FreeAddSkillsReq.skills:type_name -> user.SkillInput
	7,  // 2: user.CommonRes.DataEntry.value:type_name -> user.AnyValue
	0,  // 3: user.UserService.UserSignup:input_type -> user.SignupReq
	2,  // 4: user.UserService.VerifyingEmail:input_type -> user.VerifyReq
	4,  // 5: user.UserService.Login:input_type -> user.LoginReq
	8,  // 6: user.UserService.AdminLogin:input_type -> user.AdLoginReq
	9,  // 7: user.UserService.AddCategory:input_type -> user.CategoryReq
	10, // 8: user.UserService.AddSkill:input_type -> user.AddSkillReq
	13, // 9: user.UserService.FreelacerAddSkill:input_type -> user.FreeAddSkillsReq
	1,  // 10: user.UserService.UserSignup:output_type -> user.SignupRes
	3,  // 11: user.UserService.VerifyingEmail:output_type -> user.VerifyRes
	5,  // 12: user.UserService.Login:output_type -> user.LoginRes
	6,  // 13: user.UserService.AdminLogin:output_type -> user.CommonRes
	6,  // 14: user.UserService.AddCategory:output_type -> user.CommonRes
	6,  // 15: user.UserService.AddSkill:output_type -> user.CommonRes
	6,  // 16: user.UserService.FreelacerAddSkill:output_type -> user.CommonRes
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SignupReq); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SignupRes); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyReq); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyRes); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*LoginReq); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LoginRes); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CommonRes); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AnyValue); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*AdLoginReq); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryReq); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*AddSkillReq); i {
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
		file_user_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*EmtpyRes); i {
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
		file_user_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*SkillInput); i {
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
		file_user_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*FreeAddSkillsReq); i {
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
	file_user_proto_msgTypes[7].OneofWrappers = []any{
		(*AnyValue_StringValue)(nil),
		(*AnyValue_IntValue)(nil),
		(*AnyValue_BoolValue)(nil),
		(*AnyValue_DoubleValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
