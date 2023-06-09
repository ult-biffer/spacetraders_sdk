package requests

import (
	"fmt"
	"io"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type JettisonCargoRequest struct {
	ShipSymbol string             `json:"-"`
	Symbol     models.TradeSymbol `json:"symbol"`
	Units      int                `json:"units"`
}

func NewJettisonCargoRequest(ship string, symbol models.TradeSymbol, units int) *JettisonCargoRequest {
	return &JettisonCargoRequest{
		ShipSymbol: ship,
		Symbol:     symbol,
		Units:      units,
	}
}

func (req *JettisonCargoRequest) Method() string {
	return "GET"
}

func (req *JettisonCargoRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/jettison", req.ShipSymbol)
}

func (req *JettisonCargoRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (req *JettisonCargoRequest) AuthRequired() bool {
	return true
}

func (req *JettisonCargoRequest) AuthAllowed() bool {
	return true
}
