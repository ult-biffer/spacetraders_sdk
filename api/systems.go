package api

import (
	"github.com/ult-biffer/spacetraders_sdk/models"
	"github.com/ult-biffer/spacetraders_sdk/requests"
	"github.com/ult-biffer/spacetraders_sdk/responses"
)

func ListAllSystems() ([]models.System, error) {
	resp, err := ListSystems(1)

	if err != nil {
		return nil, err
	}

	systems := resp.Data
	pages := resp.Pages()

	for i := 2; i < pages; i++ {
		resp, err = ListSystems(i)

		if err != nil {
			return nil, err
		}

		systems = append(systems, resp.Data...)
	}

	return systems, nil
}

func ListSystems(page int) (*responses.ListSystemsResponse, error) {
	req := requests.NewListSystemsRequest(page)
	var resp responses.ListSystemsResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func WaypointsInSystem(system string) ([]models.Waypoint, error) {
	resp, err := ListWaypoints(system, 1)

	if err != nil {
		return nil, err
	}

	waypoints := resp.Data
	pages := resp.Pages()

	for i := 2; i < pages; i++ {
		resp, err = ListWaypoints(system, i)

		if err != nil {
			return nil, err
		}

		waypoints = append(waypoints, resp.Data...)
	}

	return waypoints, nil
}

func ListWaypoints(system string, page int) (*responses.ListWaypointsResponse, error) {
	req := requests.NewListWaypointsRequest(system, page)
	var resp responses.ListWaypointsResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetSystem(system string) (*models.System, error) {
	req := requests.NewGetSystemRequest(system)
	var resp responses.GetSystemResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func GetWaypoint(waypoint string) (*models.Waypoint, error) {
	req := requests.NewGetWaypointRequest(waypoint)
	var resp responses.GetWaypointResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func GetMarket(waypoint string) (*models.Market, error) {
	req := requests.NewGetMarketRequest(waypoint)
	var resp responses.GetMarketResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func GetShipyard(waypoint string) (*models.Shipyard, error) {
	req := requests.NewGetShipyardRequest(waypoint)
	var resp responses.GetShipyardResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func GetJumpGate(waypoint string) (*models.JumpGate, error) {
	req := requests.NewGetJumpGateRequest(waypoint)
	var resp responses.GetJumpGateResponse

	if err := GetClient().ExecuteRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
