package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func ListAllShips() ([]models.Ship, error) {
	resp, err := ListShips(1)

	if err != nil {
		return nil, err
	}

	ships := resp.Data
	pages := resp.Pages()

	if pages < 2 {
		return ships, nil
	}

	for i := 2; i < pages; i++ {
		resp, err := ListShips(i)

		if err != nil {
			return nil, err
		}

		ships = append(ships, resp.Data...)
		pages = resp.Pages()
	}

	return ships, nil
}

func ListShips(page int) (*responses.ListShipsResponse, error) {
	c := GetClient()
	req := requests.NewListShipsRequest(page)
	var result responses.ListShipsResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func GetShip(symbol string) (*models.Ship, error) {
	c := GetClient()
	req := requests.NewGetShipRequest(symbol)
	var result responses.GetShipResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func GetShipCargo(symbol string) (*models.ShipCargo, error) {
	c := GetClient()
	req := requests.NewGetShipCargoRequest(symbol)
	var result responses.GetShipCargoResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func GetShipCooldown(symbol string) (*models.Cooldown, error) {
	c := GetClient()
	req := requests.NewGetShipCooldownRequest(symbol)
	var result responses.GetShipCooldownResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func CreateChart(symbol string) (*responses.CreateChartResponse, error) {
	c := GetClient()
	req := requests.NewCreateChartRequest(symbol)
	var result responses.CreateChartResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func OrbitShip(symbol string) (*models.ShipNav, error) {
	c := GetClient()
	req := requests.NewOrbitShipRequest(symbol)
	var result responses.OrbitShipResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data.Nav, nil
}

func PurchaseShip(shipType models.ShipType, waypoint string) (*responses.PurchaseShipResponse, error) {
	c := GetClient()
	req := requests.NewPurchaseShipRequest(shipType, waypoint)
	var result responses.PurchaseShipResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func ShipRefine(symbol string, produce models.TradeSymbol) (*responses.ShipRefineResponse, error) {
	c := GetClient()
	req := requests.NewShipRefineRequest(symbol, produce)
	var result responses.ShipRefineResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
