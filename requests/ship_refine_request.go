package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type ShipRefineRequest struct {
	ShipSymbol string             `json:"-"`
	Produce    models.TradeSymbol `json:"produce"`
}

func NewShipRefineRequest(symbol string, produce models.TradeSymbol) *ShipRefineRequest {
	return &ShipRefineRequest{
		ShipSymbol: symbol,
		Produce:    produce,
	}
}

func (req *ShipRefineRequest) Method() string {
	return "POST"
}

func (req *ShipRefineRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/refine", req.ShipSymbol)
}

func (req *ShipRefineRequest) Body() (io.Reader, error) {
	body, err := json.Marshal(req)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}

func (req *ShipRefineRequest) AuthRequired() bool {
	return true
}

func (req *ShipRefineRequest) AuthAllowed() bool {
	return true
}
