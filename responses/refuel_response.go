package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type RefuelResponse struct {
	Data RefuelResponseData `json:"data"`
}

type RefuelResponseData struct {
	Agent       models.Agent             `json:"agent"`
	Fuel        models.ShipFuel          `json:"fuel"`
	Transaction models.MarketTransaction `json:"transaction"`
}
