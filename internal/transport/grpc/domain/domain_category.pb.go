// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: domain/domain_category.proto

package domain

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

type DomainCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Children    []*DomainCategory `protobuf:"bytes,48913922,rep,name=children,proto3" json:"children,omitempty"`
	Description string            `protobuf:"bytes,113933319,opt,name=description,proto3" json:"description,omitempty"`
	Hide        bool              `protobuf:"varint,3202370,opt,name=hide,proto3" json:"hide,omitempty"`
	Image       string            `protobuf:"bytes,100313435,opt,name=image,proto3" json:"image,omitempty"`
	ImageMini   string            `protobuf:"bytes,349353038,opt,name=image_mini,json=imageMini,proto3" json:"image_mini,omitempty"`
	Name        string            `protobuf:"bytes,3373707,opt,name=name,proto3" json:"name,omitempty"`
	ParentName  string            `protobuf:"bytes,425266290,opt,name=parent_name,json=parentName,proto3" json:"parent_name,omitempty"`
	ParentUuid  string            `protobuf:"bytes,425038658,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty"`
	Uuid        string            `protobuf:"bytes,3601339,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DomainCategory) Reset() {
	*x = DomainCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_domain_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainCategory) ProtoMessage() {}

func (x *DomainCategory) ProtoReflect() protoreflect.Message {
	mi := &file_domain_domain_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainCategory.ProtoReflect.Descriptor instead.
func (*DomainCategory) Descriptor() ([]byte, []int) {
	return file_domain_domain_category_proto_rawDescGZIP(), []int{0}
}

func (x *DomainCategory) GetChildren() []*DomainCategory {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *DomainCategory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DomainCategory) GetHide() bool {
	if x != nil {
		return x.Hide
	}
	return false
}

func (x *DomainCategory) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *DomainCategory) GetImageMini() string {
	if x != nil {
		return x.ImageMini
	}
	return ""
}

func (x *DomainCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DomainCategory) GetParentName() string {
	if x != nil {
		return x.ParentName
	}
	return ""
}

func (x *DomainCategory) GetParentUuid() string {
	if x != nil {
		return x.ParentUuid
	}
	return ""
}

func (x *DomainCategory) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_domain_domain_category_proto protoreflect.FileDescriptor

var file_domain_domain_category_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xb7, 0x02, 0x0a, 0x0e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x82, 0xbc, 0xa9, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x87, 0xf8, 0xa9, 0x36, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x04, 0x68, 0x69, 0x64, 0x65, 0x18, 0xc2, 0xba,
	0xc3, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x69, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0xdb, 0xd2, 0xea, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d,
	0x69, 0x6e, 0x69, 0x18, 0xce, 0xe8, 0xca, 0xa6, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x12, 0x15, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x8b, 0xf5, 0xcd, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xf2,
	0x98, 0xe4, 0xca, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0xc2, 0xa6, 0xd6, 0xca, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0xbb, 0xe7, 0xdb, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x30,
	0x78, 0x64, 0x65, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_domain_category_proto_rawDescOnce sync.Once
	file_domain_domain_category_proto_rawDescData = file_domain_domain_category_proto_rawDesc
)

func file_domain_domain_category_proto_rawDescGZIP() []byte {
	file_domain_domain_category_proto_rawDescOnce.Do(func() {
		file_domain_domain_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_domain_category_proto_rawDescData)
	})
	return file_domain_domain_category_proto_rawDescData
}

var file_domain_domain_category_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_domain_domain_category_proto_goTypes = []interface{}{
	(*DomainCategory)(nil), // 0: domain.DomainCategory
}
var file_domain_domain_category_proto_depIdxs = []int32{
	0, // 0: domain.DomainCategory.children:type_name -> domain.DomainCategory
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_domain_category_proto_init() }
func file_domain_domain_category_proto_init() {
	if File_domain_domain_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_domain_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainCategory); i {
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
			RawDescriptor: file_domain_domain_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_domain_category_proto_goTypes,
		DependencyIndexes: file_domain_domain_category_proto_depIdxs,
		MessageInfos:      file_domain_domain_category_proto_msgTypes,
	}.Build()
	File_domain_domain_category_proto = out.File
	file_domain_domain_category_proto_rawDesc = nil
	file_domain_domain_category_proto_goTypes = nil
	file_domain_domain_category_proto_depIdxs = nil
}
