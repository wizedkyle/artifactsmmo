package models

import "time"

const (
	Bank        string = "bank"
	BankX       int    = 4
	BankY       int    = 1
	CoalLevel   int    = 20
	Coal        string = "coal_rocks"
	CoalX       int    = 1
	CoalY       int    = 6
	Copper      string = "copper_ore"
	CopperLevel int    = 0
	CopperX     int    = 2
	CopperY     int    = 0
	GoldLevel   int    = 30
	Gold        string = "gold_rocks"
	GoldX       int    = 10
	GoldY       int    = -4
	IronLevel   int    = 10
	Iron        string = "iron_rocks"
	IronX       int    = 1
	IronY       int    = 7
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
