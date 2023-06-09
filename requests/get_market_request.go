package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/util"
)

type GetMarketRequest struct {
	SystemSymbol   string
	WaypointSymbol string
}

func NewGetMarketRequest(waypoint string) *GetMarketRequest {
	loc := util.NewLocation(waypoint)

	return &GetMarketRequest{
		SystemSymbol:   loc.System,
		WaypointSymbol: loc.Waypoint,
	}
}

func (*GetMarketRequest) Method() string {
	return "GET"
}

func (req *GetMarketRequest) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s/market", req.SystemSymbol, req.WaypointSymbol)
}

func (*GetMarketRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetMarketRequest) AuthRequired() bool {
	return false
}

func (*GetMarketRequest) AuthAllowed() bool {
	return true
}
