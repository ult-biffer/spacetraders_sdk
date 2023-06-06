package models

type RefinedTradeGood struct {
	TradeSymbol TradeSymbol `json:"tradeSymbol"`
	Units       int         `json:"units"`
}
