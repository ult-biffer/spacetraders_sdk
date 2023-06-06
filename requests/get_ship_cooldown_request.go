package requests

import (
	"fmt"
	"io"
)

type GetShipCooldownRequest struct {
	Symbol string
}

func NewGetShipCooldownRequest(symbol string) *GetShipCooldownRequest {
	return &GetShipCooldownRequest{
		Symbol: symbol,
	}
}

func (req *GetShipCooldownRequest) Method() string {
	return "GET"
}

func (req *GetShipCooldownRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/cooldown", req.Symbol)
}

func (req *GetShipCooldownRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (req *GetShipCooldownRequest) AuthRequired() bool {
	return true
}

func (req *GetShipCooldownRequest) AuthAllowed() bool {
	return true
}
