// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: graveler.proto

package graveler

import (
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

type BranchProtectionBlockedAction int32

const (
	BranchProtectionBlockedAction_STAGING_WRITE BranchProtectionBlockedAction = 0
	BranchProtectionBlockedAction_COMMIT        BranchProtectionBlockedAction = 1
)

// Enum value maps for BranchProtectionBlockedAction.
var (
	BranchProtectionBlockedAction_name = map[int32]string{
		0: "STAGING_WRITE",
		1: "COMMIT",
	}
	BranchProtectionBlockedAction_value = map[string]int32{
		"STAGING_WRITE": 0,
		"COMMIT":        1,
	}
)

func (x BranchProtectionBlockedAction) Enum() *BranchProtectionBlockedAction {
	p := new(BranchProtectionBlockedAction)
	*p = x
	return p
}

func (x BranchProtectionBlockedAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BranchProtectionBlockedAction) Descriptor() protoreflect.EnumDescriptor {
	return file_graveler_proto_enumTypes[0].Descriptor()
}

func (BranchProtectionBlockedAction) Type() protoreflect.EnumType {
	return &file_graveler_proto_enumTypes[0]
}

func (x BranchProtectionBlockedAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BranchProtectionBlockedAction.Descriptor instead.
func (BranchProtectionBlockedAction) EnumDescriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{0}
}

type RepositoryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StorageNamespace string                 `protobuf:"bytes,2,opt,name=storage_namespace,json=storageNamespace,proto3" json:"storage_namespace,omitempty"`
	DefaultBranchId  string                 `protobuf:"bytes,3,opt,name=default_branch_id,json=defaultBranchId,proto3" json:"default_branch_id,omitempty"`
	CreationDate     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
}

func (x *RepositoryData) Reset() {
	*x = RepositoryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryData) ProtoMessage() {}

func (x *RepositoryData) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryData.ProtoReflect.Descriptor instead.
func (*RepositoryData) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{0}
}

func (x *RepositoryData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RepositoryData) GetStorageNamespace() string {
	if x != nil {
		return x.StorageNamespace
	}
	return ""
}

func (x *RepositoryData) GetDefaultBranchId() string {
	if x != nil {
		return x.DefaultBranchId
	}
	return ""
}

func (x *RepositoryData) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

type BranchData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommitId     string   `protobuf:"bytes,2,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	StagingToken string   `protobuf:"bytes,3,opt,name=staging_token,json=stagingToken,proto3" json:"staging_token,omitempty"`
	SealedTokens []string `protobuf:"bytes,4,rep,name=sealed_tokens,json=sealedTokens,proto3" json:"sealed_tokens,omitempty"`
}

func (x *BranchData) Reset() {
	*x = BranchData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchData) ProtoMessage() {}

func (x *BranchData) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchData.ProtoReflect.Descriptor instead.
func (*BranchData) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{1}
}

func (x *BranchData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BranchData) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *BranchData) GetStagingToken() string {
	if x != nil {
		return x.StagingToken
	}
	return ""
}

func (x *BranchData) GetSealedTokens() []string {
	if x != nil {
		return x.SealedTokens
	}
	return nil
}

type TagData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommitId string `protobuf:"bytes,2,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
}

func (x *TagData) Reset() {
	*x = TagData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagData) ProtoMessage() {}

func (x *TagData) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagData.ProtoReflect.Descriptor instead.
func (*TagData) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{2}
}

func (x *TagData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagData) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

type CommitData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Committer    string                 `protobuf:"bytes,2,opt,name=committer,proto3" json:"committer,omitempty"`
	Message      string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	CreationDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	MetaRangeId  string                 `protobuf:"bytes,5,opt,name=meta_range_id,json=metaRangeId,proto3" json:"meta_range_id,omitempty"`
	Metadata     map[string]string      `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Parents      []string               `protobuf:"bytes,7,rep,name=parents,proto3" json:"parents,omitempty"`
	Version      int32                  `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	Generation   int32                  `protobuf:"varint,9,opt,name=generation,proto3" json:"generation,omitempty"`
}

func (x *CommitData) Reset() {
	*x = CommitData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitData) ProtoMessage() {}

func (x *CommitData) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitData.ProtoReflect.Descriptor instead.
func (*CommitData) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{3}
}

func (x *CommitData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommitData) GetCommitter() string {
	if x != nil {
		return x.Committer
	}
	return ""
}

func (x *CommitData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommitData) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *CommitData) GetMetaRangeId() string {
	if x != nil {
		return x.MetaRangeId
	}
	return ""
}

func (x *CommitData) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CommitData) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

func (x *CommitData) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CommitData) GetGeneration() int32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

type GarbageCollectionRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultRetentionDays int32            `protobuf:"varint,1,opt,name=default_retention_days,json=defaultRetentionDays,proto3" json:"default_retention_days,omitempty"`
	BranchRetentionDays  map[string]int32 `protobuf:"bytes,2,rep,name=branch_retention_days,json=branchRetentionDays,proto3" json:"branch_retention_days,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GarbageCollectionRules) Reset() {
	*x = GarbageCollectionRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarbageCollectionRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarbageCollectionRules) ProtoMessage() {}

func (x *GarbageCollectionRules) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarbageCollectionRules.ProtoReflect.Descriptor instead.
func (*GarbageCollectionRules) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{4}
}

func (x *GarbageCollectionRules) GetDefaultRetentionDays() int32 {
	if x != nil {
		return x.DefaultRetentionDays
	}
	return 0
}

func (x *GarbageCollectionRules) GetBranchRetentionDays() map[string]int32 {
	if x != nil {
		return x.BranchRetentionDays
	}
	return nil
}

type GarbageCollectionRunMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId              string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	CommitsCsvLocation string `protobuf:"bytes,2,opt,name=Commits_csv_location,json=CommitsCsvLocation,proto3" json:"Commits_csv_location,omitempty"`
	AddressLocation    string `protobuf:"bytes,3,opt,name=Address_location,json=AddressLocation,proto3" json:"Address_location,omitempty"`
}

func (x *GarbageCollectionRunMetadata) Reset() {
	*x = GarbageCollectionRunMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarbageCollectionRunMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarbageCollectionRunMetadata) ProtoMessage() {}

func (x *GarbageCollectionRunMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarbageCollectionRunMetadata.ProtoReflect.Descriptor instead.
func (*GarbageCollectionRunMetadata) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{5}
}

func (x *GarbageCollectionRunMetadata) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *GarbageCollectionRunMetadata) GetCommitsCsvLocation() string {
	if x != nil {
		return x.CommitsCsvLocation
	}
	return ""
}

func (x *GarbageCollectionRunMetadata) GetAddressLocation() string {
	if x != nil {
		return x.AddressLocation
	}
	return ""
}

type BranchProtectionBlockedActions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []BranchProtectionBlockedAction `protobuf:"varint,1,rep,packed,name=value,proto3,enum=io.treeverse.lakefs.graveler.BranchProtectionBlockedAction" json:"value,omitempty"`
}

func (x *BranchProtectionBlockedActions) Reset() {
	*x = BranchProtectionBlockedActions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchProtectionBlockedActions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchProtectionBlockedActions) ProtoMessage() {}

func (x *BranchProtectionBlockedActions) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchProtectionBlockedActions.ProtoReflect.Descriptor instead.
func (*BranchProtectionBlockedActions) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{6}
}

func (x *BranchProtectionBlockedActions) GetValue() []BranchProtectionBlockedAction {
	if x != nil {
		return x.Value
	}
	return nil
}

type BranchProtectionRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchPatternToBlockedActions map[string]*BranchProtectionBlockedActions `protobuf:"bytes,1,rep,name=branch_pattern_to_blocked_actions,json=branchPatternToBlockedActions,proto3" json:"branch_pattern_to_blocked_actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BranchProtectionRules) Reset() {
	*x = BranchProtectionRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchProtectionRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchProtectionRules) ProtoMessage() {}

func (x *BranchProtectionRules) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchProtectionRules.ProtoReflect.Descriptor instead.
func (*BranchProtectionRules) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{7}
}

func (x *BranchProtectionRules) GetBranchPatternToBlockedActions() map[string]*BranchProtectionBlockedActions {
	if x != nil {
		return x.BranchPatternToBlockedActions
	}
	return nil
}

type StagedEntryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Identity []byte `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StagedEntryData) Reset() {
	*x = StagedEntryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graveler_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StagedEntryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StagedEntryData) ProtoMessage() {}

func (x *StagedEntryData) ProtoReflect() protoreflect.Message {
	mi := &file_graveler_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StagedEntryData.ProtoReflect.Descriptor instead.
func (*StagedEntryData) Descriptor() ([]byte, []int) {
	return file_graveler_proto_rawDescGZIP(), []int{8}
}

func (x *StagedEntryData) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *StagedEntryData) GetIdentity() []byte {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *StagedEntryData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_graveler_proto protoreflect.FileDescriptor

var file_graveler_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x6c,
	0x61, 0x6b, 0x65, 0x66, 0x73, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xba, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x83, 0x01, 0x0a,
	0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x22, 0x36, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x22, 0x9e, 0x03, 0x0a, 0x0a, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x69, 0x6f, 0x2e, 0x74, 0x72,
	0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x6c, 0x61, 0x6b, 0x65, 0x66, 0x73, 0x2e, 0x67,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9a, 0x02, 0x0a, 0x16,
	0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x12, 0x81, 0x01, 0x0a,
	0x15, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x69,
	0x6f, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x6c, 0x61, 0x6b, 0x65,
	0x66, 0x73, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x72, 0x62,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x79, 0x73,
	0x1a, 0x46, 0x0a, 0x18, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x92, 0x01, 0x0a, 0x1c, 0x47, 0x61, 0x72,
	0x62, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64,
	0x12, 0x30, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x5f, 0x63, 0x73, 0x76, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x43, 0x73, 0x76, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a,
	0x1e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x51, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x3b,
	0x2e, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x6c, 0x61,
	0x6b, 0x65, 0x66, 0x73, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xcb, 0x02, 0x0a, 0x15, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0xa0, 0x01, 0x0a,
	0x21, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x5f,
	0x74, 0x6f, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x56, 0x2e, 0x69, 0x6f, 0x2e, 0x74, 0x72,
	0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x6c, 0x61, 0x6b, 0x65, 0x66, 0x73, 0x2e, 0x67,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x54, 0x6f, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x1d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x54,
	0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x8e, 0x01, 0x0a, 0x22, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x54, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x52, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x65,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x6c, 0x61, 0x6b, 0x65, 0x66, 0x73, 0x2e, 0x67, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x53, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x67, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x3e, 0x0a, 0x1d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x47, 0x49, 0x4e,
	0x47, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d,
	0x4d, 0x49, 0x54, 0x10, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x6c, 0x61,
	0x6b, 0x65, 0x66, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graveler_proto_rawDescOnce sync.Once
	file_graveler_proto_rawDescData = file_graveler_proto_rawDesc
)

func file_graveler_proto_rawDescGZIP() []byte {
	file_graveler_proto_rawDescOnce.Do(func() {
		file_graveler_proto_rawDescData = protoimpl.X.CompressGZIP(file_graveler_proto_rawDescData)
	})
	return file_graveler_proto_rawDescData
}

var file_graveler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_graveler_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_graveler_proto_goTypes = []interface{}{
	(BranchProtectionBlockedAction)(0),     // 0: io.treeverse.lakefs.graveler.BranchProtectionBlockedAction
	(*RepositoryData)(nil),                 // 1: io.treeverse.lakefs.graveler.RepositoryData
	(*BranchData)(nil),                     // 2: io.treeverse.lakefs.graveler.BranchData
	(*TagData)(nil),                        // 3: io.treeverse.lakefs.graveler.TagData
	(*CommitData)(nil),                     // 4: io.treeverse.lakefs.graveler.CommitData
	(*GarbageCollectionRules)(nil),         // 5: io.treeverse.lakefs.graveler.GarbageCollectionRules
	(*GarbageCollectionRunMetadata)(nil),   // 6: io.treeverse.lakefs.graveler.GarbageCollectionRunMetadata
	(*BranchProtectionBlockedActions)(nil), // 7: io.treeverse.lakefs.graveler.BranchProtectionBlockedActions
	(*BranchProtectionRules)(nil),          // 8: io.treeverse.lakefs.graveler.BranchProtectionRules
	(*StagedEntryData)(nil),                // 9: io.treeverse.lakefs.graveler.StagedEntryData
	nil,                                    // 10: io.treeverse.lakefs.graveler.CommitData.MetadataEntry
	nil,                                    // 11: io.treeverse.lakefs.graveler.GarbageCollectionRules.BranchRetentionDaysEntry
	nil,                                    // 12: io.treeverse.lakefs.graveler.BranchProtectionRules.BranchPatternToBlockedActionsEntry
	(*timestamppb.Timestamp)(nil),          // 13: google.protobuf.Timestamp
}
var file_graveler_proto_depIdxs = []int32{
	13, // 0: io.treeverse.lakefs.graveler.RepositoryData.creation_date:type_name -> google.protobuf.Timestamp
	13, // 1: io.treeverse.lakefs.graveler.CommitData.creation_date:type_name -> google.protobuf.Timestamp
	10, // 2: io.treeverse.lakefs.graveler.CommitData.metadata:type_name -> io.treeverse.lakefs.graveler.CommitData.MetadataEntry
	11, // 3: io.treeverse.lakefs.graveler.GarbageCollectionRules.branch_retention_days:type_name -> io.treeverse.lakefs.graveler.GarbageCollectionRules.BranchRetentionDaysEntry
	0,  // 4: io.treeverse.lakefs.graveler.BranchProtectionBlockedActions.value:type_name -> io.treeverse.lakefs.graveler.BranchProtectionBlockedAction
	12, // 5: io.treeverse.lakefs.graveler.BranchProtectionRules.branch_pattern_to_blocked_actions:type_name -> io.treeverse.lakefs.graveler.BranchProtectionRules.BranchPatternToBlockedActionsEntry
	7,  // 6: io.treeverse.lakefs.graveler.BranchProtectionRules.BranchPatternToBlockedActionsEntry.value:type_name -> io.treeverse.lakefs.graveler.BranchProtectionBlockedActions
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_graveler_proto_init() }
func file_graveler_proto_init() {
	if File_graveler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graveler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryData); i {
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
		file_graveler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchData); i {
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
		file_graveler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagData); i {
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
		file_graveler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitData); i {
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
		file_graveler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarbageCollectionRules); i {
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
		file_graveler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarbageCollectionRunMetadata); i {
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
		file_graveler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchProtectionBlockedActions); i {
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
		file_graveler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchProtectionRules); i {
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
		file_graveler_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StagedEntryData); i {
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
			RawDescriptor: file_graveler_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_graveler_proto_goTypes,
		DependencyIndexes: file_graveler_proto_depIdxs,
		EnumInfos:         file_graveler_proto_enumTypes,
		MessageInfos:      file_graveler_proto_msgTypes,
	}.Build()
	File_graveler_proto = out.File
	file_graveler_proto_rawDesc = nil
	file_graveler_proto_goTypes = nil
	file_graveler_proto_depIdxs = nil
}
