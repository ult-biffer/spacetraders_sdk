package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type GetShipResponse struct {
	Data models.Ship `json:"data"`
}
