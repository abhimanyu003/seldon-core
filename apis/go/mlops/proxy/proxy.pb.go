/*
Copyright 2022 Seldon Technologies Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: mlops/proxy/proxy.proto

package proxy

import (
	scheduler "github.com/seldonio/seldon-core/apis/go/v2/mlops/scheduler"
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

type LoadModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repeated string agents = 1;
	Request *scheduler.LoadModelRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Version uint32                      `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *LoadModelRequest) Reset() {
	*x = LoadModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_proxy_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadModelRequest) ProtoMessage() {}

func (x *LoadModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_proxy_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadModelRequest.ProtoReflect.Descriptor instead.
func (*LoadModelRequest) Descriptor() ([]byte, []int) {
	return file_mlops_proxy_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *LoadModelRequest) GetRequest() *scheduler.LoadModelRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *LoadModelRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type LoadModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadModelResponse) Reset() {
	*x = LoadModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_proxy_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadModelResponse) ProtoMessage() {}

func (x *LoadModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_proxy_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadModelResponse.ProtoReflect.Descriptor instead.
func (*LoadModelResponse) Descriptor() ([]byte, []int) {
	return file_mlops_proxy_proxy_proto_rawDescGZIP(), []int{1}
}

type UnloadModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repeated string agents = 1;
	Model   *scheduler.ModelReference `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Version uint32                    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *UnloadModelRequest) Reset() {
	*x = UnloadModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_proxy_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnloadModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnloadModelRequest) ProtoMessage() {}

func (x *UnloadModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_proxy_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnloadModelRequest.ProtoReflect.Descriptor instead.
func (*UnloadModelRequest) Descriptor() ([]byte, []int) {
	return file_mlops_proxy_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *UnloadModelRequest) GetModel() *scheduler.ModelReference {
	if x != nil {
		return x.Model
	}
	return nil
}

func (x *UnloadModelRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type UnloadModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnloadModelResponse) Reset() {
	*x = UnloadModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_proxy_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnloadModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnloadModelResponse) ProtoMessage() {}

func (x *UnloadModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_proxy_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnloadModelResponse.ProtoReflect.Descriptor instead.
func (*UnloadModelResponse) Descriptor() ([]byte, []int) {
	return file_mlops_proxy_proxy_proto_rawDescGZIP(), []int{3}
}

var File_mlops_proxy_proxy_proto protoreflect.FileDescriptor

var file_mlops_proxy_proxy_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x65, 0x6c, 0x64, 0x6f,
	0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x1f, 0x6d,
	0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70,
	0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f,
	0x70, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x61,
	0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x13, 0x0a, 0x11, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6c, 0x0a, 0x12, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x65, 0x6c,
	0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xce, 0x01, 0x0a, 0x0e, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x5a, 0x0a,
	0x09, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x6c,
	0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0b, 0x55, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f,
	0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x55, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e,
	0x69, 0x6f, 0x2f, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mlops_proxy_proxy_proto_rawDescOnce sync.Once
	file_mlops_proxy_proxy_proto_rawDescData = file_mlops_proxy_proxy_proto_rawDesc
)

func file_mlops_proxy_proxy_proto_rawDescGZIP() []byte {
	file_mlops_proxy_proxy_proto_rawDescOnce.Do(func() {
		file_mlops_proxy_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_mlops_proxy_proxy_proto_rawDescData)
	})
	return file_mlops_proxy_proxy_proto_rawDescData
}

var file_mlops_proxy_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mlops_proxy_proxy_proto_goTypes = []interface{}{
	(*LoadModelRequest)(nil),           // 0: seldon.mlops.proxy.LoadModelRequest
	(*LoadModelResponse)(nil),          // 1: seldon.mlops.proxy.LoadModelResponse
	(*UnloadModelRequest)(nil),         // 2: seldon.mlops.proxy.UnloadModelRequest
	(*UnloadModelResponse)(nil),        // 3: seldon.mlops.proxy.UnloadModelResponse
	(*scheduler.LoadModelRequest)(nil), // 4: seldon.mlops.scheduler.LoadModelRequest
	(*scheduler.ModelReference)(nil),   // 5: seldon.mlops.scheduler.ModelReference
}
var file_mlops_proxy_proxy_proto_depIdxs = []int32{
	4, // 0: seldon.mlops.proxy.LoadModelRequest.request:type_name -> seldon.mlops.scheduler.LoadModelRequest
	5, // 1: seldon.mlops.proxy.UnloadModelRequest.model:type_name -> seldon.mlops.scheduler.ModelReference
	0, // 2: seldon.mlops.proxy.SchedulerProxy.LoadModel:input_type -> seldon.mlops.proxy.LoadModelRequest
	2, // 3: seldon.mlops.proxy.SchedulerProxy.UnloadModel:input_type -> seldon.mlops.proxy.UnloadModelRequest
	1, // 4: seldon.mlops.proxy.SchedulerProxy.LoadModel:output_type -> seldon.mlops.proxy.LoadModelResponse
	3, // 5: seldon.mlops.proxy.SchedulerProxy.UnloadModel:output_type -> seldon.mlops.proxy.UnloadModelResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mlops_proxy_proxy_proto_init() }
func file_mlops_proxy_proxy_proto_init() {
	if File_mlops_proxy_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mlops_proxy_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadModelRequest); i {
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
		file_mlops_proxy_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadModelResponse); i {
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
		file_mlops_proxy_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnloadModelRequest); i {
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
		file_mlops_proxy_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnloadModelResponse); i {
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
			RawDescriptor: file_mlops_proxy_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mlops_proxy_proxy_proto_goTypes,
		DependencyIndexes: file_mlops_proxy_proxy_proto_depIdxs,
		MessageInfos:      file_mlops_proxy_proxy_proto_msgTypes,
	}.Build()
	File_mlops_proxy_proxy_proto = out.File
	file_mlops_proxy_proxy_proto_rawDesc = nil
	file_mlops_proxy_proxy_proto_goTypes = nil
	file_mlops_proxy_proxy_proto_depIdxs = nil
}
