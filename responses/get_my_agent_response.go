package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type GetMyAgentResponse struct {
	Data models.Agent `json:"data"`
}
