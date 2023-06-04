package models

type ShipCrew struct {
	Current  int              `json:"current"`
	Required int              `json:"required"`
	Capacity int              `json:"capacity"`
	Rotation ShipCrewRotation `json:"rotation"`
	Morale   int              `json:"morale"`
	Wages    int              `json:"wages"`
}
