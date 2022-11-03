// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: apiserver/apiserver.proto

package apiserver

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "lightflus/proto/common"
	stream "lightflus/proto/stream"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResourceTypeEnum int32

const (
	ResourceTypeEnum_RESOURCE_TYPE_ENUM_UNSPECIFIC ResourceTypeEnum = 0
	ResourceTypeEnum_RESOURCE_TYPE_ENUM_DATAFLOW   ResourceTypeEnum = 1
)

// Enum value maps for ResourceTypeEnum.
var (
	ResourceTypeEnum_name = map[int32]string{
		0: "RESOURCE_TYPE_ENUM_UNSPECIFIC",
		1: "RESOURCE_TYPE_ENUM_DATAFLOW",
	}
	ResourceTypeEnum_value = map[string]int32{
		"RESOURCE_TYPE_ENUM_UNSPECIFIC": 0,
		"RESOURCE_TYPE_ENUM_DATAFLOW":   1,
	}
)

func (x ResourceTypeEnum) Enum() *ResourceTypeEnum {
	p := new(ResourceTypeEnum)
	*p = x
	return p
}

func (x ResourceTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_apiserver_apiserver_proto_enumTypes[0].Descriptor()
}

func (ResourceTypeEnum) Type() protoreflect.EnumType {
	return &file_apiserver_apiserver_proto_enumTypes[0]
}

func (x ResourceTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceTypeEnum.Descriptor instead.
func (ResourceTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{0}
}

type ResourceStatusEnum int32

const (
	ResourceStatusEnum_RESOURCE_STATUS_ENUM_UNSPECIFIC ResourceStatusEnum = 0
	ResourceStatusEnum_RESOURCE_STATUS_ENUM_STARTING   ResourceStatusEnum = 1
	ResourceStatusEnum_RESOURCE_STATUS_ENUM_RUNNING    ResourceStatusEnum = 2
	ResourceStatusEnum_RESOURCE_STATUS_ENUM_FAILURE    ResourceStatusEnum = 3
	ResourceStatusEnum_RESOURCE_STATUS_ENUM_STOPPING   ResourceStatusEnum = 4
	ResourceStatusEnum_RESOURCE_STATUS_ENUM_DELETED    ResourceStatusEnum = 5
)

// Enum value maps for ResourceStatusEnum.
var (
	ResourceStatusEnum_name = map[int32]string{
		0: "RESOURCE_STATUS_ENUM_UNSPECIFIC",
		1: "RESOURCE_STATUS_ENUM_STARTING",
		2: "RESOURCE_STATUS_ENUM_RUNNING",
		3: "RESOURCE_STATUS_ENUM_FAILURE",
		4: "RESOURCE_STATUS_ENUM_STOPPING",
		5: "RESOURCE_STATUS_ENUM_DELETED",
	}
	ResourceStatusEnum_value = map[string]int32{
		"RESOURCE_STATUS_ENUM_UNSPECIFIC": 0,
		"RESOURCE_STATUS_ENUM_STARTING":   1,
		"RESOURCE_STATUS_ENUM_RUNNING":    2,
		"RESOURCE_STATUS_ENUM_FAILURE":    3,
		"RESOURCE_STATUS_ENUM_STOPPING":   4,
		"RESOURCE_STATUS_ENUM_DELETED":    5,
	}
)

func (x ResourceStatusEnum) Enum() *ResourceStatusEnum {
	p := new(ResourceStatusEnum)
	*p = x
	return p
}

func (x ResourceStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_apiserver_apiserver_proto_enumTypes[1].Descriptor()
}

func (ResourceStatusEnum) Type() protoreflect.EnumType {
	return &file_apiserver_apiserver_proto_enumTypes[1]
}

func (x ResourceStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceStatusEnum.Descriptor instead.
func (ResourceStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{1}
}

type CreateResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string           `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`                                                            // namespace
	ResourceType ResourceTypeEnum `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=apiserver.ResourceTypeEnum" json:"resource_type,omitempty"` // resource type to create
	// Types that are assignable to Options:
	//	*CreateResourceRequest_Dataflow
	Options isCreateResourceRequest_Options `protobuf_oneof:"options"`
}

func (x *CreateResourceRequest) Reset() {
	*x = CreateResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceRequest) ProtoMessage() {}

func (x *CreateResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceRequest) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{0}
}

func (x *CreateResourceRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateResourceRequest) GetResourceType() ResourceTypeEnum {
	if x != nil {
		return x.ResourceType
	}
	return ResourceTypeEnum_RESOURCE_TYPE_ENUM_UNSPECIFIC
}

func (m *CreateResourceRequest) GetOptions() isCreateResourceRequest_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (x *CreateResourceRequest) GetDataflow() *CreateDataflowOptions {
	if x, ok := x.GetOptions().(*CreateResourceRequest_Dataflow); ok {
		return x.Dataflow
	}
	return nil
}

type isCreateResourceRequest_Options interface {
	isCreateResourceRequest_Options()
}

type CreateResourceRequest_Dataflow struct {
	Dataflow *CreateDataflowOptions `protobuf:"bytes,3,opt,name=dataflow,proto3,oneof"`
}

func (*CreateResourceRequest_Dataflow) isCreateResourceRequest_Options() {}

type CreateResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     ResourceStatusEnum `protobuf:"varint,1,opt,name=status,proto3,enum=apiserver.ResourceStatusEnum" json:"status,omitempty"`
	ResourceId *common.ResourceId `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ErrorMsg   string             `protobuf:"bytes,3,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
}

func (x *CreateResourceResponse) Reset() {
	*x = CreateResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceResponse) ProtoMessage() {}

func (x *CreateResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceResponse.ProtoReflect.Descriptor instead.
func (*CreateResourceResponse) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResourceResponse) GetStatus() ResourceStatusEnum {
	if x != nil {
		return x.Status
	}
	return ResourceStatusEnum_RESOURCE_STATUS_ENUM_UNSPECIFIC
}

func (x *CreateResourceResponse) GetResourceId() *common.ResourceId {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

func (x *CreateResourceResponse) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type CreateDataflowOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataflow *stream.Dataflow `protobuf:"bytes,1,opt,name=dataflow,proto3" json:"dataflow,omitempty"`
}

func (x *CreateDataflowOptions) Reset() {
	*x = CreateDataflowOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDataflowOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDataflowOptions) ProtoMessage() {}

func (x *CreateDataflowOptions) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDataflowOptions.ProtoReflect.Descriptor instead.
func (*CreateDataflowOptions) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDataflowOptions) GetDataflow() *stream.Dataflow {
	if x != nil {
		return x.Dataflow
	}
	return nil
}

type ListResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string           `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`                                                            // namespace
	ResourceType ResourceTypeEnum `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=apiserver.ResourceTypeEnum" json:"resource_type,omitempty"` // resource type to list
}

func (x *ListResourcesRequest) Reset() {
	*x = ListResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcesRequest) ProtoMessage() {}

func (x *ListResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcesRequest.ProtoReflect.Descriptor instead.
func (*ListResourcesRequest) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{3}
}

func (x *ListResourcesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListResourcesRequest) GetResourceType() ResourceTypeEnum {
	if x != nil {
		return x.ResourceType
	}
	return ResourceTypeEnum_RESOURCE_TYPE_ENUM_UNSPECIFIC
}

type ListResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*Resource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *ListResourcesResponse) Reset() {
	*x = ListResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcesResponse) ProtoMessage() {}

func (x *ListResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcesResponse.ProtoReflect.Descriptor instead.
func (*ListResourcesResponse) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{4}
}

func (x *ListResourcesResponse) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId   *common.ResourceId `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ResourceName string             `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	ResourceType ResourceTypeEnum   `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=apiserver.ResourceTypeEnum" json:"resource_type,omitempty"` // resource type
	Status       ResourceStatusEnum `protobuf:"varint,4,opt,name=status,proto3,enum=apiserver.ResourceStatusEnum" json:"status,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{5}
}

func (x *Resource) GetResourceId() *common.ResourceId {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

func (x *Resource) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Resource) GetResourceType() ResourceTypeEnum {
	if x != nil {
		return x.ResourceType
	}
	return ResourceTypeEnum_RESOURCE_TYPE_ENUM_UNSPECIFIC
}

func (x *Resource) GetStatus() ResourceStatusEnum {
	if x != nil {
		return x.Status
	}
	return ResourceStatusEnum_RESOURCE_STATUS_ENUM_UNSPECIFIC
}

type GetResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId *common.ResourceId `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *GetResourceRequest) Reset() {
	*x = GetResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceRequest) ProtoMessage() {}

func (x *GetResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceRequest.ProtoReflect.Descriptor instead.
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{6}
}

func (x *GetResourceRequest) GetResourceId() *common.ResourceId {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

type GetResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *GetResourceResponse) Reset() {
	*x = GetResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceResponse) ProtoMessage() {}

func (x *GetResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceResponse.ProtoReflect.Descriptor instead.
func (*GetResourceResponse) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{7}
}

func (x *GetResourceResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type DeleteResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId *common.ResourceId `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteResourceRequest) Reset() {
	*x = DeleteResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceRequest) ProtoMessage() {}

func (x *DeleteResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteResourceRequest) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResourceRequest) GetResourceId() *common.ResourceId {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

type DeleteResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *DeleteResourceResponse) Reset() {
	*x = DeleteResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiserver_apiserver_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceResponse) ProtoMessage() {}

func (x *DeleteResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apiserver_apiserver_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceResponse.ProtoReflect.Descriptor instead.
func (*DeleteResourceResponse) Descriptor() ([]byte, []int) {
	return file_apiserver_apiserver_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteResourceResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_apiserver_apiserver_proto protoreflect.FileDescriptor

var file_apiserver_apiserver_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc2, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x45, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6c, 0x6f, 0x77,
	0x22, 0x76, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22,
	0x46, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x33, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2a, 0x56, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x43, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x2a, 0xe5, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x23, 0x0a, 0x1f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x43, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45,
	0x4e, 0x55, 0x4d, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x20,
	0x0a, 0x1c, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x05,
	0x42, 0x1b, 0x5a, 0x19, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x66, 0x6c, 0x75, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apiserver_apiserver_proto_rawDescOnce sync.Once
	file_apiserver_apiserver_proto_rawDescData = file_apiserver_apiserver_proto_rawDesc
)

func file_apiserver_apiserver_proto_rawDescGZIP() []byte {
	file_apiserver_apiserver_proto_rawDescOnce.Do(func() {
		file_apiserver_apiserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_apiserver_apiserver_proto_rawDescData)
	})
	return file_apiserver_apiserver_proto_rawDescData
}

var file_apiserver_apiserver_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apiserver_apiserver_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_apiserver_apiserver_proto_goTypes = []interface{}{
	(ResourceTypeEnum)(0),          // 0: apiserver.ResourceTypeEnum
	(ResourceStatusEnum)(0),        // 1: apiserver.ResourceStatusEnum
	(*CreateResourceRequest)(nil),  // 2: apiserver.CreateResourceRequest
	(*CreateResourceResponse)(nil), // 3: apiserver.CreateResourceResponse
	(*CreateDataflowOptions)(nil),  // 4: apiserver.CreateDataflowOptions
	(*ListResourcesRequest)(nil),   // 5: apiserver.ListResourcesRequest
	(*ListResourcesResponse)(nil),  // 6: apiserver.ListResourcesResponse
	(*Resource)(nil),               // 7: apiserver.Resource
	(*GetResourceRequest)(nil),     // 8: apiserver.GetResourceRequest
	(*GetResourceResponse)(nil),    // 9: apiserver.GetResourceResponse
	(*DeleteResourceRequest)(nil),  // 10: apiserver.DeleteResourceRequest
	(*DeleteResourceResponse)(nil), // 11: apiserver.DeleteResourceResponse
	(*common.ResourceId)(nil),      // 12: common.ResourceId
	(*stream.Dataflow)(nil),        // 13: common.Dataflow
}
var file_apiserver_apiserver_proto_depIdxs = []int32{
	0,  // 0: apiserver.CreateResourceRequest.resource_type:type_name -> apiserver.ResourceTypeEnum
	4,  // 1: apiserver.CreateResourceRequest.dataflow:type_name -> apiserver.CreateDataflowOptions
	1,  // 2: apiserver.CreateResourceResponse.status:type_name -> apiserver.ResourceStatusEnum
	12, // 3: apiserver.CreateResourceResponse.resource_id:type_name -> common.ResourceId
	13, // 4: apiserver.CreateDataflowOptions.dataflow:type_name -> common.Dataflow
	0,  // 5: apiserver.ListResourcesRequest.resource_type:type_name -> apiserver.ResourceTypeEnum
	7,  // 6: apiserver.ListResourcesResponse.resources:type_name -> apiserver.Resource
	12, // 7: apiserver.Resource.resource_id:type_name -> common.ResourceId
	0,  // 8: apiserver.Resource.resource_type:type_name -> apiserver.ResourceTypeEnum
	1,  // 9: apiserver.Resource.status:type_name -> apiserver.ResourceStatusEnum
	12, // 10: apiserver.GetResourceRequest.resource_id:type_name -> common.ResourceId
	7,  // 11: apiserver.GetResourceResponse.resource:type_name -> apiserver.Resource
	12, // 12: apiserver.DeleteResourceRequest.resource_id:type_name -> common.ResourceId
	7,  // 13: apiserver.DeleteResourceResponse.resource:type_name -> apiserver.Resource
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_apiserver_apiserver_proto_init() }
func file_apiserver_apiserver_proto_init() {
	if File_apiserver_apiserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apiserver_apiserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceRequest); i {
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
		file_apiserver_apiserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceResponse); i {
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
		file_apiserver_apiserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDataflowOptions); i {
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
		file_apiserver_apiserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcesRequest); i {
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
		file_apiserver_apiserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcesResponse); i {
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
		file_apiserver_apiserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_apiserver_apiserver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceRequest); i {
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
		file_apiserver_apiserver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceResponse); i {
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
		file_apiserver_apiserver_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceRequest); i {
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
		file_apiserver_apiserver_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceResponse); i {
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
	file_apiserver_apiserver_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CreateResourceRequest_Dataflow)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apiserver_apiserver_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apiserver_apiserver_proto_goTypes,
		DependencyIndexes: file_apiserver_apiserver_proto_depIdxs,
		EnumInfos:         file_apiserver_apiserver_proto_enumTypes,
		MessageInfos:      file_apiserver_apiserver_proto_msgTypes,
	}.Build()
	File_apiserver_apiserver_proto = out.File
	file_apiserver_apiserver_proto_rawDesc = nil
	file_apiserver_apiserver_proto_goTypes = nil
	file_apiserver_apiserver_proto_depIdxs = nil
}