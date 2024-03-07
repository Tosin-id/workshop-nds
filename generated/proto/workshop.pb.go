// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.2.0
// source: workshop.proto

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

type RequestHitungLuasPersegi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Panjang int32 `protobuf:"varint,1,opt,name=Panjang,proto3" json:"Panjang,omitempty"`
	Lebar   int32 `protobuf:"varint,2,opt,name=Lebar,proto3" json:"Lebar,omitempty"`
}

func (x *RequestHitungLuasPersegi) Reset() {
	*x = RequestHitungLuasPersegi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHitungLuasPersegi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHitungLuasPersegi) ProtoMessage() {}

func (x *RequestHitungLuasPersegi) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHitungLuasPersegi.ProtoReflect.Descriptor instead.
func (*RequestHitungLuasPersegi) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{0}
}

func (x *RequestHitungLuasPersegi) GetPanjang() int32 {
	if x != nil {
		return x.Panjang
	}
	return 0
}

func (x *RequestHitungLuasPersegi) GetLebar() int32 {
	if x != nil {
		return x.Lebar
	}
	return 0
}

type ResponseLuasPersegi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Luas int32 `protobuf:"varint,1,opt,name=Luas,proto3" json:"Luas,omitempty"`
}

func (x *ResponseLuasPersegi) Reset() {
	*x = ResponseLuasPersegi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLuasPersegi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLuasPersegi) ProtoMessage() {}

func (x *ResponseLuasPersegi) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLuasPersegi.ProtoReflect.Descriptor instead.
func (*ResponseLuasPersegi) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseLuasPersegi) GetLuas() int32 {
	if x != nil {
		return x.Luas
	}
	return 0
}

type RequestHitungKelilingPersegi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Panjang int32 `protobuf:"varint,1,opt,name=Panjang,proto3" json:"Panjang,omitempty"`
	Lebar   int32 `protobuf:"varint,2,opt,name=Lebar,proto3" json:"Lebar,omitempty"`
}

func (x *RequestHitungKelilingPersegi) Reset() {
	*x = RequestHitungKelilingPersegi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHitungKelilingPersegi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHitungKelilingPersegi) ProtoMessage() {}

func (x *RequestHitungKelilingPersegi) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHitungKelilingPersegi.ProtoReflect.Descriptor instead.
func (*RequestHitungKelilingPersegi) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{2}
}

func (x *RequestHitungKelilingPersegi) GetPanjang() int32 {
	if x != nil {
		return x.Panjang
	}
	return 0
}

func (x *RequestHitungKelilingPersegi) GetLebar() int32 {
	if x != nil {
		return x.Lebar
	}
	return 0
}

type ResponseKelilingPersegi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keliling int32 `protobuf:"varint,1,opt,name=Keliling,proto3" json:"Keliling,omitempty"`
}

func (x *ResponseKelilingPersegi) Reset() {
	*x = ResponseKelilingPersegi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseKelilingPersegi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseKelilingPersegi) ProtoMessage() {}

func (x *ResponseKelilingPersegi) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseKelilingPersegi.ProtoReflect.Descriptor instead.
func (*ResponseKelilingPersegi) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseKelilingPersegi) GetKeliling() int32 {
	if x != nil {
		return x.Keliling
	}
	return 0
}

type RequestHitungLuasLingkaran struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jarijari int32 `protobuf:"varint,1,opt,name=Jarijari,proto3" json:"Jarijari,omitempty"`
}

func (x *RequestHitungLuasLingkaran) Reset() {
	*x = RequestHitungLuasLingkaran{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHitungLuasLingkaran) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHitungLuasLingkaran) ProtoMessage() {}

func (x *RequestHitungLuasLingkaran) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHitungLuasLingkaran.ProtoReflect.Descriptor instead.
func (*RequestHitungLuasLingkaran) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{4}
}

func (x *RequestHitungLuasLingkaran) GetJarijari() int32 {
	if x != nil {
		return x.Jarijari
	}
	return 0
}

type ResponseLuasLingkaran struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Luas float64 `protobuf:"fixed64,1,opt,name=Luas,proto3" json:"Luas,omitempty"`
}

func (x *ResponseLuasLingkaran) Reset() {
	*x = ResponseLuasLingkaran{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLuasLingkaran) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLuasLingkaran) ProtoMessage() {}

func (x *ResponseLuasLingkaran) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLuasLingkaran.ProtoReflect.Descriptor instead.
func (*ResponseLuasLingkaran) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseLuasLingkaran) GetLuas() float64 {
	if x != nil {
		return x.Luas
	}
	return 0
}

type RequestHitungKelilingLingkaran struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jarijari float64 `protobuf:"fixed64,1,opt,name=Jarijari,proto3" json:"Jarijari,omitempty"`
}

func (x *RequestHitungKelilingLingkaran) Reset() {
	*x = RequestHitungKelilingLingkaran{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHitungKelilingLingkaran) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHitungKelilingLingkaran) ProtoMessage() {}

func (x *RequestHitungKelilingLingkaran) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHitungKelilingLingkaran.ProtoReflect.Descriptor instead.
func (*RequestHitungKelilingLingkaran) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{6}
}

func (x *RequestHitungKelilingLingkaran) GetJarijari() float64 {
	if x != nil {
		return x.Jarijari
	}
	return 0
}

type ResponseKelilingLingkaran struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keliling float64 `protobuf:"fixed64,1,opt,name=Keliling,proto3" json:"Keliling,omitempty"`
}

func (x *ResponseKelilingLingkaran) Reset() {
	*x = ResponseKelilingLingkaran{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseKelilingLingkaran) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseKelilingLingkaran) ProtoMessage() {}

func (x *ResponseKelilingLingkaran) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseKelilingLingkaran.ProtoReflect.Descriptor instead.
func (*ResponseKelilingLingkaran) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseKelilingLingkaran) GetKeliling() float64 {
	if x != nil {
		return x.Keliling
	}
	return 0
}

type RequestHitungLuasSegitiga struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alas   int32 `protobuf:"varint,1,opt,name=Alas,proto3" json:"Alas,omitempty"`
	Tinggi int32 `protobuf:"varint,2,opt,name=Tinggi,proto3" json:"Tinggi,omitempty"`
}

func (x *RequestHitungLuasSegitiga) Reset() {
	*x = RequestHitungLuasSegitiga{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHitungLuasSegitiga) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHitungLuasSegitiga) ProtoMessage() {}

func (x *RequestHitungLuasSegitiga) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHitungLuasSegitiga.ProtoReflect.Descriptor instead.
func (*RequestHitungLuasSegitiga) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{8}
}

func (x *RequestHitungLuasSegitiga) GetAlas() int32 {
	if x != nil {
		return x.Alas
	}
	return 0
}

func (x *RequestHitungLuasSegitiga) GetTinggi() int32 {
	if x != nil {
		return x.Tinggi
	}
	return 0
}

type ResponseLuasSegitiga struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Luas float64 `protobuf:"fixed64,1,opt,name=Luas,proto3" json:"Luas,omitempty"`
}

func (x *ResponseLuasSegitiga) Reset() {
	*x = ResponseLuasSegitiga{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLuasSegitiga) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLuasSegitiga) ProtoMessage() {}

func (x *ResponseLuasSegitiga) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLuasSegitiga.ProtoReflect.Descriptor instead.
func (*ResponseLuasSegitiga) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{9}
}

func (x *ResponseLuasSegitiga) GetLuas() float64 {
	if x != nil {
		return x.Luas
	}
	return 0
}

type RequestHitungKelilingSegitiga struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sisi1 int32 `protobuf:"varint,1,opt,name=Sisi1,proto3" json:"Sisi1,omitempty"`
	Sisi2 int32 `protobuf:"varint,2,opt,name=Sisi2,proto3" json:"Sisi2,omitempty"`
	Sisi3 int32 `protobuf:"varint,3,opt,name=Sisi3,proto3" json:"Sisi3,omitempty"`
}

func (x *RequestHitungKelilingSegitiga) Reset() {
	*x = RequestHitungKelilingSegitiga{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHitungKelilingSegitiga) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHitungKelilingSegitiga) ProtoMessage() {}

func (x *RequestHitungKelilingSegitiga) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHitungKelilingSegitiga.ProtoReflect.Descriptor instead.
func (*RequestHitungKelilingSegitiga) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{10}
}

func (x *RequestHitungKelilingSegitiga) GetSisi1() int32 {
	if x != nil {
		return x.Sisi1
	}
	return 0
}

func (x *RequestHitungKelilingSegitiga) GetSisi2() int32 {
	if x != nil {
		return x.Sisi2
	}
	return 0
}

func (x *RequestHitungKelilingSegitiga) GetSisi3() int32 {
	if x != nil {
		return x.Sisi3
	}
	return 0
}

type ResponseKelilingSegitiga struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keliling int32 `protobuf:"varint,1,opt,name=Keliling,proto3" json:"Keliling,omitempty"`
}

func (x *ResponseKelilingSegitiga) Reset() {
	*x = ResponseKelilingSegitiga{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workshop_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseKelilingSegitiga) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseKelilingSegitiga) ProtoMessage() {}

func (x *ResponseKelilingSegitiga) ProtoReflect() protoreflect.Message {
	mi := &file_workshop_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseKelilingSegitiga.ProtoReflect.Descriptor instead.
func (*ResponseKelilingSegitiga) Descriptor() ([]byte, []int) {
	return file_workshop_proto_rawDescGZIP(), []int{11}
}

func (x *ResponseKelilingSegitiga) GetKeliling() int32 {
	if x != nil {
		return x.Keliling
	}
	return 0
}

var File_workshop_proto protoreflect.FileDescriptor

var file_workshop_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4c, 0x75, 0x61, 0x73, 0x50, 0x65, 0x72, 0x73,
	0x65, 0x67, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x61, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x65, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65,
	0x62, 0x61, 0x72, 0x22, 0x29, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c,
	0x75, 0x61, 0x73, 0x50, 0x65, 0x72, 0x73, 0x65, 0x67, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x75,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4c, 0x75, 0x61, 0x73, 0x22, 0x4e,
	0x0a, 0x1c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4b,
	0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x73, 0x65, 0x67, 0x69, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x61, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x50, 0x61, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x62, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x62, 0x61, 0x72, 0x22, 0x35,
	0x0a, 0x17, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69,
	0x6e, 0x67, 0x50, 0x65, 0x72, 0x73, 0x65, 0x67, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x6c,
	0x69, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4b, 0x65, 0x6c,
	0x69, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x38, 0x0a, 0x1a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4c, 0x75, 0x61, 0x73, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61,
	0x72, 0x61, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4a, 0x61, 0x72, 0x69, 0x6a, 0x61, 0x72, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4a, 0x61, 0x72, 0x69, 0x6a, 0x61, 0x72, 0x69, 0x22,
	0x2b, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x75, 0x61, 0x73, 0x4c,
	0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x75, 0x61, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x4c, 0x75, 0x61, 0x73, 0x22, 0x3c, 0x0a, 0x1e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4b, 0x65, 0x6c,
	0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x4a, 0x61, 0x72, 0x69, 0x6a, 0x61, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x4a, 0x61, 0x72, 0x69, 0x6a, 0x61, 0x72, 0x69, 0x22, 0x37, 0x0a, 0x19, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x6c, 0x69, 0x6c,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4b, 0x65, 0x6c, 0x69, 0x6c,
	0x69, 0x6e, 0x67, 0x22, 0x47, 0x0a, 0x19, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69,
	0x74, 0x75, 0x6e, 0x67, 0x4c, 0x75, 0x61, 0x73, 0x53, 0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x41, 0x6c, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x41, 0x6c, 0x61, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x69, 0x6e, 0x67, 0x67, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x54, 0x69, 0x6e, 0x67, 0x67, 0x69, 0x22, 0x2a, 0x0a, 0x14,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x75, 0x61, 0x73, 0x53, 0x65, 0x67, 0x69,
	0x74, 0x69, 0x67, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x75, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x4c, 0x75, 0x61, 0x73, 0x22, 0x61, 0x0a, 0x1d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x69, 0x73,
	0x69, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x69, 0x73, 0x69, 0x31, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x69, 0x73, 0x69, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x53, 0x69, 0x73, 0x69, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x69, 0x73, 0x69, 0x33, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x69, 0x73, 0x69, 0x33, 0x22, 0x36, 0x0a, 0x18, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x6c, 0x69, 0x6c,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4b, 0x65, 0x6c, 0x69, 0x6c,
	0x69, 0x6e, 0x67, 0x32, 0xaf, 0x04, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x67, 0x75, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x50, 0x0a, 0x11, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4c, 0x75, 0x61,
	0x73, 0x50, 0x65, 0x72, 0x73, 0x65, 0x67, 0x69, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4c, 0x75,
	0x61, 0x73, 0x50, 0x65, 0x72, 0x73, 0x65, 0x67, 0x69, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x75, 0x61, 0x73, 0x50, 0x65,
	0x72, 0x73, 0x65, 0x67, 0x69, 0x12, 0x5c, 0x0a, 0x15, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4b,
	0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x73, 0x65, 0x67, 0x69, 0x12, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69,
	0x74, 0x75, 0x6e, 0x67, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x73,
	0x65, 0x67, 0x69, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x73,
	0x65, 0x67, 0x69, 0x12, 0x56, 0x0a, 0x13, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4c, 0x75, 0x61,
	0x73, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67,
	0x4c, 0x75, 0x61, 0x73, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x1a, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x75,
	0x61, 0x73, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12, 0x62, 0x0a, 0x17, 0x48,
	0x69, 0x74, 0x75, 0x6e, 0x67, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6e,
	0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4b, 0x65, 0x6c, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x1a, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65,
	0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12,
	0x53, 0x0a, 0x12, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4c, 0x75, 0x61, 0x73, 0x53, 0x65, 0x67,
	0x69, 0x74, 0x69, 0x67, 0x61, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4c, 0x75, 0x61, 0x73, 0x53,
	0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x75, 0x61, 0x73, 0x53, 0x65, 0x67, 0x69,
	0x74, 0x69, 0x67, 0x61, 0x12, 0x5f, 0x0a, 0x16, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4b, 0x65,
	0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x12, 0x24,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x69,
	0x74, 0x75, 0x6e, 0x67, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67, 0x69,
	0x74, 0x69, 0x67, 0x61, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67,
	0x69, 0x74, 0x69, 0x67, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workshop_proto_rawDescOnce sync.Once
	file_workshop_proto_rawDescData = file_workshop_proto_rawDesc
)

func file_workshop_proto_rawDescGZIP() []byte {
	file_workshop_proto_rawDescOnce.Do(func() {
		file_workshop_proto_rawDescData = protoimpl.X.CompressGZIP(file_workshop_proto_rawDescData)
	})
	return file_workshop_proto_rawDescData
}

var file_workshop_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_workshop_proto_goTypes = []interface{}{
	(*RequestHitungLuasPersegi)(nil),       // 0: proto.RequestHitungLuasPersegi
	(*ResponseLuasPersegi)(nil),            // 1: proto.ResponseLuasPersegi
	(*RequestHitungKelilingPersegi)(nil),   // 2: proto.RequestHitungKelilingPersegi
	(*ResponseKelilingPersegi)(nil),        // 3: proto.ResponseKelilingPersegi
	(*RequestHitungLuasLingkaran)(nil),     // 4: proto.RequestHitungLuasLingkaran
	(*ResponseLuasLingkaran)(nil),          // 5: proto.ResponseLuasLingkaran
	(*RequestHitungKelilingLingkaran)(nil), // 6: proto.RequestHitungKelilingLingkaran
	(*ResponseKelilingLingkaran)(nil),      // 7: proto.ResponseKelilingLingkaran
	(*RequestHitungLuasSegitiga)(nil),      // 8: proto.RequestHitungLuasSegitiga
	(*ResponseLuasSegitiga)(nil),           // 9: proto.ResponseLuasSegitiga
	(*RequestHitungKelilingSegitiga)(nil),  // 10: proto.RequestHitungKelilingSegitiga
	(*ResponseKelilingSegitiga)(nil),       // 11: proto.ResponseKelilingSegitiga
}
var file_workshop_proto_depIdxs = []int32{
	0,  // 0: proto.BangunDatar.HitungLuasPersegi:input_type -> proto.RequestHitungLuasPersegi
	2,  // 1: proto.BangunDatar.HitungKelilingPersegi:input_type -> proto.RequestHitungKelilingPersegi
	4,  // 2: proto.BangunDatar.HitungLuasLingkaran:input_type -> proto.RequestHitungLuasLingkaran
	6,  // 3: proto.BangunDatar.HitungKelilingLingkaran:input_type -> proto.RequestHitungKelilingLingkaran
	8,  // 4: proto.BangunDatar.HitungLuasSegitiga:input_type -> proto.RequestHitungLuasSegitiga
	10, // 5: proto.BangunDatar.HitungKelilingSegitiga:input_type -> proto.RequestHitungKelilingSegitiga
	1,  // 6: proto.BangunDatar.HitungLuasPersegi:output_type -> proto.ResponseLuasPersegi
	3,  // 7: proto.BangunDatar.HitungKelilingPersegi:output_type -> proto.ResponseKelilingPersegi
	5,  // 8: proto.BangunDatar.HitungLuasLingkaran:output_type -> proto.ResponseLuasLingkaran
	7,  // 9: proto.BangunDatar.HitungKelilingLingkaran:output_type -> proto.ResponseKelilingLingkaran
	9,  // 10: proto.BangunDatar.HitungLuasSegitiga:output_type -> proto.ResponseLuasSegitiga
	11, // 11: proto.BangunDatar.HitungKelilingSegitiga:output_type -> proto.ResponseKelilingSegitiga
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_workshop_proto_init() }
func file_workshop_proto_init() {
	if File_workshop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workshop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHitungLuasPersegi); i {
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
		file_workshop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLuasPersegi); i {
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
		file_workshop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHitungKelilingPersegi); i {
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
		file_workshop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseKelilingPersegi); i {
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
		file_workshop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHitungLuasLingkaran); i {
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
		file_workshop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLuasLingkaran); i {
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
		file_workshop_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHitungKelilingLingkaran); i {
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
		file_workshop_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseKelilingLingkaran); i {
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
		file_workshop_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHitungLuasSegitiga); i {
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
		file_workshop_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLuasSegitiga); i {
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
		file_workshop_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHitungKelilingSegitiga); i {
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
		file_workshop_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseKelilingSegitiga); i {
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
			RawDescriptor: file_workshop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workshop_proto_goTypes,
		DependencyIndexes: file_workshop_proto_depIdxs,
		MessageInfos:      file_workshop_proto_msgTypes,
	}.Build()
	File_workshop_proto = out.File
	file_workshop_proto_rawDesc = nil
	file_workshop_proto_goTypes = nil
	file_workshop_proto_depIdxs = nil
}
