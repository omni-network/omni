// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: monitor/xmonitor/indexer/indexer.proto

package indexer

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                      // Auto-incremented ID
	ChainId       uint64                 `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`             // Source chain ID as per https://chainlist.org
	BlockHeight   uint64                 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"` // Height of the source-chain block
	BlockHash     []byte                 `protobuf:"bytes,4,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`        // Hash of the source-chain block
	BlockJson     []byte                 `protobuf:"bytes,5,opt,name=block_json,json=blockJson,proto3" json:"block_json,omitempty"`        // xchain.Block JSON
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Block) Reset() {
	*x = Block{}
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_monitor_xmonitor_indexer_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Block) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Block) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *Block) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Block) GetBlockJson() []byte {
	if x != nil {
		return x.BlockJson
	}
	return nil
}

type MsgLink struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	IdHash         []byte                 `protobuf:"bytes,1,opt,name=id_hash,json=idHash,proto3" json:"id_hash,omitempty"` // RouteScan IDHash of the MsgID
	MsgBlockId     uint64                 `protobuf:"varint,2,opt,name=msg_block_id,json=msgBlockId,proto3" json:"msg_block_id,omitempty"`
	ReceiptBlockId uint64                 `protobuf:"varint,3,opt,name=receipt_block_id,json=receiptBlockId,proto3" json:"receipt_block_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MsgLink) Reset() {
	*x = MsgLink{}
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgLink) ProtoMessage() {}

func (x *MsgLink) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgLink.ProtoReflect.Descriptor instead.
func (*MsgLink) Descriptor() ([]byte, []int) {
	return file_monitor_xmonitor_indexer_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *MsgLink) GetIdHash() []byte {
	if x != nil {
		return x.IdHash
	}
	return nil
}

func (x *MsgLink) GetMsgBlockId() uint64 {
	if x != nil {
		return x.MsgBlockId
	}
	return 0
}

func (x *MsgLink) GetReceiptBlockId() uint64 {
	if x != nil {
		return x.ReceiptBlockId
	}
	return 0
}

type Cursor struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChainId       uint64                 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ConfLevel     uint32                 `protobuf:"varint,2,opt,name=conf_level,json=confLevel,proto3" json:"conf_level,omitempty"`
	BlockHeight   uint64                 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_monitor_xmonitor_indexer_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *Cursor) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Cursor) GetConfLevel() uint32 {
	if x != nil {
		return x.ConfLevel
	}
	return 0
}

func (x *Cursor) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

var File_monitor_xmonitor_indexer_indexer_proto protoreflect.FileDescriptor

const file_monitor_xmonitor_indexer_indexer_proto_rawDesc = "" +
	"\n" +
	"&monitor/xmonitor/indexer/indexer.proto\x12\x18monitor.xmonitor.indexer\x1a\x17cosmos/orm/v1/orm.proto\"\xcd\x01\n" +
	"\x05Block\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x19\n" +
	"\bchain_id\x18\x02 \x01(\x04R\achainId\x12!\n" +
	"\fblock_height\x18\x03 \x01(\x04R\vblockHeight\x12\x1d\n" +
	"\n" +
	"block_hash\x18\x04 \x01(\fR\tblockHash\x12\x1d\n" +
	"\n" +
	"block_json\x18\x05 \x01(\fR\tblockJson:8\xf2\x9eӎ\x032\n" +
	"\x06\n" +
	"\x02id\x10\x01\x12&\n" +
	" chain_id,block_height,block_hash\x10\x02\x18\x01\x18\x01\"\x83\x01\n" +
	"\aMsgLink\x12\x17\n" +
	"\aid_hash\x18\x01 \x01(\fR\x06idHash\x12 \n" +
	"\fmsg_block_id\x18\x02 \x01(\x04R\n" +
	"msgBlockId\x12(\n" +
	"\x10receipt_block_id\x18\x03 \x01(\x04R\x0ereceiptBlockId:\x13\xf2\x9eӎ\x03\r\n" +
	"\t\n" +
	"\aid_hash\x18\x02\"\x86\x01\n" +
	"\x06Cursor\x12\x19\n" +
	"\bchain_id\x18\x01 \x01(\x04R\achainId\x12\x1d\n" +
	"\n" +
	"conf_level\x18\x02 \x01(\rR\tconfLevel\x12!\n" +
	"\fblock_height\x18\x03 \x01(\x04R\vblockHeight:\x1f\xf2\x9eӎ\x03\x19\n" +
	"\x15\n" +
	"\x13chain_id,conf_level\x18\x03B\xe5\x01\n" +
	"\x1ccom.monitor.xmonitor.indexerB\fIndexerProtoP\x01Z5github.com/omni-network/omni/monitor/xmonitor/indexer\xa2\x02\x03MXI\xaa\x02\x18Monitor.Xmonitor.Indexer\xca\x02\x18Monitor\\Xmonitor\\Indexer\xe2\x02$Monitor\\Xmonitor\\Indexer\\GPBMetadata\xea\x02\x1aMonitor::Xmonitor::Indexerb\x06proto3"

var (
	file_monitor_xmonitor_indexer_indexer_proto_rawDescOnce sync.Once
	file_monitor_xmonitor_indexer_indexer_proto_rawDescData []byte
)

func file_monitor_xmonitor_indexer_indexer_proto_rawDescGZIP() []byte {
	file_monitor_xmonitor_indexer_indexer_proto_rawDescOnce.Do(func() {
		file_monitor_xmonitor_indexer_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_monitor_xmonitor_indexer_indexer_proto_rawDesc), len(file_monitor_xmonitor_indexer_indexer_proto_rawDesc)))
	})
	return file_monitor_xmonitor_indexer_indexer_proto_rawDescData
}

var file_monitor_xmonitor_indexer_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_monitor_xmonitor_indexer_indexer_proto_goTypes = []any{
	(*Block)(nil),   // 0: monitor.xmonitor.indexer.Block
	(*MsgLink)(nil), // 1: monitor.xmonitor.indexer.MsgLink
	(*Cursor)(nil),  // 2: monitor.xmonitor.indexer.Cursor
}
var file_monitor_xmonitor_indexer_indexer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_monitor_xmonitor_indexer_indexer_proto_init() }
func file_monitor_xmonitor_indexer_indexer_proto_init() {
	if File_monitor_xmonitor_indexer_indexer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_monitor_xmonitor_indexer_indexer_proto_rawDesc), len(file_monitor_xmonitor_indexer_indexer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_monitor_xmonitor_indexer_indexer_proto_goTypes,
		DependencyIndexes: file_monitor_xmonitor_indexer_indexer_proto_depIdxs,
		MessageInfos:      file_monitor_xmonitor_indexer_indexer_proto_msgTypes,
	}.Build()
	File_monitor_xmonitor_indexer_indexer_proto = out.File
	file_monitor_xmonitor_indexer_indexer_proto_goTypes = nil
	file_monitor_xmonitor_indexer_indexer_proto_depIdxs = nil
}
