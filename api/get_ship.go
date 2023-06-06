package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func GetShip(symbol string) (*models.Ship, error) {
	c := GetClient()
	req := requests.NewGetShipRequest(symbol)
	var result responses.GetShipResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}
