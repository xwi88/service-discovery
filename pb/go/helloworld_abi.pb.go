// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: helloworld_abi.proto

package entry

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

type OSType int32

const (
	OSType_UNKNOWN_OS OSType = 0
	OSType_IOS        OSType = 1
	OSType_ANDROID    OSType = 2
)

// Enum value maps for OSType.
var (
	OSType_name = map[int32]string{
		0: "UNKNOWN_OS",
		1: "IOS",
		2: "ANDROID",
	}
	OSType_value = map[string]int32{
		"UNKNOWN_OS": 0,
		"IOS":        1,
		"ANDROID":    2,
	}
)

func (x OSType) Enum() *OSType {
	p := new(OSType)
	*p = x
	return p
}

func (x OSType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OSType) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld_abi_proto_enumTypes[0].Descriptor()
}

func (OSType) Type() protoreflect.EnumType {
	return &file_helloworld_abi_proto_enumTypes[0]
}

func (x OSType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OSType.Descriptor instead.
func (OSType) EnumDescriptor() ([]byte, []int) {
	return file_helloworld_abi_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Ip       string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	OsType   int32  `protobuf:"varint,4,opt,name=os_type,json=osType,proto3" json:"os_type,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_abi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_abi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_abi_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *HelloRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetOsType() int32 {
	if x != nil {
		return x.OsType
	}
	return 0
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Ip       string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_abi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_abi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_helloworld_abi_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *HelloReply) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_helloworld_abi_proto protoreflect.FileDescriptor

var file_helloworld_abi_proto_rawDesc = []byte{
	0x0a, 0x14, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x61, 0x62, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x68, 0x0a, 0x0c, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6f,
	0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x53, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2e, 0x0a, 0x06, 0x4f, 0x53, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4f,
	0x53, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x02, 0x42, 0x1a, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x6d, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x5a, 0x07, 0x2e, 0x3b,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_abi_proto_rawDescOnce sync.Once
	file_helloworld_abi_proto_rawDescData = file_helloworld_abi_proto_rawDesc
)

func file_helloworld_abi_proto_rawDescGZIP() []byte {
	file_helloworld_abi_proto_rawDescOnce.Do(func() {
		file_helloworld_abi_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_abi_proto_rawDescData)
	})
	return file_helloworld_abi_proto_rawDescData
}

var file_helloworld_abi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_helloworld_abi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_helloworld_abi_proto_goTypes = []interface{}{
	(OSType)(0),          // 0: pb.OSType
	(*HelloRequest)(nil), // 1: pb.HelloRequest
	(*HelloReply)(nil),   // 2: pb.HelloReply
}
var file_helloworld_abi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_abi_proto_init() }
func file_helloworld_abi_proto_init() {
	if File_helloworld_abi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_abi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_helloworld_abi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_helloworld_abi_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_helloworld_abi_proto_goTypes,
		DependencyIndexes: file_helloworld_abi_proto_depIdxs,
		EnumInfos:         file_helloworld_abi_proto_enumTypes,
		MessageInfos:      file_helloworld_abi_proto_msgTypes,
	}.Build()
	File_helloworld_abi_proto = out.File
	file_helloworld_abi_proto_rawDesc = nil
	file_helloworld_abi_proto_goTypes = nil
	file_helloworld_abi_proto_depIdxs = nil
}
