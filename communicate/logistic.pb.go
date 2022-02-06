// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: logistic.proto

package communicate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CalulateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdCarring string `protobuf:"bytes,1,opt,name=idCarring,proto3" json:"idCarring,omitempty"`
	IdClient  string `protobuf:"bytes,2,opt,name=idClient,proto3" json:"idClient,omitempty"`
}

func (x *CalulateRequest) Reset() {
	*x = CalulateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalulateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalulateRequest) ProtoMessage() {}

func (x *CalulateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logistic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalulateRequest.ProtoReflect.Descriptor instead.
func (*CalulateRequest) Descriptor() ([]byte, []int) {
	return file_logistic_proto_rawDescGZIP(), []int{0}
}

func (x *CalulateRequest) GetIdCarring() string {
	if x != nil {
		return x.IdCarring
	}
	return ""
}

func (x *CalulateRequest) GetIdClient() string {
	if x != nil {
		return x.IdClient
	}
	return ""
}

type LatAndLong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat string `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng string `protobuf:"bytes,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *LatAndLong) Reset() {
	*x = LatAndLong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatAndLong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatAndLong) ProtoMessage() {}

func (x *LatAndLong) ProtoReflect() protoreflect.Message {
	mi := &file_logistic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatAndLong.ProtoReflect.Descriptor instead.
func (*LatAndLong) Descriptor() ([]byte, []int) {
	return file_logistic_proto_rawDescGZIP(), []int{1}
}

func (x *LatAndLong) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *LatAndLong) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin        *LatAndLong `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destiny       *LatAndLong `protobuf:"bytes,2,opt,name=destiny,proto3" json:"destiny,omitempty"`
	Meters        int64       `protobuf:"varint,3,opt,name=meters,proto3" json:"meters,omitempty"`
	Duration      int64       `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	HumanReadable string      `protobuf:"bytes,5,opt,name=HumanReadable,proto3" json:"HumanReadable,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logistic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_logistic_proto_rawDescGZIP(), []int{2}
}

func (x *CalculateResponse) GetOrigin() *LatAndLong {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *CalculateResponse) GetDestiny() *LatAndLong {
	if x != nil {
		return x.Destiny
	}
	return nil
}

func (x *CalculateResponse) GetMeters() int64 {
	if x != nil {
		return x.Meters
	}
	return 0
}

func (x *CalculateResponse) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *CalculateResponse) GetHumanReadable() string {
	if x != nil {
		return x.HumanReadable
	}
	return ""
}

type ValidateCarryingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdCarring string `protobuf:"bytes,1,opt,name=idCarring,proto3" json:"idCarring,omitempty"`
}

func (x *ValidateCarryingRequest) Reset() {
	*x = ValidateCarryingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCarryingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCarryingRequest) ProtoMessage() {}

func (x *ValidateCarryingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logistic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCarryingRequest.ProtoReflect.Descriptor instead.
func (*ValidateCarryingRequest) Descriptor() ([]byte, []int) {
	return file_logistic_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateCarryingRequest) GetIdCarring() string {
	if x != nil {
		return x.IdCarring
	}
	return ""
}

type ValidateCarryingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *ValidateCarryingResponse) Reset() {
	*x = ValidateCarryingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCarryingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCarryingResponse) ProtoMessage() {}

func (x *ValidateCarryingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logistic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCarryingResponse.ProtoReflect.Descriptor instead.
func (*ValidateCarryingResponse) Descriptor() ([]byte, []int) {
	return file_logistic_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateCarryingResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type ValidateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdClient string `protobuf:"bytes,1,opt,name=idClient,proto3" json:"idClient,omitempty"`
}

func (x *ValidateClientRequest) Reset() {
	*x = ValidateClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateClientRequest) ProtoMessage() {}

func (x *ValidateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logistic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateClientRequest.ProtoReflect.Descriptor instead.
func (*ValidateClientRequest) Descriptor() ([]byte, []int) {
	return file_logistic_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateClientRequest) GetIdClient() string {
	if x != nil {
		return x.IdClient
	}
	return ""
}

type ValidateClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *ValidateClientResponse) Reset() {
	*x = ValidateClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateClientResponse) ProtoMessage() {}

func (x *ValidateClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logistic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateClientResponse.ProtoReflect.Descriptor instead.
func (*ValidateClientResponse) Descriptor() ([]byte, []int) {
	return file_logistic_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateClientResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_logistic_proto protoreflect.FileDescriptor

var file_logistic_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4b, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x43, 0x61, 0x72, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x43, 0x61, 0x72, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x30, 0x0a,
	0x0a, 0x4c, 0x61, 0x74, 0x41, 0x6e, 0x64, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22,
	0xb9, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x61, 0x74, 0x41, 0x6e, 0x64, 0x4c, 0x6f,
	0x6e, 0x67, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x07, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x61,
	0x74, 0x41, 0x6e, 0x64, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65,
	0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x48, 0x75,
	0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x37, 0x0a, 0x17, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72, 0x79, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x43, 0x61, 0x72, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x43, 0x61, 0x72,
	0x72, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x72, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x16, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0xe9, 0x01, 0x0a, 0x13,
	0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x10, 0x2e, 0x43, 0x61, 0x6c, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72,
	0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x72, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72,
	0x79, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logistic_proto_rawDescOnce sync.Once
	file_logistic_proto_rawDescData = file_logistic_proto_rawDesc
)

func file_logistic_proto_rawDescGZIP() []byte {
	file_logistic_proto_rawDescOnce.Do(func() {
		file_logistic_proto_rawDescData = protoimpl.X.CompressGZIP(file_logistic_proto_rawDescData)
	})
	return file_logistic_proto_rawDescData
}

var file_logistic_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_logistic_proto_goTypes = []interface{}{
	(*CalulateRequest)(nil),          // 0: CalulateRequest
	(*LatAndLong)(nil),               // 1: LatAndLong
	(*CalculateResponse)(nil),        // 2: CalculateResponse
	(*ValidateCarryingRequest)(nil),  // 3: ValidateCarryingRequest
	(*ValidateCarryingResponse)(nil), // 4: ValidateCarryingResponse
	(*ValidateClientRequest)(nil),    // 5: ValidateClientRequest
	(*ValidateClientResponse)(nil),   // 6: ValidateClientResponse
}
var file_logistic_proto_depIdxs = []int32{
	1, // 0: CalculateResponse.origin:type_name -> LatAndLong
	1, // 1: CalculateResponse.destiny:type_name -> LatAndLong
	0, // 2: LogisticCommunicate.CalculateLogistic:input_type -> CalulateRequest
	3, // 3: LogisticCommunicate.ValidateCarringById:input_type -> ValidateCarryingRequest
	5, // 4: LogisticCommunicate.ValidateClientById:input_type -> ValidateClientRequest
	2, // 5: LogisticCommunicate.CalculateLogistic:output_type -> CalculateResponse
	4, // 6: LogisticCommunicate.ValidateCarringById:output_type -> ValidateCarryingResponse
	6, // 7: LogisticCommunicate.ValidateClientById:output_type -> ValidateClientResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_logistic_proto_init() }
func file_logistic_proto_init() {
	if File_logistic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logistic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalulateRequest); i {
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
		file_logistic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatAndLong); i {
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
		file_logistic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse); i {
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
		file_logistic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCarryingRequest); i {
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
		file_logistic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCarryingResponse); i {
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
		file_logistic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateClientRequest); i {
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
		file_logistic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateClientResponse); i {
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
			RawDescriptor: file_logistic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logistic_proto_goTypes,
		DependencyIndexes: file_logistic_proto_depIdxs,
		MessageInfos:      file_logistic_proto_msgTypes,
	}.Build()
	File_logistic_proto = out.File
	file_logistic_proto_rawDesc = nil
	file_logistic_proto_goTypes = nil
	file_logistic_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogisticCommunicateClient is the client API for LogisticCommunicate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogisticCommunicateClient interface {
	CalculateLogistic(ctx context.Context, in *CalulateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
	ValidateCarringById(ctx context.Context, in *ValidateCarryingRequest, opts ...grpc.CallOption) (*ValidateCarryingResponse, error)
	ValidateClientById(ctx context.Context, in *ValidateClientRequest, opts ...grpc.CallOption) (*ValidateClientResponse, error)
}

type logisticCommunicateClient struct {
	cc grpc.ClientConnInterface
}

func NewLogisticCommunicateClient(cc grpc.ClientConnInterface) LogisticCommunicateClient {
	return &logisticCommunicateClient{cc}
}

func (c *logisticCommunicateClient) CalculateLogistic(ctx context.Context, in *CalulateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/LogisticCommunicate/CalculateLogistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticCommunicateClient) ValidateCarringById(ctx context.Context, in *ValidateCarryingRequest, opts ...grpc.CallOption) (*ValidateCarryingResponse, error) {
	out := new(ValidateCarryingResponse)
	err := c.cc.Invoke(ctx, "/LogisticCommunicate/ValidateCarringById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticCommunicateClient) ValidateClientById(ctx context.Context, in *ValidateClientRequest, opts ...grpc.CallOption) (*ValidateClientResponse, error) {
	out := new(ValidateClientResponse)
	err := c.cc.Invoke(ctx, "/LogisticCommunicate/ValidateClientById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogisticCommunicateServer is the server API for LogisticCommunicate service.
type LogisticCommunicateServer interface {
	CalculateLogistic(context.Context, *CalulateRequest) (*CalculateResponse, error)
	ValidateCarringById(context.Context, *ValidateCarryingRequest) (*ValidateCarryingResponse, error)
	ValidateClientById(context.Context, *ValidateClientRequest) (*ValidateClientResponse, error)
}

// UnimplementedLogisticCommunicateServer can be embedded to have forward compatible implementations.
type UnimplementedLogisticCommunicateServer struct {
}

func (*UnimplementedLogisticCommunicateServer) CalculateLogistic(context.Context, *CalulateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateLogistic not implemented")
}
func (*UnimplementedLogisticCommunicateServer) ValidateCarringById(context.Context, *ValidateCarryingRequest) (*ValidateCarryingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCarringById not implemented")
}
func (*UnimplementedLogisticCommunicateServer) ValidateClientById(context.Context, *ValidateClientRequest) (*ValidateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateClientById not implemented")
}

func RegisterLogisticCommunicateServer(s *grpc.Server, srv LogisticCommunicateServer) {
	s.RegisterService(&_LogisticCommunicate_serviceDesc, srv)
}

func _LogisticCommunicate_CalculateLogistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalulateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticCommunicateServer).CalculateLogistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LogisticCommunicate/CalculateLogistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticCommunicateServer).CalculateLogistic(ctx, req.(*CalulateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticCommunicate_ValidateCarringById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateCarryingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticCommunicateServer).ValidateCarringById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LogisticCommunicate/ValidateCarringById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticCommunicateServer).ValidateCarringById(ctx, req.(*ValidateCarryingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticCommunicate_ValidateClientById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticCommunicateServer).ValidateClientById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LogisticCommunicate/ValidateClientById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticCommunicateServer).ValidateClientById(ctx, req.(*ValidateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogisticCommunicate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LogisticCommunicate",
	HandlerType: (*LogisticCommunicateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateLogistic",
			Handler:    _LogisticCommunicate_CalculateLogistic_Handler,
		},
		{
			MethodName: "ValidateCarringById",
			Handler:    _LogisticCommunicate_ValidateCarringById_Handler,
		},
		{
			MethodName: "ValidateClientById",
			Handler:    _LogisticCommunicate_ValidateClientById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logistic.proto",
}
