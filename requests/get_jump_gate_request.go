package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/util"
)

type GetJumpGateRequest struct {
	SystemSymbol   string
	WaypointSymbol string
}

func NewGetJumpGateRequest(waypoint string) *GetJumpGateRequest {
	loc := util.NewLocation(waypoint)

	return &GetJumpGateRequest{
		SystemSymbol:   loc.System,
		WaypointSymbol: loc.Waypoint,
	}
}

func (*GetJumpGateRequest) Method() string {
	return "GET"
}

func (req *GetJumpGateRequest) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s/jump-gate", req.SystemSymbol, req.WaypointSymbol)
}

func (*GetJumpGateRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetJumpGateRequest) AuthRequired() bool {
	return false
}

func (*GetJumpGateRequest) AuthAllowed() bool {
	return true
}
