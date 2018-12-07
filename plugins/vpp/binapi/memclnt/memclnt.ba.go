// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: /usr/share/vpp/api/memclnt.api.json

/*
 Package memclnt is a generated from VPP binary API module 'memclnt'.

 It contains following objects:
	 22 messages
	  2 types
	 13 services

*/
package memclnt

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Types */

// ModuleVersion represents the VPP binary API type 'module_version'.
//
//            "module_version",
//            [
//                "u32",
//                "major"
//            ],
//            [
//                "u32",
//                "minor"
//            ],
//            [
//                "u32",
//                "patch"
//            ],
//            [
//                "u8",
//                "name",
//                64
//            ],
//            {
//                "crc": "0x4b6da11a"
//            }
//
type ModuleVersion struct {
	Major uint32
	Minor uint32
	Patch uint32
	Name  []byte `struc:"[64]byte"`
}

func (*ModuleVersion) GetTypeName() string {
	return "module_version"
}
func (*ModuleVersion) GetCrcString() string {
	return "4b6da11a"
}

// MessageTableEntry represents the VPP binary API type 'message_table_entry'.
//
//            "message_table_entry",
//            [
//                "u16",
//                "index"
//            ],
//            [
//                "u8",
//                "name",
//                64
//            ],
//            {
//                "crc": "0x913bf1c6"
//            }
//
type MessageTableEntry struct {
	Index uint16
	Name  []byte `struc:"[64]byte"`
}

func (*MessageTableEntry) GetTypeName() string {
	return "message_table_entry"
}
func (*MessageTableEntry) GetCrcString() string {
	return "913bf1c6"
}

/* Messages */

// MemclntCreate represents the VPP binary API message 'memclnt_create'.
//
//            "memclnt_create",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "ctx_quota"
//            ],
//            [
//                "u64",
//                "input_queue"
//            ],
//            [
//                "u8",
//                "name",
//                64
//            ],
//            [
//                "u32",
//                "api_versions",
//                8
//            ],
//            {
//                "crc": "0x6d33c5ea"
//            }
//
type MemclntCreate struct {
	CtxQuota    int32
	InputQueue  uint64
	Name        []byte   `struc:"[64]byte"`
	APIVersions []uint32 `struc:"[8]uint32"`
}

func (*MemclntCreate) GetMessageName() string {
	return "memclnt_create"
}
func (*MemclntCreate) GetCrcString() string {
	return "6d33c5ea"
}
func (*MemclntCreate) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemclntCreateReply represents the VPP binary API message 'memclnt_create_reply'.
//
//            "memclnt_create_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "response"
//            ],
//            [
//                "u64",
//                "handle"
//            ],
//            [
//                "u32",
//                "index"
//            ],
//            [
//                "u64",
//                "message_table"
//            ],
//            {
//                "crc": "0x42ec4560"
//            }
//
type MemclntCreateReply struct {
	Response     int32
	Handle       uint64
	Index        uint32
	MessageTable uint64
}

func (*MemclntCreateReply) GetMessageName() string {
	return "memclnt_create_reply"
}
func (*MemclntCreateReply) GetCrcString() string {
	return "42ec4560"
}
func (*MemclntCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemclntDelete represents the VPP binary API message 'memclnt_delete'.
//
//            "memclnt_delete",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "index"
//            ],
//            [
//                "u64",
//                "handle"
//            ],
//            {
//                "crc": "0x73240f13"
//            }
//
type MemclntDelete struct {
	Index  uint32
	Handle uint64
}

func (*MemclntDelete) GetMessageName() string {
	return "memclnt_delete"
}
func (*MemclntDelete) GetCrcString() string {
	return "73240f13"
}
func (*MemclntDelete) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// MemclntDeleteReply represents the VPP binary API message 'memclnt_delete_reply'.
//
//            "memclnt_delete_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "i32",
//                "response"
//            ],
//            [
//                "u64",
//                "handle"
//            ],
//            {
//                "crc": "0x3d3b6312"
//            }
//
type MemclntDeleteReply struct {
	Response int32
	Handle   uint64
}

func (*MemclntDeleteReply) GetMessageName() string {
	return "memclnt_delete_reply"
}
func (*MemclntDeleteReply) GetCrcString() string {
	return "3d3b6312"
}
func (*MemclntDeleteReply) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// RxThreadExit represents the VPP binary API message 'rx_thread_exit'.
//
//            "rx_thread_exit",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u8",
//                "dummy"
//            ],
//            {
//                "crc": "0xc3a3a452"
//            }
//
type RxThreadExit struct {
	Dummy uint8
}

func (*RxThreadExit) GetMessageName() string {
	return "rx_thread_exit"
}
func (*RxThreadExit) GetCrcString() string {
	return "c3a3a452"
}
func (*RxThreadExit) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// MemclntRxThreadSuspend represents the VPP binary API message 'memclnt_rx_thread_suspend'.
//
//            "memclnt_rx_thread_suspend",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u8",
//                "dummy"
//            ],
//            {
//                "crc": "0xc3a3a452"
//            }
//
type MemclntRxThreadSuspend struct {
	Dummy uint8
}

func (*MemclntRxThreadSuspend) GetMessageName() string {
	return "memclnt_rx_thread_suspend"
}
func (*MemclntRxThreadSuspend) GetCrcString() string {
	return "c3a3a452"
}
func (*MemclntRxThreadSuspend) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// MemclntReadTimeout represents the VPP binary API message 'memclnt_read_timeout'.
//
//            "memclnt_read_timeout",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u8",
//                "dummy"
//            ],
//            {
//                "crc": "0xc3a3a452"
//            }
//
type MemclntReadTimeout struct {
	Dummy uint8
}

func (*MemclntReadTimeout) GetMessageName() string {
	return "memclnt_read_timeout"
}
func (*MemclntReadTimeout) GetCrcString() string {
	return "c3a3a452"
}
func (*MemclntReadTimeout) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// RPCCall represents the VPP binary API message 'rpc_call'.
//
//            "rpc_call",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u64",
//                "function"
//            ],
//            [
//                "u8",
//                "multicast"
//            ],
//            [
//                "u8",
//                "need_barrier_sync"
//            ],
//            [
//                "u8",
//                "send_reply"
//            ],
//            [
//                "u32",
//                "data_len"
//            ],
//            [
//                "u8",
//                "data",
//                0,
//                "data_len"
//            ],
//            {
//                "crc": "0x7e8a2c95"
//            }
//
type RPCCall struct {
	Function        uint64
	Multicast       uint8
	NeedBarrierSync uint8
	SendReply       uint8
	DataLen         uint32 `struc:"sizeof=Data"`
	Data            []byte
}

func (*RPCCall) GetMessageName() string {
	return "rpc_call"
}
func (*RPCCall) GetCrcString() string {
	return "7e8a2c95"
}
func (*RPCCall) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// RPCCallReply represents the VPP binary API message 'rpc_call_reply'.
//
//            "rpc_call_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type RPCCallReply struct {
	Retval int32
}

func (*RPCCallReply) GetMessageName() string {
	return "rpc_call_reply"
}
func (*RPCCallReply) GetCrcString() string {
	return "e8d4e804"
}
func (*RPCCallReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GetFirstMsgID represents the VPP binary API message 'get_first_msg_id'.
//
//            "get_first_msg_id",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "name",
//                64
//            ],
//            {
//                "crc": "0x0cb71b0e"
//            }
//
type GetFirstMsgID struct {
	Name []byte `struc:"[64]byte"`
}

func (*GetFirstMsgID) GetMessageName() string {
	return "get_first_msg_id"
}
func (*GetFirstMsgID) GetCrcString() string {
	return "0cb71b0e"
}
func (*GetFirstMsgID) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GetFirstMsgIDReply represents the VPP binary API message 'get_first_msg_id_reply'.
//
//            "get_first_msg_id_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u16",
//                "first_msg_id"
//            ],
//            {
//                "crc": "0x7d337472"
//            }
//
type GetFirstMsgIDReply struct {
	Retval     int32
	FirstMsgID uint16
}

func (*GetFirstMsgIDReply) GetMessageName() string {
	return "get_first_msg_id_reply"
}
func (*GetFirstMsgIDReply) GetCrcString() string {
	return "7d337472"
}
func (*GetFirstMsgIDReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// APIVersions represents the VPP binary API message 'api_versions'.
//
//            "api_versions",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type APIVersions struct{}

func (*APIVersions) GetMessageName() string {
	return "api_versions"
}
func (*APIVersions) GetCrcString() string {
	return "51077d14"
}
func (*APIVersions) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// APIVersionsReply represents the VPP binary API message 'api_versions_reply'.
//
//            "api_versions_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u32",
//                "count"
//            ],
//            [
//                "vl_api_module_version_t",
//                "api_versions",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0x90a39195"
//            }
//
type APIVersionsReply struct {
	Retval      int32
	Count       uint32 `struc:"sizeof=APIVersions"`
	APIVersions []ModuleVersion
}

func (*APIVersionsReply) GetMessageName() string {
	return "api_versions_reply"
}
func (*APIVersionsReply) GetCrcString() string {
	return "90a39195"
}
func (*APIVersionsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TracePluginMsgIds represents the VPP binary API message 'trace_plugin_msg_ids'.
//
//            "trace_plugin_msg_ids",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "plugin_name",
//                128
//            ],
//            [
//                "u16",
//                "first_msg_id"
//            ],
//            [
//                "u16",
//                "last_msg_id"
//            ],
//            {
//                "crc": "0x64af79f9"
//            }
//
type TracePluginMsgIds struct {
	PluginName []byte `struc:"[128]byte"`
	FirstMsgID uint16
	LastMsgID  uint16
}

func (*TracePluginMsgIds) GetMessageName() string {
	return "trace_plugin_msg_ids"
}
func (*TracePluginMsgIds) GetCrcString() string {
	return "64af79f9"
}
func (*TracePluginMsgIds) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SockclntCreate represents the VPP binary API message 'sockclnt_create'.
//
//            "sockclnt_create",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "name",
//                64
//            ],
//            {
//                "crc": "0xdf2cf94d"
//            }
//
type SockclntCreate struct {
	Name []byte `struc:"[64]byte"`
}

func (*SockclntCreate) GetMessageName() string {
	return "sockclnt_create"
}
func (*SockclntCreate) GetCrcString() string {
	return "df2cf94d"
}
func (*SockclntCreate) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SockclntCreateReply represents the VPP binary API message 'sockclnt_create_reply'.
//
//            "sockclnt_create_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "response"
//            ],
//            [
//                "u32",
//                "index"
//            ],
//            [
//                "u16",
//                "count"
//            ],
//            [
//                "vl_api_message_table_entry_t",
//                "message_table",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0xa134a8a8"
//            }
//
type SockclntCreateReply struct {
	Response     int32
	Index        uint32
	Count        uint16 `struc:"sizeof=MessageTable"`
	MessageTable []MessageTableEntry
}

func (*SockclntCreateReply) GetMessageName() string {
	return "sockclnt_create_reply"
}
func (*SockclntCreateReply) GetCrcString() string {
	return "a134a8a8"
}
func (*SockclntCreateReply) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SockclntDelete represents the VPP binary API message 'sockclnt_delete'.
//
//            "sockclnt_delete",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "index"
//            ],
//            {
//                "crc": "0x8ac76db6"
//            }
//
type SockclntDelete struct {
	Index uint32
}

func (*SockclntDelete) GetMessageName() string {
	return "sockclnt_delete"
}
func (*SockclntDelete) GetCrcString() string {
	return "8ac76db6"
}
func (*SockclntDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SockclntDeleteReply represents the VPP binary API message 'sockclnt_delete_reply'.
//
//            "sockclnt_delete_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "response"
//            ],
//            {
//                "crc": "0x8f38b1ee"
//            }
//
type SockclntDeleteReply struct {
	Response int32
}

func (*SockclntDeleteReply) GetMessageName() string {
	return "sockclnt_delete_reply"
}
func (*SockclntDeleteReply) GetCrcString() string {
	return "8f38b1ee"
}
func (*SockclntDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SockInitShm represents the VPP binary API message 'sock_init_shm'.
//
//            "sock_init_shm",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "requested_size"
//            ],
//            [
//                "u8",
//                "nitems"
//            ],
//            [
//                "u64",
//                "configs",
//                0,
//                "nitems"
//            ],
//            {
//                "crc": "0x51646d92"
//            }
//
type SockInitShm struct {
	RequestedSize uint32
	Nitems        uint8 `struc:"sizeof=Configs"`
	Configs       []uint64
}

func (*SockInitShm) GetMessageName() string {
	return "sock_init_shm"
}
func (*SockInitShm) GetCrcString() string {
	return "51646d92"
}
func (*SockInitShm) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SockInitShmReply represents the VPP binary API message 'sock_init_shm_reply'.
//
//            "sock_init_shm_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SockInitShmReply struct {
	Retval int32
}

func (*SockInitShmReply) GetMessageName() string {
	return "sock_init_shm_reply"
}
func (*SockInitShmReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SockInitShmReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemclntKeepalive represents the VPP binary API message 'memclnt_keepalive'.
//
//            "memclnt_keepalive",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type MemclntKeepalive struct{}

func (*MemclntKeepalive) GetMessageName() string {
	return "memclnt_keepalive"
}
func (*MemclntKeepalive) GetCrcString() string {
	return "51077d14"
}
func (*MemclntKeepalive) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemclntKeepaliveReply represents the VPP binary API message 'memclnt_keepalive_reply'.
//
//            "memclnt_keepalive_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type MemclntKeepaliveReply struct {
	Retval int32
}

func (*MemclntKeepaliveReply) GetMessageName() string {
	return "memclnt_keepalive_reply"
}
func (*MemclntKeepaliveReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MemclntKeepaliveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

/* Services */

type Services interface {
	APIVersions(*APIVersions) (*APIVersionsReply, error)
	GetFirstMsgID(*GetFirstMsgID) (*GetFirstMsgIDReply, error)
	MemclntCreate(*MemclntCreate) (*MemclntCreateReply, error)
	MemclntDelete(*MemclntDelete) (*MemclntDeleteReply, error)
	MemclntKeepalive(*MemclntKeepalive) (*MemclntKeepaliveReply, error)
	MemclntReadTimeout(*MemclntReadTimeout) error
	MemclntRxThreadSuspend(*MemclntRxThreadSuspend) error
	RPCCall(*RPCCall) (*RPCCallReply, error)
	RxThreadExit(*RxThreadExit) error
	SockInitShm(*SockInitShm) (*SockInitShmReply, error)
	SockclntCreate(*SockclntCreate) (*SockclntCreateReply, error)
	SockclntDelete(*SockclntDelete) (*SockclntDeleteReply, error)
	TracePluginMsgIds(*TracePluginMsgIds) error
}

func init() {
	api.RegisterMessage((*MemclntCreate)(nil), "memclnt.MemclntCreate")
	api.RegisterMessage((*MemclntCreateReply)(nil), "memclnt.MemclntCreateReply")
	api.RegisterMessage((*MemclntDelete)(nil), "memclnt.MemclntDelete")
	api.RegisterMessage((*MemclntDeleteReply)(nil), "memclnt.MemclntDeleteReply")
	api.RegisterMessage((*RxThreadExit)(nil), "memclnt.RxThreadExit")
	api.RegisterMessage((*MemclntRxThreadSuspend)(nil), "memclnt.MemclntRxThreadSuspend")
	api.RegisterMessage((*MemclntReadTimeout)(nil), "memclnt.MemclntReadTimeout")
	api.RegisterMessage((*RPCCall)(nil), "memclnt.RPCCall")
	api.RegisterMessage((*RPCCallReply)(nil), "memclnt.RPCCallReply")
	api.RegisterMessage((*GetFirstMsgID)(nil), "memclnt.GetFirstMsgID")
	api.RegisterMessage((*GetFirstMsgIDReply)(nil), "memclnt.GetFirstMsgIDReply")
	api.RegisterMessage((*APIVersions)(nil), "memclnt.APIVersions")
	api.RegisterMessage((*APIVersionsReply)(nil), "memclnt.APIVersionsReply")
	api.RegisterMessage((*TracePluginMsgIds)(nil), "memclnt.TracePluginMsgIds")
	api.RegisterMessage((*SockclntCreate)(nil), "memclnt.SockclntCreate")
	api.RegisterMessage((*SockclntCreateReply)(nil), "memclnt.SockclntCreateReply")
	api.RegisterMessage((*SockclntDelete)(nil), "memclnt.SockclntDelete")
	api.RegisterMessage((*SockclntDeleteReply)(nil), "memclnt.SockclntDeleteReply")
	api.RegisterMessage((*SockInitShm)(nil), "memclnt.SockInitShm")
	api.RegisterMessage((*SockInitShmReply)(nil), "memclnt.SockInitShmReply")
	api.RegisterMessage((*MemclntKeepalive)(nil), "memclnt.MemclntKeepalive")
	api.RegisterMessage((*MemclntKeepaliveReply)(nil), "memclnt.MemclntKeepaliveReply")
}
