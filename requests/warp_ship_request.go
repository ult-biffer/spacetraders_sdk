package requests

import (
	"fmt"
	"io"
)

type WarpShipRequest struct {
	ShipSymbol     string `json:"-"`
	WaypointSymbol string `json:"waypointSymbol"`
}

func NewWarpShipRequest(ship string, waypoint string) *WarpShipRequest {
	return &WarpShipRequest{
		ShipSymbol:     ship,
		WaypointSymbol: waypoint,
	}
}

func (*WarpShipRequest) Method() string {
	return "POST"
}

func (req *WarpShipRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/warp", req.ShipSymbol)
}

func (req *WarpShipRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*WarpShipRequest) AuthRequired() bool {
	return true
}

func (*WarpShipRequest) AuthAllowed() bool {
	return true
}
