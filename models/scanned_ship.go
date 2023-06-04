package models

type ScannedShip struct {
	Symbol       string              `json:"symbol"`
	Registration ShipRegistration    `json:"registration"`
	Nav          ShipNav             `json:"nav"`
	Engine       ScannedShipEngine   `json:"engine"`
	Frame        *ScannedShipFrame   `json:"frame,omitempty"`
	Reactor      *ScannedShipReactor `json:"reactor,omitempty"`
	Mounts       []ScannedShipMount  `json:"mounts,omitempty"`
}
