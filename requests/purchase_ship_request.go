package requests

import (
	"io"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type PurchaseShipRequest struct {
	ShipType       models.ShipType `json:"shipType"`
	WaypointSymbol string          `json:"waypointSymbol"`
}

func NewPurchaseShipRequest(st models.ShipType, wp string) *PurchaseShipRequest {
	return &PurchaseShipRequest{
		ShipType:       st,
		WaypointSymbol: wp,
	}
}

func (req *PurchaseShipRequest) Method() string {
	return "POST"
}

func (req *PurchaseShipRequest) Path() string {
	return "/my/ships"
}

func (req *PurchaseShipRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (req *PurchaseShipRequest) AuthRequired() bool {
	return true
}

func (req *PurchaseShipRequest) AuthAllowed() bool {
	return true
}
