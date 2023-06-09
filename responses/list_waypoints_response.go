package responses

import (
	"math"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type ListWaypointsResponse struct {
	Data []models.Waypoint `json:"data"`
	Meta models.Meta       `json:"meta"`
}

func (resp ListWaypointsResponse) Pages() int {
	f := float64(resp.Meta.Total) / float64(models.MAX_LIMIT)
	return int(math.Ceil(f))
}
