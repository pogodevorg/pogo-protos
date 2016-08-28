// Code generated by protoc-gen-go.
// source: data_gym.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of BuddyPokemon from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of FortSummary from map_fort.proto

// Ignoring public import of FortModifier from map_fort.proto

// Ignoring public import of FortData from map_fort.proto

// Ignoring public import of FortLureInfo from map_fort.proto

// Ignoring public import of FortSponsor from map_fort.proto

// Ignoring public import of FortRenderingType from map_fort.proto

// Ignoring public import of FortType from map_fort.proto

type GymMembership struct {
	PokemonData          *PokemonData         `protobuf:"bytes,1,opt,name=pokemon_data,json=pokemonData" json:"pokemon_data,omitempty"`
	TrainerPublicProfile *PlayerPublicProfile `protobuf:"bytes,2,opt,name=trainer_public_profile,json=trainerPublicProfile" json:"trainer_public_profile,omitempty"`
}

func (m *GymMembership) Reset()                    { *m = GymMembership{} }
func (m *GymMembership) String() string            { return proto.CompactTextString(m) }
func (*GymMembership) ProtoMessage()               {}
func (*GymMembership) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *GymMembership) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func (m *GymMembership) GetTrainerPublicProfile() *PlayerPublicProfile {
	if m != nil {
		return m.TrainerPublicProfile
	}
	return nil
}

type GymState struct {
	FortData    *FortData        `protobuf:"bytes,1,opt,name=fort_data,json=fortData" json:"fort_data,omitempty"`
	Memberships []*GymMembership `protobuf:"bytes,2,rep,name=memberships" json:"memberships,omitempty"`
}

func (m *GymState) Reset()                    { *m = GymState{} }
func (m *GymState) String() string            { return proto.CompactTextString(m) }
func (*GymState) ProtoMessage()               {}
func (*GymState) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *GymState) GetFortData() *FortData {
	if m != nil {
		return m.FortData
	}
	return nil
}

func (m *GymState) GetMemberships() []*GymMembership {
	if m != nil {
		return m.Memberships
	}
	return nil
}

func init() {
	proto.RegisterType((*GymMembership)(nil), "POGOProtos.Data.Gym.GymMembership")
	proto.RegisterType((*GymState)(nil), "POGOProtos.Data.Gym.GymState")
}

func init() { proto.RegisterFile("data_gym.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x49, 0x2c, 0x49,
	0x8c, 0x4f, 0xaf, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x0e, 0xf0, 0x77, 0xf7,
	0x0f, 0x00, 0x31, 0x8b, 0xf5, 0x5c, 0x80, 0x52, 0x7a, 0xee, 0x95, 0xb9, 0x52, 0x5c, 0x20, 0x45,
	0x10, 0x05, 0x52, 0x82, 0x60, 0x0d, 0x05, 0x39, 0x89, 0x95, 0xa9, 0x45, 0x50, 0x21, 0xbe, 0xdc,
	0xc4, 0x82, 0xf8, 0xb4, 0xfc, 0xa2, 0x12, 0x08, 0x5f, 0x69, 0x33, 0x23, 0x17, 0x2f, 0x50, 0x9b,
	0x6f, 0x6a, 0x6e, 0x52, 0x6a, 0x51, 0x71, 0x46, 0x66, 0x81, 0x90, 0x3d, 0x17, 0x4f, 0x41, 0x7e,
	0x76, 0x6a, 0x6e, 0x7e, 0x5e, 0x3c, 0x48, 0xbb, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x8c,
	0x1e, 0xba, 0x65, 0x01, 0x10, 0x45, 0x20, 0x76, 0x10, 0x77, 0x01, 0x82, 0x23, 0x94, 0xc8, 0x25,
	0x56, 0x52, 0x94, 0x98, 0x99, 0x97, 0x5a, 0x14, 0x5f, 0x50, 0x9a, 0x94, 0x93, 0x99, 0x1c, 0x0f,
	0xb4, 0x2a, 0x2d, 0x33, 0x27, 0x55, 0x82, 0x09, 0x6c, 0x94, 0x36, 0xa6, 0x51, 0x10, 0x17, 0x42,
	0xa8, 0x00, 0xb0, 0x9e, 0x00, 0x88, 0x96, 0x20, 0x11, 0xa8, 0x51, 0x28, 0xa2, 0x4a, 0x3d, 0x8c,
	0x5c, 0x1c, 0x40, 0x57, 0x07, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0x59, 0x71, 0x71, 0x82, 0x3c, 0x84,
	0xec, 0x5a, 0x59, 0x64, 0x2b, 0x7c, 0x13, 0x0b, 0xf4, 0xdc, 0x40, 0x3e, 0x06, 0x11, 0x60, 0xe7,
	0x72, 0xa4, 0x41, 0x59, 0x42, 0x2e, 0x5c, 0xdc, 0xb9, 0x70, 0xaf, 0x17, 0x03, 0x1d, 0xc8, 0x0c,
	0xd4, 0xad, 0xa4, 0x87, 0x25, 0x60, 0xf5, 0x50, 0x42, 0x29, 0x08, 0x59, 0x9b, 0x13, 0x47, 0x14,
	0x1b, 0x38, 0x34, 0x8b, 0x03, 0x18, 0x02, 0x18, 0x03, 0x98, 0x92, 0x20, 0x3c, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x90, 0x16, 0x20, 0xe9, 0xb1, 0x01, 0x00, 0x00,
}
