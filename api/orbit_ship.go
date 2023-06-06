package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func OrbitShip(symbol string) (*models.ShipNav, error) {
	c := GetClient()
	req := requests.NewOrbitShipRequest(symbol)
	var result responses.OrbitShipResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data.Nav, nil
}
