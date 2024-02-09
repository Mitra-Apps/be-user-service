// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/error.proto

package error

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

type ErrorCode int32

const (
	ErrorCode_UNKNOWN                       ErrorCode = 0
	ErrorCode_AUTH_REGISTER_USER_UNVERIFIED ErrorCode = 1
	ErrorCode_AUTH_REGISTER_USER_VERIFIED   ErrorCode = 2
	ErrorCode_AUTH_LOGIN_NOT_FOUND          ErrorCode = 3
	ErrorCode_AUTH_LOGIN_PASSWORD_INCORRECT ErrorCode = 4
	ErrorCode_AUTH_LOGIN_USER_UNVERIFIED    ErrorCode = 5
	ErrorCode_AUTH_OTP_INVALID              ErrorCode = 6
	ErrorCode_AUTH_OTP_ERROR_VERIFIED_USER  ErrorCode = 7
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "UNKNOWN",
		1: "AUTH_REGISTER_USER_UNVERIFIED",
		2: "AUTH_REGISTER_USER_VERIFIED",
		3: "AUTH_LOGIN_NOT_FOUND",
		4: "AUTH_LOGIN_PASSWORD_INCORRECT",
		5: "AUTH_LOGIN_USER_UNVERIFIED",
		6: "AUTH_OTP_INVALID",
		7: "AUTH_OTP_ERROR_VERIFIED_USER",
	}
	ErrorCode_value = map[string]int32{
		"UNKNOWN":                       0,
		"AUTH_REGISTER_USER_UNVERIFIED": 1,
		"AUTH_REGISTER_USER_VERIFIED":   2,
		"AUTH_LOGIN_NOT_FOUND":          3,
		"AUTH_LOGIN_PASSWORD_INCORRECT": 4,
		"AUTH_LOGIN_USER_UNVERIFIED":    5,
		"AUTH_OTP_INVALID":              6,
		"AUTH_OTP_ERROR_VERIFIED_USER":  7,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_error_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_proto_error_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_error_proto_rawDescGZIP(), []int{0}
}

var File_proto_error_proto protoreflect.FileDescriptor

var file_proto_error_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf1, 0x01, 0x0a, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45,
	0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x56, 0x45,
	0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x55, 0x54, 0x48,
	0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x55, 0x54,
	0x48, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x52,
	0x52, 0x45, 0x43, 0x54, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x56, 0x45, 0x52, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4f,
	0x54, 0x50, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x06, 0x12, 0x20, 0x0a, 0x1c,
	0x41, 0x55, 0x54, 0x48, 0x5f, 0x4f, 0x54, 0x50, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x07, 0x42, 0x85,
	0x01, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0a, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x69, 0x74, 0x72, 0x61, 0x2d, 0x41, 0x70, 0x70,
	0x73, 0x2f, 0x62, 0x65, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0xca, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x11, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_error_proto_rawDescOnce sync.Once
	file_proto_error_proto_rawDescData = file_proto_error_proto_rawDesc
)

func file_proto_error_proto_rawDescGZIP() []byte {
	file_proto_error_proto_rawDescOnce.Do(func() {
		file_proto_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_error_proto_rawDescData)
	})
	return file_proto_error_proto_rawDescData
}

var file_proto_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_error_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: proto.ErrorCode
}
var file_proto_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_error_proto_init() }
func file_proto_error_proto_init() {
	if File_proto_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_error_proto_goTypes,
		DependencyIndexes: file_proto_error_proto_depIdxs,
		EnumInfos:         file_proto_error_proto_enumTypes,
	}.Build()
	File_proto_error_proto = out.File
	file_proto_error_proto_rawDesc = nil
	file_proto_error_proto_goTypes = nil
	file_proto_error_proto_depIdxs = nil
}
