package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func GetShipCargo(symbol string) (*models.ShipCargo, error) {
	c := GetClient()
	req := requests.NewGetShipCargoRequest(symbol)
	var result responses.GetShipCargoResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}
