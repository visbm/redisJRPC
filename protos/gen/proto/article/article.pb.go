// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/article/article.proto

package article_proto

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

type ArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ArticleResponse) Reset() {
	*x = ArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleResponse) ProtoMessage() {}

func (x *ArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleResponse.ProtoReflect.Descriptor instead.
func (*ArticleResponse) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArticleResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ArticleResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ArticleArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*ArticleResponse `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ArticleArray) Reset() {
	*x = ArticleArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleArray) ProtoMessage() {}

func (x *ArticleArray) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleArray.ProtoReflect.Descriptor instead.
func (*ArticleArray) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleArray) GetValues() []*ArticleResponse {
	if x != nil {
		return x.Values
	}
	return nil
}

type ArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ArticleRequest) Reset() {
	*x = ArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleRequest) ProtoMessage() {}

func (x *ArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleRequest.ProtoReflect.Descriptor instead.
func (*ArticleRequest) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetArticleID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetArticleID) Reset() {
	*x = GetArticleID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleID) ProtoMessage() {}

func (x *GetArticleID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleID.ProtoReflect.Descriptor instead.
func (*GetArticleID) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_article_article_proto_rawDescGZIP(), []int{4}
}

var File_proto_article_article_proto protoreflect.FileDescriptor

var file_proto_article_article_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x49, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x40, 0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x12, 0x30, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x1e, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xc2, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x12, 0x0e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_article_article_proto_rawDescOnce sync.Once
	file_proto_article_article_proto_rawDescData = file_proto_article_article_proto_rawDesc
)

func file_proto_article_article_proto_rawDescGZIP() []byte {
	file_proto_article_article_proto_rawDescOnce.Do(func() {
		file_proto_article_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_article_article_proto_rawDescData)
	})
	return file_proto_article_article_proto_rawDescData
}

var file_proto_article_article_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_article_article_proto_goTypes = []interface{}{
	(*ArticleResponse)(nil), // 0: article.ArticleResponse
	(*ArticleArray)(nil),    // 1: article.ArticleArray
	(*ArticleRequest)(nil),  // 2: article.ArticleRequest
	(*GetArticleID)(nil),    // 3: article.GetArticleID
	(*Empty)(nil),           // 4: article.Empty
}
var file_proto_article_article_proto_depIdxs = []int32{
	0, // 0: article.ArticleArray.values:type_name -> article.ArticleResponse
	3, // 1: article.Article.GetArticle:input_type -> article.GetArticleID
	4, // 2: article.Article.GetArticles:input_type -> article.Empty
	2, // 3: article.Article.CreateArticle:input_type -> article.ArticleRequest
	0, // 4: article.Article.GetArticle:output_type -> article.ArticleResponse
	1, // 5: article.Article.GetArticles:output_type -> article.ArticleArray
	0, // 6: article.Article.CreateArticle:output_type -> article.ArticleResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_article_article_proto_init() }
func file_proto_article_article_proto_init() {
	if File_proto_article_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_article_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleResponse); i {
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
		file_proto_article_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleArray); i {
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
		file_proto_article_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleRequest); i {
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
		file_proto_article_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleID); i {
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
		file_proto_article_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_article_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_article_article_proto_goTypes,
		DependencyIndexes: file_proto_article_article_proto_depIdxs,
		MessageInfos:      file_proto_article_article_proto_msgTypes,
	}.Build()
	File_proto_article_article_proto = out.File
	file_proto_article_article_proto_rawDesc = nil
	file_proto_article_article_proto_goTypes = nil
	file_proto_article_article_proto_depIdxs = nil
}