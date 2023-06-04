package models

type ShipRequirements struct {
	Power *int `json:"power"`
	Crew  *int `json:"crew"`
	Slots *int `json:"slots"`
}
