package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type GetShipMountsResponse struct {
	Data []models.ShipMount `json:"data"`
}
