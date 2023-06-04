package models

type ShipMount struct {
	Describable[ShipMountSymbol]
	Strength     int              `json:"strength"`
	Deposits     []string         `json:"deposits"`
	Requirements ShipRequirements `json:"requirements"`
}
