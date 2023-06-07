package responses

import "spacetraders_sdk/models"

type NavigateShipResponse struct {
	Data NavigateShipResponseData `json:"data"`
}

type NavigateShipResponseData struct {
	Fuel models.ShipFuel `json:"fuel"`
	Nav  models.ShipNav  `json:"nav"`
}
