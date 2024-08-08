package models

import "time"

const (
	AshTree            string = "ash_tree"
	AshTeeX            int    = 6
	AshTreeY           int    = 1
	Bank               string = "bank"
	BankX              int    = 4
	BankY              int    = 1
	Bass               string = "bass"
	BassLevel          int    = 30
	BassX              int    = -3
	BassY              int    = 6
	BirchTree          string = "birch_tree"
	BirchTreeLevel     int    = 20
	BirchTreeX         int    = 3
	BirchTreeY         int    = 5
	CoalLevel          int    = 20
	Coal               string = "coal_rocks"
	CoalX              int    = 1
	CoalY              int    = 6
	Copper             string = "copper_ore"
	CopperX            int    = 2
	CopperY            int    = 0
	Gudgeon            string = "gudgeon"
	GudgeonX           int    = 4
	GudgeonY           int    = 2
	GoldLevel          int    = 30
	Gold               string = "gold_rocks"
	GoldX              int    = 10
	GoldY              int    = -4
	ItemTypeConsumable string = ""
	ItemTypeBodyArmor  string = ""
	IronLevel          int    = 10
	Iron               string = "iron_rocks"
	IronX              int    = 1
	IronY              int    = 7
	Shrimp             string = "shrimp"
	ShrimpLevel        int    = 10
	ShrimpX            int    = 5
	ShrimpY            int    = 2
	SpruceTree         string = "spruce_tree"
	SpruceTreeLevel    int    = 10
	SpruceTreeX        int    = 2
	SpruceTreeY        int    = 6
	Trout              string = "trout"
	TroutLevel         int    = 20
	TroutX             int    = -2
	TroutY             int    = 6
	ItemRetrieved      string = "item_retrieved"
	TaskCreated        string = "task_created"
	TaskRetrieved      string = "task_retrieved"
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
