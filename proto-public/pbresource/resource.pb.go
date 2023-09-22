// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbresource/resource.proto

// For more information, see: https://github.com/hashicorp/consul/tree/main/docs/resources

package pbresource

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// State represents the state of the condition (i.e. true/false/unknown).
type Condition_State int32

const (
	// STATE_UNKNOWN means that the state of the condition is unknown.
	//
	// buf:lint:ignore ENUM_ZERO_VALUE_SUFFIX
	Condition_STATE_UNKNOWN Condition_State = 0
	// STATE_TRUE means that the state of the condition is true.
	Condition_STATE_TRUE Condition_State = 1
	// STATE_FALSE means that the state of the condition is false.
	Condition_STATE_FALSE Condition_State = 2
)

// Enum value maps for Condition_State.
var (
	Condition_State_name = map[int32]string{
		0: "STATE_UNKNOWN",
		1: "STATE_TRUE",
		2: "STATE_FALSE",
	}
	Condition_State_value = map[string]int32{
		"STATE_UNKNOWN": 0,
		"STATE_TRUE":    1,
		"STATE_FALSE":   2,
	}
)

func (x Condition_State) Enum() *Condition_State {
	p := new(Condition_State)
	*p = x
	return p
}

func (x Condition_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition_State) Descriptor() protoreflect.EnumDescriptor {
	return file_pbresource_resource_proto_enumTypes[0].Descriptor()
}

func (Condition_State) Type() protoreflect.EnumType {
	return &file_pbresource_resource_proto_enumTypes[0]
}

func (x Condition_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition_State.Descriptor instead.
func (Condition_State) EnumDescriptor() ([]byte, []int) {
	return file_pbresource_resource_proto_rawDescGZIP(), []int{5, 0}
}

// Type describes a resource's type. It follows the GVK (Group Version Kind)
// [pattern](https://book.kubebuilder.io/cronjob-tutorial/gvks.html) established
// by Kubernetes.
type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Group describes the area of functionality to which this resource type
	// relates (e.g. "catalog", "authorization").
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// GroupVersion is incremented when sweeping or backward-incompatible changes
	// are made to the group's resource types.
	GroupVersion string `protobuf:"bytes,2,opt,name=group_version,json=groupVersion,proto3" json:"group_version,omitempty"`
	// Kind identifies the specific resource type within the group.
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbresource_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_pbresource_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_pbresource_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Type) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Type) GetGroupVersion() string {
	if x != nil {
		return x.GroupVersion
	}
	return ""
}

func (x *Type) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

// Tenancy describes the tenancy units in which the resource resides.
type Tenancy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Partition is the topmost administrative boundary within a cluster.
	// https://developer.hashicorp.com/consul/docs/enterprise/admin-partitions
	//
	// When using the List and WatchList endpoints, provide the wildcard value "*"
	// to list resources across all partitions.
	Partition string `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	// Namespace further isolates resources within a partition.
	// https://developer.hashicorp.com/consul/docs/enterprise/namespaces
	//
	// When using the List and WatchList endpoints, provide the wildcard value "*"
	// to list resources across all namespaces.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// PeerName identifies which peer the resource is imported from.
	// https://developer.hashicorp.com/consul/docs/connect/cluster-peering
	//
	// When using the List and WatchList endpoints, provide the wildcard value "*"
	// to list resources across all peers.
	PeerName string `protobuf:"bytes,3,opt,name=peer_name,json=peerName,proto3" json:"peer_name,omitempty"`
}

func (x *Tenancy) Reset() {
	*x = Tenancy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbresource_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenancy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenancy) ProtoMessage() {}

func (x *Tenancy) ProtoReflect() protoreflect.Message {
	mi := &file_pbresource_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenancy.ProtoReflect.Descriptor instead.
func (*Tenancy) Descriptor() ([]byte, []int) {
	return file_pbresource_resource_proto_rawDescGZIP(), []int{1}
}

func (x *Tenancy) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *Tenancy) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Tenancy) GetPeerName() string {
	if x != nil {
		return x.PeerName
	}
	return ""
}

// ID uniquely identifies a resource.
type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Uid is the unique internal identifier we gave to the resource.
	//
	// It is primarily used to tell the difference between the current resource
	// and previous deleted resources with the same user-given name.
	//
	// Concretely, Uid is a [ULID](https://github.com/ulid/spec) and you can treat
	// its timestamp component as the resource's creation time.
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// Name is the user-given name of the resource (e.g. the "billing" service).
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Type identifies the resource's type.
	Type *Type `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Tenancy identifies the tenancy units (i.e. partition, namespace) in which
	// the resource resides.
	Tenancy *Tenancy `protobuf:"bytes,4,opt,name=tenancy,proto3" json:"tenancy,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbresource_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_pbresource_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_pbresource_resource_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ID) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ID) GetTenancy() *Tenancy {
	if x != nil {
		return x.Tenancy
	}
	return nil
}

// Resource describes a resource of a known type managed by Consul.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID uniquely identifies the resource.
	Id *ID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Owner (optionally) describes which resource "owns" this resource, it is
	// immutable and can only be set on resource creation. Owned resources will
	// be automatically deleted when their owner is deleted.
	Owner *ID `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// Version is the low-level version identifier used by the storage backend
	// in CAS (Compare-And-Swap) operations. It will change when the resource is
	// modified in any way, including status updates.
	//
	// When calling the Write endpoint, providing a non-blank version will perform
	// a CAS (Compare-And-Swap) write, which will result in an Aborted error code
	// if the given version doesn't match what is stored.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Generation is incremented whenever the resource's content (i.e. not its
	// status) is modified. You can think of it as being the "user version".
	//
	// Concretely, Generation is a [ULID](https://github.com/ulid/spec) and you
	// can treat its timestamp component as the resource's modification time.
	Generation string `protobuf:"bytes,4,opt,name=generation,proto3" json:"generation,omitempty"`
	// Metadata contains key/value pairs of arbitrary metadata about the resource.
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Status is used by controllers to communicate the result of attempting to
	// reconcile and apply the resource (e.g. surface semantic validation errors)
	// with users and other controllers. Each status is identified by a unique key
	// and should only ever be updated by one controller.
	//
	// Status can only be updated via the WriteStatus endpoint. Attempting to do
	// so via the Write endpoint will result in an InvalidArgument error code.
	Status map[string]*Status `protobuf:"bytes,6,rep,name=status,proto3" json:"status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Data contains the resource's type-specific content.
	Data *anypb.Any `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbresource_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_pbresource_resource_proto_msgTypes[3]
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
	return file_pbresource_resource_proto_rawDescGZIP(), []int{3}
}

func (x *Resource) GetId() *ID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Resource) GetOwner() *ID {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Resource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Resource) GetGeneration() string {
	if x != nil {
		return x.Generation
	}
	return ""
}

func (x *Resource) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Resource) GetStatus() map[string]*Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Resource) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

// Status is used by controllers to communicate the result of attempting to
// reconcile and apply a resource (e.g. surface semantic validation errors)
// with users and other controllers.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ObservedGeneration identifies which generation of a resource this status
	// related to. It can be used to determine whether the current generation of
	// a resource has been reconciled.
	ObservedGeneration string `protobuf:"bytes,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Conditions contains a set of discreet observations about the resource in
	// relation to the current state of the system (e.g. it is semantically valid).
	Conditions []*Condition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// UpdatedAt is the time at which the status was last written.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbresource_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_pbresource_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_pbresource_resource_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetObservedGeneration() string {
	if x != nil {
		return x.ObservedGeneration
	}
	return ""
}

func (x *Status) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Status) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Condition represents a discreet observation about a resource in relation to
// the current state of the system.
//
// It is heavily inspired by Kubernetes' [conditions](https://bit.ly/3H9Y6IK)
// and the Gateway API [types and reasons](https://bit.ly/3n2PPiP).
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type identifies the type of condition (e.g. "Invalid", "ResolvedRefs").
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// State represents the state of the condition (i.e. true/false/unknown).
	State Condition_State `protobuf:"varint,2,opt,name=state,proto3,enum=hashicorp.consul.resource.Condition_State" json:"state,omitempty"`
	// Reason provides more machine-readable details about the condition (e.g.
	// "InvalidProtocol").
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	// Message contains a human-friendly description of the status.
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Resource identifies which resource this condition relates to, when it is
	// not the core resource itself.
	Resource *Reference `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbresource_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_pbresource_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_pbresource_resource_proto_rawDescGZIP(), []int{5}
}

func (x *Condition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Condition) GetState() Condition_State {
	if x != nil {
		return x.State
	}
	return Condition_STATE_UNKNOWN
}

func (x *Condition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Condition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Condition) GetResource() *Reference {
	if x != nil {
		return x.Resource
	}
	return nil
}

// Reference identifies which resource a condition relates to, when it is not
// the core resource itself.
type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type identifies the resource's type.
	Type *Type `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Tenancy identifies the tenancy units (i.e. partition, namespace) in which
	// the resource resides.
	Tenancy *Tenancy `protobuf:"bytes,2,opt,name=tenancy,proto3" json:"tenancy,omitempty"`
	// Name is the user-given name of the resource (e.g. the "billing" service).
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Section identifies which part of the resource the condition relates to.
	Section string `protobuf:"bytes,4,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbresource_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_pbresource_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_pbresource_resource_proto_rawDescGZIP(), []int{6}
}

func (x *Reference) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Reference) GetTenancy() *Tenancy {
	if x != nil {
		return x.Tenancy
	}
	return nil
}

func (x *Reference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Reference) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

var File_pbresource_resource_proto protoreflect.FileDescriptor

var file_pbresource_resource_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x55, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x23, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x62, 0x0a, 0x07, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9d, 0x01,
	0x0a, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x3c, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x79, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x22, 0x85, 0x04,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xba, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x92, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x52, 0x55, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x02, 0x22, 0xac, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x52,
	0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xe9, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xa2, 0x02,
	0x03, 0x48, 0x43, 0x52, 0xaa, 0x02, 0x19, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0xca, 0x02, 0x19, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xe2, 0x02, 0x25, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbresource_resource_proto_rawDescOnce sync.Once
	file_pbresource_resource_proto_rawDescData = file_pbresource_resource_proto_rawDesc
)

func file_pbresource_resource_proto_rawDescGZIP() []byte {
	file_pbresource_resource_proto_rawDescOnce.Do(func() {
		file_pbresource_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbresource_resource_proto_rawDescData)
	})
	return file_pbresource_resource_proto_rawDescData
}

var file_pbresource_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pbresource_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pbresource_resource_proto_goTypes = []interface{}{
	(Condition_State)(0),          // 0: hashicorp.consul.resource.Condition.State
	(*Type)(nil),                  // 1: hashicorp.consul.resource.Type
	(*Tenancy)(nil),               // 2: hashicorp.consul.resource.Tenancy
	(*ID)(nil),                    // 3: hashicorp.consul.resource.ID
	(*Resource)(nil),              // 4: hashicorp.consul.resource.Resource
	(*Status)(nil),                // 5: hashicorp.consul.resource.Status
	(*Condition)(nil),             // 6: hashicorp.consul.resource.Condition
	(*Reference)(nil),             // 7: hashicorp.consul.resource.Reference
	nil,                           // 8: hashicorp.consul.resource.Resource.MetadataEntry
	nil,                           // 9: hashicorp.consul.resource.Resource.StatusEntry
	(*anypb.Any)(nil),             // 10: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_pbresource_resource_proto_depIdxs = []int32{
	1,  // 0: hashicorp.consul.resource.ID.type:type_name -> hashicorp.consul.resource.Type
	2,  // 1: hashicorp.consul.resource.ID.tenancy:type_name -> hashicorp.consul.resource.Tenancy
	3,  // 2: hashicorp.consul.resource.Resource.id:type_name -> hashicorp.consul.resource.ID
	3,  // 3: hashicorp.consul.resource.Resource.owner:type_name -> hashicorp.consul.resource.ID
	8,  // 4: hashicorp.consul.resource.Resource.metadata:type_name -> hashicorp.consul.resource.Resource.MetadataEntry
	9,  // 5: hashicorp.consul.resource.Resource.status:type_name -> hashicorp.consul.resource.Resource.StatusEntry
	10, // 6: hashicorp.consul.resource.Resource.data:type_name -> google.protobuf.Any
	6,  // 7: hashicorp.consul.resource.Status.conditions:type_name -> hashicorp.consul.resource.Condition
	11, // 8: hashicorp.consul.resource.Status.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 9: hashicorp.consul.resource.Condition.state:type_name -> hashicorp.consul.resource.Condition.State
	7,  // 10: hashicorp.consul.resource.Condition.resource:type_name -> hashicorp.consul.resource.Reference
	1,  // 11: hashicorp.consul.resource.Reference.type:type_name -> hashicorp.consul.resource.Type
	2,  // 12: hashicorp.consul.resource.Reference.tenancy:type_name -> hashicorp.consul.resource.Tenancy
	5,  // 13: hashicorp.consul.resource.Resource.StatusEntry.value:type_name -> hashicorp.consul.resource.Status
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_pbresource_resource_proto_init() }
func file_pbresource_resource_proto_init() {
	if File_pbresource_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbresource_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
		file_pbresource_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenancy); i {
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
		file_pbresource_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_pbresource_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbresource_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_pbresource_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_pbresource_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
			RawDescriptor: file_pbresource_resource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbresource_resource_proto_goTypes,
		DependencyIndexes: file_pbresource_resource_proto_depIdxs,
		EnumInfos:         file_pbresource_resource_proto_enumTypes,
		MessageInfos:      file_pbresource_resource_proto_msgTypes,
	}.Build()
	File_pbresource_resource_proto = out.File
	file_pbresource_resource_proto_rawDesc = nil
	file_pbresource_resource_proto_goTypes = nil
	file_pbresource_resource_proto_depIdxs = nil
}
