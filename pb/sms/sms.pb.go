// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: sms.proto

package sms

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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type SmsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile          string            `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`                                                                                         //接受方手机号
	Template        string            `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`                                                                                     //模板id
	Params          map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //模板参数
	Type            string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`                                                                                             //查询条件短信类型
	SignName        string            `protobuf:"bytes,5,opt,name=signName,proto3" json:"signName,omitempty"`                                                                                     //阿里云短信签名
	Status          int32             `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`                                                                                        //查询条件状态
	AccessKeyId     string            `protobuf:"bytes,7,opt,name=AccessKeyId,proto3" json:"AccessKeyId,omitempty"`
	AccessKeySecret string            `protobuf:"bytes,8,opt,name=AccessKeySecret,proto3" json:"AccessKeySecret,omitempty"`
	Endpoint        string            `protobuf:"bytes,9,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
}

func (x *SmsParams) Reset() {
	*x = SmsParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsParams) ProtoMessage() {}

func (x *SmsParams) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsParams.ProtoReflect.Descriptor instead.
func (*SmsParams) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{1}
}

func (x *SmsParams) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SmsParams) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *SmsParams) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *SmsParams) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SmsParams) GetSignName() string {
	if x != nil {
		return x.SignName
	}
	return ""
}

func (x *SmsParams) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SmsParams) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *SmsParams) GetAccessKeySecret() string {
	if x != nil {
		return x.AccessKeySecret
	}
	return ""
}

func (x *SmsParams) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type SmsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RequestID string `protobuf:"bytes,3,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Path      string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Data      *Data  `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SmsResp) Reset() {
	*x = SmsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsResp) ProtoMessage() {}

func (x *SmsResp) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsResp.ProtoReflect.Descriptor instead.
func (*SmsResp) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{2}
}

func (x *SmsResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SmsResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SmsResp) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *SmsResp) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SmsResp) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_sms_proto protoreflect.FileDescriptor

var file_sms_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x6d, 0x73,
	0x22, 0x1a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xde, 0x02, 0x0a,
	0x09, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x32,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x01,
	0x0a, 0x07, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x62, 0x0a, 0x0d, 0x53, 0x6d, 0x73, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x27, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x12, 0x0e, 0x2e, 0x73,
	0x6d, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x73,
	0x6d, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x08, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x6d, 0x73, 0x12, 0x0e, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53, 0x6d, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x73, 0x6d, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_proto_rawDescOnce sync.Once
	file_sms_proto_rawDescData = file_sms_proto_rawDesc
)

func file_sms_proto_rawDescGZIP() []byte {
	file_sms_proto_rawDescOnce.Do(func() {
		file_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_proto_rawDescData)
	})
	return file_sms_proto_rawDescData
}

var file_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sms_proto_goTypes = []interface{}{
	(*Data)(nil),      // 0: sms.Data
	(*SmsParams)(nil), // 1: sms.SmsParams
	(*SmsResp)(nil),   // 2: sms.SmsResp
	nil,               // 3: sms.SmsParams.ParamsEntry
}
var file_sms_proto_depIdxs = []int32{
	3, // 0: sms.SmsParams.params:type_name -> sms.SmsParams.ParamsEntry
	0, // 1: sms.SmsResp.data:type_name -> sms.Data
	1, // 2: sms.SmsRpcService.SendSms:input_type -> sms.SmsParams
	1, // 3: sms.SmsRpcService.CheckSms:input_type -> sms.SmsParams
	2, // 4: sms.SmsRpcService.SendSms:output_type -> sms.SmsResp
	2, // 5: sms.SmsRpcService.CheckSms:output_type -> sms.SmsResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sms_proto_init() }
func file_sms_proto_init() {
	if File_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsParams); i {
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
		file_sms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsResp); i {
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
			RawDescriptor: file_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_proto_goTypes,
		DependencyIndexes: file_sms_proto_depIdxs,
		MessageInfos:      file_sms_proto_msgTypes,
	}.Build()
	File_sms_proto = out.File
	file_sms_proto_rawDesc = nil
	file_sms_proto_goTypes = nil
	file_sms_proto_depIdxs = nil
}
