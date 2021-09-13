// Copyright 2021 The Rode Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/v1alpha1/new-collector-template.proto

package v1alpha1

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

type CreateEventOccurrenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateEventOccurrenceRequest) Reset() {
	*x = CreateEventOccurrenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_new_collector_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventOccurrenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventOccurrenceRequest) ProtoMessage() {}

func (x *CreateEventOccurrenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_new_collector_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventOccurrenceRequest.ProtoReflect.Descriptor instead.
func (*CreateEventOccurrenceRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_new_collector_template_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEventOccurrenceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateEventOccurrenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateEventOccurrenceResponse) Reset() {
	*x = CreateEventOccurrenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_new_collector_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventOccurrenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventOccurrenceResponse) ProtoMessage() {}

func (x *CreateEventOccurrenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_new_collector_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventOccurrenceResponse.ProtoReflect.Descriptor instead.
func (*CreateEventOccurrenceResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_new_collector_template_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventOccurrenceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_v1alpha1_new_collector_template_proto protoreflect.FileDescriptor

var file_proto_v1alpha1_new_collector_template_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x6e, 0x65, 0x77, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6e,
	0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x1c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2f, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f,
	0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xcc, 0x01, 0x0a, 0x14, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0xb3, 0x01, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x3d, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x65, 0x77, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_v1alpha1_new_collector_template_proto_rawDescOnce sync.Once
	file_proto_v1alpha1_new_collector_template_proto_rawDescData = file_proto_v1alpha1_new_collector_template_proto_rawDesc
)

func file_proto_v1alpha1_new_collector_template_proto_rawDescGZIP() []byte {
	file_proto_v1alpha1_new_collector_template_proto_rawDescOnce.Do(func() {
		file_proto_v1alpha1_new_collector_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1alpha1_new_collector_template_proto_rawDescData)
	})
	return file_proto_v1alpha1_new_collector_template_proto_rawDescData
}

var file_proto_v1alpha1_new_collector_template_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v1alpha1_new_collector_template_proto_goTypes = []interface{}{
	(*CreateEventOccurrenceRequest)(nil),  // 0: new_collector_template.v1alpha1.CreateEventOccurrenceRequest
	(*CreateEventOccurrenceResponse)(nil), // 1: new_collector_template.v1alpha1.CreateEventOccurrenceResponse
}
var file_proto_v1alpha1_new_collector_template_proto_depIdxs = []int32{
	0, // 0: new_collector_template.v1alpha1.NewCollectorTemplate.CreateEventOccurrence:input_type -> new_collector_template.v1alpha1.CreateEventOccurrenceRequest
	1, // 1: new_collector_template.v1alpha1.NewCollectorTemplate.CreateEventOccurrence:output_type -> new_collector_template.v1alpha1.CreateEventOccurrenceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1alpha1_new_collector_template_proto_init() }
func file_proto_v1alpha1_new_collector_template_proto_init() {
	if File_proto_v1alpha1_new_collector_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1alpha1_new_collector_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventOccurrenceRequest); i {
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
		file_proto_v1alpha1_new_collector_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventOccurrenceResponse); i {
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
			RawDescriptor: file_proto_v1alpha1_new_collector_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1alpha1_new_collector_template_proto_goTypes,
		DependencyIndexes: file_proto_v1alpha1_new_collector_template_proto_depIdxs,
		MessageInfos:      file_proto_v1alpha1_new_collector_template_proto_msgTypes,
	}.Build()
	File_proto_v1alpha1_new_collector_template_proto = out.File
	file_proto_v1alpha1_new_collector_template_proto_rawDesc = nil
	file_proto_v1alpha1_new_collector_template_proto_goTypes = nil
	file_proto_v1alpha1_new_collector_template_proto_depIdxs = nil
}
