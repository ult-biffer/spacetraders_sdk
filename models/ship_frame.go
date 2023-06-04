package models

type ShipFrame struct {
	Describable[ShipFrameSymbol]
	Condition      int              `json:"condition"`
	ModuleSlots    int              `json:"moduleSlots"`
	MountingPoints int              `json:"mountingPoints"`
	FuelCapacity   int              `json:"fuelCapacity"`
	Requirements   ShipRequirements `json:"requirements"`
}
