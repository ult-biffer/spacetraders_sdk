package responses

import "spacetraders_sdk/models"

type JumpShipResponse struct {
	Data JumpShipResponseData `json:"data"`
}

type JumpShipResponseData struct {
	Cooldown models.Cooldown `json:"cooldown"`
	Nav      models.ShipNav  `json:"nav"`
}
