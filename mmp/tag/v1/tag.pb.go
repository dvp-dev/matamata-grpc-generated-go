// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: mmp/tag/v1/tag.proto

package tag

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

type GetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetOneRequest) Reset() {
	*x = GetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmp_tag_v1_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneRequest) ProtoMessage() {}

func (x *GetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmp_tag_v1_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneRequest.ProtoReflect.Descriptor instead.
func (*GetOneRequest) Descriptor() ([]byte, []int) {
	return file_mmp_tag_v1_tag_proto_rawDescGZIP(), []int{0}
}

func (x *GetOneRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetOneResponse) Reset() {
	*x = GetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmp_tag_v1_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneResponse) ProtoMessage() {}

func (x *GetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmp_tag_v1_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneResponse.ProtoReflect.Descriptor instead.
func (*GetOneResponse) Descriptor() ([]byte, []int) {
	return file_mmp_tag_v1_tag_proto_rawDescGZIP(), []int{1}
}

func (x *GetOneResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetOneResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_mmp_tag_v1_tag_proto protoreflect.FileDescriptor

var file_mmp_tag_v1_tag_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x67, 0x2e,
	0x76, 0x31, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x4d, 0x0a,
	0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6d, 0x6d, 0x70, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x5a, 0x4e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x76, 0x70, 0x2d, 0x64,
	0x65, 0x76, 0x2f, 0x6d, 0x61, 0x74, 0x61, 0x6d, 0x61, 0x74, 0x61, 0x2d, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2d, 0x43, 0x6f, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x75, 0x6d, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f,
	0x6d, 0x6d, 0x70, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmp_tag_v1_tag_proto_rawDescOnce sync.Once
	file_mmp_tag_v1_tag_proto_rawDescData = file_mmp_tag_v1_tag_proto_rawDesc
)

func file_mmp_tag_v1_tag_proto_rawDescGZIP() []byte {
	file_mmp_tag_v1_tag_proto_rawDescOnce.Do(func() {
		file_mmp_tag_v1_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmp_tag_v1_tag_proto_rawDescData)
	})
	return file_mmp_tag_v1_tag_proto_rawDescData
}

var file_mmp_tag_v1_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mmp_tag_v1_tag_proto_goTypes = []interface{}{
	(*GetOneRequest)(nil),  // 0: mmp.tag.v1.GetOneRequest
	(*GetOneResponse)(nil), // 1: mmp.tag.v1.GetOneResponse
}
var file_mmp_tag_v1_tag_proto_depIdxs = []int32{
	0, // 0: mmp.tag.v1.TagService.GetOne:input_type -> mmp.tag.v1.GetOneRequest
	1, // 1: mmp.tag.v1.TagService.GetOne:output_type -> mmp.tag.v1.GetOneResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mmp_tag_v1_tag_proto_init() }
func file_mmp_tag_v1_tag_proto_init() {
	if File_mmp_tag_v1_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmp_tag_v1_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneRequest); i {
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
		file_mmp_tag_v1_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneResponse); i {
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
			RawDescriptor: file_mmp_tag_v1_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mmp_tag_v1_tag_proto_goTypes,
		DependencyIndexes: file_mmp_tag_v1_tag_proto_depIdxs,
		MessageInfos:      file_mmp_tag_v1_tag_proto_msgTypes,
	}.Build()
	File_mmp_tag_v1_tag_proto = out.File
	file_mmp_tag_v1_tag_proto_rawDesc = nil
	file_mmp_tag_v1_tag_proto_goTypes = nil
	file_mmp_tag_v1_tag_proto_depIdxs = nil
}
