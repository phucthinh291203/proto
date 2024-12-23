// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: proto/storage.proto

package protobufpb

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

type GetPresignedPutObjectURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PictureName string `protobuf:"bytes,1,opt,name=pictureName,proto3" json:"pictureName,omitempty"` // anhSinhVien1.png
}

func (x *GetPresignedPutObjectURLRequest) Reset() {
	*x = GetPresignedPutObjectURLRequest{}
	mi := &file_proto_storage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPresignedPutObjectURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPresignedPutObjectURLRequest) ProtoMessage() {}

func (x *GetPresignedPutObjectURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPresignedPutObjectURLRequest.ProtoReflect.Descriptor instead.
func (*GetPresignedPutObjectURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{0}
}

func (x *GetPresignedPutObjectURLRequest) GetPictureName() string {
	if x != nil {
		return x.PictureName
	}
	return ""
}

type GetPresignedPutObjectURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"` // token_anhSinhVien1.png
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`   // https://localhost:9000/put ....
}

func (x *GetPresignedPutObjectURLResponse) Reset() {
	*x = GetPresignedPutObjectURLResponse{}
	mi := &file_proto_storage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPresignedPutObjectURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPresignedPutObjectURLResponse) ProtoMessage() {}

func (x *GetPresignedPutObjectURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPresignedPutObjectURLResponse.ProtoReflect.Descriptor instead.
func (*GetPresignedPutObjectURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{1}
}

func (x *GetPresignedPutObjectURLResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetPresignedPutObjectURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetPresignedGetObjectURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"` // token_anhSinhVien1.png
}

func (x *GetPresignedGetObjectURLRequest) Reset() {
	*x = GetPresignedGetObjectURLRequest{}
	mi := &file_proto_storage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPresignedGetObjectURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPresignedGetObjectURLRequest) ProtoMessage() {}

func (x *GetPresignedGetObjectURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPresignedGetObjectURLRequest.ProtoReflect.Descriptor instead.
func (*GetPresignedGetObjectURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{2}
}

func (x *GetPresignedGetObjectURLRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetPresignedGetObjectURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"` // https://localhost:9000/get
}

func (x *GetPresignedGetObjectURLResponse) Reset() {
	*x = GetPresignedGetObjectURLResponse{}
	mi := &file_proto_storage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPresignedGetObjectURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPresignedGetObjectURLResponse) ProtoMessage() {}

func (x *GetPresignedGetObjectURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPresignedGetObjectURLResponse.ProtoReflect.Descriptor instead.
func (*GetPresignedGetObjectURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{3}
}

func (x *GetPresignedGetObjectURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CheckObjectExistResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CheckObjectExistResquest) Reset() {
	*x = CheckObjectExistResquest{}
	mi := &file_proto_storage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckObjectExistResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckObjectExistResquest) ProtoMessage() {}

func (x *CheckObjectExistResquest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckObjectExistResquest.ProtoReflect.Descriptor instead.
func (*CheckObjectExistResquest) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{4}
}

func (x *CheckObjectExistResquest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type CheckObjectExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *CheckObjectExistResponse) Reset() {
	*x = CheckObjectExistResponse{}
	mi := &file_proto_storage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckObjectExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckObjectExistResponse) ProtoMessage() {}

func (x *CheckObjectExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckObjectExistResponse.ProtoReflect.Descriptor instead.
func (*CheckObjectExistResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{5}
}

func (x *CheckObjectExistResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

var File_proto_storage_proto protoreflect.FileDescriptor

var file_proto_storage_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x20, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x35, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x34, 0x0a, 0x20, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x47, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x2e, 0x0a, 0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x30, 0x0a, 0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x32, 0x98, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x72, 0x76, 0x12, 0x5f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x20,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x75, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50,
	0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x12,
	0x20, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x47, 0x65,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x11,
	0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_storage_proto_rawDescOnce sync.Once
	file_proto_storage_proto_rawDescData = file_proto_storage_proto_rawDesc
)

func file_proto_storage_proto_rawDescGZIP() []byte {
	file_proto_storage_proto_rawDescOnce.Do(func() {
		file_proto_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_storage_proto_rawDescData)
	})
	return file_proto_storage_proto_rawDescData
}

var file_proto_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_storage_proto_goTypes = []any{
	(*GetPresignedPutObjectURLRequest)(nil),  // 0: GetPresignedPutObjectURLRequest
	(*GetPresignedPutObjectURLResponse)(nil), // 1: GetPresignedPutObjectURLResponse
	(*GetPresignedGetObjectURLRequest)(nil),  // 2: GetPresignedGetObjectURLRequest
	(*GetPresignedGetObjectURLResponse)(nil), // 3: GetPresignedGetObjectURLResponse
	(*CheckObjectExistResquest)(nil),         // 4: CheckObjectExistResquest
	(*CheckObjectExistResponse)(nil),         // 5: CheckObjectExistResponse
}
var file_proto_storage_proto_depIdxs = []int32{
	0, // 0: StorageSrv.GetPresignedPutObjectURL:input_type -> GetPresignedPutObjectURLRequest
	2, // 1: StorageSrv.GetPresignedGetObjectURL:input_type -> GetPresignedGetObjectURLRequest
	4, // 2: StorageSrv.CheckObjectExist:input_type -> CheckObjectExistResquest
	1, // 3: StorageSrv.GetPresignedPutObjectURL:output_type -> GetPresignedPutObjectURLResponse
	3, // 4: StorageSrv.GetPresignedGetObjectURL:output_type -> GetPresignedGetObjectURLResponse
	5, // 5: StorageSrv.CheckObjectExist:output_type -> CheckObjectExistResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_storage_proto_init() }
func file_proto_storage_proto_init() {
	if File_proto_storage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_storage_proto_goTypes,
		DependencyIndexes: file_proto_storage_proto_depIdxs,
		MessageInfos:      file_proto_storage_proto_msgTypes,
	}.Build()
	File_proto_storage_proto = out.File
	file_proto_storage_proto_rawDesc = nil
	file_proto_storage_proto_goTypes = nil
	file_proto_storage_proto_depIdxs = nil
}
