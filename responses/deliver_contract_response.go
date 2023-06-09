package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type DeliverContractResponse struct {
	Data DeliverContractResponseData `json:"data"`
}

type DeliverContractResponseData struct {
	Contract models.Contract  `json:"contract"`
	Cargo    models.ShipCargo `json:"cargo"`
}
