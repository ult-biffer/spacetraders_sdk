package models

type ShipEngine struct {
	Describable[ShipEngineSymbol]
	Condition    int              `json:"condition"`
	Speed        int              `json:"speed"`
	Requirements ShipRequirements `json:"requirements"`
}
