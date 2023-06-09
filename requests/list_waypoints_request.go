package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type ListWaypointsRequest struct {
	SystemSymbol string
	Page         int
}

func NewListWaypointsRequest(system string, page int) *ListWaypointsRequest {
	return &ListWaypointsRequest{
		SystemSymbol: system,
		Page:         page,
	}
}

func (*ListWaypointsRequest) Method() string {
	return "GET"
}

func (req *ListWaypointsRequest) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints?limit=%d&page=%d", req.SystemSymbol, models.MAX_LIMIT, req.Page)
}

func (*ListWaypointsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*ListWaypointsRequest) AuthRequired() bool {
	return false
}

func (*ListWaypointsRequest) AuthAllowed() bool {
	return true
}
