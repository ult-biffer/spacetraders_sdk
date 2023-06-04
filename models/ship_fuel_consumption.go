package models

import "time"

type ShipFuelConsumption struct {
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
