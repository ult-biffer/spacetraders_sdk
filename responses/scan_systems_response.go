package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type ScanSystemsResponse struct {
	Data ScanSystemsResponseData `json:"data"`
}

type ScanSystemsResponseData struct {
	Cooldown models.Cooldown        `json:"cooldown"`
	Systems  []models.ScannedSystem `json:"systems"`
}
