package models

import "time"

type ActionMove struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type CharacterGatheringResponse struct {
	Data CharacterGathering `json:"data"`
}

type CharacterGathering struct {
	Cooldown  Cooldown        `json:"cooldown"`
	Details   CraftingDetails `json:"details"`
	Character Character       `json:"character"`
}

type CharacterMovementResponse struct {
	Data CharacterMovement `json:"data"`
}

type CharacterMovement struct {
	Cooldown    Cooldown    `json:"cooldown"`
	Destination Destination `json:"destination"`
	Character   Character   `json:"character"`
}

type Character struct {
	Name                    string          `json:"name"`
	Skin                    string          `json:"skin"`
	Level                   int             `json:"level"`
	XP                      int             `json:"xp"`
	MaxXp                   int             `json:"max_xp"`
	TotalXp                 int             `json:"total_xp"`
	Gold                    int             `json:"gold"`
	Speed                   int             `json:"speed"`
	MiningLevel             int             `json:"mining_level"`
	MiningXp                int             `json:"mining_xp"`
	MiningMaxXp             int             `json:"mining_max_xp"`
	WoodcuttingLevel        int             `json:"woodcutting_level"`
	WoodcuttingMaxXp        int             `json:"woodcutting_max_xp"`
	FishingLevel            int             `json:"fishing_level"`
	FishingXp               int             `json:"fishing_xp"`
	FishingMaxXp            int             `json:"fishing_max_xp"`
	WeaponCraftingLevel     int             `json:"weaponcrafting_level"`
	WeaponCraftingXp        int             `json:"weaponcrafting_xp"`
	WeaponCraftingMaxXp     int             `json:"weaponcrafting_max_xp"`
	GearCraftingLevel       int             `json:"gearcrafting_level"`
	GearCraftingXp          int             `json:"gearcrafting_xp"`
	GearCraftingMaxXp       int             `json:"gearcrafting_max_xp"`
	JewelryCraftingLevel    int             `json:"jewelrycrafting_level"`
	JewelryCraftingXp       int             `json:"jewelrycrafting_xp"`
	JewelryCraftingMaxXp    int             `json:"jewelrycrafting_max_xp"`
	CookingLevel            int             `json:"cooking_level"`
	CookingXp               int             `json:"cooking_xp"`
	CookingMaxXp            int             `json:"cooking_max_xp"`
	Hp                      int             `json:"hp"`
	Haste                   int             `json:"haste"`
	CriticalStrike          int             `json:"critical_strike"`
	Stamina                 int             `json:"stamina"`
	AttackFire              int             `json:"attack_fire"`
	AttackEarth             int             `json:"attack_earth"`
	AttackWater             int             `json:"attack_water"`
	AttackAir               int             `json:"attack_air"`
	DmgFire                 int             `json:"dmg_fire"`
	DmgEarth                int             `json:"dmg_earth"`
	DmgWater                int             `json:"dmg_water"`
	DmgAir                  int             `json:"dmg_air"`
	ResFire                 int             `json:"res_fire"`
	ResEarth                int             `json:"res_earth"`
	ResWater                int             `json:"res_water"`
	ResAir                  int             `json:"res_air"`
	X                       int             `json:"x"`
	Y                       int             `json:"y"`
	Cooldown                int             `json:"cooldown"`
	CooldownExpiration      time.Time       `json:"cooldown_expiration"`
	WeaponSlot              string          `json:"weapon_slot"`
	ShieldSlot              string          `json:"shield_slot"`
	HelmetSlot              string          `json:"helmet_slot"`
	BodyArmorSlot           string          `json:"body_armor_slot"`
	LegArmorSlot            string          `json:"leg_armor_slot"`
	BootsSlot               string          `json:"boots_slot"`
	Ring1Slot               string          `json:"ring1_slot"`
	Ring2Slot               string          `json:"ring2_slot"`
	AmuletSlot              string          `json:"amulet_slot"`
	Artifact1Slot           string          `json:"artifact1_slot"`
	Artifact2Slot           string          `json:"artifact2_slot"`
	Artifact3Slot           string          `json:"artifact3_slot"`
	Consumable1Slot         string          `json:"consumable1_slot"`
	Consumable1SlotQuantity int             `json:"consumable1_slot_quantity"`
	Consumable2Slot         string          `json:"consumable2_slot"`
	Consumable2SlotQuantity int             `json:"consumable2_slot_quantity"`
	Task                    string          `json:"task"`
	TaskType                string          `json:"task_type"`
	TaskProgress            int             `json:"task_progress"`
	TaskTotal               int             `json:"task_total"`
	InventoryMaxItems       int             `json:"inventory_max_items"`
	Inventory               []InventorySlot `json:"inventory"`
}

type CraftingDetails struct {
	Xp    int    `json:"xp"`
	Items []Item `json:"items"`
}

type Item struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Content struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type Destination struct {
	Name    string  `json:"name"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	Content Content `json:"content,omitempty"`
}

type InventorySlot struct {
	Slot     int    `json:"slot"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}
