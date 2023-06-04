package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Request interface {
	Method() string
	Path() string
	Body() (io.Reader, error)
	AuthRequired() bool
}

type RequestDoer interface {
	Do(*http.Request) (*http.Response, error)
}

func Execute(r Request, d RequestDoer, token *string) (*http.Response, error) {
	req := requestObject(r)

	if req == nil {
		return nil, fmt.Errorf("could not build request %v", r)
	}

	if r.AuthRequired() {
		if token == nil {
			return nil, fmt.Errorf("must supply token for authenticated request")
		}

		auth := fmt.Sprintf("Bearer %s", *token)
		req.Header.Add("Authorization", auth)
	}

	resp, err := d.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return handleHttpError(resp, r, d, token)
	}

	return resp, nil
}

func requestObject(r Request) *http.Request {
	url := urlFromPath(r.Path())

	body, err := r.Body()

	if err != nil {
		return nil
	}

	req, err := http.NewRequest(r.Method(), url, body)

	if err != nil {
		return nil
	}

	return req
}

func urlFromPath(path string) string {
	return fmt.Sprintf("https://api.spacetraders.io/v2%s", path)
}

func handleHttpError(resp *http.Response, r Request, d RequestDoer, token *string) (*http.Response, error) {
	switch resp.StatusCode {
	case 429:
		return retryAfter(resp, r, d, token)
	case 502:
		return recoverFromDdos(r, d, token)
	default:
		return resp, NewGenericHttpError(resp)
	}
}

func retryAfter(resp *http.Response, r Request, d RequestDoer, token *string) (*http.Response, error) {
	var body map[string]interface{}

	s, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(s, &body); err != nil {
		return nil, err
	}

	e := body["error"].(map[string]interface{})
	data := e["data"].(map[string]interface{})
	timeout := data["retryAfter"].(int)

	timeToWait := time.Duration(timeout) * time.Second
	time.Sleep(timeToWait)

	return Execute(r, d, token)
}

func recoverFromDdos(r Request, d RequestDoer, token *string) (*http.Response, error) {
	timeToWait := time.Duration(10) * time.Second
	time.Sleep(timeToWait)

	return Execute(r, d, token)
}
