package models

type ShipReactor struct {
	Describable[ShipReactorSymbol]
	Condition    int              `json:"condition"`
	PowerOutput  int              `json:"powerOutput"`
	Requirements ShipRequirements `json:"requirements"`
}
