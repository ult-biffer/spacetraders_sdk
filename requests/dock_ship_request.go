package requests

import (
	"fmt"
	"io"
)

type DockShipRequest struct {
	Symbol string
}

func NewDockShipRequest(symbol string) *DockShipRequest {
	return &DockShipRequest{
		Symbol: symbol,
	}
}

func (req *DockShipRequest) Method() string {
	return "POST"
}

func (req *DockShipRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/dock", req.Symbol)
}

func (req *DockShipRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (req *DockShipRequest) AuthRequired() bool {
	return true
}

func (req *DockShipRequest) AuthAllowed() bool {
	return true
}
