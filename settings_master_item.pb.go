// Code generated by protoc-gen-go.
// source: settings_master_item.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of EggIncubators from inventory.proto

// Ignoring public import of InventoryUpgrade from inventory.proto

// Ignoring public import of InventoryItem from inventory.proto

// Ignoring public import of DeletedItem from inventory.proto

// Ignoring public import of AppliedItem from inventory.proto

// Ignoring public import of EggIncubator from inventory.proto

// Ignoring public import of InventoryDelta from inventory.proto

// Ignoring public import of InventoryUpgrades from inventory.proto

// Ignoring public import of AppliedItems from inventory.proto

// Ignoring public import of InventoryItemData from inventory.proto

// Ignoring public import of Candy from inventory.proto

// Ignoring public import of EggIncubatorType from inventory.proto

// Ignoring public import of InventoryUpgradeType from inventory.proto

type FortModifierAttributes struct {
	ModifierLifetimeSeconds   int32 `protobuf:"varint,1,opt,name=modifier_lifetime_seconds" json:"modifier_lifetime_seconds,omitempty"`
	TroyDiskNumPokemonSpawned int32 `protobuf:"varint,2,opt,name=troy_disk_num_pokemon_spawned" json:"troy_disk_num_pokemon_spawned,omitempty"`
}

func (m *FortModifierAttributes) Reset()         { *m = FortModifierAttributes{} }
func (m *FortModifierAttributes) String() string { return proto.CompactTextString(m) }
func (*FortModifierAttributes) ProtoMessage()    {}

type ReviveAttributes struct {
	StaPercent float32 `protobuf:"fixed32,1,opt,name=sta_percent" json:"sta_percent,omitempty"`
}

func (m *ReviveAttributes) Reset()         { *m = ReviveAttributes{} }
func (m *ReviveAttributes) String() string { return proto.CompactTextString(m) }
func (*ReviveAttributes) ProtoMessage()    {}

type PokeballAttributes struct {
	ItemEffect         ItemEffect `protobuf:"varint,1,opt,name=item_effect,enum=POGOProtos.Enums.ItemEffect" json:"item_effect,omitempty"`
	CaptureMulti       float32    `protobuf:"fixed32,2,opt,name=capture_multi" json:"capture_multi,omitempty"`
	CaptureMultiEffect float32    `protobuf:"fixed32,3,opt,name=capture_multi_effect" json:"capture_multi_effect,omitempty"`
	ItemEffectMod      float32    `protobuf:"fixed32,4,opt,name=item_effect_mod" json:"item_effect_mod,omitempty"`
}

func (m *PokeballAttributes) Reset()         { *m = PokeballAttributes{} }
func (m *PokeballAttributes) String() string { return proto.CompactTextString(m) }
func (*PokeballAttributes) ProtoMessage()    {}

type PotionAttributes struct {
	StaPercent float32 `protobuf:"fixed32,1,opt,name=sta_percent" json:"sta_percent,omitempty"`
	StaAmount  int32   `protobuf:"varint,2,opt,name=sta_amount" json:"sta_amount,omitempty"`
}

func (m *PotionAttributes) Reset()         { *m = PotionAttributes{} }
func (m *PotionAttributes) String() string { return proto.CompactTextString(m) }
func (*PotionAttributes) ProtoMessage()    {}

type IncenseAttributes struct {
	IncenseLifetimeSeconds                   int32         `protobuf:"varint,1,opt,name=incense_lifetime_seconds" json:"incense_lifetime_seconds,omitempty"`
	PokemonType                              []PokemonType `protobuf:"varint,2,rep,name=pokemon_type,enum=POGOProtos.Enums.PokemonType" json:"pokemon_type,omitempty"`
	PokemonIncenseTypeProbability            float32       `protobuf:"fixed32,3,opt,name=pokemon_incense_type_probability" json:"pokemon_incense_type_probability,omitempty"`
	StandingTimeBetweenEncountersSeconds     int32         `protobuf:"varint,4,opt,name=standing_time_between_encounters_seconds" json:"standing_time_between_encounters_seconds,omitempty"`
	MovingTimeBetweenEncounterSeconds        int32         `protobuf:"varint,5,opt,name=moving_time_between_encounter_seconds" json:"moving_time_between_encounter_seconds,omitempty"`
	DistanceRequiredForShorterIntervalMeters int32         `protobuf:"varint,6,opt,name=distance_required_for_shorter_interval_meters" json:"distance_required_for_shorter_interval_meters,omitempty"`
	PokemonAttractedLengthSec                int32         `protobuf:"varint,7,opt,name=pokemon_attracted_length_sec" json:"pokemon_attracted_length_sec,omitempty"`
}

func (m *IncenseAttributes) Reset()         { *m = IncenseAttributes{} }
func (m *IncenseAttributes) String() string { return proto.CompactTextString(m) }
func (*IncenseAttributes) ProtoMessage()    {}

type InventoryUpgradeAttributes struct {
	AdditionalStorage int32                `protobuf:"varint,1,opt,name=additional_storage" json:"additional_storage,omitempty"`
	UpgradeType       InventoryUpgradeType `protobuf:"varint,2,opt,name=upgrade_type,enum=POGOProtos.Inventory.InventoryUpgradeType" json:"upgrade_type,omitempty"`
}

func (m *InventoryUpgradeAttributes) Reset()         { *m = InventoryUpgradeAttributes{} }
func (m *InventoryUpgradeAttributes) String() string { return proto.CompactTextString(m) }
func (*InventoryUpgradeAttributes) ProtoMessage()    {}

type ExperienceBoostAttributes struct {
	XpMultiplier    float32 `protobuf:"fixed32,1,opt,name=xp_multiplier" json:"xp_multiplier,omitempty"`
	BoostDurationMs int32   `protobuf:"varint,2,opt,name=boost_duration_ms" json:"boost_duration_ms,omitempty"`
}

func (m *ExperienceBoostAttributes) Reset()         { *m = ExperienceBoostAttributes{} }
func (m *ExperienceBoostAttributes) String() string { return proto.CompactTextString(m) }
func (*ExperienceBoostAttributes) ProtoMessage()    {}

type FoodAttributes struct {
	ItemEffect        []ItemEffect `protobuf:"varint,1,rep,name=item_effect,enum=POGOProtos.Enums.ItemEffect" json:"item_effect,omitempty"`
	ItemEffectPercent []float32    `protobuf:"fixed32,2,rep,name=item_effect_percent" json:"item_effect_percent,omitempty"`
	GrowthPercent     float32      `protobuf:"fixed32,3,opt,name=growth_percent" json:"growth_percent,omitempty"`
}

func (m *FoodAttributes) Reset()         { *m = FoodAttributes{} }
func (m *FoodAttributes) String() string { return proto.CompactTextString(m) }
func (*FoodAttributes) ProtoMessage()    {}

type BattleAttributes struct {
	StaPercent float32 `protobuf:"fixed32,1,opt,name=sta_percent" json:"sta_percent,omitempty"`
}

func (m *BattleAttributes) Reset()         { *m = BattleAttributes{} }
func (m *BattleAttributes) String() string { return proto.CompactTextString(m) }
func (*BattleAttributes) ProtoMessage()    {}

type EggIncubatorAttributes struct {
	IncubatorType      EggIncubatorType `protobuf:"varint,1,opt,name=incubator_type,enum=POGOProtos.Inventory.EggIncubatorType" json:"incubator_type,omitempty"`
	Uses               int32            `protobuf:"varint,2,opt,name=uses" json:"uses,omitempty"`
	DistanceMultiplier float32          `protobuf:"fixed32,3,opt,name=distance_multiplier" json:"distance_multiplier,omitempty"`
}

func (m *EggIncubatorAttributes) Reset()         { *m = EggIncubatorAttributes{} }
func (m *EggIncubatorAttributes) String() string { return proto.CompactTextString(m) }
func (*EggIncubatorAttributes) ProtoMessage()    {}

func init() {
}
