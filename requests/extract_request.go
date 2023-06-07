package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type ExtractRequest struct {
	Symbol string         `json:"-"`
	Survey *models.Survey `json:"survey"`
}

func NewExtractRequest(symbol string, survey *models.Survey) *ExtractRequest {
	return &ExtractRequest{
		Symbol: symbol,
		Survey: survey,
	}
}

func (req *ExtractRequest) Method() string {
	return "POST"
}

func (req *ExtractRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/extract", req.Symbol)
}

func (req *ExtractRequest) Body() (io.Reader, error) {
	if req.Survey == nil {
		return nil, nil
	}

	body, err := json.Marshal(req)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}

func (req *ExtractRequest) AuthRequired() bool {
	return true
}

func (req *ExtractRequest) AuthAllowed() bool {
	return true
}
