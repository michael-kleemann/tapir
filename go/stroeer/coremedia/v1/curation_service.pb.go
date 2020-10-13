// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: stroeer/coremedia/v1/curation_service.proto

package coremedia

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetCuratedListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId string `protobuf:"bytes,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
}

func (x *GetCuratedListRequest) Reset() {
	*x = GetCuratedListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_coremedia_v1_curation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCuratedListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCuratedListRequest) ProtoMessage() {}

func (x *GetCuratedListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_coremedia_v1_curation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCuratedListRequest.ProtoReflect.Descriptor instead.
func (*GetCuratedListRequest) Descriptor() ([]byte, []int) {
	return file_stroeer_coremedia_v1_curation_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCuratedListRequest) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

type GetCuratedListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CuratedArticle `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetCuratedListResponse) Reset() {
	*x = GetCuratedListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_coremedia_v1_curation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCuratedListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCuratedListResponse) ProtoMessage() {}

func (x *GetCuratedListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_coremedia_v1_curation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCuratedListResponse.ProtoReflect.Descriptor instead.
func (*GetCuratedListResponse) Descriptor() ([]byte, []int) {
	return file_stroeer_coremedia_v1_curation_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCuratedListResponse) GetItems() []*CuratedArticle {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_stroeer_coremedia_v1_curation_service_proto protoreflect.FileDescriptor

var file_stroeer_coremedia_v1_curation_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49,
	0x64, 0x22, 0x54, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x72,
	0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x80, 0x01, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e,
	0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x74, 0x72,
	0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x57, 0x0a, 0x17, 0x64, 0x65,
	0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69,
	0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_coremedia_v1_curation_service_proto_rawDescOnce sync.Once
	file_stroeer_coremedia_v1_curation_service_proto_rawDescData = file_stroeer_coremedia_v1_curation_service_proto_rawDesc
)

func file_stroeer_coremedia_v1_curation_service_proto_rawDescGZIP() []byte {
	file_stroeer_coremedia_v1_curation_service_proto_rawDescOnce.Do(func() {
		file_stroeer_coremedia_v1_curation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_coremedia_v1_curation_service_proto_rawDescData)
	})
	return file_stroeer_coremedia_v1_curation_service_proto_rawDescData
}

var file_stroeer_coremedia_v1_curation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stroeer_coremedia_v1_curation_service_proto_goTypes = []interface{}{
	(*GetCuratedListRequest)(nil),  // 0: stroeer.coremedia.v1.GetCuratedListRequest
	(*GetCuratedListResponse)(nil), // 1: stroeer.coremedia.v1.GetCuratedListResponse
	(*CuratedArticle)(nil),         // 2: stroeer.coremedia.v1.CuratedArticle
}
var file_stroeer_coremedia_v1_curation_service_proto_depIdxs = []int32{
	2, // 0: stroeer.coremedia.v1.GetCuratedListResponse.items:type_name -> stroeer.coremedia.v1.CuratedArticle
	0, // 1: stroeer.coremedia.v1.CurationService.GetCuratedList:input_type -> stroeer.coremedia.v1.GetCuratedListRequest
	1, // 2: stroeer.coremedia.v1.CurationService.GetCuratedList:output_type -> stroeer.coremedia.v1.GetCuratedListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stroeer_coremedia_v1_curation_service_proto_init() }
func file_stroeer_coremedia_v1_curation_service_proto_init() {
	if File_stroeer_coremedia_v1_curation_service_proto != nil {
		return
	}
	file_stroeer_coremedia_v1_curated_article_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stroeer_coremedia_v1_curation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCuratedListRequest); i {
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
		file_stroeer_coremedia_v1_curation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCuratedListResponse); i {
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
			RawDescriptor: file_stroeer_coremedia_v1_curation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stroeer_coremedia_v1_curation_service_proto_goTypes,
		DependencyIndexes: file_stroeer_coremedia_v1_curation_service_proto_depIdxs,
		MessageInfos:      file_stroeer_coremedia_v1_curation_service_proto_msgTypes,
	}.Build()
	File_stroeer_coremedia_v1_curation_service_proto = out.File
	file_stroeer_coremedia_v1_curation_service_proto_rawDesc = nil
	file_stroeer_coremedia_v1_curation_service_proto_goTypes = nil
	file_stroeer_coremedia_v1_curation_service_proto_depIdxs = nil
}