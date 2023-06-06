package requests

import (
	"fmt"
	"io"
)

type GetShipRequest struct {
	Symbol string
}

func NewGetShipRequest(symbol string) *GetShipRequest {
	return &GetShipRequest{
		Symbol: symbol,
	}
}

func (req *GetShipRequest) Method() string {
	return "GET"
}

func (req *GetShipRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s", req.Symbol)
}

func (req *GetShipRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (req *GetShipRequest) AuthRequired() bool {
	return true
}

func (req *GetShipRequest) AuthAllowed() bool {
	return true
}
