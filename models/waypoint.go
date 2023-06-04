package models

type Waypoint struct {
	ShipNavRouteWaypoint
	Orbitals []WaypointOrbital `json:"orbitals"`
	Faction  WaypointFaction   `json:"faction"`
	Traits   []WaypointTrait   `json:"traits"`
	Chart    *Chart            `json:"chart"`
}
