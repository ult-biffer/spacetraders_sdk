package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func PurchaseShip(shipType models.ShipType, waypoint string) (*responses.PurchaseShipResponse, error) {
	c := GetClient()
	req := requests.NewPurchaseShipRequest(shipType, waypoint)
	var result responses.PurchaseShipResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
