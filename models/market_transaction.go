package models

import "time"

type MarketTransaction struct {
	WaypointSymbol string      `json:"waypointSymbol"`
	ShipSymbol     string      `json:"shipSymbol"`
	TradeSymbol    TradeSymbol `json:"tradeSymbol"`
	Type           string      `json:"type"`
	Units          int         `json:"units"`
	PricePerUnit   int         `json:"pricePerUnit"`
	TotalPrice     int         `json:"totalPrice"`
	Timestamp      time.Time   `json:"timestamp"`
}
