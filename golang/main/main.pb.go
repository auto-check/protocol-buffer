// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.15.8
// source: proto/main/main.proto

package mainproto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type GetLogListWithIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLogListWithIDRequest) Reset() {
	*x = GetLogListWithIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogListWithIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogListWithIDRequest) ProtoMessage() {}

func (x *GetLogListWithIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogListWithIDRequest.ProtoReflect.Descriptor instead.
func (*GetLogListWithIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_main_proto_rawDescGZIP(), []int{0}
}

type GetLogListWithIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*Log `protobuf:"bytes,1,rep,name=Logs,proto3" json:"Logs,omitempty"`
}

func (x *GetLogListWithIDResponse) Reset() {
	*x = GetLogListWithIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogListWithIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogListWithIDResponse) ProtoMessage() {}

func (x *GetLogListWithIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogListWithIDResponse.ProtoReflect.Descriptor instead.
func (*GetLogListWithIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_main_proto_rawDescGZIP(), []int{1}
}

func (x *GetLogListWithIDResponse) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type GetMacroStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gcn    string `protobuf:"bytes,1,opt,name=gcn,proto3" json:"gcn,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	School string `protobuf:"bytes,3,opt,name=school,proto3" json:"school,omitempty"`
	Onoff  bool   `protobuf:"varint,4,opt,name=onoff,proto3" json:"onoff,omitempty"`
}

func (x *GetMacroStatusResponse) Reset() {
	*x = GetMacroStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMacroStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMacroStatusResponse) ProtoMessage() {}

func (x *GetMacroStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMacroStatusResponse.ProtoReflect.Descriptor instead.
func (*GetMacroStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_main_proto_rawDescGZIP(), []int{2}
}

func (x *GetMacroStatusResponse) GetGcn() string {
	if x != nil {
		return x.Gcn
	}
	return ""
}

func (x *GetMacroStatusResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetMacroStatusResponse) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *GetMacroStatusResponse) GetOnoff() bool {
	if x != nil {
		return x.Onoff
	}
	return false
}

type GetMacroSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Birth    string `protobuf:"bytes,2,opt,name=birth,proto3" json:"birth,omitempty"`
}

func (x *GetMacroSecretResponse) Reset() {
	*x = GetMacroSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMacroSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMacroSecretResponse) ProtoMessage() {}

func (x *GetMacroSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMacroSecretResponse.ProtoReflect.Descriptor instead.
func (*GetMacroSecretResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_main_proto_rawDescGZIP(), []int{3}
}

func (x *GetMacroSecretResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetMacroSecretResponse) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

type ControlMacroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Onoff bool `protobuf:"varint,2,opt,name=onoff,proto3" json:"onoff,omitempty"`
}

func (x *ControlMacroRequest) Reset() {
	*x = ControlMacroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlMacroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlMacroRequest) ProtoMessage() {}

func (x *ControlMacroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlMacroRequest.ProtoReflect.Descriptor instead.
func (*ControlMacroRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_main_proto_rawDescGZIP(), []int{4}
}

func (x *ControlMacroRequest) GetOnoff() bool {
	if x != nil {
		return x.Onoff
	}
	return false
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_main_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_proto_main_main_proto_rawDescGZIP(), []int{5}
}

var File_proto_main_main_proto protoreflect.FileDescriptor

var file_proto_main_main_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x73,
	0x22, 0x6c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x63,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x63, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x6e, 0x6f, 0x66,
	0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x22, 0x4a,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x22, 0x05, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x32, 0xee,
	0x02, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x4c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x63, 0x72, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63,
	0x72, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x72, 0x6f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_main_main_proto_rawDescOnce sync.Once
	file_proto_main_main_proto_rawDescData = file_proto_main_main_proto_rawDesc
)

func file_proto_main_main_proto_rawDescGZIP() []byte {
	file_proto_main_main_proto_rawDescOnce.Do(func() {
		file_proto_main_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_main_main_proto_rawDescData)
	})
	return file_proto_main_main_proto_rawDescData
}

var file_proto_main_main_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_main_main_proto_goTypes = []interface{}{
	(*GetLogListWithIDRequest)(nil),  // 0: main.GetLogListWithIDRequest
	(*GetLogListWithIDResponse)(nil), // 1: main.GetLogListWithIDResponse
	(*GetMacroStatusResponse)(nil),   // 2: main.GetMacroStatusResponse
	(*GetMacroSecretResponse)(nil),   // 3: main.GetMacroSecretResponse
	(*ControlMacroRequest)(nil),      // 4: main.ControlMacroRequest
	(*Log)(nil),                      // 5: main.Log
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_proto_main_main_proto_depIdxs = []int32{
	5, // 0: main.GetLogListWithIDResponse.Logs:type_name -> main.Log
	6, // 1: main.Main.GetLogListWithID:input_type -> google.protobuf.Empty
	6, // 2: main.Main.CreateMacro:input_type -> google.protobuf.Empty
	6, // 3: main.Main.GetMacroStatus:input_type -> google.protobuf.Empty
	6, // 4: main.Main.GetMacroSecret:input_type -> google.protobuf.Empty
	4, // 5: main.Main.ControlMacro:input_type -> main.ControlMacroRequest
	1, // 6: main.Main.GetLogListWithID:output_type -> main.GetLogListWithIDResponse
	6, // 7: main.Main.CreateMacro:output_type -> google.protobuf.Empty
	2, // 8: main.Main.GetMacroStatus:output_type -> main.GetMacroStatusResponse
	3, // 9: main.Main.GetMacroSecret:output_type -> main.GetMacroSecretResponse
	6, // 10: main.Main.ControlMacro:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_main_main_proto_init() }
func file_proto_main_main_proto_init() {
	if File_proto_main_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_main_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogListWithIDRequest); i {
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
		file_proto_main_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogListWithIDResponse); i {
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
		file_proto_main_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMacroStatusResponse); i {
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
		file_proto_main_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMacroSecretResponse); i {
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
		file_proto_main_main_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlMacroRequest); i {
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
		file_proto_main_main_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_proto_main_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_main_main_proto_goTypes,
		DependencyIndexes: file_proto_main_main_proto_depIdxs,
		MessageInfos:      file_proto_main_main_proto_msgTypes,
	}.Build()
	File_proto_main_main_proto = out.File
	file_proto_main_main_proto_rawDesc = nil
	file_proto_main_main_proto_goTypes = nil
	file_proto_main_main_proto_depIdxs = nil
}
