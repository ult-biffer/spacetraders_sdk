package models

type MarketTradeGood struct {
	Symbol        TradeSymbol `json:"symbol"`
	TradeVolume   int         `json:"tradeVolume"`
	Supply        string      `json:"supply"`
	PurchasePrice int         `json:"purchasePrice"`
	SellPrice     int         `json:"sellPrice"`
}
