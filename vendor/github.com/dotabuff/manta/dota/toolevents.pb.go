// Code generated by protoc-gen-go.
// source: toolevents.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ChangeMapToolEvent struct {
	Mapname          *string `protobuf:"bytes,1,opt,name=mapname" json:"mapname,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChangeMapToolEvent) Reset()                    { *m = ChangeMapToolEvent{} }
func (m *ChangeMapToolEvent) String() string            { return proto.CompactTextString(m) }
func (*ChangeMapToolEvent) ProtoMessage()               {}
func (*ChangeMapToolEvent) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{0} }

func (m *ChangeMapToolEvent) GetMapname() string {
	if m != nil && m.Mapname != nil {
		return *m.Mapname
	}
	return ""
}

type TraceRayServerToolEvent struct {
	Start            *CMsgVector `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End              *CMsgVector `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TraceRayServerToolEvent) Reset()                    { *m = TraceRayServerToolEvent{} }
func (m *TraceRayServerToolEvent) String() string            { return proto.CompactTextString(m) }
func (*TraceRayServerToolEvent) ProtoMessage()               {}
func (*TraceRayServerToolEvent) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{1} }

func (m *TraceRayServerToolEvent) GetStart() *CMsgVector {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TraceRayServerToolEvent) GetEnd() *CMsgVector {
	if m != nil {
		return m.End
	}
	return nil
}

type ToolTraceRayResult struct {
	Hit              *bool       `protobuf:"varint,1,opt,name=hit" json:"hit,omitempty"`
	Impact           *CMsgVector `protobuf:"bytes,2,opt,name=impact" json:"impact,omitempty"`
	Normal           *CMsgVector `protobuf:"bytes,3,opt,name=normal" json:"normal,omitempty"`
	Distance         *float32    `protobuf:"fixed32,4,opt,name=distance" json:"distance,omitempty"`
	Fraction         *float32    `protobuf:"fixed32,5,opt,name=fraction" json:"fraction,omitempty"`
	Ehandle          *int32      `protobuf:"varint,6,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ToolTraceRayResult) Reset()                    { *m = ToolTraceRayResult{} }
func (m *ToolTraceRayResult) String() string            { return proto.CompactTextString(m) }
func (*ToolTraceRayResult) ProtoMessage()               {}
func (*ToolTraceRayResult) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{2} }

func (m *ToolTraceRayResult) GetHit() bool {
	if m != nil && m.Hit != nil {
		return *m.Hit
	}
	return false
}

func (m *ToolTraceRayResult) GetImpact() *CMsgVector {
	if m != nil {
		return m.Impact
	}
	return nil
}

func (m *ToolTraceRayResult) GetNormal() *CMsgVector {
	if m != nil {
		return m.Normal
	}
	return nil
}

func (m *ToolTraceRayResult) GetDistance() float32 {
	if m != nil && m.Distance != nil {
		return *m.Distance
	}
	return 0
}

func (m *ToolTraceRayResult) GetFraction() float32 {
	if m != nil && m.Fraction != nil {
		return *m.Fraction
	}
	return 0
}

func (m *ToolTraceRayResult) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type SpawnEntityToolEvent struct {
	EntityKeyvalues  []byte `protobuf:"bytes,1,opt,name=entity_keyvalues" json:"entity_keyvalues,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SpawnEntityToolEvent) Reset()                    { *m = SpawnEntityToolEvent{} }
func (m *SpawnEntityToolEvent) String() string            { return proto.CompactTextString(m) }
func (*SpawnEntityToolEvent) ProtoMessage()               {}
func (*SpawnEntityToolEvent) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{3} }

func (m *SpawnEntityToolEvent) GetEntityKeyvalues() []byte {
	if m != nil {
		return m.EntityKeyvalues
	}
	return nil
}

func (m *SpawnEntityToolEvent) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type SpawnEntityToolEventResult struct {
	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SpawnEntityToolEventResult) Reset()                    { *m = SpawnEntityToolEventResult{} }
func (m *SpawnEntityToolEventResult) String() string            { return proto.CompactTextString(m) }
func (*SpawnEntityToolEventResult) ProtoMessage()               {}
func (*SpawnEntityToolEventResult) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{4} }

func (m *SpawnEntityToolEventResult) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type DestroyEntityToolEvent struct {
	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DestroyEntityToolEvent) Reset()                    { *m = DestroyEntityToolEvent{} }
func (m *DestroyEntityToolEvent) String() string            { return proto.CompactTextString(m) }
func (*DestroyEntityToolEvent) ProtoMessage()               {}
func (*DestroyEntityToolEvent) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{5} }

func (m *DestroyEntityToolEvent) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type DestroyAllEntitiesToolEvent struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DestroyAllEntitiesToolEvent) Reset()                    { *m = DestroyAllEntitiesToolEvent{} }
func (m *DestroyAllEntitiesToolEvent) String() string            { return proto.CompactTextString(m) }
func (*DestroyAllEntitiesToolEvent) ProtoMessage()               {}
func (*DestroyAllEntitiesToolEvent) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{6} }

type RestartMapToolEvent struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RestartMapToolEvent) Reset()                    { *m = RestartMapToolEvent{} }
func (m *RestartMapToolEvent) String() string            { return proto.CompactTextString(m) }
func (*RestartMapToolEvent) ProtoMessage()               {}
func (*RestartMapToolEvent) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{7} }

type ToolEvent_GetEntityInfo struct {
	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToolEvent_GetEntityInfo) Reset()                    { *m = ToolEvent_GetEntityInfo{} }
func (m *ToolEvent_GetEntityInfo) String() string            { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInfo) ProtoMessage()               {}
func (*ToolEvent_GetEntityInfo) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{8} }

func (m *ToolEvent_GetEntityInfo) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_GetEntityInfo) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type ToolEvent_GetEntityInfoResult struct {
	Cppclass         *string     `protobuf:"bytes,1,opt,name=cppclass,def=shithead" json:"cppclass,omitempty"`
	Classname        *string     `protobuf:"bytes,2,opt,name=classname" json:"classname,omitempty"`
	Name             *string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Origin           *CMsgVector `protobuf:"bytes,4,opt,name=origin" json:"origin,omitempty"`
	Mins             *CMsgVector `protobuf:"bytes,5,opt,name=mins" json:"mins,omitempty"`
	Maxs             *CMsgVector `protobuf:"bytes,6,opt,name=maxs" json:"maxs,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ToolEvent_GetEntityInfoResult) Reset()                    { *m = ToolEvent_GetEntityInfoResult{} }
func (m *ToolEvent_GetEntityInfoResult) String() string            { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInfoResult) ProtoMessage()               {}
func (*ToolEvent_GetEntityInfoResult) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{9} }

const Default_ToolEvent_GetEntityInfoResult_Cppclass string = "shithead"

func (m *ToolEvent_GetEntityInfoResult) GetCppclass() string {
	if m != nil && m.Cppclass != nil {
		return *m.Cppclass
	}
	return Default_ToolEvent_GetEntityInfoResult_Cppclass
}

func (m *ToolEvent_GetEntityInfoResult) GetClassname() string {
	if m != nil && m.Classname != nil {
		return *m.Classname
	}
	return ""
}

func (m *ToolEvent_GetEntityInfoResult) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ToolEvent_GetEntityInfoResult) GetOrigin() *CMsgVector {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *ToolEvent_GetEntityInfoResult) GetMins() *CMsgVector {
	if m != nil {
		return m.Mins
	}
	return nil
}

func (m *ToolEvent_GetEntityInfoResult) GetMaxs() *CMsgVector {
	if m != nil {
		return m.Maxs
	}
	return nil
}

type ToolEvent_GetEntityInputs struct {
	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToolEvent_GetEntityInputs) Reset()                    { *m = ToolEvent_GetEntityInputs{} }
func (m *ToolEvent_GetEntityInputs) String() string            { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInputs) ProtoMessage()               {}
func (*ToolEvent_GetEntityInputs) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{10} }

func (m *ToolEvent_GetEntityInputs) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_GetEntityInputs) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type ToolEvent_GetEntityInputsResult struct {
	InputList        []string `protobuf:"bytes,1,rep,name=input_list" json:"input_list,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ToolEvent_GetEntityInputsResult) Reset()         { *m = ToolEvent_GetEntityInputsResult{} }
func (m *ToolEvent_GetEntityInputsResult) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInputsResult) ProtoMessage()    {}
func (*ToolEvent_GetEntityInputsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor42, []int{11}
}

func (m *ToolEvent_GetEntityInputsResult) GetInputList() []string {
	if m != nil {
		return m.InputList
	}
	return nil
}

type ToolEvent_FireEntityInput struct {
	Ehandle          *int32  `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool   `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	InputName        *string `protobuf:"bytes,3,opt,name=input_name" json:"input_name,omitempty"`
	InputParam       *string `protobuf:"bytes,4,opt,name=input_param" json:"input_param,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ToolEvent_FireEntityInput) Reset()                    { *m = ToolEvent_FireEntityInput{} }
func (m *ToolEvent_FireEntityInput) String() string            { return proto.CompactTextString(m) }
func (*ToolEvent_FireEntityInput) ProtoMessage()               {}
func (*ToolEvent_FireEntityInput) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{12} }

func (m *ToolEvent_FireEntityInput) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_FireEntityInput) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

func (m *ToolEvent_FireEntityInput) GetInputName() string {
	if m != nil && m.InputName != nil {
		return *m.InputName
	}
	return ""
}

func (m *ToolEvent_FireEntityInput) GetInputParam() string {
	if m != nil && m.InputParam != nil {
		return *m.InputParam
	}
	return ""
}

type ToolEvent_SFMRecordingStateChanged struct {
	Isrecording      *bool  `protobuf:"varint,1,opt,name=isrecording" json:"isrecording,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToolEvent_SFMRecordingStateChanged) Reset()         { *m = ToolEvent_SFMRecordingStateChanged{} }
func (m *ToolEvent_SFMRecordingStateChanged) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_SFMRecordingStateChanged) ProtoMessage()    {}
func (*ToolEvent_SFMRecordingStateChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor42, []int{13}
}

func (m *ToolEvent_SFMRecordingStateChanged) GetIsrecording() bool {
	if m != nil && m.Isrecording != nil {
		return *m.Isrecording
	}
	return false
}

type ToolEvent_SFMToolActiveStateChanged struct {
	Isactive         *bool  `protobuf:"varint,1,opt,name=isactive" json:"isactive,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToolEvent_SFMToolActiveStateChanged) Reset()         { *m = ToolEvent_SFMToolActiveStateChanged{} }
func (m *ToolEvent_SFMToolActiveStateChanged) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_SFMToolActiveStateChanged) ProtoMessage()    {}
func (*ToolEvent_SFMToolActiveStateChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor42, []int{14}
}

func (m *ToolEvent_SFMToolActiveStateChanged) GetIsactive() bool {
	if m != nil && m.Isactive != nil {
		return *m.Isactive
	}
	return false
}

func init() {
	proto.RegisterType((*ChangeMapToolEvent)(nil), "dota.ChangeMapToolEvent")
	proto.RegisterType((*TraceRayServerToolEvent)(nil), "dota.TraceRayServerToolEvent")
	proto.RegisterType((*ToolTraceRayResult)(nil), "dota.ToolTraceRayResult")
	proto.RegisterType((*SpawnEntityToolEvent)(nil), "dota.SpawnEntityToolEvent")
	proto.RegisterType((*SpawnEntityToolEventResult)(nil), "dota.SpawnEntityToolEventResult")
	proto.RegisterType((*DestroyEntityToolEvent)(nil), "dota.DestroyEntityToolEvent")
	proto.RegisterType((*DestroyAllEntitiesToolEvent)(nil), "dota.DestroyAllEntitiesToolEvent")
	proto.RegisterType((*RestartMapToolEvent)(nil), "dota.RestartMapToolEvent")
	proto.RegisterType((*ToolEvent_GetEntityInfo)(nil), "dota.ToolEvent_GetEntityInfo")
	proto.RegisterType((*ToolEvent_GetEntityInfoResult)(nil), "dota.ToolEvent_GetEntityInfoResult")
	proto.RegisterType((*ToolEvent_GetEntityInputs)(nil), "dota.ToolEvent_GetEntityInputs")
	proto.RegisterType((*ToolEvent_GetEntityInputsResult)(nil), "dota.ToolEvent_GetEntityInputsResult")
	proto.RegisterType((*ToolEvent_FireEntityInput)(nil), "dota.ToolEvent_FireEntityInput")
	proto.RegisterType((*ToolEvent_SFMRecordingStateChanged)(nil), "dota.ToolEvent_SFMRecordingStateChanged")
	proto.RegisterType((*ToolEvent_SFMToolActiveStateChanged)(nil), "dota.ToolEvent_SFMToolActiveStateChanged")
}

func init() { proto.RegisterFile("toolevents.proto", fileDescriptor42) }

var fileDescriptor42 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x75, 0xda, 0xda, 0x93, 0x4a, 0x0d, 0x2e, 0xb4, 0x26, 0x28, 0x34, 0x32, 0x42, 0x0a,
	0x07, 0x72, 0x40, 0x42, 0x08, 0x6e, 0xa5, 0x6d, 0x10, 0x48, 0xb9, 0x24, 0x15, 0x12, 0xa7, 0x68,
	0xb0, 0xa7, 0xc9, 0xaa, 0xf6, 0xae, 0xb5, 0x3b, 0x49, 0xc9, 0x8d, 0xef, 0xe0, 0x63, 0xf8, 0x36,
	0xb4, 0xeb, 0xb4, 0x49, 0x69, 0x82, 0xd4, 0xe3, 0xcc, 0xbc, 0x79, 0x6f, 0x67, 0xde, 0xd8, 0xd0,
	0x60, 0xa5, 0x72, 0x9a, 0x91, 0x64, 0xd3, 0x2d, 0xb5, 0x62, 0x15, 0xd5, 0x32, 0xc5, 0xd8, 0x3c,
	0x94, 0xc4, 0xd7, 0x4a, 0x5f, 0xfd, 0x40, 0x43, 0x3c, 0x2f, 0x69, 0x51, 0x4d, 0x5e, 0x41, 0x74,
	0x3a, 0x41, 0x39, 0xa6, 0x3e, 0x96, 0x17, 0x4a, 0xe5, 0xe7, 0xb6, 0x35, 0xda, 0x87, 0xdd, 0x02,
	0x4b, 0x89, 0x05, 0xc5, 0x5e, 0xdb, 0xeb, 0x84, 0xc9, 0x77, 0x38, 0xba, 0xd0, 0x98, 0xd2, 0x00,
	0xe7, 0x43, 0xd2, 0x33, 0xd2, 0x4b, 0xec, 0x31, 0x6c, 0x1b, 0x46, 0xcd, 0x0e, 0x59, 0x7f, 0xdb,
	0xe8, 0x5a, 0xbd, 0xee, 0x69, 0xdf, 0x8c, 0xbf, 0x51, 0xca, 0x4a, 0x47, 0x2d, 0xf0, 0x49, 0x66,
	0xf1, 0xd6, 0xfa, 0x72, 0xf2, 0xdb, 0x83, 0xc8, 0xb2, 0xdd, 0xf0, 0x0f, 0xc8, 0x4c, 0x73, 0x8e,
	0xea, 0xe0, 0x4f, 0x44, 0x45, 0x1a, 0x44, 0x6d, 0xd8, 0x11, 0x45, 0x89, 0x29, 0x6f, 0x62, 0xb1,
	0x08, 0xa9, 0x74, 0x81, 0x79, 0xec, 0x6f, 0x40, 0x34, 0x20, 0xc8, 0x84, 0x61, 0x94, 0x29, 0xc5,
	0xb5, 0xb6, 0xd7, 0xd9, 0xb2, 0x99, 0x4b, 0x8d, 0x29, 0x0b, 0x25, 0xe3, 0x6d, 0x97, 0xd9, 0x87,
	0x5d, 0x9a, 0xa0, 0xcc, 0x72, 0x8a, 0x77, 0xda, 0x5e, 0x67, 0x3b, 0xf9, 0x0a, 0x4f, 0x86, 0x25,
	0x5e, 0xcb, 0x73, 0xc9, 0x82, 0xe7, 0xcb, 0xa1, 0x63, 0x68, 0x90, 0x4b, 0x8d, 0xae, 0x68, 0x3e,
	0xc3, 0x7c, 0x4a, 0xc6, 0x3d, 0x75, 0xcf, 0x56, 0xd2, 0x5c, 0xd8, 0xfd, 0x8b, 0x8c, 0x2a, 0x8c,
	0x7b, 0x74, 0x90, 0xbc, 0x81, 0xe6, 0x3a, 0xae, 0xc5, 0xbc, 0x2b, 0xd2, 0x9e, 0x93, 0x7e, 0x0d,
	0x87, 0x67, 0x64, 0x58, 0xab, 0xf9, 0xbf, 0xe2, 0xf7, 0xa0, 0x2d, 0x78, 0xbe, 0x80, 0x9e, 0xe4,
	0xb9, 0x43, 0x0b, 0x32, 0xb7, 0xf8, 0xe4, 0x29, 0x1c, 0x0c, 0xc8, 0x79, 0xb4, 0x6a, 0x72, 0x72,
	0x06, 0x47, 0xb7, 0xc1, 0xe8, 0x33, 0x71, 0x25, 0xf3, 0x45, 0x5e, 0xaa, 0x7b, 0x0a, 0xff, 0x99,
	0xea, 0x8f, 0x07, 0xad, 0x0d, 0x34, 0x8b, 0xc9, 0x9a, 0x10, 0xa4, 0x65, 0x99, 0xe6, 0x68, 0xaa,
	0x1d, 0x85, 0x1f, 0x03, 0x33, 0x11, 0x3c, 0x21, 0xcc, 0xa2, 0xc7, 0x10, 0xba, 0x82, 0x3b, 0x35,
	0x4b, 0x18, 0x46, 0x7b, 0x50, 0x73, 0x91, 0xef, 0xa2, 0x36, 0xec, 0x28, 0x2d, 0xc6, 0x42, 0x3a,
	0xcf, 0xd6, 0xf9, 0xfa, 0x02, 0x6a, 0x85, 0x90, 0xc6, 0x39, 0xb8, 0xa9, 0x8e, 0x3f, 0x8d, 0x33,
	0x74, 0xdd, 0xfd, 0xf5, 0xe0, 0xd9, 0xda, 0xf7, 0x97, 0x53, 0x36, 0x0f, 0x59, 0xc4, 0x3b, 0x38,
	0xde, 0xc8, 0xb3, 0xd8, 0x44, 0x04, 0x20, 0x6c, 0x3c, 0xca, 0x85, 0xb1, 0xa7, 0xed, 0x77, 0xc2,
	0x44, 0xad, 0xca, 0xf7, 0x84, 0xa6, 0x95, 0xbe, 0x07, 0xc8, 0x2f, 0xb9, 0x57, 0x96, 0x77, 0x00,
	0xf5, 0x2a, 0x57, 0xa2, 0xc6, 0xc2, 0x6d, 0x30, 0x4c, 0x3e, 0x40, 0xb2, 0x14, 0x1c, 0xf6, 0xfa,
	0x03, 0x4a, 0x95, 0xce, 0x84, 0x1c, 0x0f, 0x19, 0x99, 0xaa, 0x9f, 0x41, 0xe6, 0x5a, 0x8d, 0xbe,
	0x29, 0x55, 0x9f, 0x61, 0xf2, 0x1e, 0x5e, 0xde, 0x69, 0xb5, 0xc1, 0x49, 0xca, 0x62, 0x46, 0x77,
	0x7a, 0x1b, 0x10, 0x08, 0x83, 0x2e, 0x5f, 0x35, 0x7e, 0xf2, 0x7f, 0x79, 0x8f, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x07, 0x35, 0x3d, 0x85, 0x9b, 0x04, 0x00, 0x00,
}