package responses

import "spacetraders_sdk/models"

type GetWaypointResponse struct {
	Data models.Waypoint `json:"data"`
}
