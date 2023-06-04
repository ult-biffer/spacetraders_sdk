package models

import "time"

type ShipNavRoute struct {
	Destination   ShipNavRouteWaypoint `json:"destination"`
	Departure     ShipNavRouteWaypoint `json:"departure"`
	DepartureTime time.Time            `json:"departureTime"`
	Arrival       time.Time            `json:"arrival"`
}
