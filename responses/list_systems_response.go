package responses

import (
	"math"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type ListSystemsResponse struct {
	Data []models.System `json:"data"`
	Meta models.Meta     `json:"meta"`
}

func (resp ListSystemsResponse) Pages() int {
	f := float64(resp.Meta.Total) / float64(models.MAX_LIMIT)
	return int(math.Ceil(f))
}
