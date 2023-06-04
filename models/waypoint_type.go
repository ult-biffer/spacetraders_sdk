package models

type WaypointType string

const (
	WP_PLANET          WaypointType = "PLANET"
	WP_GAS_GIANT       WaypointType = "GAS_GIANT"
	WP_MOON            WaypointType = "MOON"
	WP_ORBITAL_STATION WaypointType = "ORBITAL_STATION"
	WP_JUMP_GATE       WaypointType = "JUMP_GATE"
	WP_ASTEROID_FIELD  WaypointType = "ASTEROID_FIELD"
	WP_NEBULA          WaypointType = "NEBULA"
	WP_DEBRIS_FIELD    WaypointType = "DEBRIS_FIELD"
	WP_GRAVITY_WELL    WaypointType = "GRAVITY_WELL"
)
