// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: paginator/paginator.proto

package paginator

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

type Paginator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalRecord int64 `protobuf:"varint,1,opt,name=total_record,json=totalRecord,proto3" json:"total_record,omitempty"`
	TotalPage   int64 `protobuf:"varint,2,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	Offset      int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit       int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	CurrPage    int64 `protobuf:"varint,5,opt,name=curr_page,json=currPage,proto3" json:"curr_page,omitempty"`
	PrevPage    int64 `protobuf:"varint,6,opt,name=prev_page,json=prevPage,proto3" json:"prev_page,omitempty"`
	NextPage    int64 `protobuf:"varint,7,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *Paginator) Reset() {
	*x = Paginator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paginator_paginator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paginator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paginator) ProtoMessage() {}

func (x *Paginator) ProtoReflect() protoreflect.Message {
	mi := &file_paginator_paginator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paginator.ProtoReflect.Descriptor instead.
func (*Paginator) Descriptor() ([]byte, []int) {
	return file_paginator_paginator_proto_rawDescGZIP(), []int{0}
}

func (x *Paginator) GetTotalRecord() int64 {
	if x != nil {
		return x.TotalRecord
	}
	return 0
}

func (x *Paginator) GetTotalPage() int64 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *Paginator) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Paginator) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Paginator) GetCurrPage() int64 {
	if x != nil {
		return x.CurrPage
	}
	return 0
}

func (x *Paginator) GetPrevPage() int64 {
	if x != nil {
		return x.PrevPage
	}
	return 0
}

func (x *Paginator) GetNextPage() int64 {
	if x != nil {
		return x.NextPage
	}
	return 0
}

var File_paginator_paginator_proto protoreflect.FileDescriptor

var file_paginator_paginator_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x73,
	0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xd2, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x72, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x42, 0x20, 0x5a, 0x1e,
	0x6c, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x72, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paginator_paginator_proto_rawDescOnce sync.Once
	file_paginator_paginator_proto_rawDescData = file_paginator_paginator_proto_rawDesc
)

func file_paginator_paginator_proto_rawDescGZIP() []byte {
	file_paginator_paginator_proto_rawDescOnce.Do(func() {
		file_paginator_paginator_proto_rawDescData = protoimpl.X.CompressGZIP(file_paginator_paginator_proto_rawDescData)
	})
	return file_paginator_paginator_proto_rawDescData
}

var file_paginator_paginator_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_paginator_paginator_proto_goTypes = []any{
	(*Paginator)(nil), // 0: postclient.Paginator
}
var file_paginator_paginator_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_paginator_paginator_proto_init() }
func file_paginator_paginator_proto_init() {
	if File_paginator_paginator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paginator_paginator_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Paginator); i {
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
			RawDescriptor: file_paginator_paginator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_paginator_paginator_proto_goTypes,
		DependencyIndexes: file_paginator_paginator_proto_depIdxs,
		MessageInfos:      file_paginator_paginator_proto_msgTypes,
	}.Build()
	File_paginator_paginator_proto = out.File
	file_paginator_paginator_proto_rawDesc = nil
	file_paginator_paginator_proto_goTypes = nil
	file_paginator_paginator_proto_depIdxs = nil
}
