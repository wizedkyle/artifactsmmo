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
	Name        string           `json:"name"`
	Code        string           `json:"code"`
	Level       int              `json:"level"`
	Type        string           `json:"type"`
	SubType     string           `json:"subType"`
	Description string           `json:"description"`
	Effects     []ItemEffect     `json:"effects,omitempty"`
	Craft       CraftInformation `json:"craft,omitempty"`
}

type ItemEffect struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type CraftInformation struct {
	Skill    string `json:"skill,omitempty"`
	Level    int    `json:"level,omitempty"`
	Items    []Item `json:"items,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}
