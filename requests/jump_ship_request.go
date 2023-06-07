package requests

import (
	"fmt"
	"io"
)

type JumpShipRequest struct {
	ShipSymbol   string `json:"-"`
	SystemSymbol string `json:"systemSymbol"`
}

func NewJumpShipRequest(ship string, system string) *JumpShipRequest {
	return &JumpShipRequest{
		ShipSymbol:   ship,
		SystemSymbol: system,
	}
}

func (*JumpShipRequest) Method() string {
	return "POST"
}

func (req *JumpShipRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/jump", req.ShipSymbol)
}

func (req *JumpShipRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*JumpShipRequest) AuthRequired() bool {
	return true
}

func (*JumpShipRequest) AuthAllowed() bool {
	return true
}
