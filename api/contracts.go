package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func ListAllContracts() ([]models.Contract, error) {
	resp, err := ListContracts(1)

	if err != nil {
		return nil, err
	}

	contracts := resp.Data
	pages := resp.Pages()

	if pages < 2 {
		return contracts, nil
	}

	for i := 2; i < pages; i++ {
		resp, err = ListContracts(i)

		if err != nil {
			return nil, err
		}

		contracts = append(contracts, resp.Data...)
	}

	return contracts, nil
}

func ListContracts(page int) (*responses.ListContractsResponse, error) {
	req := requests.NewListContractsRequest(page)
	var resp responses.ListContractsResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetContract(id string) (*models.Contract, error) {
	req := requests.NewGetContractRequest(id)
	var resp responses.GetContractResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func AcceptContract(id string) (*responses.AcceptContractResponse, error) {
	req := requests.NewAcceptContractRequest(id)
	var resp responses.AcceptContractResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func DeliverContract(id, ship string, item models.TradeSymbol, units int) (*responses.DeliverContractResponse, error) {
	req := requests.NewDeliverContractRequest(id, ship, item, units)
	var resp responses.DeliverContractResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func FulfillContract(id string) (*responses.AcceptContractResponse, error) {
	req := requests.NewFulfillContractRequest(id)
	var resp responses.AcceptContractResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
