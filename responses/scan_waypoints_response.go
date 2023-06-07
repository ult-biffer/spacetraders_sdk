package responses

import "spacetraders_sdk/models"

type ScanWaypointsResponse struct {
	Data ScanWaypointsResponseData `json:"data"`
}

type ScanWaypointsResponseData struct {
	Cooldown  models.Cooldown          `json:"cooldown"`
	Waypoints []models.ScannedWaypoint `json:"waypoints"`
}
