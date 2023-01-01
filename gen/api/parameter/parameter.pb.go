// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/parameter/parameter.proto

package parameter

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

type ParameterType int32

const (
	ParameterType_PARAMETER_TYPE_ENV_VAR       ParameterType = 0
	ParameterType_PARAMETER_TYPE_FILE          ParameterType = 1
	ParameterType_PARAMETER_TYPE_TEMPLATE_FILE ParameterType = 2
)

// Enum value maps for ParameterType.
var (
	ParameterType_name = map[int32]string{
		0: "PARAMETER_TYPE_ENV_VAR",
		1: "PARAMETER_TYPE_FILE",
		2: "PARAMETER_TYPE_TEMPLATE_FILE",
	}
	ParameterType_value = map[string]int32{
		"PARAMETER_TYPE_ENV_VAR":       0,
		"PARAMETER_TYPE_FILE":          1,
		"PARAMETER_TYPE_TEMPLATE_FILE": 2,
	}
)

func (x ParameterType) Enum() *ParameterType {
	p := new(ParameterType)
	*p = x
	return p
}

func (x ParameterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParameterType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_parameter_parameter_proto_enumTypes[0].Descriptor()
}

func (ParameterType) Type() protoreflect.EnumType {
	return &file_api_parameter_parameter_proto_enumTypes[0]
}

func (x ParameterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParameterType.Descriptor instead.
func (ParameterType) EnumDescriptor() ([]byte, []int) {
	return file_api_parameter_parameter_proto_rawDescGZIP(), []int{0}
}

type WellKnown int32

const (
	WellKnown_WELL_KNOWN_NOOP            WellKnown = 0
	WellKnown_WELL_KNOWN_SERVICE_CLUSTER WellKnown = 1
	WellKnown_WELL_KNOWN_INGRESS         WellKnown = 3
)

// Enum value maps for WellKnown.
var (
	WellKnown_name = map[int32]string{
		0: "WELL_KNOWN_NOOP",
		1: "WELL_KNOWN_SERVICE_CLUSTER",
		3: "WELL_KNOWN_INGRESS",
	}
	WellKnown_value = map[string]int32{
		"WELL_KNOWN_NOOP":            0,
		"WELL_KNOWN_SERVICE_CLUSTER": 1,
		"WELL_KNOWN_INGRESS":         3,
	}
)

func (x WellKnown) Enum() *WellKnown {
	p := new(WellKnown)
	*p = x
	return p
}

func (x WellKnown) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WellKnown) Descriptor() protoreflect.EnumDescriptor {
	return file_api_parameter_parameter_proto_enumTypes[1].Descriptor()
}

func (WellKnown) Type() protoreflect.EnumType {
	return &file_api_parameter_parameter_proto_enumTypes[1]
}

func (x WellKnown) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WellKnown.Descriptor instead.
func (WellKnown) EnumDescriptor() ([]byte, []int) {
	return file_api_parameter_parameter_proto_rawDescGZIP(), []int{1}
}

var File_api_parameter_parameter_proto protoreflect.FileDescriptor

var file_api_parameter_parameter_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x6f, 0x70, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2a, 0x66, 0x0a, 0x0d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x56,
	0x5f, 0x56, 0x41, 0x52, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45,
	0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x20, 0x0a, 0x1c, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10,
	0x02, 0x2a, 0x58, 0x0a, 0x09, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x13,
	0x0a, 0x0f, 0x57, 0x45, 0x4c, 0x4c, 0x5f, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4e, 0x4f, 0x4f,
	0x50, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x57, 0x45, 0x4c, 0x4c, 0x5f, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x57, 0x45, 0x4c, 0x4c, 0x5f, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x03, 0x42, 0xee, 0x01, 0x0a, 0x1c,
	0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x70, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x42, 0x0e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x70, 0x70, 0x65, 0x72,
	0x2d, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x2f, 0x6f, 0x70, 0x73, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x4f,
	0x41, 0x50, 0xaa, 0x02, 0x18, 0x4f, 0x70, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x41, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0xca, 0x02, 0x18,
	0x4f, 0x70, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0xe2, 0x02, 0x24, 0x4f, 0x70, 0x73, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1a, 0x4f, 0x70, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_parameter_parameter_proto_rawDescOnce sync.Once
	file_api_parameter_parameter_proto_rawDescData = file_api_parameter_parameter_proto_rawDesc
)

func file_api_parameter_parameter_proto_rawDescGZIP() []byte {
	file_api_parameter_parameter_proto_rawDescOnce.Do(func() {
		file_api_parameter_parameter_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_parameter_parameter_proto_rawDescData)
	})
	return file_api_parameter_parameter_proto_rawDescData
}

var file_api_parameter_parameter_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_parameter_parameter_proto_goTypes = []interface{}{
	(ParameterType)(0), // 0: opscontrol.api.parameter.ParameterType
	(WellKnown)(0),     // 1: opscontrol.api.parameter.WellKnown
}
var file_api_parameter_parameter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_parameter_parameter_proto_init() }
func file_api_parameter_parameter_proto_init() {
	if File_api_parameter_parameter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_parameter_parameter_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_parameter_parameter_proto_goTypes,
		DependencyIndexes: file_api_parameter_parameter_proto_depIdxs,
		EnumInfos:         file_api_parameter_parameter_proto_enumTypes,
	}.Build()
	File_api_parameter_parameter_proto = out.File
	file_api_parameter_parameter_proto_rawDesc = nil
	file_api_parameter_parameter_proto_goTypes = nil
	file_api_parameter_parameter_proto_depIdxs = nil
}