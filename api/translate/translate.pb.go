// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: translate/translate.proto

package translate

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

type TranslateTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *TranslateTextRequest) Reset() {
	*x = TranslateTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_translate_translate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateTextRequest) ProtoMessage() {}

func (x *TranslateTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_translate_translate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateTextRequest.ProtoReflect.Descriptor instead.
func (*TranslateTextRequest) Descriptor() ([]byte, []int) {
	return file_translate_translate_proto_rawDescGZIP(), []int{0}
}

func (x *TranslateTextRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type TranslateTextReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *TranslateTextReply) Reset() {
	*x = TranslateTextReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_translate_translate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateTextReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateTextReply) ProtoMessage() {}

func (x *TranslateTextReply) ProtoReflect() protoreflect.Message {
	mi := &file_translate_translate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateTextReply.ProtoReflect.Descriptor instead.
func (*TranslateTextReply) Descriptor() ([]byte, []int) {
	return file_translate_translate_proto_rawDescGZIP(), []int{1}
}

func (x *TranslateTextReply) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_translate_translate_proto protoreflect.FileDescriptor

var file_translate_translate_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x22, 0x2c, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0x68, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x5b, 0x0a,
	0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x1e, 0x5a, 0x1c, 0x77, 0x6c,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2f,
	0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_translate_translate_proto_rawDescOnce sync.Once
	file_translate_translate_proto_rawDescData = file_translate_translate_proto_rawDesc
)

func file_translate_translate_proto_rawDescGZIP() []byte {
	file_translate_translate_proto_rawDescOnce.Do(func() {
		file_translate_translate_proto_rawDescData = protoimpl.X.CompressGZIP(file_translate_translate_proto_rawDescData)
	})
	return file_translate_translate_proto_rawDescData
}

var file_translate_translate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_translate_translate_proto_goTypes = []interface{}{
	(*TranslateTextRequest)(nil), // 0: translate.TranslateTextRequest
	(*TranslateTextReply)(nil),   // 1: translate.TranslateTextReply
}
var file_translate_translate_proto_depIdxs = []int32{
	0, // 0: translate.Translate.Text:input_type -> translate.TranslateTextRequest
	1, // 1: translate.Translate.Text:output_type -> translate.TranslateTextReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_translate_translate_proto_init() }
func file_translate_translate_proto_init() {
	if File_translate_translate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_translate_translate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateTextRequest); i {
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
		file_translate_translate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateTextReply); i {
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
			RawDescriptor: file_translate_translate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_translate_translate_proto_goTypes,
		DependencyIndexes: file_translate_translate_proto_depIdxs,
		MessageInfos:      file_translate_translate_proto_msgTypes,
	}.Build()
	File_translate_translate_proto = out.File
	file_translate_translate_proto_rawDesc = nil
	file_translate_translate_proto_goTypes = nil
	file_translate_translate_proto_depIdxs = nil
}
