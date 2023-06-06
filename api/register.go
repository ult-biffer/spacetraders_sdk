package api

import (
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func Register(symbol string, faction string, email *string) (*responses.RegisterResponse, error) {
	c := GetClient()
	req := requests.NewRegisterRequest(symbol, faction, email)
	var result responses.RegisterResponse

	if err := c.ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	c.SetToken(result.Data.Token)
	return &result, nil
}
