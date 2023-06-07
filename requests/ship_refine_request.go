package requests

import (
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
	return marshal(req)
}

func (req *ShipRefineRequest) AuthRequired() bool {
	return true
}

func (req *ShipRefineRequest) AuthAllowed() bool {
	return true
}
