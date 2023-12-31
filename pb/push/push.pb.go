// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.1
// source: push.proto

package push

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

type IosPushReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Msg     string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	IsDebug bool     `protobuf:"varint,3,opt,name=IsDebug,proto3" json:"IsDebug,omitempty"`
	Token   []string `protobuf:"bytes,4,rep,name=Token,proto3" json:"Token,omitempty"`
	MsgType string   `protobuf:"bytes,5,opt,name=msgType,proto3" json:"msgType,omitempty"` // 推送类型
	Scheme  string   `protobuf:"bytes,6,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Extra   string   `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *IosPushReq) Reset() {
	*x = IosPushReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IosPushReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IosPushReq) ProtoMessage() {}

func (x *IosPushReq) ProtoReflect() protoreflect.Message {
	mi := &file_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IosPushReq.ProtoReflect.Descriptor instead.
func (*IosPushReq) Descriptor() ([]byte, []int) {
	return file_push_proto_rawDescGZIP(), []int{0}
}

func (x *IosPushReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IosPushReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IosPushReq) GetIsDebug() bool {
	if x != nil {
		return x.IsDebug
	}
	return false
}

func (x *IosPushReq) GetToken() []string {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *IosPushReq) GetMsgType() string {
	if x != nil {
		return x.MsgType
	}
	return ""
}

func (x *IosPushReq) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *IosPushReq) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type AndroidPushReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	IsDebug bool     `protobuf:"varint,2,opt,name=IsDebug,proto3" json:"IsDebug,omitempty"`
	Msg     string   `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Token   []string `protobuf:"bytes,4,rep,name=Token,proto3" json:"Token,omitempty"`
	MsgType string   `protobuf:"bytes,5,opt,name=msgType,proto3" json:"msgType,omitempty"` //推送类型
	Type    string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`       //消息类型：launcher->首页 web->跳转web intent->跳转内部指定页面
	Action  string   `protobuf:"bytes,7,opt,name=action,proto3" json:"action,omitempty"`
	Scheme  string   `protobuf:"bytes,8,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Extra   string   `protobuf:"bytes,9,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *AndroidPushReq) Reset() {
	*x = AndroidPushReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidPushReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidPushReq) ProtoMessage() {}

func (x *AndroidPushReq) ProtoReflect() protoreflect.Message {
	mi := &file_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidPushReq.ProtoReflect.Descriptor instead.
func (*AndroidPushReq) Descriptor() ([]byte, []int) {
	return file_push_proto_rawDescGZIP(), []int{1}
}

func (x *AndroidPushReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AndroidPushReq) GetIsDebug() bool {
	if x != nil {
		return x.IsDebug
	}
	return false
}

func (x *AndroidPushReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AndroidPushReq) GetToken() []string {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *AndroidPushReq) GetMsgType() string {
	if x != nil {
		return x.MsgType
	}
	return ""
}

func (x *AndroidPushReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AndroidPushReq) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AndroidPushReq) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *AndroidPushReq) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type PushResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RequestID string        `protobuf:"bytes,3,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Path      string        `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Data      *PushRespData `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PushResp) Reset() {
	*x = PushResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResp) ProtoMessage() {}

func (x *PushResp) ProtoReflect() protoreflect.Message {
	mi := &file_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResp.ProtoReflect.Descriptor instead.
func (*PushResp) Descriptor() ([]byte, []int) {
	return file_push_proto_rawDescGZIP(), []int{2}
}

func (x *PushResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PushResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PushResp) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *PushResp) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PushResp) GetData() *PushRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PushRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key          string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Total        int64    `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Success      int64    `protobuf:"varint,3,opt,name=Success,proto3" json:"Success,omitempty"`
	Fail         int64    `protobuf:"varint,4,opt,name=Fail,proto3" json:"Fail,omitempty"`
	IsDebug      bool     `protobuf:"varint,5,opt,name=IsDebug,proto3" json:"IsDebug,omitempty"`
	SuccessToken []string `protobuf:"bytes,6,rep,name=SuccessToken,proto3" json:"SuccessToken,omitempty"`
	FailToken    []string `protobuf:"bytes,7,rep,name=FailToken,proto3" json:"FailToken,omitempty"`
}

func (x *PushRespData) Reset() {
	*x = PushRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRespData) ProtoMessage() {}

func (x *PushRespData) ProtoReflect() protoreflect.Message {
	mi := &file_push_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRespData.ProtoReflect.Descriptor instead.
func (*PushRespData) Descriptor() ([]byte, []int) {
	return file_push_proto_rawDescGZIP(), []int{3}
}

func (x *PushRespData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PushRespData) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PushRespData) GetSuccess() int64 {
	if x != nil {
		return x.Success
	}
	return 0
}

func (x *PushRespData) GetFail() int64 {
	if x != nil {
		return x.Fail
	}
	return 0
}

func (x *PushRespData) GetIsDebug() bool {
	if x != nil {
		return x.IsDebug
	}
	return false
}

func (x *PushRespData) GetSuccessToken() []string {
	if x != nil {
		return x.SuccessToken
	}
	return nil
}

func (x *PushRespData) GetFailToken() []string {
	if x != nil {
		return x.FailToken
	}
	return nil
}

var File_push_proto protoreflect.FileDescriptor

var file_push_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x75,
	0x73, 0x68, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x49, 0x6f, 0x73, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0xd8, 0x01,
	0x0a, 0x0e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x44, 0x65, 0x62, 0x75, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x26, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc0, 0x01, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x49, 0x73, 0x44, 0x65, 0x62, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49,
	0x73, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x61,
	0x69, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x46,
	0x61, 0x69, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa5, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x49,
	0x6f, 0x73, 0x50, 0x75, 0x73, 0x68, 0x12, 0x10, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x49, 0x6f,
	0x73, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0a, 0x48, 0x75, 0x61, 0x77,
	0x65, 0x69, 0x50, 0x75, 0x73, 0x68, 0x12, 0x14, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x41, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70,
	0x75, 0x73, 0x68, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0a,
	0x58, 0x69, 0x61, 0x6f, 0x6d, 0x69, 0x50, 0x75, 0x73, 0x68, 0x12, 0x14, 0x2e, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x1a, 0x0e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_push_proto_rawDescOnce sync.Once
	file_push_proto_rawDescData = file_push_proto_rawDesc
)

func file_push_proto_rawDescGZIP() []byte {
	file_push_proto_rawDescOnce.Do(func() {
		file_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_proto_rawDescData)
	})
	return file_push_proto_rawDescData
}

var file_push_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_push_proto_goTypes = []interface{}{
	(*IosPushReq)(nil),     // 0: push.IosPushReq
	(*AndroidPushReq)(nil), // 1: push.AndroidPushReq
	(*PushResp)(nil),       // 2: push.PushResp
	(*PushRespData)(nil),   // 3: push.PushRespData
}
var file_push_proto_depIdxs = []int32{
	3, // 0: push.PushResp.data:type_name -> push.PushRespData
	0, // 1: push.PushRpcService.IosPush:input_type -> push.IosPushReq
	1, // 2: push.PushRpcService.HuaweiPush:input_type -> push.AndroidPushReq
	1, // 3: push.PushRpcService.XiaomiPush:input_type -> push.AndroidPushReq
	2, // 4: push.PushRpcService.IosPush:output_type -> push.PushResp
	2, // 5: push.PushRpcService.HuaweiPush:output_type -> push.PushResp
	2, // 6: push.PushRpcService.XiaomiPush:output_type -> push.PushResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_push_proto_init() }
func file_push_proto_init() {
	if File_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IosPushReq); i {
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
		file_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidPushReq); i {
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
		file_push_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResp); i {
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
		file_push_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRespData); i {
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
			RawDescriptor: file_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_push_proto_goTypes,
		DependencyIndexes: file_push_proto_depIdxs,
		MessageInfos:      file_push_proto_msgTypes,
	}.Build()
	File_push_proto = out.File
	file_push_proto_rawDesc = nil
	file_push_proto_goTypes = nil
	file_push_proto_depIdxs = nil
}
