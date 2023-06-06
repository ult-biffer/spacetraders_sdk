package responses

import "spacetraders_sdk/models"

type GetShipResponse struct {
	Data models.Ship `json:"data"`
}
