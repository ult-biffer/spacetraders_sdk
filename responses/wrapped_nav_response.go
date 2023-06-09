package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type WrappedNavResponse struct {
	Data WrappedNavResponseData `json:"data"`
}

type WrappedNavResponseData struct {
	Nav models.ShipNav `json:"nav"`
}
