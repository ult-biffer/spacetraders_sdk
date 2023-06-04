package models

import "time"

type ShipyardTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	Price          int       `json:"price"`
	AgentSymbol    string    `json:"agentSymbol"`
	Timestamp      time.Time `json:"timestamp"`
}
