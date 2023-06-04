package models

type ShipNavFlightMode string

const (
	CRUISE  ShipNavFlightMode = "CRUISE"
	DRIFT   ShipNavFlightMode = "DRIFT"
	STEALTH ShipNavFlightMode = "STEALTH"
	BURN    ShipNavFlightMode = "BURN"
)
