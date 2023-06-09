package responses

import (
	"math"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type ListFactionsResponse struct {
	Data []models.Faction `json:"data"`
	Meta models.Meta      `json:"meta"`
}

func (resp ListFactionsResponse) Pages() int {
	f := float64(resp.Meta.Total) / float64(models.MAX_LIMIT)
	return int(math.Ceil(f))
}
