package responses

import "spacetraders_sdk/models"

type GetShipMountsResponse struct {
	Data []models.ShipMount `json:"data"`
}
