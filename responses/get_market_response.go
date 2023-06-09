package responses

import "spacetraders_sdk/models"

type GetMarketResponse struct {
	Data models.Market `json:"data"`
}
