// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: lib/usdt0/db.proto

package usdt0

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Match layer zero message status types
type MsgStatus int32

const (
	MsgStatus_MSG_STATUS_UNKNOWN        MsgStatus = 0
	MsgStatus_MSG_STATUS_CONFIRMING     MsgStatus = 1
	MsgStatus_MSG_STATUS_INFLIGHT       MsgStatus = 2
	MsgStatus_MSG_STATUS_DELIVERED      MsgStatus = 3
	MsgStatus_MSG_STATUS_FAILED         MsgStatus = 4
	MsgStatus_MSG_STATUS_PAYLOAD_STORED MsgStatus = 6
)

// Enum value maps for MsgStatus.
var (
	MsgStatus_name = map[int32]string{
		0: "MSG_STATUS_UNKNOWN",
		1: "MSG_STATUS_CONFIRMING",
		2: "MSG_STATUS_INFLIGHT",
		3: "MSG_STATUS_DELIVERED",
		4: "MSG_STATUS_FAILED",
		6: "MSG_STATUS_PAYLOAD_STORED",
	}
	MsgStatus_value = map[string]int32{
		"MSG_STATUS_UNKNOWN":        0,
		"MSG_STATUS_CONFIRMING":     1,
		"MSG_STATUS_INFLIGHT":       2,
		"MSG_STATUS_DELIVERED":      3,
		"MSG_STATUS_FAILED":         4,
		"MSG_STATUS_PAYLOAD_STORED": 6,
	}
)

func (x MsgStatus) Enum() *MsgStatus {
	p := new(MsgStatus)
	*p = x
	return p
}

func (x MsgStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_lib_usdt0_db_proto_enumTypes[0].Descriptor()
}

func (MsgStatus) Type() protoreflect.EnumType {
	return &file_lib_usdt0_db_proto_enumTypes[0]
}

func (x MsgStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgStatus.Descriptor instead.
func (MsgStatus) EnumDescriptor() ([]byte, []int) {
	return file_lib_usdt0_db_proto_rawDescGZIP(), []int{0}
}

type MsgSendUSDT0 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxHash        []byte                 `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`                   // Source tx hash
	BlockHeight   uint64                 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`   // Height of soruce chain block
	SrcChainId    uint64                 `protobuf:"varint,3,opt,name=src_chain_id,json=srcChainId,proto3" json:"src_chain_id,omitempty"`    // Source chain ID
	DestChainId   uint64                 `protobuf:"varint,4,opt,name=dest_chain_id,json=destChainId,proto3" json:"dest_chain_id,omitempty"` // Destination chain ID
	Amount        []byte                 `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`                                 // Amount of USDT0 sent
	Status        int32                  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`                                // Message status
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`          // Creation timestamp
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MsgSendUSDT0) Reset() {
	*x = MsgSendUSDT0{}
	mi := &file_lib_usdt0_db_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgSendUSDT0) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSendUSDT0) ProtoMessage() {}

func (x *MsgSendUSDT0) ProtoReflect() protoreflect.Message {
	mi := &file_lib_usdt0_db_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSendUSDT0.ProtoReflect.Descriptor instead.
func (*MsgSendUSDT0) Descriptor() ([]byte, []int) {
	return file_lib_usdt0_db_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSendUSDT0) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *MsgSendUSDT0) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *MsgSendUSDT0) GetSrcChainId() uint64 {
	if x != nil {
		return x.SrcChainId
	}
	return 0
}

func (x *MsgSendUSDT0) GetDestChainId() uint64 {
	if x != nil {
		return x.DestChainId
	}
	return 0
}

func (x *MsgSendUSDT0) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *MsgSendUSDT0) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MsgSendUSDT0) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_lib_usdt0_db_proto protoreflect.FileDescriptor

const file_lib_usdt0_db_proto_rawDesc = "" +
	"\n" +
	"\x12lib/usdt0/db.proto\x12\tlib.usdt0\x1a\x17cosmos/orm/v1/orm.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x90\x02\n" +
	"\fMsgSendUSDT0\x12\x17\n" +
	"\atx_hash\x18\x01 \x01(\fR\x06txHash\x12!\n" +
	"\fblock_height\x18\x02 \x01(\x04R\vblockHeight\x12 \n" +
	"\fsrc_chain_id\x18\x03 \x01(\x04R\n" +
	"srcChainId\x12\"\n" +
	"\rdest_chain_id\x18\x04 \x01(\x04R\vdestChainId\x12\x16\n" +
	"\x06amount\x18\x05 \x01(\fR\x06amount\x12\x16\n" +
	"\x06status\x18\x06 \x01(\x05R\x06status\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt:\x13\xf2\x9eӎ\x03\r\n" +
	"\t\n" +
	"\atx_hash\x18\x01*\xa7\x01\n" +
	"\tMsgStatus\x12\x16\n" +
	"\x12MSG_STATUS_UNKNOWN\x10\x00\x12\x19\n" +
	"\x15MSG_STATUS_CONFIRMING\x10\x01\x12\x17\n" +
	"\x13MSG_STATUS_INFLIGHT\x10\x02\x12\x18\n" +
	"\x14MSG_STATUS_DELIVERED\x10\x03\x12\x15\n" +
	"\x11MSG_STATUS_FAILED\x10\x04\x12\x1d\n" +
	"\x19MSG_STATUS_PAYLOAD_STORED\x10\x06B\x85\x01\n" +
	"\rcom.lib.usdt0B\aDbProtoP\x01Z&github.com/omni-network/omni/lib/usdt0\xa2\x02\x03LUX\xaa\x02\tLib.Usdt0\xca\x02\tLib\\Usdt0\xe2\x02\x15Lib\\Usdt0\\GPBMetadata\xea\x02\n" +
	"Lib::Usdt0b\x06proto3"

var (
	file_lib_usdt0_db_proto_rawDescOnce sync.Once
	file_lib_usdt0_db_proto_rawDescData []byte
)

func file_lib_usdt0_db_proto_rawDescGZIP() []byte {
	file_lib_usdt0_db_proto_rawDescOnce.Do(func() {
		file_lib_usdt0_db_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lib_usdt0_db_proto_rawDesc), len(file_lib_usdt0_db_proto_rawDesc)))
	})
	return file_lib_usdt0_db_proto_rawDescData
}

var file_lib_usdt0_db_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lib_usdt0_db_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lib_usdt0_db_proto_goTypes = []any{
	(MsgStatus)(0),                // 0: lib.usdt0.MsgStatus
	(*MsgSendUSDT0)(nil),          // 1: lib.usdt0.MsgSendUSDT0
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_lib_usdt0_db_proto_depIdxs = []int32{
	2, // 0: lib.usdt0.MsgSendUSDT0.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lib_usdt0_db_proto_init() }
func file_lib_usdt0_db_proto_init() {
	if File_lib_usdt0_db_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lib_usdt0_db_proto_rawDesc), len(file_lib_usdt0_db_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lib_usdt0_db_proto_goTypes,
		DependencyIndexes: file_lib_usdt0_db_proto_depIdxs,
		EnumInfos:         file_lib_usdt0_db_proto_enumTypes,
		MessageInfos:      file_lib_usdt0_db_proto_msgTypes,
	}.Build()
	File_lib_usdt0_db_proto = out.File
	file_lib_usdt0_db_proto_goTypes = nil
	file_lib_usdt0_db_proto_depIdxs = nil
}
