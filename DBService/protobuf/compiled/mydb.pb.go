// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: protobuf/raw/mydb.proto

package compiled

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

type AllRows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *AllRows) Reset() {
	*x = AllRows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_raw_mydb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRows) ProtoMessage() {}

func (x *AllRows) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_raw_mydb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRows.ProtoReflect.Descriptor instead.
func (*AllRows) Descriptor() ([]byte, []int) {
	return file_protobuf_raw_mydb_proto_rawDescGZIP(), []int{0}
}

func (x *AllRows) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data map[string]*Value `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_raw_mydb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_raw_mydb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_protobuf_raw_mydb_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Row) GetData() map[string]*Value {
	if x != nil {
		return x.Data
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Unique bool   `protobuf:"varint,2,opt,name=unique,proto3" json:"unique,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_raw_mydb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_raw_mydb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_protobuf_raw_mydb_proto_rawDescGZIP(), []int{2}
}

func (x *Value) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Value) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

type GetByKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetByKeyRequest) Reset() {
	*x = GetByKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_raw_mydb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByKeyRequest) ProtoMessage() {}

func (x *GetByKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_raw_mydb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByKeyRequest.ProtoReflect.Descriptor instead.
func (*GetByKeyRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_raw_mydb_proto_rawDescGZIP(), []int{3}
}

func (x *GetByKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_raw_mydb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_raw_mydb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_raw_mydb_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

type DBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DBResponse) Reset() {
	*x = DBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_raw_mydb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBResponse) ProtoMessage() {}

func (x *DBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_raw_mydb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBResponse.ProtoReflect.Descriptor instead.
func (*DBResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_raw_mydb_proto_rawDescGZIP(), []int{5}
}

func (x *DBResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteByKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteByKeyRequest) Reset() {
	*x = DeleteByKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_raw_mydb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByKeyRequest) ProtoMessage() {}

func (x *DeleteByKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_raw_mydb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteByKeyRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_raw_mydb_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteByKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_protobuf_raw_mydb_proto protoreflect.FileDescriptor

var file_protobuf_raw_mydb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x61, 0x77, 0x2f, 0x6d,
	0x79, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x41, 0x6c, 0x6c,
	0x52, 0x6f, 0x77, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x04, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x7c,
	0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3f, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x35, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x26, 0x0a, 0x0a, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x26, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32,
	0xce, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12,
	0x10, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x04, 0x2e, 0x52, 0x6f, 0x77, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x77, 0x73, 0x22, 0x00, 0x12,
	0x1d, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x04, 0x2e, 0x52, 0x6f, 0x77, 0x1a,
	0x0b, 0x2e, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1d,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x04, 0x2e, 0x52, 0x6f, 0x77, 0x1a, 0x0b,
	0x2e, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x13, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_raw_mydb_proto_rawDescOnce sync.Once
	file_protobuf_raw_mydb_proto_rawDescData = file_protobuf_raw_mydb_proto_rawDesc
)

func file_protobuf_raw_mydb_proto_rawDescGZIP() []byte {
	file_protobuf_raw_mydb_proto_rawDescOnce.Do(func() {
		file_protobuf_raw_mydb_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_raw_mydb_proto_rawDescData)
	})
	return file_protobuf_raw_mydb_proto_rawDescData
}

var file_protobuf_raw_mydb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protobuf_raw_mydb_proto_goTypes = []interface{}{
	(*AllRows)(nil),            // 0: AllRows
	(*Row)(nil),                // 1: Row
	(*Value)(nil),              // 2: Value
	(*GetByKeyRequest)(nil),    // 3: GetByKeyRequest
	(*GetAllRequest)(nil),      // 4: GetAllRequest
	(*DBResponse)(nil),         // 5: DBResponse
	(*DeleteByKeyRequest)(nil), // 6: DeleteByKeyRequest
	nil,                        // 7: Row.DataEntry
}
var file_protobuf_raw_mydb_proto_depIdxs = []int32{
	1, // 0: AllRows.rows:type_name -> Row
	7, // 1: Row.data:type_name -> Row.DataEntry
	2, // 2: Row.DataEntry.value:type_name -> Value
	3, // 3: DatabaseService.GetByKey:input_type -> GetByKeyRequest
	4, // 4: DatabaseService.GetAll:input_type -> GetAllRequest
	1, // 5: DatabaseService.Insert:input_type -> Row
	1, // 6: DatabaseService.Update:input_type -> Row
	6, // 7: DatabaseService.DeleteByKey:input_type -> DeleteByKeyRequest
	1, // 8: DatabaseService.GetByKey:output_type -> Row
	0, // 9: DatabaseService.GetAll:output_type -> AllRows
	5, // 10: DatabaseService.Insert:output_type -> DBResponse
	5, // 11: DatabaseService.Update:output_type -> DBResponse
	5, // 12: DatabaseService.DeleteByKey:output_type -> DBResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_raw_mydb_proto_init() }
func file_protobuf_raw_mydb_proto_init() {
	if File_protobuf_raw_mydb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_raw_mydb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRows); i {
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
		file_protobuf_raw_mydb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_protobuf_raw_mydb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_protobuf_raw_mydb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByKeyRequest); i {
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
		file_protobuf_raw_mydb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_protobuf_raw_mydb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBResponse); i {
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
		file_protobuf_raw_mydb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByKeyRequest); i {
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
			RawDescriptor: file_protobuf_raw_mydb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_raw_mydb_proto_goTypes,
		DependencyIndexes: file_protobuf_raw_mydb_proto_depIdxs,
		MessageInfos:      file_protobuf_raw_mydb_proto_msgTypes,
	}.Build()
	File_protobuf_raw_mydb_proto = out.File
	file_protobuf_raw_mydb_proto_rawDesc = nil
	file_protobuf_raw_mydb_proto_goTypes = nil
	file_protobuf_raw_mydb_proto_depIdxs = nil
}
