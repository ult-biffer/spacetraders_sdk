package requests

import (
	"fmt"
	"io"
)

type NegotiateContractRequest struct {
	ShipSymbol string
}

func NewNegotiateContractRequest(ship string) *NegotiateContractRequest {
	return &NegotiateContractRequest{
		ShipSymbol: ship,
	}
}

func (*NegotiateContractRequest) Method() string {
	return "POST"
}

func (req *NegotiateContractRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/negotiate/contract", req.ShipSymbol)
}

func (*NegotiateContractRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*NegotiateContractRequest) AuthRequired() bool {
	return true
}

func (*NegotiateContractRequest) AuthAllowed() bool {
	return true
}
