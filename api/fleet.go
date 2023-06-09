package api

import (
	"spacetraders_sdk/models"
	"spacetraders_sdk/requests"
	"spacetraders_sdk/responses"
)

func ListAllShips() ([]models.Ship, error) {
	resp, err := ListShips(1)

	if err != nil {
		return nil, err
	}

	ships := resp.Data
	pages := resp.Pages()

	for i := 2; i < pages; i++ {
		resp, err := ListShips(i)

		if err != nil {
			return nil, err
		}

		ships = append(ships, resp.Data...)
		pages = resp.Pages()
	}

	return ships, nil
}

func ListShips(page int) (*responses.ListShipsResponse, error) {
	req := requests.NewListShipsRequest(page)
	var result responses.ListShipsResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func GetShip(symbol string) (*models.Ship, error) {
	req := requests.NewGetShipRequest(symbol)
	var result responses.GetShipResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func GetShipCargo(symbol string) (*models.ShipCargo, error) {
	req := requests.NewGetShipCargoRequest(symbol)
	var result responses.GetShipCargoResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func GetShipCooldown(symbol string) (*models.Cooldown, error) {
	req := requests.NewGetShipCooldownRequest(symbol)
	var result responses.GetShipCooldownResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func GetShipNav(symbol string) (*models.ShipNav, error) {
	req := requests.NewGetShipNavRequest(symbol)
	var result responses.ShipNavResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func GetShipMounts(symbol string) ([]models.ShipMount, error) {
	req := requests.NewGetShipMountsRequest(symbol)
	var result responses.GetShipMountsResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func InstallShipMount(ship string, symbol models.ShipMountSymbol) (*responses.AlterShipMountsResponse, error) {
	req := requests.NewInstallShipMountRequest(ship, symbol)
	var result responses.AlterShipMountsResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func RemoveShipMount(ship string, symbol models.ShipMountSymbol) (*responses.AlterShipMountsResponse, error) {
	req := requests.NewRemoveShipMountRequest(ship, symbol)
	var result responses.AlterShipMountsResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func DockShip(symbol string) (*models.ShipNav, error) {
	req := requests.NewDockShipRequest(symbol)
	var result responses.DockShipResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data.Nav, nil
}

func JumpShip(ship string, system string) (*responses.JumpShipResponse, error) {
	req := requests.NewJumpShipRequest(ship, system)
	var result responses.JumpShipResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func NavigateShip(ship string, waypoint string) (*responses.NavigateShipResponse, error) {
	req := requests.NewNavigateShipRequest(ship, waypoint)
	var result responses.NavigateShipResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func OrbitShip(symbol string) (*models.ShipNav, error) {
	req := requests.NewOrbitShipRequest(symbol)
	var result responses.OrbitShipResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data.Nav, nil
}

func WarpShip(ship string, waypoint string) (*responses.WarpShipResponse, error) {
	req := requests.NewWarpShipRequest(ship, waypoint)
	var result responses.WarpShipResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateChart(symbol string) (*responses.CreateChartResponse, error) {
	req := requests.NewCreateChartRequest(symbol)
	var result responses.CreateChartResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateSurvey(symbol string) (*responses.CreateSurveyResponse, error) {
	req := requests.NewCreateSurveyRequest(symbol)
	var result responses.CreateSurveyResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func Extract(symbol string, survey *models.Survey) (*responses.ExtractResponse, error) {
	req := requests.NewExtractRequest(symbol, survey)
	var result responses.ExtractResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func JettisonCargo(ship string, symbol models.TradeSymbol, units int) (*models.ShipCargo, error) {
	req := requests.NewJettisonCargoRequest(ship, symbol, units)
	var result responses.JettisonCargoResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data.Cargo, nil
}

func NegotiateContract(ship string) (*models.Contract, error) {
	req := requests.NewNegotiateContractRequest(ship)
	var result responses.NegotiateContractResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data.Contract, nil
}

func PatchShipNav(symbol string, mode models.ShipNavFlightMode) (*models.ShipNav, error) {
	req := requests.NewPatchShipNavRequest(symbol, mode)
	var result responses.ShipNavResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func PurchaseCargo(ship string, symbol models.TradeSymbol, units int) (*responses.CargoTransactionResponse, error) {
	req := requests.NewPurchaseCargoRequest(ship, symbol, units)
	var result responses.CargoTransactionResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func PurchaseShip(shipType models.ShipType, waypoint string) (*responses.PurchaseShipResponse, error) {
	req := requests.NewPurchaseShipRequest(shipType, waypoint)
	var result responses.PurchaseShipResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func RefuelShip(ship string) (*responses.RefuelResponse, error) {
	req := requests.NewRefuelRequest(ship)
	var result responses.RefuelResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func ScanShips(ship string) (*responses.ScanShipsResponse, error) {
	req := requests.NewScanShipsRequest(ship)
	var result responses.ScanShipsResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func ScanSystems(ship string) (*responses.ScanSystemsResponse, error) {
	req := requests.NewScanSystemsRequest(ship)
	var result responses.ScanSystemsResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func ScanWaypoints(ship string) (*responses.ScanWaypointsResponse, error) {
	req := requests.NewScanWaypointsRequest(ship)
	var result responses.ScanWaypointsResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func SellCargo(ship string, symbol models.TradeSymbol, units int) (*responses.CargoTransactionResponse, error) {
	req := requests.NewSellCargoRequest(ship, symbol, units)
	var result responses.CargoTransactionResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func ShipRefine(symbol string, produce models.TradeSymbol) (*responses.ShipRefineResponse, error) {
	req := requests.NewShipRefineRequest(symbol, produce)
	var result responses.ShipRefineResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func TransferCargo(src, dest string, symbol models.TradeSymbol, units int) (*responses.WrappedCargoResponse, error) {
	req := requests.NewTransferCargoRequest(src, dest, symbol, units)
	var result responses.WrappedCargoResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
