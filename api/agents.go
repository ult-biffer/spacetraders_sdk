package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func GetMyAgent() (*models.Agent, error) {
	req := &requests.GetMyAgentRequest{}
	var resp responses.GetMyAgentResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
