package responses

import "spacetraders_sdk/models"

type GetFactionResponse struct {
	Data models.Faction `json:"data"`
}
