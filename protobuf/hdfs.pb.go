// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: hdfs.proto

package protobuf

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

// Request and Response messages for NameNodeService
type RegisterDataNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatanodeAddress string `protobuf:"bytes,1,opt,name=datanode_address,json=datanodeAddress,proto3" json:"datanode_address,omitempty"`
}

func (x *RegisterDataNodeRequest) Reset() {
	*x = RegisterDataNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdfs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDataNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDataNodeRequest) ProtoMessage() {}

func (x *RegisterDataNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hdfs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDataNodeRequest.ProtoReflect.Descriptor instead.
func (*RegisterDataNodeRequest) Descriptor() ([]byte, []int) {
	return file_hdfs_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterDataNodeRequest) GetDatanodeAddress() string {
	if x != nil {
		return x.DatanodeAddress
	}
	return ""
}

type RegisterDataNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	DatanodeId string `protobuf:"bytes,2,opt,name=datanode_id,json=datanodeId,proto3" json:"datanode_id,omitempty"` // The ID assigned by the NameNode
}

func (x *RegisterDataNodeResponse) Reset() {
	*x = RegisterDataNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdfs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDataNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDataNodeResponse) ProtoMessage() {}

func (x *RegisterDataNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hdfs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDataNodeResponse.ProtoReflect.Descriptor instead.
func (*RegisterDataNodeResponse) Descriptor() ([]byte, []int) {
	return file_hdfs_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterDataNodeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterDataNodeResponse) GetDatanodeId() string {
	if x != nil {
		return x.DatanodeId
	}
	return ""
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatanodeId string `protobuf:"bytes,1,opt,name=datanode_id,json=datanodeId,proto3" json:"datanode_id,omitempty"` // Identifier for the DataNode
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdfs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hdfs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_hdfs_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatRequest) GetDatanodeId() string {
	if x != nil {
		return x.DatanodeId
	}
	return ""
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdfs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hdfs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_hdfs_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request and Response messages for DataNodeService
type StoreBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId   string `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	BlockData []byte `protobuf:"bytes,2,opt,name=block_data,json=blockData,proto3" json:"block_data,omitempty"`
}

func (x *StoreBlockRequest) Reset() {
	*x = StoreBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdfs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreBlockRequest) ProtoMessage() {}

func (x *StoreBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hdfs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreBlockRequest.ProtoReflect.Descriptor instead.
func (*StoreBlockRequest) Descriptor() ([]byte, []int) {
	return file_hdfs_proto_rawDescGZIP(), []int{4}
}

func (x *StoreBlockRequest) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *StoreBlockRequest) GetBlockData() []byte {
	if x != nil {
		return x.BlockData
	}
	return nil
}

type StoreBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StoreBlockResponse) Reset() {
	*x = StoreBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdfs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreBlockResponse) ProtoMessage() {}

func (x *StoreBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hdfs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreBlockResponse.ProtoReflect.Descriptor instead.
func (*StoreBlockResponse) Descriptor() ([]byte, []int) {
	return file_hdfs_proto_rawDescGZIP(), []int{5}
}

func (x *StoreBlockResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RetrieveBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId string `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
}

func (x *RetrieveBlockRequest) Reset() {
	*x = RetrieveBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdfs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveBlockRequest) ProtoMessage() {}

func (x *RetrieveBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hdfs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveBlockRequest.ProtoReflect.Descriptor instead.
func (*RetrieveBlockRequest) Descriptor() ([]byte, []int) {
	return file_hdfs_proto_rawDescGZIP(), []int{6}
}

func (x *RetrieveBlockRequest) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

type RetrieveBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	BlockData []byte `protobuf:"bytes,2,opt,name=block_data,json=blockData,proto3" json:"block_data,omitempty"` // The data of the block being retrieved
}

func (x *RetrieveBlockResponse) Reset() {
	*x = RetrieveBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdfs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveBlockResponse) ProtoMessage() {}

func (x *RetrieveBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hdfs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveBlockResponse.ProtoReflect.Descriptor instead.
func (*RetrieveBlockResponse) Descriptor() ([]byte, []int) {
	return file_hdfs_proto_rawDescGZIP(), []int{7}
}

func (x *RetrieveBlockResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RetrieveBlockResponse) GetBlockData() []byte {
	if x != nil {
		return x.BlockData
	}
	return nil
}

var File_hdfs_proto protoreflect.FileDescriptor

var file_hdfs_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x64,
	0x66, 0x73, 0x22, 0x44, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x55, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22,
	0x33, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x4d, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x31, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x15, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x32, 0xaa, 0x01, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x12, 0x16, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x64, 0x66, 0x73,
	0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0xa0, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x68,
	0x64, 0x66, 0x73, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x61, 0x72, 0x72, 0x61, 0x73, 0x73, 0x65, 0x61, 0x79,
	0x6f, 0x75, 0x62, 0x30, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hdfs_proto_rawDescOnce sync.Once
	file_hdfs_proto_rawDescData = file_hdfs_proto_rawDesc
)

func file_hdfs_proto_rawDescGZIP() []byte {
	file_hdfs_proto_rawDescOnce.Do(func() {
		file_hdfs_proto_rawDescData = protoimpl.X.CompressGZIP(file_hdfs_proto_rawDescData)
	})
	return file_hdfs_proto_rawDescData
}

var file_hdfs_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_hdfs_proto_goTypes = []interface{}{
	(*RegisterDataNodeRequest)(nil),  // 0: hdfs.RegisterDataNodeRequest
	(*RegisterDataNodeResponse)(nil), // 1: hdfs.RegisterDataNodeResponse
	(*HeartbeatRequest)(nil),         // 2: hdfs.HeartbeatRequest
	(*HeartbeatResponse)(nil),        // 3: hdfs.HeartbeatResponse
	(*StoreBlockRequest)(nil),        // 4: hdfs.StoreBlockRequest
	(*StoreBlockResponse)(nil),       // 5: hdfs.StoreBlockResponse
	(*RetrieveBlockRequest)(nil),     // 6: hdfs.RetrieveBlockRequest
	(*RetrieveBlockResponse)(nil),    // 7: hdfs.RetrieveBlockResponse
}
var file_hdfs_proto_depIdxs = []int32{
	0, // 0: hdfs.NameNodeService.RegisterDataNode:input_type -> hdfs.RegisterDataNodeRequest
	2, // 1: hdfs.NameNodeService.SendHeartbeat:input_type -> hdfs.HeartbeatRequest
	4, // 2: hdfs.DataNodeService.StoreBlock:input_type -> hdfs.StoreBlockRequest
	6, // 3: hdfs.DataNodeService.RetrieveBlock:input_type -> hdfs.RetrieveBlockRequest
	1, // 4: hdfs.NameNodeService.RegisterDataNode:output_type -> hdfs.RegisterDataNodeResponse
	3, // 5: hdfs.NameNodeService.SendHeartbeat:output_type -> hdfs.HeartbeatResponse
	5, // 6: hdfs.DataNodeService.StoreBlock:output_type -> hdfs.StoreBlockResponse
	7, // 7: hdfs.DataNodeService.RetrieveBlock:output_type -> hdfs.RetrieveBlockResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hdfs_proto_init() }
func file_hdfs_proto_init() {
	if File_hdfs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hdfs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDataNodeRequest); i {
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
		file_hdfs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDataNodeResponse); i {
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
		file_hdfs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_hdfs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_hdfs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreBlockRequest); i {
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
		file_hdfs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreBlockResponse); i {
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
		file_hdfs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveBlockRequest); i {
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
		file_hdfs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveBlockResponse); i {
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
			RawDescriptor: file_hdfs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_hdfs_proto_goTypes,
		DependencyIndexes: file_hdfs_proto_depIdxs,
		MessageInfos:      file_hdfs_proto_msgTypes,
	}.Build()
	File_hdfs_proto = out.File
	file_hdfs_proto_rawDesc = nil
	file_hdfs_proto_goTypes = nil
	file_hdfs_proto_depIdxs = nil
}
