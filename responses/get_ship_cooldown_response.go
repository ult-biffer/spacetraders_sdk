package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type GetShipCooldownResponse struct {
	Data models.Cooldown `json:"data"`
}
