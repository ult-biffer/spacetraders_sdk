package responses

import "spacetraders_sdk/models"

type GetMyAgentResponse struct {
	Data models.Agent `json:"data"`
}
