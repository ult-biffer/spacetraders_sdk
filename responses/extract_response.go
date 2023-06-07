package responses

import "spacetraders_sdk/models"

type ExtractResponse struct {
	Data ExtractResponseData `json:"data"`
}

type ExtractResponseData struct {
	Cooldown   models.Cooldown   `json:"cooldown"`
	Extraction models.Extraction `json:"extraction"`
	Cargo      models.ShipCargo  `json:"cargo"`
}
