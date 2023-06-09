package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type AcceptContractResponse struct {
	Data AcceptContractResponseData `json:"data"`
}

type AcceptContractResponseData struct {
	Agent    models.Agent    `json:"agent"`
	Contract models.Contract `json:"contract"`
}
