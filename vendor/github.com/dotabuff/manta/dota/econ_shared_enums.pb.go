// Code generated by protoc-gen-go.
// source: econ_shared_enums.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EGCEconBaseMsg int32

const (
	EGCEconBaseMsg_k_EMsgGCGenericResult EGCEconBaseMsg = 2579
)

var EGCEconBaseMsg_name = map[int32]string{
	2579: "k_EMsgGCGenericResult",
}
var EGCEconBaseMsg_value = map[string]int32{
	"k_EMsgGCGenericResult": 2579,
}

func (x EGCEconBaseMsg) Enum() *EGCEconBaseMsg {
	p := new(EGCEconBaseMsg)
	*p = x
	return p
}
func (x EGCEconBaseMsg) String() string {
	return proto.EnumName(EGCEconBaseMsg_name, int32(x))
}
func (x *EGCEconBaseMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCEconBaseMsg_value, data, "EGCEconBaseMsg")
	if err != nil {
		return err
	}
	*x = EGCEconBaseMsg(value)
	return nil
}
func (EGCEconBaseMsg) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

type EGCMsgResponse int32

const (
	EGCMsgResponse_k_EGCMsgResponseOK           EGCMsgResponse = 0
	EGCMsgResponse_k_EGCMsgResponseDenied       EGCMsgResponse = 1
	EGCMsgResponse_k_EGCMsgResponseServerError  EGCMsgResponse = 2
	EGCMsgResponse_k_EGCMsgResponseTimeout      EGCMsgResponse = 3
	EGCMsgResponse_k_EGCMsgResponseInvalid      EGCMsgResponse = 4
	EGCMsgResponse_k_EGCMsgResponseNoMatch      EGCMsgResponse = 5
	EGCMsgResponse_k_EGCMsgResponseUnknownError EGCMsgResponse = 6
	EGCMsgResponse_k_EGCMsgResponseNotLoggedOn  EGCMsgResponse = 7
	EGCMsgResponse_k_EGCMsgFailedToCreate       EGCMsgResponse = 8
)

var EGCMsgResponse_name = map[int32]string{
	0: "k_EGCMsgResponseOK",
	1: "k_EGCMsgResponseDenied",
	2: "k_EGCMsgResponseServerError",
	3: "k_EGCMsgResponseTimeout",
	4: "k_EGCMsgResponseInvalid",
	5: "k_EGCMsgResponseNoMatch",
	6: "k_EGCMsgResponseUnknownError",
	7: "k_EGCMsgResponseNotLoggedOn",
	8: "k_EGCMsgFailedToCreate",
}
var EGCMsgResponse_value = map[string]int32{
	"k_EGCMsgResponseOK":           0,
	"k_EGCMsgResponseDenied":       1,
	"k_EGCMsgResponseServerError":  2,
	"k_EGCMsgResponseTimeout":      3,
	"k_EGCMsgResponseInvalid":      4,
	"k_EGCMsgResponseNoMatch":      5,
	"k_EGCMsgResponseUnknownError": 6,
	"k_EGCMsgResponseNotLoggedOn":  7,
	"k_EGCMsgFailedToCreate":       8,
}

func (x EGCMsgResponse) Enum() *EGCMsgResponse {
	p := new(EGCMsgResponse)
	*p = x
	return p
}
func (x EGCMsgResponse) String() string {
	return proto.EnumName(EGCMsgResponse_name, int32(x))
}
func (x *EGCMsgResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCMsgResponse_value, data, "EGCMsgResponse")
	if err != nil {
		return err
	}
	*x = EGCMsgResponse(value)
	return nil
}
func (EGCMsgResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{1} }

type EGCPartnerRequestResponse int32

const (
	EGCPartnerRequestResponse_k_EPartnerRequestOK                     EGCPartnerRequestResponse = 1
	EGCPartnerRequestResponse_k_EPartnerRequestBadAccount             EGCPartnerRequestResponse = 2
	EGCPartnerRequestResponse_k_EPartnerRequestNotLinked              EGCPartnerRequestResponse = 3
	EGCPartnerRequestResponse_k_EPartnerRequestUnsupportedPartnerType EGCPartnerRequestResponse = 4
)

var EGCPartnerRequestResponse_name = map[int32]string{
	1: "k_EPartnerRequestOK",
	2: "k_EPartnerRequestBadAccount",
	3: "k_EPartnerRequestNotLinked",
	4: "k_EPartnerRequestUnsupportedPartnerType",
}
var EGCPartnerRequestResponse_value = map[string]int32{
	"k_EPartnerRequestOK":                     1,
	"k_EPartnerRequestBadAccount":             2,
	"k_EPartnerRequestNotLinked":              3,
	"k_EPartnerRequestUnsupportedPartnerType": 4,
}

func (x EGCPartnerRequestResponse) Enum() *EGCPartnerRequestResponse {
	p := new(EGCPartnerRequestResponse)
	*p = x
	return p
}
func (x EGCPartnerRequestResponse) String() string {
	return proto.EnumName(EGCPartnerRequestResponse_name, int32(x))
}
func (x *EGCPartnerRequestResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCPartnerRequestResponse_value, data, "EGCPartnerRequestResponse")
	if err != nil {
		return err
	}
	*x = EGCPartnerRequestResponse(value)
	return nil
}
func (EGCPartnerRequestResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{2} }

type EGCMsgUseItemResponse int32

const (
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed                    EGCMsgUseItemResponse = 0
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_GiftNoOtherPlayers          EGCMsgUseItemResponse = 1
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ServerError                 EGCMsgUseItemResponse = 2
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_MiniGameAlreadyStarted      EGCMsgUseItemResponse = 3
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted       EGCMsgUseItemResponse = 4
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted EGCMsgUseItemResponse = 5
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_NotInLowPriorityPool        EGCMsgUseItemResponse = 6
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_NotHighEnoughLevel          EGCMsgUseItemResponse = 7
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EventNotActive              EGCMsgUseItemResponse = 8
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted EGCMsgUseItemResponse = 9
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_MissingRequirement          EGCMsgUseItemResponse = 10
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew        EGCMsgUseItemResponse = 11
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EmoticonUnlock_Complete     EGCMsgUseItemResponse = 12
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_Compendium         EGCMsgUseItemResponse = 13
)

var EGCMsgUseItemResponse_name = map[int32]string{
	0:  "k_EGCMsgUseItemResponse_ItemUsed",
	1:  "k_EGCMsgUseItemResponse_GiftNoOtherPlayers",
	2:  "k_EGCMsgUseItemResponse_ServerError",
	3:  "k_EGCMsgUseItemResponse_MiniGameAlreadyStarted",
	4:  "k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted",
	5:  "k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted",
	6:  "k_EGCMsgUseItemResponse_NotInLowPriorityPool",
	7:  "k_EGCMsgUseItemResponse_NotHighEnoughLevel",
	8:  "k_EGCMsgUseItemResponse_EventNotActive",
	9:  "k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted",
	10: "k_EGCMsgUseItemResponse_MissingRequirement",
	11: "k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew",
	12: "k_EGCMsgUseItemResponse_EmoticonUnlock_Complete",
	13: "k_EGCMsgUseItemResponse_ItemUsed_Compendium",
}
var EGCMsgUseItemResponse_value = map[string]int32{
	"k_EGCMsgUseItemResponse_ItemUsed":                    0,
	"k_EGCMsgUseItemResponse_GiftNoOtherPlayers":          1,
	"k_EGCMsgUseItemResponse_ServerError":                 2,
	"k_EGCMsgUseItemResponse_MiniGameAlreadyStarted":      3,
	"k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted":       4,
	"k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted": 5,
	"k_EGCMsgUseItemResponse_NotInLowPriorityPool":        6,
	"k_EGCMsgUseItemResponse_NotHighEnoughLevel":          7,
	"k_EGCMsgUseItemResponse_EventNotActive":              8,
	"k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted": 9,
	"k_EGCMsgUseItemResponse_MissingRequirement":          10,
	"k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew":        11,
	"k_EGCMsgUseItemResponse_EmoticonUnlock_Complete":     12,
	"k_EGCMsgUseItemResponse_ItemUsed_Compendium":         13,
}

func (x EGCMsgUseItemResponse) Enum() *EGCMsgUseItemResponse {
	p := new(EGCMsgUseItemResponse)
	*p = x
	return p
}
func (x EGCMsgUseItemResponse) String() string {
	return proto.EnumName(EGCMsgUseItemResponse_name, int32(x))
}
func (x *EGCMsgUseItemResponse) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCMsgUseItemResponse_value, data, "EGCMsgUseItemResponse")
	if err != nil {
		return err
	}
	*x = EGCMsgUseItemResponse(value)
	return nil
}
func (EGCMsgUseItemResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{3} }

type CMsgGenericResult struct {
	Eresult          *uint32 `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	DebugMessage     *string `protobuf:"bytes,2,opt,name=debug_message" json:"debug_message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgGenericResult) Reset()                    { *m = CMsgGenericResult{} }
func (m *CMsgGenericResult) String() string            { return proto.CompactTextString(m) }
func (*CMsgGenericResult) ProtoMessage()               {}
func (*CMsgGenericResult) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

const Default_CMsgGenericResult_Eresult uint32 = 2

func (m *CMsgGenericResult) GetEresult() uint32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgGenericResult_Eresult
}

func (m *CMsgGenericResult) GetDebugMessage() string {
	if m != nil && m.DebugMessage != nil {
		return *m.DebugMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*CMsgGenericResult)(nil), "dota.CMsgGenericResult")
	proto.RegisterEnum("dota.EGCEconBaseMsg", EGCEconBaseMsg_name, EGCEconBaseMsg_value)
	proto.RegisterEnum("dota.EGCMsgResponse", EGCMsgResponse_name, EGCMsgResponse_value)
	proto.RegisterEnum("dota.EGCPartnerRequestResponse", EGCPartnerRequestResponse_name, EGCPartnerRequestResponse_value)
	proto.RegisterEnum("dota.EGCMsgUseItemResponse", EGCMsgUseItemResponse_name, EGCMsgUseItemResponse_value)
}

func init() { proto.RegisterFile("econ_shared_enums.proto", fileDescriptor26) }

var fileDescriptor26 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x4f, 0x13, 0x41,
	0x14, 0xc5, 0xd9, 0xd2, 0xf2, 0x67, 0x14, 0x33, 0x8e, 0xf2, 0x47, 0x30, 0xda, 0xa8, 0x11, 0x52,
	0x10, 0x14, 0x1e, 0x4c, 0x7c, 0x30, 0xa1, 0x65, 0x2d, 0x04, 0xfa, 0x27, 0x85, 0x3e, 0x6f, 0xc6,
	0xdd, 0xeb, 0x76, 0xd2, 0xdd, 0x7b, 0xd7, 0x99, 0xd9, 0x92, 0xbe, 0xf9, 0x1d, 0x7c, 0xf4, 0xcd,
	0xef, 0xe2, 0xf7, 0x32, 0x5b, 0xa0, 0xb1, 0x85, 0x16, 0xdf, 0x66, 0xe7, 0x77, 0xce, 0xcc, 0x39,
	0x77, 0x96, 0xad, 0x82, 0x4f, 0xe8, 0x99, 0x8e, 0xd4, 0x10, 0x78, 0x80, 0x69, 0x6c, 0x76, 0x13,
	0x4d, 0x96, 0x44, 0x3e, 0x20, 0x2b, 0x5f, 0x7d, 0x66, 0x8f, 0x2b, 0x35, 0x13, 0x56, 0x01, 0x41,
	0x2b, 0xbf, 0x05, 0x26, 0x8d, 0xac, 0x10, 0x6c, 0x1e, 0xf4, 0x60, 0xb9, 0xe6, 0x14, 0x9d, 0xad,
	0xa5, 0x4f, 0xce, 0xbe, 0x58, 0x66, 0x4b, 0x01, 0x7c, 0x4d, 0x43, 0x2f, 0x06, 0x63, 0x64, 0x08,
	0x6b, 0xb9, 0xa2, 0xb3, 0xb5, 0x58, 0xda, 0x61, 0x8f, 0xdc, 0x6a, 0xc5, 0xf5, 0x09, 0xcb, 0xd2,
	0x40, 0xcd, 0x84, 0x62, 0x9d, 0x2d, 0x77, 0x3d, 0x37, 0x3b, 0xb3, 0x32, 0x72, 0x2a, 0xff, 0xf9,
	0xb4, 0xf4, 0x2b, 0x37, 0x90, 0xd7, 0x4c, 0xd8, 0x02, 0x93, 0x10, 0x1a, 0x10, 0x2b, 0x4c, 0x74,
	0xbd, 0xd1, 0xbd, 0xc6, 0x29, 0x9f, 0x11, 0xeb, 0x6c, 0x65, 0x7c, 0xff, 0x08, 0x50, 0x41, 0xc0,
	0x1d, 0xf1, 0x92, 0x6d, 0x8c, 0xb3, 0x73, 0xd0, 0x3d, 0xd0, 0xae, 0xd6, 0xa4, 0x79, 0x4e, 0x6c,
	0xb0, 0xd5, 0x71, 0xc1, 0x85, 0x8a, 0x81, 0x52, 0xcb, 0x67, 0xef, 0x82, 0x27, 0xd8, 0x93, 0x91,
	0x0a, 0x78, 0xfe, 0x2e, 0x58, 0xa7, 0x9a, 0xb4, 0x7e, 0x87, 0x17, 0x44, 0x91, 0x3d, 0x1f, 0x87,
	0x6d, 0xec, 0x22, 0x5d, 0xe2, 0xd5, 0xc5, 0x73, 0x77, 0x25, 0xab, 0x93, 0x3d, 0xa3, 0x30, 0x84,
	0xa0, 0x81, 0x7c, 0xfe, 0xdf, 0x5a, 0x5f, 0xa4, 0x8a, 0x20, 0xb8, 0xa0, 0x8a, 0x06, 0x69, 0x81,
	0x2f, 0x94, 0x7e, 0x3b, 0xec, 0x99, 0x5b, 0xad, 0x34, 0xa5, 0xb6, 0x08, 0xba, 0x05, 0xdf, 0x53,
	0x30, 0x76, 0x38, 0xa8, 0x55, 0xf6, 0xa4, 0xeb, 0xb9, 0xa3, 0xb0, 0x71, 0x3a, 0x9c, 0xc6, 0x28,
	0x28, 0xcb, 0xe0, 0xd0, 0xf7, 0x29, 0x45, 0xcb, 0x73, 0xe2, 0x05, 0x5b, 0xbf, 0x25, 0xc8, 0x52,
	0x29, 0xec, 0x42, 0xc0, 0x67, 0xc5, 0x36, 0xdb, 0xbc, 0xc5, 0xdb, 0x68, 0xd2, 0x24, 0x21, 0x6d,
	0x21, 0xb8, 0x06, 0x17, 0xfd, 0x04, 0x78, 0xbe, 0xf4, 0xa7, 0xc0, 0x96, 0xaf, 0xf2, 0xb7, 0x0d,
	0x9c, 0x58, 0x88, 0x87, 0x01, 0xdf, 0xb0, 0xe2, 0x4d, 0xb5, 0x31, 0xe4, 0x65, 0x1f, 0x6d, 0x03,
	0x01, 0x9f, 0x11, 0xbb, 0xac, 0x34, 0x49, 0x55, 0x55, 0xdf, 0x6c, 0x9d, 0x1a, 0xb6, 0x03, 0xba,
	0x19, 0xc9, 0x3e, 0x68, 0xc3, 0x1d, 0xb1, 0xc9, 0x5e, 0x4f, 0xd2, 0x8f, 0xbe, 0xf9, 0x3e, 0xdb,
	0x9d, 0x24, 0xac, 0x29, 0x54, 0x55, 0x19, 0xc3, 0x61, 0xa4, 0x41, 0x06, 0xfd, 0x73, 0x2b, 0xb3,
	0x5e, 0x7c, 0x56, 0x7c, 0x60, 0xef, 0xee, 0x8b, 0x3c, 0x58, 0x98, 0xaa, 0x96, 0x98, 0x59, 0xf2,
	0xe2, 0x23, 0x3b, 0x98, 0x64, 0x39, 0xd2, 0x94, 0xb4, 0xa4, 0x85, 0x32, 0x61, 0x6a, 0xae, 0xef,
	0xba, 0x31, 0x16, 0xc4, 0x7b, 0xb6, 0x33, 0xc9, 0x58, 0x27, 0x7b, 0x82, 0x67, 0x74, 0xd9, 0xd4,
	0x8a, 0xb4, 0xb2, 0xfd, 0x26, 0x51, 0xc4, 0xe7, 0xa6, 0x8d, 0xaa, 0x4e, 0xf6, 0x58, 0x85, 0x1d,
	0x17, 0x29, 0x0d, 0x3b, 0x67, 0xd0, 0x83, 0x88, 0xcf, 0x8b, 0x12, 0x7b, 0x3b, 0x49, 0xef, 0xf6,
	0x00, 0xb3, 0x37, 0x3f, 0xf4, 0xad, 0xea, 0x01, 0x5f, 0x98, 0x56, 0x63, 0xd8, 0x7c, 0x60, 0x6a,
	0x92, 0x42, 0x3b, 0xec, 0xbf, 0x38, 0x2d, 0x54, 0x4d, 0x19, 0xa3, 0x30, 0xcc, 0xfe, 0x20, 0xa5,
	0x21, 0x06, 0xb4, 0x9c, 0x4d, 0xab, 0xed, 0xc6, 0x64, 0x95, 0x4f, 0xd8, 0xc6, 0x88, 0xfc, 0xae,
	0x57, 0xa7, 0x3a, 0x5c, 0xf2, 0x07, 0xe2, 0x80, 0xed, 0xfd, 0xa7, 0xa3, 0x42, 0x71, 0x12, 0x81,
	0x05, 0xfe, 0x50, 0xec, 0xb1, 0xed, 0x7b, 0xfb, 0x64, 0x72, 0xc0, 0x40, 0xa5, 0x31, 0x5f, 0x2a,
	0x17, 0x8e, 0x9d, 0x1f, 0xce, 0xcc, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0xef, 0x79, 0x01,
	0x1f, 0x05, 0x00, 0x00,
}
