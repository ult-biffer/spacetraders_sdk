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
