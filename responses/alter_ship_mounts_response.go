package responses

import "github.com/ult-biffer/spacetraders_sdk/models"

type AlterShipMountsResponse struct {
	Data AlterShipMountsResponseData `json:"data"`
}

type AlterShipMountsResponseData struct {
	Agent       models.Agent            `json:"agent"`
	Mounts      []models.ShipMount      `json:"mounts"`
	Cargo       models.ShipCargo        `json:"cargo"`
	Transaction models.MountTransaction `json:"transaction"`
}
