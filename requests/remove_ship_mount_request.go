package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type RemoveShipMountRequest struct {
	ShipSymbol string                 `json:"-"`
	Symbol     models.ShipMountSymbol `json:"symbol"`
}

func NewRemoveShipMountRequest(ship string, symbol models.ShipMountSymbol) *RemoveShipMountRequest {
	return &RemoveShipMountRequest{
		ShipSymbol: ship,
		Symbol:     symbol,
	}
}

func (*RemoveShipMountRequest) Method() string {
	return "POST"
}

func (req *RemoveShipMountRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/mounts/remove", req.ShipSymbol)
}

func (req *RemoveShipMountRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*RemoveShipMountRequest) AuthRequired() bool {
	return true
}

func (*RemoveShipMountRequest) AuthAllowed() bool {
	return true
}
