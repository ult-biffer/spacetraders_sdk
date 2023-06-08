package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func ListAllFactions() ([]models.Faction, error) {
	resp, err := ListFactions(1)

	if err != nil {
		return nil, err
	}

	factions := resp.Data
	pages := resp.Pages()

	if pages < 2 {
		return factions, nil
	}

	for i := 2; i < pages; i++ {
		resp, err := ListFactions(i)

		if err != nil {
			return nil, err
		}

		factions = append(factions, resp.Data...)
		pages = resp.Pages()
	}

	return factions, nil
}

func ListFactions(page int) (*responses.ListFactionsResponse, error) {
	req := requests.NewListFactionsRequest(page)
	var result responses.ListFactionsResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
