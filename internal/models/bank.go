package models

type ActionDepositBank struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type ActionDepositBankResponse struct {
	Data BankDepositResponse `json:"data"`
}

type BankDepositResponse struct {
	Cooldown  Cooldown    `json:"cooldown"`
	Item      ItemDetails `json:"item"`
	Bank      []Item      `json:"bank"`
	Character Character   `json:"character"`
}

type ItemDetails struct {
	Name        string           `bson:"name" json:"name"`
	Code        string           `bson:"code" json:"code"`
	Level       int              `bson:"level" json:"level"`
	Type        string           `bson:"type" json:"type"`
	SubType     string           `bson:"subType" json:"subType"`
	Description string           `bson:"description" json:"description"`
	Effects     []ItemEffect     `bson:"effects,omitempty" json:"effects,omitempty"`
	Craft       CraftInformation `bson:"craft,omitempty" json:"craft,omitempty"`
}

type ItemEffect struct {
	Name  string `bson:"name" json:"name"`
	Value int    `bson:"value" json:"value"`
}

type CraftInformation struct {
	Skill    string `bson:"skill,omitempty" json:"skill,omitempty"`
	Level    int    `bson:"level,omitempty" json:"level,omitempty"`
	Items    []Item `bson:"items,omitempty" json:"items,omitempty"`
	Quantity int    `bson:"quantity,omitempty" json:"quantity,omitempty"`
}
