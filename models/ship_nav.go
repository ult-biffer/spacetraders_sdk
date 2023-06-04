package models

type ShipNav struct {
	SystemSymbol   string            `json:"systemSymbol"`
	WaypointSymbol string            `json:"waypointSymbol"`
	Route          ShipNavRoute      `json:"route"`
	Status         ShipNavStatus     `json:"status"`
	FlightMode     ShipNavFlightMode `json:"flightMode"`
}
