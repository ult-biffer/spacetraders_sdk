package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/util"
)

type GetShipyardRequest struct {
	SystemSymbol   string
	WaypointSymbol string
}

func NewGetShipyardRequest(waypoint string) *GetShipyardRequest {
	loc := util.NewLocation(waypoint)

	return &GetShipyardRequest{
		SystemSymbol:   loc.System,
		WaypointSymbol: loc.Waypoint,
	}
}

func (*GetShipyardRequest) Method() string {
	return "GET"
}

func (req *GetShipyardRequest) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s/shipyard", req.SystemSymbol, req.WaypointSymbol)
}

func (*GetShipyardRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetShipyardRequest) AuthRequired() bool {
	return false
}

func (*GetShipyardRequest) AuthAllowed() bool {
	return true
}
