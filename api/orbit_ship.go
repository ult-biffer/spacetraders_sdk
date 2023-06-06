package api

import (
	"encoding/json"
	"io"
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func OrbitShip(symbol string) (*models.ShipNav, error) {
	c := GetClient()
	req := requests.NewOrbitShipRequest(symbol)
	resp, err := requests.Execute(req, c.Http, c.Token)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var result responses.OrbitShipResponse

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result.Data.Nav, nil
}
