package models

type ShipNavRouteWaypoint struct {
	SystemWaypoint
	SystemSymbol string `json:"systemSymbol"`
}
