package responses

import "spacetraders_sdk/models"

type GetContractResponse struct {
	Data models.Contract `json:"data"`
}
