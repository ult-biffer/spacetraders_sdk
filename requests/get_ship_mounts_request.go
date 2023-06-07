package requests

import (
	"fmt"
	"io"
)

type GetShipMountsRequest struct {
	Symbol string
}

func NewGetShipMountsRequest(symbol string) *GetShipMountsRequest {
	return &GetShipMountsRequest{
		Symbol: symbol,
	}
}

func (*GetShipMountsRequest) Method() string {
	return "GET"
}

func (req *GetShipMountsRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/mounts", req.Symbol)
}

func (*GetShipMountsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetShipMountsRequest) AuthRequired() bool {
	return true
}

func (*GetShipMountsRequest) AuthAllowed() bool {
	return true
}
