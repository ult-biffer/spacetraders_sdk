package models

import "time"

type Cooldown struct {
	ShipSymbol       string     `json:"shipSymbol"`
	TotalSeconds     int        `json:"totalSeconds"`
	RemainingSeconds int        `json:"remainingSeconds"`
	Expiration       *time.Time `json:"expiration"`
}
