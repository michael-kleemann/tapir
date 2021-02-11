// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: stroeer/page/section/v1/section_page.proto

package section

import (
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/stroeer/tapir/go/stroeer/core/v1"
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

// Info to render the width of an item, relative to the page grid.
// These values represent "desktop-first" widths.
// `FULL` respects the grid gutter.
// `FULL_PAGE` breaks out of the regular grid (no gutter).
type StageItemWidth int32

const (
	StageItemWidth_STAGE_ITEM_WIDTH_UNSPECIFIED StageItemWidth = 0
	StageItemWidth_STAGE_ITEM_WIDTH_FULL        StageItemWidth = 1
	StageItemWidth_STAGE_ITEM_WIDTH_ONE_HALF    StageItemWidth = 2
	StageItemWidth_STAGE_ITEM_WIDTH_ONE_QUARTER StageItemWidth = 4
	StageItemWidth_STAGE_ITEM_WIDTH_ONE_EIGHTHS StageItemWidth = 8
	StageItemWidth_STAGE_ITEM_WIDTH_FULL_PAGE   StageItemWidth = 9
)

// Enum value maps for StageItemWidth.
var (
	StageItemWidth_name = map[int32]string{
		0: "STAGE_ITEM_WIDTH_UNSPECIFIED",
		1: "STAGE_ITEM_WIDTH_FULL",
		2: "STAGE_ITEM_WIDTH_ONE_HALF",
		4: "STAGE_ITEM_WIDTH_ONE_QUARTER",
		8: "STAGE_ITEM_WIDTH_ONE_EIGHTHS",
		9: "STAGE_ITEM_WIDTH_FULL_PAGE",
	}
	StageItemWidth_value = map[string]int32{
		"STAGE_ITEM_WIDTH_UNSPECIFIED": 0,
		"STAGE_ITEM_WIDTH_FULL":        1,
		"STAGE_ITEM_WIDTH_ONE_HALF":    2,
		"STAGE_ITEM_WIDTH_ONE_QUARTER": 4,
		"STAGE_ITEM_WIDTH_ONE_EIGHTHS": 8,
		"STAGE_ITEM_WIDTH_FULL_PAGE":   9,
	}
)

func (x StageItemWidth) Enum() *StageItemWidth {
	p := new(StageItemWidth)
	*p = x
	return p
}

func (x StageItemWidth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StageItemWidth) Descriptor() protoreflect.EnumDescriptor {
	return file_stroeer_page_section_v1_section_page_proto_enumTypes[0].Descriptor()
}

func (StageItemWidth) Type() protoreflect.EnumType {
	return &file_stroeer_page_section_v1_section_page_proto_enumTypes[0]
}

func (x StageItemWidth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StageItemWidth.Descriptor instead.
func (StageItemWidth) EnumDescriptor() ([]byte, []int) {
	return file_stroeer_page_section_v1_section_page_proto_rawDescGZIP(), []int{0}
}

// All data to render a section page like the homepage or "/politik/".
type SectionPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stages []*Stage `protobuf:"bytes,1,rep,name=stages,proto3" json:"stages,omitempty"`
}

func (x *SectionPage) Reset() {
	*x = SectionPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionPage) ProtoMessage() {}

func (x *SectionPage) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionPage.ProtoReflect.Descriptor instead.
func (*SectionPage) Descriptor() ([]byte, []int) {
	return file_stroeer_page_section_v1_section_page_proto_rawDescGZIP(), []int{0}
}

func (x *SectionPage) GetStages() []*Stage {
	if x != nil {
		return x.Stages
	}
	return nil
}

// A block of teasers - can consist of editorial articles, advertisement and/or stages.
type Stage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageItems         []*StageItem        `protobuf:"bytes,1,rep,name=stage_items,json=stageItems,proto3" json:"stage_items,omitempty"`
	StageConfiguration *StageConfiguration `protobuf:"bytes,2,opt,name=stage_configuration,json=stageConfiguration,proto3" json:"stage_configuration,omitempty"`
}

func (x *Stage) Reset() {
	*x = Stage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stage) ProtoMessage() {}

func (x *Stage) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stage.ProtoReflect.Descriptor instead.
func (*Stage) Descriptor() ([]byte, []int) {
	return file_stroeer_page_section_v1_section_page_proto_rawDescGZIP(), []int{1}
}

func (x *Stage) GetStageItems() []*StageItem {
	if x != nil {
		return x.StageItems
	}
	return nil
}

func (x *Stage) GetStageConfiguration() *StageConfiguration {
	if x != nil {
		return x.StageConfiguration
	}
	return nil
}

// Rendering data for a stage, currently only a fields map.
// Coordinate stage layouts, especially the `default` layout with UI department.
// todo: currently doesn't contain the link for a stage, like "Mehr aus Sport".
type StageConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Entry: `stage_title` represents an optional label to be rendered above a Stage, like "CORONAVIRUS".
	// Entry: `stage_layout` represents a visual layout, which may affect teasers, colors, and other stylings.
	// If entry `stage_layout` is missing, use layout `default`.
	Fields map[string]string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StageConfiguration) Reset() {
	*x = StageConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageConfiguration) ProtoMessage() {}

func (x *StageConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageConfiguration.ProtoReflect.Descriptor instead.
func (*StageConfiguration) Descriptor() ([]byte, []int) {
	return file_stroeer_page_section_v1_section_page_proto_rawDescGZIP(), []int{2}
}

func (x *StageConfiguration) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

// A single generic item of a stage.
type StageItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Width instructions for rendering the single item.
	StageItemWidth StageItemWidth `protobuf:"varint,1,opt,name=stage_item_width,json=stageItemWidth,proto3,enum=stroeer.page.section.v1.StageItemWidth" json:"stage_item_width,omitempty"`
	// At most one field will be set at the same time.
	// The delegate of the StageItem can be another Stage, an editorial article or a Commercial.
	//
	// Types that are assignable to StageItemType:
	//	*StageItem_Stage
	//	*StageItem_ArticleTeaser
	//	*StageItem_Commercial
	StageItemType isStageItem_StageItemType `protobuf_oneof:"stage_item_type"`
}

func (x *StageItem) Reset() {
	*x = StageItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageItem) ProtoMessage() {}

func (x *StageItem) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageItem.ProtoReflect.Descriptor instead.
func (*StageItem) Descriptor() ([]byte, []int) {
	return file_stroeer_page_section_v1_section_page_proto_rawDescGZIP(), []int{3}
}

func (x *StageItem) GetStageItemWidth() StageItemWidth {
	if x != nil {
		return x.StageItemWidth
	}
	return StageItemWidth_STAGE_ITEM_WIDTH_UNSPECIFIED
}

func (m *StageItem) GetStageItemType() isStageItem_StageItemType {
	if m != nil {
		return m.StageItemType
	}
	return nil
}

func (x *StageItem) GetStage() *Stage {
	if x, ok := x.GetStageItemType().(*StageItem_Stage); ok {
		return x.Stage
	}
	return nil
}

func (x *StageItem) GetArticleTeaser() *ArticleTeaser {
	if x, ok := x.GetStageItemType().(*StageItem_ArticleTeaser); ok {
		return x.ArticleTeaser
	}
	return nil
}

func (x *StageItem) GetCommercial() *Commercial {
	if x, ok := x.GetStageItemType().(*StageItem_Commercial); ok {
		return x.Commercial
	}
	return nil
}

type isStageItem_StageItemType interface {
	isStageItem_StageItemType()
}

type StageItem_Stage struct {
	Stage *Stage `protobuf:"bytes,2,opt,name=stage,proto3,oneof"`
}

type StageItem_ArticleTeaser struct {
	ArticleTeaser *ArticleTeaser `protobuf:"bytes,3,opt,name=article_teaser,json=articleTeaser,proto3,oneof"`
}

type StageItem_Commercial struct {
	Commercial *Commercial `protobuf:"bytes,4,opt,name=commercial,proto3,oneof"`
}

func (*StageItem_Stage) isStageItem_StageItemType() {}

func (*StageItem_ArticleTeaser) isStageItem_StageItemType() {}

func (*StageItem_Commercial) isStageItem_StageItemType() {}

// Data to render a Commercial as part of a Stage.
// todo: requirements for this feature are currently missing.
type Commercial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Commercial) Reset() {
	*x = Commercial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commercial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commercial) ProtoMessage() {}

func (x *Commercial) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commercial.ProtoReflect.Descriptor instead.
func (*Commercial) Descriptor() ([]byte, []int) {
	return file_stroeer_page_section_v1_section_page_proto_rawDescGZIP(), []int{4}
}

func (x *Commercial) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

// Bundles the data needed to render a core article as a teaser.
// Coordinate teaser layouts, especially the `default` layout with UI department.
type ArticleTeaser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `Article.body` is set to `null` to reduce overhead.
	Article *v1.Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	// `fields` may contain additional rendering information.
	// Entry: `teaser_layout`, marker to use a particular layout in the rendering teaser template.
	// If entry is missing use layout `default`.
	Fields map[string]string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ArticleTeaser) Reset() {
	*x = ArticleTeaser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleTeaser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleTeaser) ProtoMessage() {}

func (x *ArticleTeaser) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_section_v1_section_page_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleTeaser.ProtoReflect.Descriptor instead.
func (*ArticleTeaser) Descriptor() ([]byte, []int) {
	return file_stroeer_page_section_v1_section_page_proto_rawDescGZIP(), []int{5}
}

func (x *ArticleTeaser) GetArticle() *v1.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

func (x *ArticleTeaser) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_stroeer_page_section_v1_section_page_proto protoreflect.FileDescriptor

var file_stroeer_page_section_v1_section_page_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x72,
	0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x5c, 0x0a, 0x13, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65,
	0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x73, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4f, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc1, 0x02, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x51, 0x0a, 0x10, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x64, 0x74, 0x68, 0x52, 0x0e, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x74, 0x65, 0x61, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65,
	0x61, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54,
	0x65, 0x61, 0x73, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x48, 0x00,
	0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x42, 0x11, 0x0a, 0x0f,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x90, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x47,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x69, 0x61, 0x6c, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xca, 0x01, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65,
	0x61, 0x73, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65,
	0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x73, 0x65, 0x72,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a,
	0xd0, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x57, 0x49, 0x44, 0x54, 0x48, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x57, 0x49, 0x44, 0x54, 0x48, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x01, 0x12,
	0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x57, 0x49,
	0x44, 0x54, 0x48, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x02, 0x12, 0x20,
	0x0a, 0x1c, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x57, 0x49, 0x44,
	0x54, 0x48, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x10, 0x04,
	0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x57,
	0x49, 0x44, 0x54, 0x48, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x45, 0x49, 0x47, 0x48, 0x54, 0x48, 0x53,
	0x10, 0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x57, 0x49, 0x44, 0x54, 0x48, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x50, 0x41, 0x47, 0x45,
	0x10, 0x09, 0x42, 0x5b, 0x0a, 0x1a, 0x64, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_page_section_v1_section_page_proto_rawDescOnce sync.Once
	file_stroeer_page_section_v1_section_page_proto_rawDescData = file_stroeer_page_section_v1_section_page_proto_rawDesc
)

func file_stroeer_page_section_v1_section_page_proto_rawDescGZIP() []byte {
	file_stroeer_page_section_v1_section_page_proto_rawDescOnce.Do(func() {
		file_stroeer_page_section_v1_section_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_page_section_v1_section_page_proto_rawDescData)
	})
	return file_stroeer_page_section_v1_section_page_proto_rawDescData
}

var file_stroeer_page_section_v1_section_page_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stroeer_page_section_v1_section_page_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_stroeer_page_section_v1_section_page_proto_goTypes = []interface{}{
	(StageItemWidth)(0),        // 0: stroeer.page.section.v1.StageItemWidth
	(*SectionPage)(nil),        // 1: stroeer.page.section.v1.SectionPage
	(*Stage)(nil),              // 2: stroeer.page.section.v1.Stage
	(*StageConfiguration)(nil), // 3: stroeer.page.section.v1.StageConfiguration
	(*StageItem)(nil),          // 4: stroeer.page.section.v1.StageItem
	(*Commercial)(nil),         // 5: stroeer.page.section.v1.Commercial
	(*ArticleTeaser)(nil),      // 6: stroeer.page.section.v1.ArticleTeaser
	nil,                        // 7: stroeer.page.section.v1.StageConfiguration.FieldsEntry
	nil,                        // 8: stroeer.page.section.v1.Commercial.FieldsEntry
	nil,                        // 9: stroeer.page.section.v1.ArticleTeaser.FieldsEntry
	(*v1.Article)(nil),         // 10: stroeer.core.v1.Article
}
var file_stroeer_page_section_v1_section_page_proto_depIdxs = []int32{
	2,  // 0: stroeer.page.section.v1.SectionPage.stages:type_name -> stroeer.page.section.v1.Stage
	4,  // 1: stroeer.page.section.v1.Stage.stage_items:type_name -> stroeer.page.section.v1.StageItem
	3,  // 2: stroeer.page.section.v1.Stage.stage_configuration:type_name -> stroeer.page.section.v1.StageConfiguration
	7,  // 3: stroeer.page.section.v1.StageConfiguration.fields:type_name -> stroeer.page.section.v1.StageConfiguration.FieldsEntry
	0,  // 4: stroeer.page.section.v1.StageItem.stage_item_width:type_name -> stroeer.page.section.v1.StageItemWidth
	2,  // 5: stroeer.page.section.v1.StageItem.stage:type_name -> stroeer.page.section.v1.Stage
	6,  // 6: stroeer.page.section.v1.StageItem.article_teaser:type_name -> stroeer.page.section.v1.ArticleTeaser
	5,  // 7: stroeer.page.section.v1.StageItem.commercial:type_name -> stroeer.page.section.v1.Commercial
	8,  // 8: stroeer.page.section.v1.Commercial.fields:type_name -> stroeer.page.section.v1.Commercial.FieldsEntry
	10, // 9: stroeer.page.section.v1.ArticleTeaser.article:type_name -> stroeer.core.v1.Article
	9,  // 10: stroeer.page.section.v1.ArticleTeaser.fields:type_name -> stroeer.page.section.v1.ArticleTeaser.FieldsEntry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_stroeer_page_section_v1_section_page_proto_init() }
func file_stroeer_page_section_v1_section_page_proto_init() {
	if File_stroeer_page_section_v1_section_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stroeer_page_section_v1_section_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionPage); i {
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
		file_stroeer_page_section_v1_section_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stage); i {
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
		file_stroeer_page_section_v1_section_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StageConfiguration); i {
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
		file_stroeer_page_section_v1_section_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StageItem); i {
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
		file_stroeer_page_section_v1_section_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commercial); i {
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
		file_stroeer_page_section_v1_section_page_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleTeaser); i {
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
	file_stroeer_page_section_v1_section_page_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*StageItem_Stage)(nil),
		(*StageItem_ArticleTeaser)(nil),
		(*StageItem_Commercial)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stroeer_page_section_v1_section_page_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stroeer_page_section_v1_section_page_proto_goTypes,
		DependencyIndexes: file_stroeer_page_section_v1_section_page_proto_depIdxs,
		EnumInfos:         file_stroeer_page_section_v1_section_page_proto_enumTypes,
		MessageInfos:      file_stroeer_page_section_v1_section_page_proto_msgTypes,
	}.Build()
	File_stroeer_page_section_v1_section_page_proto = out.File
	file_stroeer_page_section_v1_section_page_proto_rawDesc = nil
	file_stroeer_page_section_v1_section_page_proto_goTypes = nil
	file_stroeer_page_section_v1_section_page_proto_depIdxs = nil
}