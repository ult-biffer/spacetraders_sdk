package responses

import "spacetraders_sdk/models"

type GetShipCooldownResponse struct {
	Data models.Cooldown `json:"data"`
}
