// Code generated by protoc-gen-go.
// source: data_capture.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of Gender from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

type CaptureAward struct {
	ActivityType []ActivityType `protobuf:"varint,1,rep,packed,name=activity_type,json=activityType,enum=POGOProtos.Enums.ActivityType" json:"activity_type,omitempty"`
	Xp           []int32        `protobuf:"varint,2,rep,packed,name=xp" json:"xp,omitempty"`
	Candy        []int32        `protobuf:"varint,3,rep,packed,name=candy" json:"candy,omitempty"`
	Stardust     []int32        `protobuf:"varint,4,rep,packed,name=stardust" json:"stardust,omitempty"`
}

func (m *CaptureAward) Reset()                    { *m = CaptureAward{} }
func (m *CaptureAward) String() string            { return proto.CompactTextString(m) }
func (*CaptureAward) ProtoMessage()               {}
func (*CaptureAward) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type CaptureProbability struct {
	PokeballType           []ItemId  `protobuf:"varint,1,rep,packed,name=pokeball_type,json=pokeballType,enum=POGOProtos.Inventory.Item.ItemId" json:"pokeball_type,omitempty"`
	CaptureProbability     []float32 `protobuf:"fixed32,2,rep,packed,name=capture_probability,json=captureProbability" json:"capture_probability,omitempty"`
	ReticleDifficultyScale float64   `protobuf:"fixed64,12,opt,name=reticle_difficulty_scale,json=reticleDifficultyScale" json:"reticle_difficulty_scale,omitempty"`
}

func (m *CaptureProbability) Reset()                    { *m = CaptureProbability{} }
func (m *CaptureProbability) String() string            { return proto.CompactTextString(m) }
func (*CaptureProbability) ProtoMessage()               {}
func (*CaptureProbability) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func init() {
	proto.RegisterType((*CaptureAward)(nil), "POGOProtos.Data.Capture.CaptureAward")
	proto.RegisterType((*CaptureProbability)(nil), "POGOProtos.Data.Capture.CaptureProbability")
}

func init() { proto.RegisterFile("data_capture.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0xc4, 0x29, 0x45, 0x95, 0x49, 0x11, 0x32, 0x08, 0xac, 0x0e, 0x55, 0xe9, 0xd4, 0x29, 0x03,
	0x5d, 0x58, 0x5b, 0x8a, 0xaa, 0xb0, 0x34, 0x2a, 0x4c, 0x2c, 0x91, 0xe3, 0xb8, 0x92, 0x85, 0x9b,
	0x58, 0xb1, 0x53, 0xc8, 0x07, 0xf1, 0x3d, 0xfc, 0x12, 0x8e, 0x9d, 0x86, 0x48, 0x5d, 0xac, 0xe7,
	0xbb, 0x7b, 0xbe, 0x7b, 0xcf, 0x10, 0xa5, 0x44, 0x93, 0x98, 0x12, 0xa9, 0xcb, 0x82, 0x05, 0xb2,
	0xc8, 0x75, 0x8e, 0xee, 0xa3, 0xcd, 0x7a, 0x13, 0xd5, 0xa5, 0x0a, 0x56, 0x86, 0x0e, 0x9e, 0x1d,
	0x3d, 0xba, 0x64, 0x59, 0xb9, 0x57, 0x4e, 0x35, 0xba, 0xe5, 0xd9, 0x81, 0x65, 0x3a, 0x2f, 0xaa,
	0x98, 0x6b, 0xb6, 0x77, 0xe8, 0xf4, 0x07, 0x40, 0xbf, 0x91, 0x2f, 0xbe, 0x48, 0x91, 0xa2, 0x35,
	0x1c, 0x12, 0xaa, 0xf9, 0x81, 0xeb, 0x2a, 0xd6, 0x95, 0x64, 0x18, 0x4c, 0x7a, 0xb3, 0xab, 0xc7,
	0x71, 0xd0, 0x31, 0x79, 0xb1, 0xcf, 0x2e, 0x1a, 0xd9, 0xbb, 0x51, 0x2d, 0xbd, 0x6b, 0xb0, 0xf5,
	0x49, 0x07, 0x41, 0x08, 0x7a, 0xdf, 0x12, 0x7b, 0xa6, 0xbb, 0x6f, 0x59, 0x73, 0x43, 0x18, 0xf6,
	0x29, 0xc9, 0xd2, 0x0a, 0xf7, 0x5a, 0xd8, 0x01, 0x68, 0x0c, 0x07, 0x4a, 0x1b, 0xff, 0x52, 0x69,
	0x7c, 0xde, 0x92, 0x2d, 0x36, 0xfd, 0x05, 0x10, 0x35, 0x39, 0x4d, 0x88, 0x84, 0x24, 0x5c, 0x18,
	0x23, 0xf4, 0x0a, 0x87, 0x32, 0xff, 0x64, 0x09, 0x11, 0xa2, 0x9b, 0xf6, 0xa1, 0x9b, 0x36, 0x3c,
	0xce, 0x1d, 0x84, 0xf5, 0xdc, 0xf5, 0x11, 0xa6, 0x2e, 0xf0, 0xb1, 0xd7, 0x06, 0x9e, 0xc3, 0x9b,
	0x66, 0xaf, 0xb1, 0xfc, 0xb7, 0xb0, 0x13, 0x78, 0x56, 0x8e, 0xe8, 0x69, 0x80, 0x27, 0x88, 0x0b,
	0xa6, 0x39, 0x15, 0x2c, 0x4e, 0xf9, 0x6e, 0xc7, 0x69, 0x29, 0xcc, 0xe2, 0x14, 0x25, 0x82, 0x61,
	0x7f, 0x02, 0x66, 0x60, 0x7b, 0xd7, 0xf0, 0xab, 0x96, 0x7e, 0xab, 0xd9, 0xe5, 0xe0, 0xe3, 0xc2,
	0x7e, 0x81, 0x8a, 0xce, 0x22, 0x90, 0xb8, 0x7a, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xe7,
	0xaf, 0xf3, 0xe0, 0x01, 0x00, 0x00,
}
