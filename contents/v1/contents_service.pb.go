// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: contents/v1/contents_service.proto

package contents_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_contents_v1_contents_service_proto protoreflect.FileDescriptor

var file_contents_v1_contents_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdf, 0x16,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x1d, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x47,
	0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0e,
	0x54, 0x61, 0x67, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x22,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x67, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x23,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x23, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x32, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x33, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x33, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x31, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x28, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x31, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x47,
	0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6b, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x47,
	0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x32, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x47, 0x65, 0x74, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6b, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x47, 0x65, 0x74, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x47,
	0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x32, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x33, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x33, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x31, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x31, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x32, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x32, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x33, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5c, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x31,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c,
	0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x33, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61,
	0x66, 0x69, 0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66,
	0x69, 0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72,
	0x61, 0x66, 0x69, 0x6b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72,
	0x61, 0x66, 0x69, 0x6b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x17, 0x49, 0x6e,
	0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67,
	0x72, 0x61, 0x66, 0x69, 0x6b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x5f, 0x0a, 0x10,
	0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a,
	0x10, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f,
	0x0a, 0x10, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x67, 0x72, 0x61, 0x66, 0x69,
	0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x76,
	0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x61, 0x74, 0x61, 0x6d, 0x61, 0x74, 0x61, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_contents_v1_contents_service_proto_goTypes = []interface{}{
	(*TagGetOneRequest)(nil),                // 0: contents.v1.TagGetOneRequest
	(*TagGetMultipleRequest)(nil),           // 1: contents.v1.TagGetMultipleRequest
	(*TagCreateRequest)(nil),                // 2: contents.v1.TagCreateRequest
	(*TagUpdateRequest)(nil),                // 3: contents.v1.TagUpdateRequest
	(*TagDeleteRequest)(nil),                // 4: contents.v1.TagDeleteRequest
	(*Category1GetOneRequest)(nil),          // 5: contents.v1.Category1GetOneRequest
	(*Category2GetOneRequest)(nil),          // 6: contents.v1.Category2GetOneRequest
	(*Category3GetOneRequest)(nil),          // 7: contents.v1.Category3GetOneRequest
	(*Category1GetMultipleRequest)(nil),     // 8: contents.v1.Category1GetMultipleRequest
	(*Category2GetMultipleRequest)(nil),     // 9: contents.v1.Category2GetMultipleRequest
	(*Category3GetMultipleRequest)(nil),     // 10: contents.v1.Category3GetMultipleRequest
	(*Category1CreateRequest)(nil),          // 11: contents.v1.Category1CreateRequest
	(*Category2CreateRequest)(nil),          // 12: contents.v1.Category2CreateRequest
	(*Category3CreateRequest)(nil),          // 13: contents.v1.Category3CreateRequest
	(*Category1UpdateRequest)(nil),          // 14: contents.v1.Category1UpdateRequest
	(*Category2UpdateRequest)(nil),          // 15: contents.v1.Category2UpdateRequest
	(*Category3UpdateRequest)(nil),          // 16: contents.v1.Category3UpdateRequest
	(*Category1DeleteRequest)(nil),          // 17: contents.v1.Category1DeleteRequest
	(*Category2DeleteRequest)(nil),          // 18: contents.v1.Category2DeleteRequest
	(*Category3DeleteRequest)(nil),          // 19: contents.v1.Category3DeleteRequest
	(*ArticleGetOneRequest)(nil),            // 20: contents.v1.ArticleGetOneRequest
	(*ArticleGetListRequest)(nil),           // 21: contents.v1.ArticleGetListRequest
	(*ArticleCreateRequest)(nil),            // 22: contents.v1.ArticleCreateRequest
	(*ArticleUpdateRequest)(nil),            // 23: contents.v1.ArticleUpdateRequest
	(*ArticleDeleteRequest)(nil),            // 24: contents.v1.ArticleDeleteRequest
	(*InfografikGetOneRequest)(nil),         // 25: contents.v1.InfografikGetOneRequest
	(*InfografikGetListRequest)(nil),        // 26: contents.v1.InfografikGetListRequest
	(*InfografikCreateRequest)(nil),         // 27: contents.v1.InfografikCreateRequest
	(*InfografikUpdateRequest)(nil),         // 28: contents.v1.InfografikUpdateRequest
	(*InfografikDeleteRequest)(nil),         // 29: contents.v1.InfografikDeleteRequest
	(*TagGetOneResponse)(nil),               // 30: contents.v1.TagGetOneResponse
	(*TagGetMultipleResponse)(nil),          // 31: contents.v1.TagGetMultipleResponse
	(*TagCreateResponse)(nil),               // 32: contents.v1.TagCreateResponse
	(*TagUpdateResponse)(nil),               // 33: contents.v1.TagUpdateResponse
	(*TagDeleteResponse)(nil),               // 34: contents.v1.TagDeleteResponse
	(*Category1GetOneResponse)(nil),         // 35: contents.v1.Category1GetOneResponse
	(*Category2GetOneResponse)(nil),         // 36: contents.v1.Category2GetOneResponse
	(*Category3GetOneResponse)(nil),         // 37: contents.v1.Category3GetOneResponse
	(*Category1GetMultipleResponse)(nil),    // 38: contents.v1.Category1GetMultipleResponse
	(*Category2GetMultipleResponse)(nil),    // 39: contents.v1.Category2GetMultipleResponse
	(*Category3GetMultipleResponse)(nil),    // 40: contents.v1.Category3GetMultipleResponse
	(*Category1CreateResponse)(nil),         // 41: contents.v1.Category1CreateResponse
	(*Category2CreateResponse)(nil),         // 42: contents.v1.Category2CreateResponse
	(*Category3CreateResponse)(nil),         // 43: contents.v1.Category3CreateResponse
	(*Category1UpdateResponse)(nil),         // 44: contents.v1.Category1UpdateResponse
	(*Category2UpdateResponse)(nil),         // 45: contents.v1.Category2UpdateResponse
	(*Category3UpdateResponse)(nil),         // 46: contents.v1.Category3UpdateResponse
	(*Category1DeleteResponse)(nil),         // 47: contents.v1.Category1DeleteResponse
	(*Category2DeleteResponse)(nil),         // 48: contents.v1.Category2DeleteResponse
	(*Category3DeleteResponse)(nil),         // 49: contents.v1.Category3DeleteResponse
	(*ArticleGetOneResponse)(nil),           // 50: contents.v1.ArticleGetOneResponse
	(*ArticleGetListResponse)(nil),          // 51: contents.v1.ArticleGetListResponse
	(*ArticleCreateResponse)(nil),           // 52: contents.v1.ArticleCreateResponse
	(*ArticleUpdateResponse)(nil),           // 53: contents.v1.ArticleUpdateResponse
	(*ArticleDeleteResponse)(nil),           // 54: contents.v1.ArticleDeleteResponse
	(*InfografikGetOneResponse)(nil),        // 55: contents.v1.InfografikGetOneResponse
	(*InfografikGetListResponse)(nil),       // 56: contents.v1.InfografikGetListResponse
	(*InfografikGetListStreamResponse)(nil), // 57: contents.v1.InfografikGetListStreamResponse
	(*InfografikCreateResponse)(nil),        // 58: contents.v1.InfografikCreateResponse
	(*InfografikUpdateResponse)(nil),        // 59: contents.v1.InfografikUpdateResponse
	(*InfografikDeleteResponse)(nil),        // 60: contents.v1.InfografikDeleteResponse
}
var file_contents_v1_contents_service_proto_depIdxs = []int32{
	0,  // 0: contents.v1.ContentService.TagGetOne:input_type -> contents.v1.TagGetOneRequest
	1,  // 1: contents.v1.ContentService.TagGetMultiple:input_type -> contents.v1.TagGetMultipleRequest
	2,  // 2: contents.v1.ContentService.TagCreate:input_type -> contents.v1.TagCreateRequest
	3,  // 3: contents.v1.ContentService.TagUpdate:input_type -> contents.v1.TagUpdateRequest
	4,  // 4: contents.v1.ContentService.TagDelete:input_type -> contents.v1.TagDeleteRequest
	5,  // 5: contents.v1.ContentService.Category1GetOne:input_type -> contents.v1.Category1GetOneRequest
	6,  // 6: contents.v1.ContentService.Category2GetOne:input_type -> contents.v1.Category2GetOneRequest
	7,  // 7: contents.v1.ContentService.Category3GetOne:input_type -> contents.v1.Category3GetOneRequest
	8,  // 8: contents.v1.ContentService.Category1GetMultiple:input_type -> contents.v1.Category1GetMultipleRequest
	9,  // 9: contents.v1.ContentService.Category2GetMultiple:input_type -> contents.v1.Category2GetMultipleRequest
	10, // 10: contents.v1.ContentService.Category3GetMultiple:input_type -> contents.v1.Category3GetMultipleRequest
	11, // 11: contents.v1.ContentService.Category1Create:input_type -> contents.v1.Category1CreateRequest
	12, // 12: contents.v1.ContentService.Category2Create:input_type -> contents.v1.Category2CreateRequest
	13, // 13: contents.v1.ContentService.Category3Create:input_type -> contents.v1.Category3CreateRequest
	14, // 14: contents.v1.ContentService.Category1Update:input_type -> contents.v1.Category1UpdateRequest
	15, // 15: contents.v1.ContentService.Category2Update:input_type -> contents.v1.Category2UpdateRequest
	16, // 16: contents.v1.ContentService.Category3Update:input_type -> contents.v1.Category3UpdateRequest
	17, // 17: contents.v1.ContentService.Category1Delete:input_type -> contents.v1.Category1DeleteRequest
	18, // 18: contents.v1.ContentService.Category2Delete:input_type -> contents.v1.Category2DeleteRequest
	19, // 19: contents.v1.ContentService.Category3Delete:input_type -> contents.v1.Category3DeleteRequest
	20, // 20: contents.v1.ContentService.ArticleGetOne:input_type -> contents.v1.ArticleGetOneRequest
	21, // 21: contents.v1.ContentService.ArticleGetList:input_type -> contents.v1.ArticleGetListRequest
	22, // 22: contents.v1.ContentService.ArticleCreate:input_type -> contents.v1.ArticleCreateRequest
	23, // 23: contents.v1.ContentService.ArticleUpdate:input_type -> contents.v1.ArticleUpdateRequest
	24, // 24: contents.v1.ContentService.ArticleDelete:input_type -> contents.v1.ArticleDeleteRequest
	25, // 25: contents.v1.ContentService.InfografikGetOne:input_type -> contents.v1.InfografikGetOneRequest
	26, // 26: contents.v1.ContentService.InfografikGetList:input_type -> contents.v1.InfografikGetListRequest
	26, // 27: contents.v1.ContentService.InfografikGetListStream:input_type -> contents.v1.InfografikGetListRequest
	27, // 28: contents.v1.ContentService.InfografikCreate:input_type -> contents.v1.InfografikCreateRequest
	28, // 29: contents.v1.ContentService.InfografikUpdate:input_type -> contents.v1.InfografikUpdateRequest
	29, // 30: contents.v1.ContentService.InfografikDelete:input_type -> contents.v1.InfografikDeleteRequest
	30, // 31: contents.v1.ContentService.TagGetOne:output_type -> contents.v1.TagGetOneResponse
	31, // 32: contents.v1.ContentService.TagGetMultiple:output_type -> contents.v1.TagGetMultipleResponse
	32, // 33: contents.v1.ContentService.TagCreate:output_type -> contents.v1.TagCreateResponse
	33, // 34: contents.v1.ContentService.TagUpdate:output_type -> contents.v1.TagUpdateResponse
	34, // 35: contents.v1.ContentService.TagDelete:output_type -> contents.v1.TagDeleteResponse
	35, // 36: contents.v1.ContentService.Category1GetOne:output_type -> contents.v1.Category1GetOneResponse
	36, // 37: contents.v1.ContentService.Category2GetOne:output_type -> contents.v1.Category2GetOneResponse
	37, // 38: contents.v1.ContentService.Category3GetOne:output_type -> contents.v1.Category3GetOneResponse
	38, // 39: contents.v1.ContentService.Category1GetMultiple:output_type -> contents.v1.Category1GetMultipleResponse
	39, // 40: contents.v1.ContentService.Category2GetMultiple:output_type -> contents.v1.Category2GetMultipleResponse
	40, // 41: contents.v1.ContentService.Category3GetMultiple:output_type -> contents.v1.Category3GetMultipleResponse
	41, // 42: contents.v1.ContentService.Category1Create:output_type -> contents.v1.Category1CreateResponse
	42, // 43: contents.v1.ContentService.Category2Create:output_type -> contents.v1.Category2CreateResponse
	43, // 44: contents.v1.ContentService.Category3Create:output_type -> contents.v1.Category3CreateResponse
	44, // 45: contents.v1.ContentService.Category1Update:output_type -> contents.v1.Category1UpdateResponse
	45, // 46: contents.v1.ContentService.Category2Update:output_type -> contents.v1.Category2UpdateResponse
	46, // 47: contents.v1.ContentService.Category3Update:output_type -> contents.v1.Category3UpdateResponse
	47, // 48: contents.v1.ContentService.Category1Delete:output_type -> contents.v1.Category1DeleteResponse
	48, // 49: contents.v1.ContentService.Category2Delete:output_type -> contents.v1.Category2DeleteResponse
	49, // 50: contents.v1.ContentService.Category3Delete:output_type -> contents.v1.Category3DeleteResponse
	50, // 51: contents.v1.ContentService.ArticleGetOne:output_type -> contents.v1.ArticleGetOneResponse
	51, // 52: contents.v1.ContentService.ArticleGetList:output_type -> contents.v1.ArticleGetListResponse
	52, // 53: contents.v1.ContentService.ArticleCreate:output_type -> contents.v1.ArticleCreateResponse
	53, // 54: contents.v1.ContentService.ArticleUpdate:output_type -> contents.v1.ArticleUpdateResponse
	54, // 55: contents.v1.ContentService.ArticleDelete:output_type -> contents.v1.ArticleDeleteResponse
	55, // 56: contents.v1.ContentService.InfografikGetOne:output_type -> contents.v1.InfografikGetOneResponse
	56, // 57: contents.v1.ContentService.InfografikGetList:output_type -> contents.v1.InfografikGetListResponse
	57, // 58: contents.v1.ContentService.InfografikGetListStream:output_type -> contents.v1.InfografikGetListStreamResponse
	58, // 59: contents.v1.ContentService.InfografikCreate:output_type -> contents.v1.InfografikCreateResponse
	59, // 60: contents.v1.ContentService.InfografikUpdate:output_type -> contents.v1.InfografikUpdateResponse
	60, // 61: contents.v1.ContentService.InfografikDelete:output_type -> contents.v1.InfografikDeleteResponse
	31, // [31:62] is the sub-list for method output_type
	0,  // [0:31] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_contents_v1_contents_service_proto_init() }
func file_contents_v1_contents_service_proto_init() {
	if File_contents_v1_contents_service_proto != nil {
		return
	}
	file_contents_v1_tag_proto_init()
	file_contents_v1_category_proto_init()
	file_contents_v1_article_proto_init()
	file_contents_v1_infografik_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contents_v1_contents_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contents_v1_contents_service_proto_goTypes,
		DependencyIndexes: file_contents_v1_contents_service_proto_depIdxs,
	}.Build()
	File_contents_v1_contents_service_proto = out.File
	file_contents_v1_contents_service_proto_rawDesc = nil
	file_contents_v1_contents_service_proto_goTypes = nil
	file_contents_v1_contents_service_proto_depIdxs = nil
}
