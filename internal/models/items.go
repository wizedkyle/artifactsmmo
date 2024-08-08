package models

type GetAllItemsQueryParameters struct {
	CraftMaterial string `json:"craft_material"`
	CraftSkill    string `json:"craft_skill"`
	MaxLevel      int    `json:"max_level"`
	MinLevel      int    `json:"min_level"`
	Name          string `json:"name"`
	Page          int    `json:"page"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

type ListItemParameters struct {
	Type    string `json:"type"`
	SubType string `json:"subtype"`
	Level   int    `json:"level"`
}

type ItemsResponse struct {
	Data  []ItemDetails `json:"data"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Pages int           `json:"pages"`
}

type ItemResponse struct {
	Data ItemDetails             `json:"data"`
	Ge   GrandExchangeItemDetail `json:"ge,omitempty"`
}
