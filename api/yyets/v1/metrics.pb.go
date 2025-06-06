// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: yyets/v1/metrics.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateMetricsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMetricsRequest) Reset() {
	*x = CreateMetricsRequest{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetricsRequest) ProtoMessage() {}

func (x *CreateMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetricsRequest.ProtoReflect.Descriptor instead.
func (*CreateMetricsRequest) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMetricsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateMetricsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CreateMetricsReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMetricsReply) Reset() {
	*x = CreateMetricsReply{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMetricsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetricsReply) ProtoMessage() {}

func (x *CreateMetricsReply) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetricsReply.ProtoReflect.Descriptor instead.
func (*CreateMetricsReply) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{1}
}

type UpdateMetricsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateMetricsRequest) Reset() {
	*x = UpdateMetricsRequest{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMetricsRequest) ProtoMessage() {}

func (x *UpdateMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMetricsRequest.ProtoReflect.Descriptor instead.
func (*UpdateMetricsRequest) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{2}
}

type UpdateMetricsReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateMetricsReply) Reset() {
	*x = UpdateMetricsReply{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMetricsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMetricsReply) ProtoMessage() {}

func (x *UpdateMetricsReply) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMetricsReply.ProtoReflect.Descriptor instead.
func (*UpdateMetricsReply) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{3}
}

type DeleteMetricsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteMetricsRequest) Reset() {
	*x = DeleteMetricsRequest{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMetricsRequest) ProtoMessage() {}

func (x *DeleteMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMetricsRequest.ProtoReflect.Descriptor instead.
func (*DeleteMetricsRequest) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{4}
}

type DeleteMetricsReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteMetricsReply) Reset() {
	*x = DeleteMetricsReply{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMetricsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMetricsReply) ProtoMessage() {}

func (x *DeleteMetricsReply) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMetricsReply.ProtoReflect.Descriptor instead.
func (*DeleteMetricsReply) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{5}
}

type GetMetricsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMetricsRequest) Reset() {
	*x = GetMetricsRequest{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsRequest) ProtoMessage() {}

func (x *GetMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{6}
}

type GetMetricsReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMetricsReply) Reset() {
	*x = GetMetricsReply{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsReply) ProtoMessage() {}

func (x *GetMetricsReply) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsReply.ProtoReflect.Descriptor instead.
func (*GetMetricsReply) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{7}
}

type ListMetricsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMetricsRequest) Reset() {
	*x = ListMetricsRequest{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMetricsRequest) ProtoMessage() {}

func (x *ListMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMetricsRequest.ProtoReflect.Descriptor instead.
func (*ListMetricsRequest) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{8}
}

type ListMetricsReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMetricsReply) Reset() {
	*x = ListMetricsReply{}
	mi := &file_yyets_v1_metrics_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMetricsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMetricsReply) ProtoMessage() {}

func (x *ListMetricsReply) ProtoReflect() protoreflect.Message {
	mi := &file_yyets_v1_metrics_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMetricsReply.ProtoReflect.Descriptor instead.
func (*ListMetricsReply) Descriptor() ([]byte, []int) {
	return file_yyets_v1_metrics_proto_rawDescGZIP(), []int{9}
}

var File_yyets_v1_metrics_proto protoreflect.FileDescriptor

const file_yyets_v1_metrics_proto_rawDesc = "" +
	"\n" +
	"\x16yyets/v1/metrics.proto\x12\fapi.yyets.v1\x1a\x1cgoogle/api/annotations.proto\":\n" +
	"\x14CreateMetricsRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\"\x14\n" +
	"\x12CreateMetricsReply\"\x16\n" +
	"\x14UpdateMetricsRequest\"\x14\n" +
	"\x12UpdateMetricsReply\"\x16\n" +
	"\x14DeleteMetricsRequest\"\x14\n" +
	"\x12DeleteMetricsReply\"\x13\n" +
	"\x11GetMetricsRequest\"\x11\n" +
	"\x0fGetMetricsReply\"\x14\n" +
	"\x12ListMetricsRequest\"\x12\n" +
	"\x10ListMetricsReply2\xc6\x03\n" +
	"\aMetrics\x12n\n" +
	"\rCreateMetrics\x12\".api.yyets.v1.CreateMetricsRequest\x1a .api.yyets.v1.CreateMetricsReply\"\x17\x82\xd3\xe4\x93\x02\x11:\x01*\"\f/api/metrics\x12U\n" +
	"\rUpdateMetrics\x12\".api.yyets.v1.UpdateMetricsRequest\x1a .api.yyets.v1.UpdateMetricsReply\x12U\n" +
	"\rDeleteMetrics\x12\".api.yyets.v1.DeleteMetricsRequest\x1a .api.yyets.v1.DeleteMetricsReply\x12L\n" +
	"\n" +
	"GetMetrics\x12\x1f.api.yyets.v1.GetMetricsRequest\x1a\x1d.api.yyets.v1.GetMetricsReply\x12O\n" +
	"\vListMetrics\x12 .api.yyets.v1.ListMetricsRequest\x1a\x1e.api.yyets.v1.ListMetricsReplyB-\n" +
	"\fapi.yyets.v1P\x01Z\x1bYYeTsBot-Go/api/yyets/v1;v1b\x06proto3"

var (
	file_yyets_v1_metrics_proto_rawDescOnce sync.Once
	file_yyets_v1_metrics_proto_rawDescData []byte
)

func file_yyets_v1_metrics_proto_rawDescGZIP() []byte {
	file_yyets_v1_metrics_proto_rawDescOnce.Do(func() {
		file_yyets_v1_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yyets_v1_metrics_proto_rawDesc), len(file_yyets_v1_metrics_proto_rawDesc)))
	})
	return file_yyets_v1_metrics_proto_rawDescData
}

var file_yyets_v1_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_yyets_v1_metrics_proto_goTypes = []any{
	(*CreateMetricsRequest)(nil), // 0: api.yyets.v1.CreateMetricsRequest
	(*CreateMetricsReply)(nil),   // 1: api.yyets.v1.CreateMetricsReply
	(*UpdateMetricsRequest)(nil), // 2: api.yyets.v1.UpdateMetricsRequest
	(*UpdateMetricsReply)(nil),   // 3: api.yyets.v1.UpdateMetricsReply
	(*DeleteMetricsRequest)(nil), // 4: api.yyets.v1.DeleteMetricsRequest
	(*DeleteMetricsReply)(nil),   // 5: api.yyets.v1.DeleteMetricsReply
	(*GetMetricsRequest)(nil),    // 6: api.yyets.v1.GetMetricsRequest
	(*GetMetricsReply)(nil),      // 7: api.yyets.v1.GetMetricsReply
	(*ListMetricsRequest)(nil),   // 8: api.yyets.v1.ListMetricsRequest
	(*ListMetricsReply)(nil),     // 9: api.yyets.v1.ListMetricsReply
}
var file_yyets_v1_metrics_proto_depIdxs = []int32{
	0, // 0: api.yyets.v1.Metrics.CreateMetrics:input_type -> api.yyets.v1.CreateMetricsRequest
	2, // 1: api.yyets.v1.Metrics.UpdateMetrics:input_type -> api.yyets.v1.UpdateMetricsRequest
	4, // 2: api.yyets.v1.Metrics.DeleteMetrics:input_type -> api.yyets.v1.DeleteMetricsRequest
	6, // 3: api.yyets.v1.Metrics.GetMetrics:input_type -> api.yyets.v1.GetMetricsRequest
	8, // 4: api.yyets.v1.Metrics.ListMetrics:input_type -> api.yyets.v1.ListMetricsRequest
	1, // 5: api.yyets.v1.Metrics.CreateMetrics:output_type -> api.yyets.v1.CreateMetricsReply
	3, // 6: api.yyets.v1.Metrics.UpdateMetrics:output_type -> api.yyets.v1.UpdateMetricsReply
	5, // 7: api.yyets.v1.Metrics.DeleteMetrics:output_type -> api.yyets.v1.DeleteMetricsReply
	7, // 8: api.yyets.v1.Metrics.GetMetrics:output_type -> api.yyets.v1.GetMetricsReply
	9, // 9: api.yyets.v1.Metrics.ListMetrics:output_type -> api.yyets.v1.ListMetricsReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yyets_v1_metrics_proto_init() }
func file_yyets_v1_metrics_proto_init() {
	if File_yyets_v1_metrics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yyets_v1_metrics_proto_rawDesc), len(file_yyets_v1_metrics_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yyets_v1_metrics_proto_goTypes,
		DependencyIndexes: file_yyets_v1_metrics_proto_depIdxs,
		MessageInfos:      file_yyets_v1_metrics_proto_msgTypes,
	}.Build()
	File_yyets_v1_metrics_proto = out.File
	file_yyets_v1_metrics_proto_goTypes = nil
	file_yyets_v1_metrics_proto_depIdxs = nil
}
