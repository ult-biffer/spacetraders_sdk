package requests

import (
	"fmt"
	"io"
)

type NavigateShipRequest struct {
	ShipSymbol     string `json:"-"`
	WaypointSymbol string `json:"waypointSymbol"`
}

func NewNavigateShipRequest(ship string, waypoint string) *NavigateShipRequest {
	return &NavigateShipRequest{
		ShipSymbol:     ship,
		WaypointSymbol: waypoint,
	}
}

func (*NavigateShipRequest) Method() string {
	return "POST"
}

func (req *NavigateShipRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/navigate", req.ShipSymbol)
}

func (req *NavigateShipRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*NavigateShipRequest) AuthRequired() bool {
	return true
}

func (*NavigateShipRequest) AuthAllowed() bool {
	return true
}
