package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type CargoTransactionResponse struct {
	Data CargoTransactionResponseData `json:"data"`
}

type CargoTransactionResponseData struct {
	Agent       models.Agent             `json:"agent"`
	Cargo       models.ShipCargo         `json:"cargo"`
	Transaction models.MarketTransaction `json:"transaction"`
}
