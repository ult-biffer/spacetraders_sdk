package requests

import (
	"fmt"
	"io"
)

type GetShipNavRequest struct {
	Symbol string
}

func NewGetShipNavRequest(symbol string) *GetShipNavRequest {
	return &GetShipNavRequest{
		Symbol: symbol,
	}
}

func (*GetShipNavRequest) Method() string {
	return "GET"
}

func (req *GetShipNavRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/nav", req.Symbol)
}

func (*GetShipNavRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetShipNavRequest) AuthRequired() bool {
	return true
}

func (*GetShipNavRequest) AuthAllowed() bool {
	return true
}
