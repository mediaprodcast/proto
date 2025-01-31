// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: orchestrator/v1/configurations/input_config.proto

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

// ----------------------------------------------------input config---------------------------------------
// Represents the type of an input stream.
type InputType int32

const (
	// A track from a file. Usable only with VOD.
	InputType_FILE InputType = 0
	// A track from a file, looped forever by FFmpeg.
	// Usable only with live.
	// Does not support media_type of 'text'.
	InputType_LOOPED_FILE InputType = 1
	// A webcam device.
	// Usable only with live.
	// The device path should be given in the name field.
	// For example, on Linux, this might be /dev/video0.
	// Only supports media_type of 'video'.
	InputType_WEBCAM InputType = 2
	// A microphone device.
	// Usable only with live.
	// The device path should given in the name field.
	// For example, on Linux, this might be "default".
	// Only supports media_type of 'audio'.
	InputType_MICROPHONE InputType = 3
	// An external command that generates a stream of audio or video.
	// The command should be given in the name field, using shell quoting rules.
	// The command should send its generated output to the path in the environment variable $SHAKA_STREAMER_EXTERNAL_COMMAND_OUTPUT, which Shaka Streamer set to the path to the output pipe.
	// May require the user of extra_input_args if FFmpeg can't guess the format or framerate.
	// Does not support media_type of 'text'.
	InputType_EXTERNAL_COMMAND InputType = 4
)

// Enum value maps for InputType.
var (
	InputType_name = map[int32]string{
		0: "FILE",
		1: "LOOPED_FILE",
		2: "WEBCAM",
		3: "MICROPHONE",
		4: "EXTERNAL_COMMAND",
	}
	InputType_value = map[string]int32{
		"FILE":             0,
		"LOOPED_FILE":      1,
		"WEBCAM":           2,
		"MICROPHONE":       3,
		"EXTERNAL_COMMAND": 4,
	}
)

func (x InputType) Enum() *InputType {
	p := new(InputType)
	*p = x
	return p
}

func (x InputType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InputType) Descriptor() protoreflect.EnumDescriptor {
	return file_orchestrator_v1_configurations_input_config_proto_enumTypes[0].Descriptor()
}

func (InputType) Type() protoreflect.EnumType {
	return &file_orchestrator_v1_configurations_input_config_proto_enumTypes[0]
}

func (x InputType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InputType.Descriptor instead.
func (InputType) EnumDescriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_input_config_proto_rawDescGZIP(), []int{0}
}

// Represents the media type of an input stream.
type MediaType int32

const (
	MediaType_AUDIO MediaType = 0
	MediaType_VIDEO MediaType = 1
	MediaType_TEXT  MediaType = 2
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0: "AUDIO",
		1: "VIDEO",
		2: "TEXT",
	}
	MediaType_value = map[string]int32{
		"AUDIO": 0,
		"VIDEO": 1,
		"TEXT":  2,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_orchestrator_v1_configurations_input_config_proto_enumTypes[1].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_orchestrator_v1_configurations_input_config_proto_enumTypes[1]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_input_config_proto_rawDescGZIP(), []int{1}
}

// An object representing a single input stream to Shaka Streamer.
type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the input.
	InputType InputType `protobuf:"varint,1,opt,name=input_type,json=inputType,proto3,enum=configurations.v1.InputType" json:"input_type,omitempty"`
	// Name of the input.
	// With input_type set to 'file', this is a path to a file name.
	// With input_type set to 'looped_file', this is a path to a file name to be looped indefinitely in FFmpeg.
	// With input_type set to 'webcam', this is which webcam.
	// On Linux, this is a path to the device node for the webcam, such as '/dev/video0'.
	// On macOS, this is a device name, such as 'default'.
	// With input_type set to 'external_command', this is an external command that generates a stream of audio or video. The command will be parsed using shell quoting rules. The command should send its generated output to the path in the environment variable $SHAKA_STREAMER_EXTERNAL_COMMAND_OUTPUT, which Shaka Streamer set to the path to the output pipe.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Extra input arguments needed by FFmpeg to understand the input.
	// This allows you to take inputs that cannot be understand or detected automatically by FFmpeg.
	// This string will be parsed using shell quoting rules.
	ExtraInputArgs *string `protobuf:"bytes,3,opt,name=extra_input_args,json=extraInputArgs,proto3,oneof" json:"extra_input_args,omitempty"`
	// The media type of the input stream.
	MediaType MediaType `protobuf:"varint,4,opt,name=media_type,json=mediaType,proto3,enum=configurations.v1.MediaType" json:"media_type,omitempty"`
	// The frame rate of the input stream, in frames per second.
	// Only valid for media_type of 'video'.
	// Can be auto-detected for some input types, but may be required for others.
	// For example, required for input_type of 'external_command'.
	FrameRate *float64 `protobuf:"fixed64,5,opt,name=frame_rate,json=frameRate,proto3,oneof" json:"frame_rate,omitempty"`
	// The name of the input resolution (1080p, etc).
	// Only valid for media_type of 'video'.
	// Can be auto-detected for some input types, but may be required for others.
	// For example, required for input_type of 'external_command'.
	Resolution *string `protobuf:"bytes,6,opt,name=resolution,proto3,oneof" json:"resolution,omitempty"`
	// The name of the input channel layout (stereo, surround, etc).
	ChannelLayout *string `protobuf:"bytes,7,opt,name=channel_layout,json=channelLayout,proto3,oneof" json:"channel_layout,omitempty"`
	// The track number of the input.
	// The track number is specific to the media_type. For example, if there is one video track and two audio tracks, media_type of 'audio' and track_num of '0' indicates the first audio track, not the first track overall in that file.
	// If unspecified, track_num will default to 0, meaning the first track matching the media_type field will be used.
	TrackNum int32 `protobuf:"varint,8,opt,name=track_num,json=trackNum,proto3" json:"track_num,omitempty"`
	// True if the input video is interlaced.
	// Only valid for media_type of 'video'.
	// If true, the video will be deinterlaced during transcoding.
	// Can be auto-detected for some input types, but may be default to False for others.
	// For example, an input_type of 'external_command', it will default to false.
	IsInterlaced *bool `protobuf:"varint,9,opt,name=is_interlaced,json=isInterlaced,proto3,oneof" json:"is_interlaced,omitempty"`
	// The language of an audio or text stream.
	// With input_type set to 'file' or 'looped_file', this will be auto-detected.
	// Otherwise, it will default to 'und' (undetermined).
	Language *string `protobuf:"bytes,10,opt,name=language,proto3,oneof" json:"language,omitempty"`
	// The start time of the slice of the input to use.
	// Only valid for VOD and with input_type set to 'file'.
	// Not supported with media_type of 'text'.
	StartTime *string `protobuf:"bytes,11,opt,name=start_time,json=startTime,proto3,oneof" json:"start_time,omitempty"`
	// The end time of the slice of the input to use.
	// Only valid for VOD and with input_type set to 'file'.
	// Not supported with media_type of 'text'.
	EndTime *string `protobuf:"bytes,12,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
	// Optional value for a custom DRM label, which defines the encryption key applied to the stream. If not provided, the DRM label is derived from stream type (video, audio), resolutions, etc. Note that it is case sensitive. Applies to 'raw' encryption_mode only.
	DrmLabel *string `protobuf:"bytes,13,opt,name=drm_label,json=drmLabel,proto3,oneof" json:"drm_label,omitempty"`
	// If set, no encryption of the stream will be made
	SkipEncryption *int32 `protobuf:"varint,14,opt,name=skip_encryption,json=skipEncryption,proto3,oneof" json:"skip_encryption,omitempty"`
	// A list of FFmpeg filter strings to add to the transcoding of this input.
	// Each filter is a single string. For example, 'pad=1280:720:20:20'.
	// Not supported with media_type of 'text'.
	Filters []string `protobuf:"bytes,15,rep,name=filters,proto3" json:"filters,omitempty"`
	// number of frames
	Frames *int32 `protobuf:"varint,16,opt,name=frames,proto3,oneof" json:"frames,omitempty"`
	// input duration
	Duration *float64 `protobuf:"fixed64,17,opt,name=duration,proto3,oneof" json:"duration,omitempty"`
	// Indicates that the subtitle is forced.
	// Only supported with media_type of 'text'.
	ForcedSubtitle *bool `protobuf:"varint,18,opt,name=forced_subtitle,json=forcedSubtitle,proto3,oneof" json:"forced_subtitle,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_v1_configurations_input_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_v1_configurations_input_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_input_config_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetInputType() InputType {
	if x != nil {
		return x.InputType
	}
	return InputType_FILE
}

func (x *Input) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Input) GetExtraInputArgs() string {
	if x != nil && x.ExtraInputArgs != nil {
		return *x.ExtraInputArgs
	}
	return ""
}

func (x *Input) GetMediaType() MediaType {
	if x != nil {
		return x.MediaType
	}
	return MediaType_AUDIO
}

func (x *Input) GetFrameRate() float64 {
	if x != nil && x.FrameRate != nil {
		return *x.FrameRate
	}
	return 0
}

func (x *Input) GetResolution() string {
	if x != nil && x.Resolution != nil {
		return *x.Resolution
	}
	return ""
}

func (x *Input) GetChannelLayout() string {
	if x != nil && x.ChannelLayout != nil {
		return *x.ChannelLayout
	}
	return ""
}

func (x *Input) GetTrackNum() int32 {
	if x != nil {
		return x.TrackNum
	}
	return 0
}

func (x *Input) GetIsInterlaced() bool {
	if x != nil && x.IsInterlaced != nil {
		return *x.IsInterlaced
	}
	return false
}

func (x *Input) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *Input) GetStartTime() string {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return ""
}

func (x *Input) GetEndTime() string {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return ""
}

func (x *Input) GetDrmLabel() string {
	if x != nil && x.DrmLabel != nil {
		return *x.DrmLabel
	}
	return ""
}

func (x *Input) GetSkipEncryption() int32 {
	if x != nil && x.SkipEncryption != nil {
		return *x.SkipEncryption
	}
	return 0
}

func (x *Input) GetFilters() []string {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Input) GetFrames() int32 {
	if x != nil && x.Frames != nil {
		return *x.Frames
	}
	return 0
}

func (x *Input) GetDuration() float64 {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return 0
}

func (x *Input) GetForcedSubtitle() bool {
	if x != nil && x.ForcedSubtitle != nil {
		return *x.ForcedSubtitle
	}
	return false
}

// * An object representing a single period in a multiperiod inputs list.
type SinglePeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs []*Input `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
}

func (x *SinglePeriod) Reset() {
	*x = SinglePeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_v1_configurations_input_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SinglePeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SinglePeriod) ProtoMessage() {}

func (x *SinglePeriod) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_v1_configurations_input_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SinglePeriod.ProtoReflect.Descriptor instead.
func (*SinglePeriod) Descriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_input_config_proto_rawDescGZIP(), []int{1}
}

func (x *SinglePeriod) GetInputs() []*Input {
	if x != nil {
		return x.Inputs
	}
	return nil
}

// An object representing the entire input config
type InputConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of SinglePeriod objects
	MultiperiodInputsList []*SinglePeriod `protobuf:"bytes,1,rep,name=multiperiod_inputs_list,json=multiperiodInputsList,proto3" json:"multiperiod_inputs_list,omitempty"`
	// A list of Input objects
	Inputs []*Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
}

func (x *InputConfiguration) Reset() {
	*x = InputConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_v1_configurations_input_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputConfiguration) ProtoMessage() {}

func (x *InputConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_v1_configurations_input_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputConfiguration.ProtoReflect.Descriptor instead.
func (*InputConfiguration) Descriptor() ([]byte, []int) {
	return file_orchestrator_v1_configurations_input_config_proto_rawDescGZIP(), []int{2}
}

func (x *InputConfiguration) GetMultiperiodInputsList() []*SinglePeriod {
	if x != nil {
		return x.MultiperiodInputsList
	}
	return nil
}

func (x *InputConfiguration) GetInputs() []*Input {
	if x != nil {
		return x.Inputs
	}
	return nil
}

var File_orchestrator_v1_configurations_input_config_proto protoreflect.FileDescriptor

var file_orchestrator_v1_configurations_input_config_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x8a, 0x07, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x3b, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x72, 0x67, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x3b, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x01, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12,
	0x28, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6c, 0x61, 0x63, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x0c, 0x69, 0x73, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x07, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x09, 0x64, 0x72, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x08, 0x52, 0x08, 0x64, 0x72, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x09, 0x52, 0x0e, 0x73, 0x6b, 0x69,
	0x70, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0a, 0x52, 0x06, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x48, 0x0b, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64,
	0x5f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x0c, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x69,
	0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x72, 0x6d, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x0c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x06, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x17,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x15,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2a, 0x58, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x4c, 0x4f, 0x4f, 0x50, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x57, 0x45, 0x42, 0x43, 0x41, 0x4d, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4d,
	0x49, 0x43, 0x52, 0x4f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10,
	0x04, 0x2a, 0x2b, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44,
	0x45, 0x4f, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x42, 0xea,
	0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72,
	0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02,
	0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_orchestrator_v1_configurations_input_config_proto_rawDescOnce sync.Once
	file_orchestrator_v1_configurations_input_config_proto_rawDescData = file_orchestrator_v1_configurations_input_config_proto_rawDesc
)

func file_orchestrator_v1_configurations_input_config_proto_rawDescGZIP() []byte {
	file_orchestrator_v1_configurations_input_config_proto_rawDescOnce.Do(func() {
		file_orchestrator_v1_configurations_input_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestrator_v1_configurations_input_config_proto_rawDescData)
	})
	return file_orchestrator_v1_configurations_input_config_proto_rawDescData
}

var file_orchestrator_v1_configurations_input_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_orchestrator_v1_configurations_input_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orchestrator_v1_configurations_input_config_proto_goTypes = []interface{}{
	(InputType)(0),             // 0: configurations.v1.InputType
	(MediaType)(0),             // 1: configurations.v1.MediaType
	(*Input)(nil),              // 2: configurations.v1.Input
	(*SinglePeriod)(nil),       // 3: configurations.v1.SinglePeriod
	(*InputConfiguration)(nil), // 4: configurations.v1.InputConfiguration
}
var file_orchestrator_v1_configurations_input_config_proto_depIdxs = []int32{
	0, // 0: configurations.v1.Input.input_type:type_name -> configurations.v1.InputType
	1, // 1: configurations.v1.Input.media_type:type_name -> configurations.v1.MediaType
	2, // 2: configurations.v1.SinglePeriod.inputs:type_name -> configurations.v1.Input
	3, // 3: configurations.v1.InputConfiguration.multiperiod_inputs_list:type_name -> configurations.v1.SinglePeriod
	2, // 4: configurations.v1.InputConfiguration.inputs:type_name -> configurations.v1.Input
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orchestrator_v1_configurations_input_config_proto_init() }
func file_orchestrator_v1_configurations_input_config_proto_init() {
	if File_orchestrator_v1_configurations_input_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orchestrator_v1_configurations_input_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_orchestrator_v1_configurations_input_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SinglePeriod); i {
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
		file_orchestrator_v1_configurations_input_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputConfiguration); i {
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
	file_orchestrator_v1_configurations_input_config_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orchestrator_v1_configurations_input_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orchestrator_v1_configurations_input_config_proto_goTypes,
		DependencyIndexes: file_orchestrator_v1_configurations_input_config_proto_depIdxs,
		EnumInfos:         file_orchestrator_v1_configurations_input_config_proto_enumTypes,
		MessageInfos:      file_orchestrator_v1_configurations_input_config_proto_msgTypes,
	}.Build()
	File_orchestrator_v1_configurations_input_config_proto = out.File
	file_orchestrator_v1_configurations_input_config_proto_rawDesc = nil
	file_orchestrator_v1_configurations_input_config_proto_goTypes = nil
	file_orchestrator_v1_configurations_input_config_proto_depIdxs = nil
}
