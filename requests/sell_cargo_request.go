package requests

import (
	"fmt"
	"io"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type SellCargoRequest struct {
	ShipSymbol string             `json:"-"`
	Symbol     models.TradeSymbol `json:"symbol"`
	Units      int                `json:"units"`
}

func NewSellCargoRequest(ship string, symbol models.TradeSymbol, units int) *SellCargoRequest {
	return &SellCargoRequest{
		ShipSymbol: ship,
		Symbol:     symbol,
		Units:      units,
	}
}

func (*SellCargoRequest) Method() string {
	return "POST"
}

func (req *SellCargoRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/sell", req.ShipSymbol)
}

func (req *SellCargoRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*SellCargoRequest) AuthRequired() bool {
	return true
}

func (*SellCargoRequest) AuthAllowed() bool {
	return true
}
