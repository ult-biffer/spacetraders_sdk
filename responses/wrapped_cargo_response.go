package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type WrappedCargoResponse struct {
	Data WrappedCargoResponseData `json:"data"`
}

type WrappedCargoResponseData struct {
	Cargo models.ShipCargo `json:"cargo"`
}
