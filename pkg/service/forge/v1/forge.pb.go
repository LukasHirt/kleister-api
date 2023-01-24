// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: forge/v1/forge.proto

package forgev1

import (
	v1 "github.com/kleister/kleister-api/pkg/service/types/v1"
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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forge_v1_forge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forge_v1_forge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_forge_v1_forge_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*v1.Forge `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forge_v1_forge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forge_v1_forge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_forge_v1_forge_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetResult() []*v1.Forge {
	if x != nil {
		return x.Result
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forge_v1_forge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forge_v1_forge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_forge_v1_forge_proto_rawDescGZIP(), []int{2}
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forge_v1_forge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forge_v1_forge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_forge_v1_forge_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListBuildsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ListBuildsRequest) Reset() {
	*x = ListBuildsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forge_v1_forge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildsRequest) ProtoMessage() {}

func (x *ListBuildsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forge_v1_forge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildsRequest.ProtoReflect.Descriptor instead.
func (*ListBuildsRequest) Descriptor() ([]byte, []int) {
	return file_forge_v1_forge_proto_rawDescGZIP(), []int{4}
}

func (x *ListBuildsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type ListBuildsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Builds []*v1.Build `protobuf:"bytes,1,rep,name=builds,proto3" json:"builds,omitempty"`
}

func (x *ListBuildsResponse) Reset() {
	*x = ListBuildsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forge_v1_forge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildsResponse) ProtoMessage() {}

func (x *ListBuildsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forge_v1_forge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildsResponse.ProtoReflect.Descriptor instead.
func (*ListBuildsResponse) Descriptor() ([]byte, []int) {
	return file_forge_v1_forge_proto_rawDescGZIP(), []int{5}
}

func (x *ListBuildsResponse) GetBuilds() []*v1.Build {
	if x != nil {
		return x.Builds
	}
	return nil
}

var File_forge_v1_forge_proto protoreflect.FileDescriptor

var file_forge_v1_forge_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x72, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x39, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x0f,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x2a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x3d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x06, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x32, 0xd2, 0x01, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x72,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x2e, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x12,
	0x1b, 0x2e, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66,
	0x6f, 0x72, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x9a, 0x01, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x2e, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x46, 0x6f, 0x72,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6c, 0x65, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6b,
	0x6c, 0x65, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02,
	0x08, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x46, 0x6f, 0x72, 0x67,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x46, 0x6f,
	0x72, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_forge_v1_forge_proto_rawDescOnce sync.Once
	file_forge_v1_forge_proto_rawDescData = file_forge_v1_forge_proto_rawDesc
)

func file_forge_v1_forge_proto_rawDescGZIP() []byte {
	file_forge_v1_forge_proto_rawDescOnce.Do(func() {
		file_forge_v1_forge_proto_rawDescData = protoimpl.X.CompressGZIP(file_forge_v1_forge_proto_rawDescData)
	})
	return file_forge_v1_forge_proto_rawDescData
}

var file_forge_v1_forge_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_forge_v1_forge_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),      // 0: forge.v1.SearchRequest
	(*SearchResponse)(nil),     // 1: forge.v1.SearchResponse
	(*UpdateRequest)(nil),      // 2: forge.v1.UpdateRequest
	(*UpdateResponse)(nil),     // 3: forge.v1.UpdateResponse
	(*ListBuildsRequest)(nil),  // 4: forge.v1.ListBuildsRequest
	(*ListBuildsResponse)(nil), // 5: forge.v1.ListBuildsResponse
	(*v1.Forge)(nil),           // 6: types.v1.Forge
	(*v1.Build)(nil),           // 7: types.v1.Build
}
var file_forge_v1_forge_proto_depIdxs = []int32{
	6, // 0: forge.v1.SearchResponse.result:type_name -> types.v1.Forge
	7, // 1: forge.v1.ListBuildsResponse.builds:type_name -> types.v1.Build
	0, // 2: forge.v1.ForgesService.Search:input_type -> forge.v1.SearchRequest
	2, // 3: forge.v1.ForgesService.Update:input_type -> forge.v1.UpdateRequest
	4, // 4: forge.v1.ForgesService.ListBuilds:input_type -> forge.v1.ListBuildsRequest
	1, // 5: forge.v1.ForgesService.Search:output_type -> forge.v1.SearchResponse
	3, // 6: forge.v1.ForgesService.Update:output_type -> forge.v1.UpdateResponse
	5, // 7: forge.v1.ForgesService.ListBuilds:output_type -> forge.v1.ListBuildsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_forge_v1_forge_proto_init() }
func file_forge_v1_forge_proto_init() {
	if File_forge_v1_forge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_forge_v1_forge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_forge_v1_forge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_forge_v1_forge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_forge_v1_forge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_forge_v1_forge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildsRequest); i {
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
		file_forge_v1_forge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildsResponse); i {
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
			RawDescriptor: file_forge_v1_forge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_forge_v1_forge_proto_goTypes,
		DependencyIndexes: file_forge_v1_forge_proto_depIdxs,
		MessageInfos:      file_forge_v1_forge_proto_msgTypes,
	}.Build()
	File_forge_v1_forge_proto = out.File
	file_forge_v1_forge_proto_rawDesc = nil
	file_forge_v1_forge_proto_goTypes = nil
	file_forge_v1_forge_proto_depIdxs = nil
}
