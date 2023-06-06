package api

import (
	"encoding/json"
	"io"
	"net/http"
	"spacetraders_sdk/requests"
	"time"
)

type ApiClient struct {
	Http  *http.Client
	Token *string
}

func NewApiClient() *ApiClient {
	return &ApiClient{
		Http: NewThrottledDefaultClient(time.Second, 2),
	}
}

var client *ApiClient = NewApiClient()

func GetClient() *ApiClient {
	return client
}

func (c *ApiClient) SetToken(token string) {
	c.Token = &token
}

func (c *ApiClient) ExecuteRequest(req requests.Request, resp any) error {
	r, err := requests.Execute(req, c.Http, c.Token)

	if err != nil {
		return err
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return err
	}

	return nil
}
