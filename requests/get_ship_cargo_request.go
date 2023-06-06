package requests

import (
	"fmt"
	"io"
)

type GetShipCargoRequest struct {
	Symbol string
}

func NewGetShipCargoRequest(symbol string) *GetShipCargoRequest {
	return &GetShipCargoRequest{
		Symbol: symbol,
	}
}

func (req *GetShipCargoRequest) Method() string {
	return "GET"
}

func (req *GetShipCargoRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/cargo", req.Symbol)
}

func (req *GetShipCargoRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (req *GetShipCargoRequest) AuthRequired() bool {
	return true
}

func (req *GetShipCargoRequest) AuthAllowed() bool {
	return true
}
