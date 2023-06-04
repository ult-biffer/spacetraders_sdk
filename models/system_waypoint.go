package models

type SystemWaypoint struct {
	Coordinates
	Symbol string       `json:"symbol"`
	Type   WaypointType `json:"type"`
}
