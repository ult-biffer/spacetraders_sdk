package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type CreateChartResponse struct {
	Data CreateChartResponseData `json:"data"`
}

type CreateChartResponseData struct {
	Chart    models.Chart    `json:"chart"`
	Waypoint models.Waypoint `json:"waypoint"`
}
