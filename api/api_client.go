package api

import (
	"net/http"
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

func SetToken(token string) {
	client.Token = &token
}
