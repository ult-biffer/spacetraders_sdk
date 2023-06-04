package models

type Market struct {
	Symbol       string              `json:"symbol"`
	Exports      []TradeGood         `json:"exports"`
	Imports      []TradeGood         `json:"imports"`
	Exchange     []TradeGood         `json:"exchange"`
	Transactions []MarketTransaction `json:"transactions"`
	TradeGoods   []MarketTradeGood   `json:"tradeGoods"`
}
