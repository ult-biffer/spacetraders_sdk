package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type ShipRefineResponse struct {
	Data ShipRefineResponseData `json:"data"`
}

type ShipRefineResponseData struct {
	Cargo    models.ShipCargo        `json:"cargo"`
	Cooldown models.Cooldown         `json:"cooldown"`
	Produced models.RefinedTradeGood `json:"produced"`
	Consumed models.RefinedTradeGood `json:"consumed"`
}
