package api

import (
	"encoding/json"
	"io"
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func PurchaseShip(shipType models.ShipType, waypoint string) (*responses.PurchaseShipResponse, error) {
	c := GetClient()
	req := requests.NewPurchaseShipRequest(shipType, waypoint)
	resp, err := requests.Execute(req, c.Http, c.Token)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result responses.PurchaseShipResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
