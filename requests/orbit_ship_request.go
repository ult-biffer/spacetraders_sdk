package requests

import (
	"fmt"
	"io"
)

type OrbitShipRequest struct {
	Symbol string
}

func NewOrbitShipRequest(symbol string) *OrbitShipRequest {
	return &OrbitShipRequest{
		Symbol: symbol,
	}
}

func (req *OrbitShipRequest) Method() string {
	return "POST"
}

func (req *OrbitShipRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/orbit", req.Symbol)
}

func (req *OrbitShipRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (req *OrbitShipRequest) AuthRequired() bool {
	return true
}

func (req *OrbitShipRequest) AuthAllowed() bool {
	return true
}
