// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: permission.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_permission_proto protoreflect.FileDescriptor

var file_permission_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb4, 0x03,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x20, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x1a, 0x1e, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x13,
	0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1e, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a,
	0x13, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b,
	0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x1a, 0x13, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x52, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x2d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_permission_proto_goTypes = []interface{}{
	(*PermissionEntities)(nil),     // 0: ncs.protobuf.permissionEntities
	(*PermissionKey)(nil),          // 1: ncs.protobuf.permissionKey
	(*Empty)(nil),                  // 2: ncs.protobuf.empty
	(*PermissionEntity)(nil),       // 3: ncs.protobuf.permissionEntity
	(*PermissionCheckRequest)(nil), // 4: ncs.protobuf.permissionCheckRequest
	(*PermissionCheckResult)(nil),  // 5: ncs.protobuf.permissionCheckResult
}
var file_permission_proto_depIdxs = []int32{
	0, // 0: ncs.protobuf.Permission.Create:input_type -> ncs.protobuf.permissionEntities
	1, // 1: ncs.protobuf.Permission.FindById:input_type -> ncs.protobuf.permissionKey
	2, // 2: ncs.protobuf.Permission.FindAll:input_type -> ncs.protobuf.empty
	3, // 3: ncs.protobuf.Permission.Update:input_type -> ncs.protobuf.permissionEntity
	1, // 4: ncs.protobuf.Permission.Delete:input_type -> ncs.protobuf.permissionKey
	4, // 5: ncs.protobuf.Permission.Check:input_type -> ncs.protobuf.permissionCheckRequest
	0, // 6: ncs.protobuf.Permission.Create:output_type -> ncs.protobuf.permissionEntities
	3, // 7: ncs.protobuf.Permission.FindById:output_type -> ncs.protobuf.permissionEntity
	0, // 8: ncs.protobuf.Permission.FindAll:output_type -> ncs.protobuf.permissionEntities
	2, // 9: ncs.protobuf.Permission.Update:output_type -> ncs.protobuf.empty
	2, // 10: ncs.protobuf.Permission.Delete:output_type -> ncs.protobuf.empty
	5, // 11: ncs.protobuf.Permission.Check:output_type -> ncs.protobuf.permissionCheckResult
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_permission_proto_init() }
func file_permission_proto_init() {
	if File_permission_proto != nil {
		return
	}
	file_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_permission_proto_goTypes,
		DependencyIndexes: file_permission_proto_depIdxs,
	}.Build()
	File_permission_proto = out.File
	file_permission_proto_rawDesc = nil
	file_permission_proto_goTypes = nil
	file_permission_proto_depIdxs = nil
}
