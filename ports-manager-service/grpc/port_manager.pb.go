// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.4
// source: grpc/port_manager.proto

package grpc

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

// The request message containing the port's details
type PortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortCode  string  `protobuf:"bytes,1,opt,name=portCode,proto3" json:"portCode,omitempty"`
	City      string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Province  string  `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	Timezone  string  `protobuf:"bytes,6,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Code      string  `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	Country   string  `protobuf:"bytes,8,opt,name=country,proto3" json:"country,omitempty"`
	Name      string  `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PortRequest) Reset() {
	*x = PortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_port_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortRequest) ProtoMessage() {}

func (x *PortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_port_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortRequest.ProtoReflect.Descriptor instead.
func (*PortRequest) Descriptor() ([]byte, []int) {
	return file_grpc_port_manager_proto_rawDescGZIP(), []int{0}
}

func (x *PortRequest) GetPortCode() string {
	if x != nil {
		return x.PortCode
	}
	return ""
}

func (x *PortRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PortRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *PortRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *PortRequest) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *PortRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *PortRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PortRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PortRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PortCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *PortCode) Reset() {
	*x = PortCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_port_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortCode) ProtoMessage() {}

func (x *PortCode) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_port_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortCode.ProtoReflect.Descriptor instead.
func (*PortCode) Descriptor() ([]byte, []int) {
	return file_grpc_port_manager_proto_rawDescGZIP(), []int{1}
}

func (x *PortCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortCode  string  `protobuf:"bytes,2,opt,name=portCode,proto3" json:"portCode,omitempty"`
	City      string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Longitude float64 `protobuf:"fixed64,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Province  string  `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty"`
	Timezone  string  `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Code      string  `protobuf:"bytes,8,opt,name=code,proto3" json:"code,omitempty"`
	Country   string  `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	Name      string  `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_port_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_port_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_grpc_port_manager_proto_rawDescGZIP(), []int{2}
}

func (x *Port) GetPortCode() string {
	if x != nil {
		return x.PortCode
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Port) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Port) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the write success
type PortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PortResponse) Reset() {
	*x = PortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_port_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortResponse) ProtoMessage() {}

func (x *PortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_port_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortResponse.ProtoReflect.Descriptor instead.
func (*PortResponse) Descriptor() ([]byte, []int) {
	return file_grpc_port_manager_proto_rawDescGZIP(), []int{3}
}

func (x *PortResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PortResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_grpc_port_manager_proto protoreflect.FileDescriptor

var file_grpc_port_manager_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x0b, 0x50, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a,
	0x08, 0x50, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xea, 0x01,
	0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x0c, 0x50, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x54, 0x0a, 0x0b, 0x50, 0x6f,
	0x72, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x0c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x09,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x05, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x42, 0x1a, 0x5a, 0x18, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_port_manager_proto_rawDescOnce sync.Once
	file_grpc_port_manager_proto_rawDescData = file_grpc_port_manager_proto_rawDesc
)

func file_grpc_port_manager_proto_rawDescGZIP() []byte {
	file_grpc_port_manager_proto_rawDescOnce.Do(func() {
		file_grpc_port_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_port_manager_proto_rawDescData)
	})
	return file_grpc_port_manager_proto_rawDescData
}

var file_grpc_port_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_port_manager_proto_goTypes = []interface{}{
	(*PortRequest)(nil),  // 0: PortRequest
	(*PortCode)(nil),     // 1: PortCode
	(*Port)(nil),         // 2: Port
	(*PortResponse)(nil), // 3: PortResponse
}
var file_grpc_port_manager_proto_depIdxs = []int32{
	0, // 0: PortManager.AddPort:input_type -> PortRequest
	1, // 1: PortManager.GetPort:input_type -> PortCode
	3, // 2: PortManager.AddPort:output_type -> PortResponse
	2, // 3: PortManager.GetPort:output_type -> Port
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_port_manager_proto_init() }
func file_grpc_port_manager_proto_init() {
	if File_grpc_port_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_port_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortRequest); i {
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
		file_grpc_port_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortCode); i {
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
		file_grpc_port_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_grpc_port_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortResponse); i {
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
			RawDescriptor: file_grpc_port_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_port_manager_proto_goTypes,
		DependencyIndexes: file_grpc_port_manager_proto_depIdxs,
		MessageInfos:      file_grpc_port_manager_proto_msgTypes,
	}.Build()
	File_grpc_port_manager_proto = out.File
	file_grpc_port_manager_proto_rawDesc = nil
	file_grpc_port_manager_proto_goTypes = nil
	file_grpc_port_manager_proto_depIdxs = nil
}
