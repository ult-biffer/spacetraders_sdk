package responses

import (
	"math"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type ListShipsResponse struct {
	Data []models.Ship `json:"data"`
	Meta models.Meta   `json:"meta"`
}

func (resp ListShipsResponse) Pages() int {
	f := float64(resp.Meta.Total) / float64(models.MAX_LIMIT)
	return int(math.Ceil(f))
}
