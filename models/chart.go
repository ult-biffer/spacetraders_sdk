package models

import "time"

type Chart struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	SubmittedBy    string    `json:"submittedBy"`
	SubmittedOn    time.Time `json:"submittedOn"`
}
