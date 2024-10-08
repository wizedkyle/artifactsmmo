package models

import "time"

const (
	AshTree                  string = "ash_tree"
	AshTeeX                  int    = 6
	AshTreeY                 int    = 1
	Bank                     string = "bank"
	BankX                    int    = 4
	BankY                    int    = 1
	SouthBank                string = "south_bank"
	SouthBankX               int    = 7
	SouthBankY               int    = 13
	Bass                     string = "bass"
	BassLevel                int    = 30
	BassX                    int    = 6
	BassY                    int    = 12
	BirchTree                string = "birch_tree"
	BirchTreeLevel           int    = 20
	BirchTreeX               int    = 3
	BirchTreeY               int    = 5
	Crafting                 string = "crafting"
	CoalLevel                int    = 20
	Coal                     string = "coal_rocks"
	CoalX                    int    = 1
	CoalY                    int    = 6
	CookingWorkshop          string = "cooking_workshop"
	CookingWorkshopX         int    = 1
	CookingWorkshopY         int    = 1
	Copper                   string = "copper_ore"
	CopperX                  int    = 2
	CopperY                  int    = 0
	Combat                   string = "combat"
	DeadTree                 string = "dead_tree"
	DeadTreeLevel            int    = 30
	DeadTreeX                int    = 9
	DeadTreeY                int    = 8
	ForgeWorkshop            string = "forge_workshop"
	GearcraftingWorkshop     string = "gearcrafting_workshop"
	GearcraftingWorkshopX    int    = 3
	GearcraftingWorkshopY    int    = 1
	Gudgeon                  string = "gudgeon"
	GudgeonX                 int    = 4
	GudgeonY                 int    = 2
	GoldLevel                int    = 30
	Gold                     string = "gold_rocks"
	GoldX                    int    = 10
	GoldY                    int    = -4
	InventorySize            int    = 100
	ItemTypeConsumable       string = ""
	ItemTypeBodyArmor        string = ""
	IronLevel                int    = 10
	Iron                     string = "iron_rocks"
	IronX                    int    = 1
	IronY                    int    = 7
	JewelrycraftingWorkshop  string = "jewelrycrafting_workshop"
	JewelrycraftingWorkshopX int    = 1
	JewelrycraftingWorkshopY int    = 3
	MiningWorkshop           string = "mining_workshop"
	MiningWorkshopX          int    = 1
	MiningWorkshopY          int    = 5
	MagicTree                string = "magic_tree"
	Shrimp                   string = "shrimp"
	ShrimpLevel              int    = 10
	ShrimpX                  int    = 5
	ShrimpY                  int    = 2
	SpruceTree               string = "spruce_tree"
	SpruceTreeLevel          int    = 10
	SpruceTreeX              int    = 2
	SpruceTreeY              int    = 6
	StrangeRocks             string = "strange_rocks"
	Trout                    string = "trout"
	TroutLevel               int    = 20
	TroutX                   int    = 7
	TroutY                   int    = 12
	ItemRetrieved            string = "item_retrieved"
	TaskCancelled            string = "task_cancelled"
	TaskCreated              string = "task_created"
	TaskRetrieved            string = "task_retrieved"
	WeaponcraftingWorkshop   string = "weaponcrafting_workshop"
	WeaponcraftingWorkshopX  int    = 2
	WeaponcraftingWorkshopY  int    = 1
	WoodcuttingWorkshop      string = "woodcutting_workshop"
	WoodcuttingWorkshopX     int    = -2
	WoodcuttingWorkshopY     int    = -3
)

type Cooldown struct {
	TotalSeconds     int       `json:"total_seconds"`
	RemainingSeconds int       `json:"remaining_seconds"`
	StartedAt        time.Time `json:"started_at"`
	Expiration       time.Time `json:"expiration"`
	Reason           string    `json:"reason"`
}

type Credentials struct {
	CharacterName string `json:"characterName"`
	Token         string `json:"token"`
}
