package models

type ShipCargo struct {
	Capacity  int             `json:"capacity"`
	Units     int             `json:"units"`
	Inventory []ShipCargoItem `json:"inventory"`
}
