// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: orchestrator/v1/configurations/bitrate_config.proto

package configurationsv1

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

// Represents an audio codec.
type AudioCodec int32

const (
	AudioCodec_AAC  AudioCodec = 0 // Advanced Audio Coding
	AudioCodec_OPUS AudioCodec = 1 // Opus codec
	AudioCodec_AC3  AudioCodec = 2 // Dolby Digital (AC-3)
	AudioCodec_EAC3 AudioCodec = 3 // Dolby Digital Plus (E-AC-3)
)

// Enum value maps for AudioCodec.
var (
	AudioCodec_name = map[int32]string{
		0: "AAC",
		1: "OPUS",
		2: "AC3",
		3: "EAC3",
	}
	AudioCodec_value = map[string]int32{
		"AAC":  0,
		"OPUS": 1,
		"AC3":  2,
		"EAC3": 3,
	}
)

func (x AudioCodec) Enum() *AudioCodec {
	p := new(AudioCodec)
	*p = x
	return p
}

func (x AudioCodec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudioCodec) Descriptor() protoreflect.EnumDescriptor {
	return file_orchestrator_v1_configurations_bitrate_config_proto_enumTypes[0].Descriptor()
}

func (AudioCodec) Type() protoreflect.EnumType {
	return &file_orchestrator_v1_configurations_bitrate_config_proto_enumTypes[0]
}

func (x AudioCodec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudioCodec.Descriptor instead.
func (AudioCodec) EnumDescriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_bitrate_config_proto_rawDescGZIP(), []int{0}
}

// Represents a video codec.
type VideoCodec int32

const (
	VideoCodec_H264 VideoCodec = 0 // H.264/AVC
	VideoCodec_VP9  VideoCodec = 1 // VP9
	VideoCodec_AV1  VideoCodec = 2 // AV1
	VideoCodec_HEVC VideoCodec = 3 // HEVC/H.265
)

// Enum value maps for VideoCodec.
var (
	VideoCodec_name = map[int32]string{
		0: "H264",
		1: "VP9",
		2: "AV1",
		3: "HEVC",
	}
	VideoCodec_value = map[string]int32{
		"H264": 0,
		"VP9":  1,
		"AV1":  2,
		"HEVC": 3,
	}
)

func (x VideoCodec) Enum() *VideoCodec {
	p := new(VideoCodec)
	*p = x
	return p
}

func (x VideoCodec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VideoCodec) Descriptor() protoreflect.EnumDescriptor {
	return file_orchestrator_v1_configurations_bitrate_config_proto_enumTypes[1].Descriptor()
}

func (VideoCodec) Type() protoreflect.EnumType {
	return &file_orchestrator_v1_configurations_bitrate_config_proto_enumTypes[1]
}

func (x VideoCodec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VideoCodec.Descriptor instead.
func (VideoCodec) EnumDescriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_bitrate_config_proto_rawDescGZIP(), []int{1}
}

// Represents an audio channel layout.
type AudioChannelLayout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxChannels int32             `protobuf:"varint,1,opt,name=max_channels,json=maxChannels,proto3" json:"max_channels,omitempty"`                                                               // The maximum number of channels.
	Bitrates    map[string]string `protobuf:"bytes,2,rep,name=bitrates,proto3" json:"bitrates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of audio codecs to target bitrates (e.g., "500k", "7.5M").
}

func (x *AudioChannelLayout) Reset() {
	*x = AudioChannelLayout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioChannelLayout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioChannelLayout) ProtoMessage() {}

func (x *AudioChannelLayout) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioChannelLayout.ProtoReflect.Descriptor instead.
func (*AudioChannelLayout) Descriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_bitrate_config_proto_rawDescGZIP(), []int{0}
}

func (x *AudioChannelLayout) GetMaxChannels() int32 {
	if x != nil {
		return x.MaxChannels
	}
	return 0
}

func (x *AudioChannelLayout) GetBitrates() map[string]string {
	if x != nil {
		return x.Bitrates
	}
	return nil
}

// Represents a video resolution.
type VideoResolution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                                                                                 // Name of the resolution (e.g., "1080p").
	MaxWidth     int32             `protobuf:"varint,2,opt,name=max_width,json=maxWidth,proto3" json:"max_width,omitempty"`                                                                        // Maximum width in pixels.
	MaxHeight    int32             `protobuf:"varint,3,opt,name=max_height,json=maxHeight,proto3" json:"max_height,omitempty"`                                                                     // Maximum height in pixels.
	MaxFrameRate int32             `protobuf:"varint,4,opt,name=max_frame_rate,json=maxFrameRate,proto3" json:"max_frame_rate,omitempty"`                                                          // Maximum frame rate in frames per second (optional).
	Bitrates     map[string]string `protobuf:"bytes,5,rep,name=bitrates,proto3" json:"bitrates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of video codecs to target bitrates (e.g., "500k", "7.5M").
}

func (x *VideoResolution) Reset() {
	*x = VideoResolution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoResolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoResolution) ProtoMessage() {}

func (x *VideoResolution) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoResolution.ProtoReflect.Descriptor instead.
func (*VideoResolution) Descriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_bitrate_config_proto_rawDescGZIP(), []int{1}
}

func (x *VideoResolution) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VideoResolution) GetMaxWidth() int32 {
	if x != nil {
		return x.MaxWidth
	}
	return 0
}

func (x *VideoResolution) GetMaxHeight() int32 {
	if x != nil {
		return x.MaxHeight
	}
	return 0
}

func (x *VideoResolution) GetMaxFrameRate() int32 {
	if x != nil {
		return x.MaxFrameRate
	}
	return 0
}

func (x *VideoResolution) GetBitrates() map[string]string {
	if x != nil {
		return x.Bitrates
	}
	return nil
}

// Configuration for bitrates, including audio channel layouts and video resolutions.
type BitrateConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioChannelLayouts map[string]*AudioChannelLayout `protobuf:"bytes,1,rep,name=audio_channel_layouts,json=audioChannelLayouts,proto3" json:"audio_channel_layouts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of named audio channel layouts.
	VideoResolutions    map[string]*VideoResolution    `protobuf:"bytes,2,rep,name=video_resolutions,json=videoResolutions,proto3" json:"video_resolutions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`            // Map of named video resolutions.
}

func (x *BitrateConfiguration) Reset() {
	*x = BitrateConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BitrateConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BitrateConfiguration) ProtoMessage() {}

func (x *BitrateConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BitrateConfiguration.ProtoReflect.Descriptor instead.
func (*BitrateConfiguration) Descriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_bitrate_config_proto_rawDescGZIP(), []int{2}
}

func (x *BitrateConfiguration) GetAudioChannelLayouts() map[string]*AudioChannelLayout {
	if x != nil {
		return x.AudioChannelLayouts
	}
	return nil
}

func (x *BitrateConfiguration) GetVideoResolutions() map[string]*VideoResolution {
	if x != nil {
		return x.VideoResolutions
	}
	return nil
}

var File_orchestrator_v1_configurations_bitrate_config_proto protoreflect.FileDescriptor

var file_orchestrator_v1_configurations_bitrate_config_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xc5, 0x01, 0x0a, 0x12, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x12, 0x4f, 0x0a, 0x08, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x42, 0x69, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x69, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x92, 0x02, 0x0a, 0x0f, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x62, 0x69,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x42, 0x69, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd0, 0x03, 0x0a, 0x14, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x74,
	0x0a, 0x15, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x13, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x73, 0x12, 0x6a, 0x0a, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x6d, 0x0a, 0x18, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x67, 0x0a, 0x15, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x32, 0x0a, 0x0a, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x41, 0x43, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4f, 0x50, 0x55, 0x53, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x33,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x41, 0x43, 0x33, 0x10, 0x03, 0x2a, 0x32, 0x0a, 0x0a,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x32,
	0x36, 0x34, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x56, 0x50, 0x39, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x56, 0x31, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x56, 0x43, 0x10, 0x03,
	0x42, 0xec, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x42, 0x69, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x70, 0x72, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x58, 0x58, 0xaa, 0x02, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orchestrator_v1_configurations_bitrate_config_proto_rawDescOnce sync.Once
	file_orchestrator_v1_configurations_bitrate_config_proto_rawDescData = file_orchestrator_v1_configurations_bitrate_config_proto_rawDesc
)

func file_orchestrator_v1_configurations_bitrate_config_proto_rawDescGZIP() []byte {
	file_orchestrator_v1_configurations_bitrate_config_proto_rawDescOnce.Do(func() {
		file_orchestrator_v1_configurations_bitrate_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestrator_v1_configurations_bitrate_config_proto_rawDescData)
	})
	return file_orchestrator_v1_configurations_bitrate_config_proto_rawDescData
}

var file_orchestrator_v1_configurations_bitrate_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_orchestrator_v1_configurations_bitrate_config_proto_goTypes = []interface{}{
	(AudioCodec)(0),              // 0: configurations.v1.AudioCodec
	(VideoCodec)(0),              // 1: configurations.v1.VideoCodec
	(*AudioChannelLayout)(nil),   // 2: configurations.v1.AudioChannelLayout
	(*VideoResolution)(nil),      // 3: configurations.v1.VideoResolution
	(*BitrateConfiguration)(nil), // 4: configurations.v1.BitrateConfiguration
	nil,                          // 5: configurations.v1.AudioChannelLayout.BitratesEntry
	nil,                          // 6: configurations.v1.VideoResolution.BitratesEntry
	nil,                          // 7: configurations.v1.BitrateConfiguration.AudioChannelLayoutsEntry
	nil,                          // 8: configurations.v1.BitrateConfiguration.VideoResolutionsEntry
}
var file_orchestrator_v1_configurations_bitrate_config_proto_depIdxs = []int32{
	5, // 0: configurations.v1.AudioChannelLayout.bitrates:type_name -> configurations.v1.AudioChannelLayout.BitratesEntry
	6, // 1: configurations.v1.VideoResolution.bitrates:type_name -> configurations.v1.VideoResolution.BitratesEntry
	7, // 2: configurations.v1.BitrateConfiguration.audio_channel_layouts:type_name -> configurations.v1.BitrateConfiguration.AudioChannelLayoutsEntry
	8, // 3: configurations.v1.BitrateConfiguration.video_resolutions:type_name -> configurations.v1.BitrateConfiguration.VideoResolutionsEntry
	2, // 4: configurations.v1.BitrateConfiguration.AudioChannelLayoutsEntry.value:type_name -> configurations.v1.AudioChannelLayout
	3, // 5: configurations.v1.BitrateConfiguration.VideoResolutionsEntry.value:type_name -> configurations.v1.VideoResolution
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_orchestrator_v1_configurations_bitrate_config_proto_init() }
func file_orchestrator_v1_configurations_bitrate_config_proto_init() {
	if File_orchestrator_v1_configurations_bitrate_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioChannelLayout); i {
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
		file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoResolution); i {
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
		file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BitrateConfiguration); i {
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
			RawDescriptor: file_orchestrator_v1_configurations_bitrate_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orchestrator_v1_configurations_bitrate_config_proto_goTypes,
		DependencyIndexes: file_orchestrator_v1_configurations_bitrate_config_proto_depIdxs,
		EnumInfos:         file_orchestrator_v1_configurations_bitrate_config_proto_enumTypes,
		MessageInfos:      file_orchestrator_v1_configurations_bitrate_config_proto_msgTypes,
	}.Build()
	File_orchestrator_v1_configurations_bitrate_config_proto = out.File
	file_orchestrator_v1_configurations_bitrate_config_proto_rawDesc = nil
	file_orchestrator_v1_configurations_bitrate_config_proto_goTypes = nil
	file_orchestrator_v1_configurations_bitrate_config_proto_depIdxs = nil
}
