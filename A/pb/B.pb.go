// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: B.proto

package pb

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

type Posts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Code    string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Image   string `protobuf:"bytes,5,opt,name=Image,proto3" json:"Image,omitempty"`
	User    string `protobuf:"bytes,6,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *Posts) Reset() {
	*x = Posts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Posts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Posts) ProtoMessage() {}

func (x *Posts) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Posts.ProtoReflect.Descriptor instead.
func (*Posts) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{0}
}

func (x *Posts) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Posts) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Posts) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Posts) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Posts) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Posts) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type InsertPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Code    string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Image   string `protobuf:"bytes,5,opt,name=Image,proto3" json:"Image,omitempty"`
	User    string `protobuf:"bytes,6,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *InsertPostRequest) Reset() {
	*x = InsertPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertPostRequest) ProtoMessage() {}

func (x *InsertPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertPostRequest.ProtoReflect.Descriptor instead.
func (*InsertPostRequest) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{1}
}

func (x *InsertPostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InsertPostRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *InsertPostRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *InsertPostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *InsertPostRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *InsertPostRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type InsertPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Posts `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *InsertPostResponse) Reset() {
	*x = InsertPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertPostResponse) ProtoMessage() {}

func (x *InsertPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertPostResponse.ProtoReflect.Descriptor instead.
func (*InsertPostResponse) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{2}
}

func (x *InsertPostResponse) GetPost() *Posts {
	if x != nil {
		return x.Post
	}
	return nil
}

type FindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllRequest) Reset() {
	*x = FindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRequest) ProtoMessage() {}

func (x *FindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRequest.ProtoReflect.Descriptor instead.
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{3}
}

type FindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post []*Posts `protobuf:"bytes,1,rep,name=post,proto3" json:"post,omitempty"`
}

func (x *FindAllResponse) Reset() {
	*x = FindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllResponse) ProtoMessage() {}

func (x *FindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllResponse.ProtoReflect.Descriptor instead.
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllResponse) GetPost() []*Posts {
	if x != nil {
		return x.Post
	}
	return nil
}

type SearchPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *SearchPostRequest) Reset() {
	*x = SearchPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPostRequest) ProtoMessage() {}

func (x *SearchPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPostRequest.ProtoReflect.Descriptor instead.
func (*SearchPostRequest) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{5}
}

func (x *SearchPostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type SearchPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post []*Posts `protobuf:"bytes,1,rep,name=post,proto3" json:"post,omitempty"`
}

func (x *SearchPostResponse) Reset() {
	*x = SearchPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPostResponse) ProtoMessage() {}

func (x *SearchPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPostResponse.ProtoReflect.Descriptor instead.
func (*SearchPostResponse) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{6}
}

func (x *SearchPostResponse) GetPost() []*Posts {
	if x != nil {
		return x.Post
	}
	return nil
}

type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdatePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *UpdatePostResponse) Reset() {
	*x = UpdatePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostResponse) ProtoMessage() {}

func (x *UpdatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostResponse.ProtoReflect.Descriptor instead.
func (*UpdatePostResponse) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePostResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdatePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Code    string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Image   string `protobuf:"bytes,5,opt,name=Image,proto3" json:"Image,omitempty"`
	User    string `protobuf:"bytes,6,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UpdatePost) Reset() {
	*x = UpdatePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePost) ProtoMessage() {}

func (x *UpdatePost) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePost.ProtoReflect.Descriptor instead.
func (*UpdatePost) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePost) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdatePost) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdatePost) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdatePost) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UpdatePost) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{10}
}

func (x *DeletePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type DeletePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *DeletePostResponse) Reset() {
	*x = DeletePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResponse) ProtoMessage() {}

func (x *DeletePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostResponse.ProtoReflect.Descriptor instead.
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{11}
}

func (x *DeletePostResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateNew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatePost    *UpdatePost        `protobuf:"bytes,1,opt,name=updatePost,proto3" json:"updatePost,omitempty"`
	UpdateRequest *UpdatePostRequest `protobuf:"bytes,2,opt,name=updateRequest,proto3" json:"updateRequest,omitempty"`
}

func (x *UpdateNew) Reset() {
	*x = UpdateNew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_B_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNew) ProtoMessage() {}

func (x *UpdateNew) ProtoReflect() protoreflect.Message {
	mi := &file_B_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNew.ProtoReflect.Descriptor instead.
func (*UpdateNew) Descriptor() ([]byte, []int) {
	return file_B_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateNew) GetUpdatePost() *UpdatePost {
	if x != nil {
		return x.UpdatePost
	}
	return nil
}

func (x *UpdateNew) GetUpdateRequest() *UpdatePostRequest {
	if x != nil {
		return x.UpdateRequest
	}
	return nil
}

var File_B_proto protoreflect.FileDescriptor

var file_B_proto_rawDesc = []byte{
	0x0a, 0x07, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x22, 0x8b, 0x01, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x97,
	0x01, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x33, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x22, 0x36, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x7a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x29, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x22, 0x7e, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x77, 0x12, 0x31, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0xd4, 0x02, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x1a, 0x19, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_B_proto_rawDescOnce sync.Once
	file_B_proto_rawDescData = file_B_proto_rawDesc
)

func file_B_proto_rawDescGZIP() []byte {
	file_B_proto_rawDescOnce.Do(func() {
		file_B_proto_rawDescData = protoimpl.X.CompressGZIP(file_B_proto_rawDescData)
	})
	return file_B_proto_rawDescData
}

var file_B_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_B_proto_goTypes = []interface{}{
	(*Posts)(nil),              // 0: Hello.Posts
	(*InsertPostRequest)(nil),  // 1: Hello.InsertPostRequest
	(*InsertPostResponse)(nil), // 2: Hello.InsertPostResponse
	(*FindAllRequest)(nil),     // 3: Hello.FindAllRequest
	(*FindAllResponse)(nil),    // 4: Hello.FindAllResponse
	(*SearchPostRequest)(nil),  // 5: Hello.SearchPostRequest
	(*SearchPostResponse)(nil), // 6: Hello.SearchPostResponse
	(*UpdatePostRequest)(nil),  // 7: Hello.UpdatePostRequest
	(*UpdatePostResponse)(nil), // 8: Hello.UpdatePostResponse
	(*UpdatePost)(nil),         // 9: Hello.UpdatePost
	(*DeletePostRequest)(nil),  // 10: Hello.DeletePostRequest
	(*DeletePostResponse)(nil), // 11: Hello.DeletePostResponse
	(*UpdateNew)(nil),          // 12: Hello.UpdateNew
}
var file_B_proto_depIdxs = []int32{
	0,  // 0: Hello.InsertPostResponse.post:type_name -> Hello.Posts
	0,  // 1: Hello.FindAllResponse.post:type_name -> Hello.Posts
	0,  // 2: Hello.SearchPostResponse.post:type_name -> Hello.Posts
	9,  // 3: Hello.UpdateNew.updatePost:type_name -> Hello.UpdatePost
	7,  // 4: Hello.UpdateNew.updateRequest:type_name -> Hello.UpdatePostRequest
	1,  // 5: Hello.PostServices.InsertPosts:input_type -> Hello.InsertPostRequest
	5,  // 6: Hello.PostServices.SearchPosts:input_type -> Hello.SearchPostRequest
	3,  // 7: Hello.PostServices.FindAllPosts:input_type -> Hello.FindAllRequest
	12, // 8: Hello.PostServices.UpdatePosts:input_type -> Hello.UpdateNew
	10, // 9: Hello.PostServices.DeletePosts:input_type -> Hello.DeletePostRequest
	1,  // 10: Hello.PostServices.InsertPosts:output_type -> Hello.InsertPostRequest
	6,  // 11: Hello.PostServices.SearchPosts:output_type -> Hello.SearchPostResponse
	4,  // 12: Hello.PostServices.FindAllPosts:output_type -> Hello.FindAllResponse
	8,  // 13: Hello.PostServices.UpdatePosts:output_type -> Hello.UpdatePostResponse
	11, // 14: Hello.PostServices.DeletePosts:output_type -> Hello.DeletePostResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_B_proto_init() }
func file_B_proto_init() {
	if File_B_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_B_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Posts); i {
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
		file_B_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertPostRequest); i {
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
		file_B_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertPostResponse); i {
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
		file_B_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRequest); i {
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
		file_B_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllResponse); i {
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
		file_B_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPostRequest); i {
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
		file_B_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPostResponse); i {
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
		file_B_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
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
		file_B_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostResponse); i {
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
		file_B_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePost); i {
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
		file_B_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
		file_B_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostResponse); i {
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
		file_B_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNew); i {
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
			RawDescriptor: file_B_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_B_proto_goTypes,
		DependencyIndexes: file_B_proto_depIdxs,
		MessageInfos:      file_B_proto_msgTypes,
	}.Build()
	File_B_proto = out.File
	file_B_proto_rawDesc = nil
	file_B_proto_goTypes = nil
	file_B_proto_depIdxs = nil
}