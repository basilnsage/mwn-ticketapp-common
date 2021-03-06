// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: orderCreated.proto

package events

import (
	subjects "github.com/basilnsage/mwn-ticketapp-common/subjects"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type OrderCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject subjects.Subject `protobuf:"varint,1,opt,name=subject,proto3,enum=Subject" json:"subject,omitempty"`
	Data    *CreatedData     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrderCreated) Reset() {
	*x = OrderCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderCreated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreated) ProtoMessage() {}

func (x *OrderCreated) ProtoReflect() protoreflect.Message {
	mi := &file_orderCreated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreated.ProtoReflect.Descriptor instead.
func (*OrderCreated) Descriptor() ([]byte, []int) {
	return file_orderCreated_proto_rawDescGZIP(), []int{0}
}

func (x *OrderCreated) GetSubject() subjects.Subject {
	if x != nil {
		return x.Subject
	}
	return subjects.Subject_UNKNOWN_SUBJECT
}

func (x *OrderCreated) GetData() *CreatedData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreatedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    Status                 `protobuf:"varint,2,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	UserId    string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Ticket    *CreatedData_Ticket    `protobuf:"bytes,5,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *CreatedData) Reset() {
	*x = CreatedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderCreated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedData) ProtoMessage() {}

func (x *CreatedData) ProtoReflect() protoreflect.Message {
	mi := &file_orderCreated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedData.ProtoReflect.Descriptor instead.
func (*CreatedData) Descriptor() ([]byte, []int) {
	return file_orderCreated_proto_rawDescGZIP(), []int{1}
}

func (x *CreatedData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatedData) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Created
}

func (x *CreatedData) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreatedData) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *CreatedData) GetTicket() *CreatedData_Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type CreatedData_Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreatedData_Ticket) Reset() {
	*x = CreatedData_Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderCreated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedData_Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedData_Ticket) ProtoMessage() {}

func (x *CreatedData_Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_orderCreated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedData_Ticket.ProtoReflect.Descriptor instead.
func (*CreatedData_Ticket) Descriptor() ([]byte, []int) {
	return file_orderCreated_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreatedData_Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatedData_Ticket) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_orderCreated_proto protoreflect.FileDescriptor

var file_orderCreated_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6e, 0x61, 0x74, 0x73, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xef, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x1a, 0x2e, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x6c, 0x6e, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x77,
	0x6e, 0x2d, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_orderCreated_proto_rawDescOnce sync.Once
	file_orderCreated_proto_rawDescData = file_orderCreated_proto_rawDesc
)

func file_orderCreated_proto_rawDescGZIP() []byte {
	file_orderCreated_proto_rawDescOnce.Do(func() {
		file_orderCreated_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderCreated_proto_rawDescData)
	})
	return file_orderCreated_proto_rawDescData
}

var file_orderCreated_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orderCreated_proto_goTypes = []interface{}{
	(*OrderCreated)(nil),          // 0: OrderCreated
	(*CreatedData)(nil),           // 1: CreatedData
	(*CreatedData_Ticket)(nil),    // 2: CreatedData.Ticket
	(subjects.Subject)(0),         // 3: Subject
	(Status)(0),                   // 4: Status
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_orderCreated_proto_depIdxs = []int32{
	3, // 0: OrderCreated.subject:type_name -> Subject
	1, // 1: OrderCreated.data:type_name -> CreatedData
	4, // 2: CreatedData.status:type_name -> Status
	5, // 3: CreatedData.expires_at:type_name -> google.protobuf.Timestamp
	2, // 4: CreatedData.ticket:type_name -> CreatedData.Ticket
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orderCreated_proto_init() }
func file_orderCreated_proto_init() {
	if File_orderCreated_proto != nil {
		return
	}
	file_orderStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_orderCreated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreated); i {
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
		file_orderCreated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedData); i {
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
		file_orderCreated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedData_Ticket); i {
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
			RawDescriptor: file_orderCreated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderCreated_proto_goTypes,
		DependencyIndexes: file_orderCreated_proto_depIdxs,
		MessageInfos:      file_orderCreated_proto_msgTypes,
	}.Build()
	File_orderCreated_proto = out.File
	file_orderCreated_proto_rawDesc = nil
	file_orderCreated_proto_goTypes = nil
	file_orderCreated_proto_depIdxs = nil
}
