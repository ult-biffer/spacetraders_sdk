package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func ShipRefine(symbol string, produce models.TradeSymbol) (*responses.ShipRefineResponse, error) {
	c := GetClient()
	req := requests.NewShipRefineRequest(symbol, produce)
	var result responses.ShipRefineResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
