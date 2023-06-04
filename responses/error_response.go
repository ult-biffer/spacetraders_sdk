package responses

import (
	"encoding/json"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error ErrorResponseData `json:"error"`
}

type ErrorResponseData struct {
	Message string                 `json:"message"`
	Code    int                    `json:"code"`
	Data    map[string]interface{} `json:"data"`
}

func NewErrorResponse(resp *http.Response) (*ErrorResponse, error) {
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result *ErrorResponse

	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewFakeErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorResponseData{
			Message: message,
		},
	}
}
