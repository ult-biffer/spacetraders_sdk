package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type NegotiateContractResponse struct {
	Data NegotiateContractResponseData `json:"data"`
}

type NegotiateContractResponseData struct {
	Contract models.Contract `json:"contract"`
}
