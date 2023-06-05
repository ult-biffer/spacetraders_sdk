package responses

import "spacetraders_sdk/models"

type PurchaseShipResponse struct {
	Data PurchaseShipResponseData `json:"data"`
}

type PurchaseShipResponseData struct {
	Agent       models.Agent               `json:"agent"`
	Ship        models.Ship                `json:"ship"`
	Transaction models.ShipyardTransaction `json:"transaction"`
}
