package models

type GrandExchangeItemDetail struct {
	Code      string `json:"code"`
	Stock     int    `json:"stock"`
	SellPrice int    `json:"sell_price"`
	BuyPrice  int    `json:"buy_price"`
}
