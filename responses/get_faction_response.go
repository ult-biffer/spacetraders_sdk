package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type GetFactionResponse struct {
	Data models.Faction `json:"data"`
}
