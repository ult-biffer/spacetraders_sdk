package models

type ShipyardShip struct {
	Type          ShipType     `json:"type"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	PurchasePrice int          `json:"purchasePrice"`
	Frame         ShipFrame    `json:"frame"`
	Reactor       ShipReactor  `json:"reactor"`
	Engine        ShipEngine   `json:"engine"`
	Modules       []ShipModule `json:"module"`
	Mounts        []ShipMount  `json:"mounts"`
}
