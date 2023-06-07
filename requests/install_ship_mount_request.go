package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type InstallShipMountRequest struct {
	ShipSymbol string                 `json:"-"`
	Symbol     models.ShipMountSymbol `json:"symbol"`
}

func NewInstallShipMountRequest(ship string, symbol models.ShipMountSymbol) *InstallShipMountRequest {
	return &InstallShipMountRequest{
		ShipSymbol: ship,
		Symbol:     symbol,
	}
}

func (*InstallShipMountRequest) Method() string {
	return "POST"
}

func (req *InstallShipMountRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/mounts/install", req.ShipSymbol)
}

func (req *InstallShipMountRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*InstallShipMountRequest) AuthRequired() bool {
	return true
}

func (*InstallShipMountRequest) AuthAllowed() bool {
	return true
}
