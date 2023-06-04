package api

import (
	"encoding/json"
	"io"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func Register(symbol string, faction string, email *string) (*responses.RegisterResponse, error) {
	client := GetClient()
	req := requests.NewRegisterRequest(symbol, faction, email)

	resp, err := requests.Execute(req, client.Http, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result *responses.RegisterResponse
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}

	SetToken(result.Data.Token)
	return result, nil
}
