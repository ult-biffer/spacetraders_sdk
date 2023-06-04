package models

type ShipModule struct {
	Describable[ShipModuleSymbol]
	Capacity     *int             `json:"capacity,omitempty"`
	Range        *int             `json:"range,omitempty"`
	Requirements ShipRequirements `json:"requirements"`
}
