package api

import (
	"github.com/ult-biffer/spacetraders_sdk/models"
	"github.com/ult-biffer/spacetraders_sdk/requests"
	"github.com/ult-biffer/spacetraders_sdk/responses"
)

func GetMyAgent() (*models.Agent, error) {
	req := &requests.GetMyAgentRequest{}
	var resp responses.GetMyAgentResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
