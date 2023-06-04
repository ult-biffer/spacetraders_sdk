package models

type ShipFuel struct {
	Current  int                 `json:"current"`
	Capacity int                 `json:"capacity"`
	Consumed ShipFuelConsumption `json:"consumed"`
}
