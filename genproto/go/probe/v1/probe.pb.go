// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: probe/v1/probe.proto

package probe

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probe_v1_probe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_probe_v1_probe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_probe_v1_probe_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Metadata) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *GetMetadataRequest) Reset() {
	*x = GetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probe_v1_probe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetadataRequest) ProtoMessage() {}

func (x *GetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_probe_v1_probe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_probe_v1_probe_proto_rawDescGZIP(), []int{1}
}

func (x *GetMetadataRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type GetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *GetMetadataResponse) Reset() {
	*x = GetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probe_v1_probe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetadataResponse) ProtoMessage() {}

func (x *GetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_probe_v1_probe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_probe_v1_probe_proto_rawDescGZIP(), []int{2}
}

func (x *GetMetadataResponse) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_probe_v1_probe_proto protoreflect.FileDescriptor

var file_probe_v1_probe_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x83, 0x01, 0x0a, 0x0c,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21,
	0x3a, 0x01, 0x2a, 0x62, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x12, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x96, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x0a, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x70, 0x72, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0xa2, 0x02, 0x03, 0x50, 0x58,
	0x58, 0xaa, 0x02, 0x08, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x09, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_probe_v1_probe_proto_rawDescOnce sync.Once
	file_probe_v1_probe_proto_rawDescData = file_probe_v1_probe_proto_rawDesc
)

func file_probe_v1_probe_proto_rawDescGZIP() []byte {
	file_probe_v1_probe_proto_rawDescOnce.Do(func() {
		file_probe_v1_probe_proto_rawDescData = protoimpl.X.CompressGZIP(file_probe_v1_probe_proto_rawDescData)
	})
	return file_probe_v1_probe_proto_rawDescData
}

var file_probe_v1_probe_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_probe_v1_probe_proto_goTypes = []interface{}{
	(*Metadata)(nil),            // 0: probe.v1.Metadata
	(*GetMetadataRequest)(nil),  // 1: probe.v1.GetMetadataRequest
	(*GetMetadataResponse)(nil), // 2: probe.v1.GetMetadataResponse
}
var file_probe_v1_probe_proto_depIdxs = []int32{
	0, // 0: probe.v1.GetMetadataResponse.metadata:type_name -> probe.v1.Metadata
	1, // 1: probe.v1.ProbeService.GetMetadata:input_type -> probe.v1.GetMetadataRequest
	2, // 2: probe.v1.ProbeService.GetMetadata:output_type -> probe.v1.GetMetadataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_probe_v1_probe_proto_init() }
func file_probe_v1_probe_proto_init() {
	if File_probe_v1_probe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_probe_v1_probe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_probe_v1_probe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetadataRequest); i {
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
		file_probe_v1_probe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetadataResponse); i {
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
			RawDescriptor: file_probe_v1_probe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_probe_v1_probe_proto_goTypes,
		DependencyIndexes: file_probe_v1_probe_proto_depIdxs,
		MessageInfos:      file_probe_v1_probe_proto_msgTypes,
	}.Build()
	File_probe_v1_probe_proto = out.File
	file_probe_v1_probe_proto_rawDesc = nil
	file_probe_v1_probe_proto_goTypes = nil
	file_probe_v1_probe_proto_depIdxs = nil
}
