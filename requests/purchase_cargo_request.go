package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type PurchaseCargoRequest struct {
	ShipSymbol string             `json:"-"`
	Symbol     models.TradeSymbol `json:"symbol"`
	Units      int                `json:"units"`
}

func NewPurchaseCargoRequest(ship string, symbol models.TradeSymbol, units int) *PurchaseCargoRequest {
	return &PurchaseCargoRequest{
		ShipSymbol: ship,
		Symbol:     symbol,
		Units:      units,
	}
}

func (*PurchaseCargoRequest) Method() string {
	return "POST"
}

func (req *PurchaseCargoRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/purchase", req.ShipSymbol)
}

func (req *PurchaseCargoRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*PurchaseCargoRequest) AuthRequired() bool {
	return true
}

func (*PurchaseCargoRequest) AuthAllowed() bool {
	return true
}
