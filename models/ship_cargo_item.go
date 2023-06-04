package models

type ShipCargoItem struct {
	TradeGood
	Units int `json:"units"`
}
