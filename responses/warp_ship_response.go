package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type WarpShipResponse struct {
	Data WarpShipResponseData `json:"data"`
}

type WarpShipResponseData struct {
	Fuel models.ShipFuel `json:"fuel"`
	Nav  models.ShipNav  `json:"nav"`
}
