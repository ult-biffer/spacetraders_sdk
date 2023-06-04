package models

type System struct {
	SystemDetails
	Waypoints []SystemWaypoint `json:"waypoints"`
	Factions  []SystemFaction  `json:"factions"`
}
