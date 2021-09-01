//*
// @FileArticle Reference

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: stroeer/core/v1/shared.proto

package core

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

//*
// A Reference represents a link to another entity, for example an `Article`,
// a `Section` or an external website, or a whole tree structure, for example
// a section tree or navigation.
// ```protobuf
// message Reference {}
// ```
type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The `type` is used for filtering in a list of references.
	// It describes a use-case, which usually has a defined render position.
	//
	// Example entries:
	// * `unspecified`
	// * `stage_title`
	// * `stage_themenbereiche`
	// * `stage_header_links`
	// * `stage_top_themen`
	// * `stage_tag_category`
	// ```protobuf
	// string type = 1;
	// ```
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	//*
	// The text or label of the reference.
	// ```protobuf
	// string label = 2;
	// ```
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	//*
	// The href of the reference. It can be relative or absolute.
	// ```protobuf
	// string href = 3;
	// ```
	Href string `protobuf:"bytes,3,opt,name=href,proto3" json:"href,omitempty"`
	//*
	// Contains all optional attributes of the reference.
	// * `target`
	// * `rel`
	// * `flag:internal`
	// * `layout`
	//
	// Clients must be resilient to unknown or missing entry sets.
	// ```protobuf
	// map<string, string> fields = 4;
	// ```
	Fields map[string]string `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//*
	// Hierarchically structured references for representing a navigation or tree.
	// ```protobuf
	// repeated Reference children = 5;
	// ```
	Children []*Reference `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_core_v1_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_core_v1_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_stroeer_core_v1_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Reference) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Reference) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Reference) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *Reference) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Reference) GetChildren() []*Reference {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_stroeer_core_v1_shared_proto protoreflect.FileDescriptor

var file_stroeer_core_v1_shared_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0xfc, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x12, 0x3e, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x48,
	0x0a, 0x12, 0x64, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_core_v1_shared_proto_rawDescOnce sync.Once
	file_stroeer_core_v1_shared_proto_rawDescData = file_stroeer_core_v1_shared_proto_rawDesc
)

func file_stroeer_core_v1_shared_proto_rawDescGZIP() []byte {
	file_stroeer_core_v1_shared_proto_rawDescOnce.Do(func() {
		file_stroeer_core_v1_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_core_v1_shared_proto_rawDescData)
	})
	return file_stroeer_core_v1_shared_proto_rawDescData
}

var file_stroeer_core_v1_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stroeer_core_v1_shared_proto_goTypes = []interface{}{
	(*Reference)(nil), // 0: stroeer.core.v1.Reference
	nil,               // 1: stroeer.core.v1.Reference.FieldsEntry
}
var file_stroeer_core_v1_shared_proto_depIdxs = []int32{
	1, // 0: stroeer.core.v1.Reference.fields:type_name -> stroeer.core.v1.Reference.FieldsEntry
	0, // 1: stroeer.core.v1.Reference.children:type_name -> stroeer.core.v1.Reference
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stroeer_core_v1_shared_proto_init() }
func file_stroeer_core_v1_shared_proto_init() {
	if File_stroeer_core_v1_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stroeer_core_v1_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
			RawDescriptor: file_stroeer_core_v1_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stroeer_core_v1_shared_proto_goTypes,
		DependencyIndexes: file_stroeer_core_v1_shared_proto_depIdxs,
		MessageInfos:      file_stroeer_core_v1_shared_proto_msgTypes,
	}.Build()
	File_stroeer_core_v1_shared_proto = out.File
	file_stroeer_core_v1_shared_proto_rawDesc = nil
	file_stroeer_core_v1_shared_proto_goTypes = nil
	file_stroeer_core_v1_shared_proto_depIdxs = nil
}
