package responses

import "spacetraders_sdk/models"

type ScanShipsResponse struct {
	Data ScanShipsResponseData `json:"data"`
}

type ScanShipsResponseData struct {
	Cooldown models.Cooldown      `json:"cooldown"`
	Ships    []models.ScannedShip `json:"ships"`
}
