// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: protobuf/raw/mydb.proto

package compiled

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protobuf_raw_mydb_proto protoreflect.FileDescriptor

var file_protobuf_raw_mydb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x61, 0x77, 0x2f, 0x6d,
	0x79, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x61, 0x77, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xda, 0x01, 0x0a, 0x09, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x05, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x1c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x05,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x1c,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x05, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x0e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x05, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protobuf_raw_mydb_proto_goTypes = []interface{}{
	(*User)(nil),          // 0: User
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
	(*UserList)(nil),      // 2: UserList
}
var file_protobuf_raw_mydb_proto_depIdxs = []int32{
	0, // 0: DBService.GetUserById:input_type -> User
	1, // 1: DBService.GetUserList:input_type -> google.protobuf.Empty
	0, // 2: DBService.CreateUser:input_type -> User
	0, // 3: DBService.UpdateUser:input_type -> User
	0, // 4: DBService.DeleteUser:input_type -> User
	0, // 5: DBService.CheckUserExist:input_type -> User
	0, // 6: DBService.GetUserById:output_type -> User
	2, // 7: DBService.GetUserList:output_type -> UserList
	0, // 8: DBService.CreateUser:output_type -> User
	0, // 9: DBService.UpdateUser:output_type -> User
	0, // 10: DBService.DeleteUser:output_type -> User
	0, // 11: DBService.CheckUserExist:output_type -> User
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_raw_mydb_proto_init() }
func file_protobuf_raw_mydb_proto_init() {
	if File_protobuf_raw_mydb_proto != nil {
		return
	}
	file_protobuf_raw_main_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_raw_mydb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_raw_mydb_proto_goTypes,
		DependencyIndexes: file_protobuf_raw_mydb_proto_depIdxs,
	}.Build()
	File_protobuf_raw_mydb_proto = out.File
	file_protobuf_raw_mydb_proto_rawDesc = nil
	file_protobuf_raw_mydb_proto_goTypes = nil
	file_protobuf_raw_mydb_proto_depIdxs = nil
}
