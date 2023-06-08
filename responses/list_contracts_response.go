package responses

import (
	"math"
	"spacetraders_sdk/models"
)

type ListContractsResponse struct {
	Data []models.Contract `json:"data"`
	Meta models.Meta       `json:"meta"`
}

func (resp ListContractsResponse) Pages() int {
	f := float64(resp.Meta.Total) / float64(models.MAX_LIMIT)
	return int(math.Ceil(f))
}
