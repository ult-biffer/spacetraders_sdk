package responses

import "spacetraders_sdk/models"

type WrappedCargoResponse struct {
	Data WrappedCargoResponseData `json:"data"`
}

type WrappedCargoResponseData struct {
	Cargo models.ShipCargo `json:"cargo"`
}
