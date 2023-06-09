package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/util"
)

type GetWaypointRequest struct {
	SystemSymbol   string
	WaypointSymbol string
}

func NewGetWaypointRequest(waypoint string) *GetWaypointRequest {
	loc := util.NewLocation(waypoint)

	return &GetWaypointRequest{
		SystemSymbol:   loc.System,
		WaypointSymbol: loc.Waypoint,
	}
}

func (*GetWaypointRequest) Method() string {
	return "GET"
}

func (req *GetWaypointRequest) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s", req.SystemSymbol, req.WaypointSymbol)
}

func (*GetWaypointRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetWaypointRequest) AuthRequired() bool {
	return false
}

func (*GetWaypointRequest) AuthAllowed() bool {
	return true
}
