package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type GetMarketResponse struct {
	Data models.Market `json:"data"`
}
