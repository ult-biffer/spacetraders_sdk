package responses

import "spacetraders_sdk/models"

type GetShipCargoResponse struct {
	Data models.ShipCargo `json:"data"`
}
