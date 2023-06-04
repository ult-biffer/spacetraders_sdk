package api

import (
	"encoding/json"
	"io"
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
	client := GetClient()
	req := requests.NewListShipsRequest(page)
	resp, err := requests.Execute(req, client.Http, client.Token)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result *responses.ListShipsResponse
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}

	return result, nil
}
