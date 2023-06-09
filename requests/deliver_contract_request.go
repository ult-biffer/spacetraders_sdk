package requests

import (
	"fmt"
	"io"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type DeliverContractRequest struct {
	Id          string             `json:"-"`
	ShipSymbol  string             `json:"shipSymbol"`
	TradeSymbol models.TradeSymbol `json:"tradeSymbol"`
	Units       int                `json:"units"`
}

func NewDeliverContractRequest(id, ship string, item models.TradeSymbol, units int) *DeliverContractRequest {
	return &DeliverContractRequest{
		Id:          id,
		ShipSymbol:  ship,
		TradeSymbol: item,
		Units:       units,
	}
}

func (*DeliverContractRequest) Method() string {
	return "POST"
}

func (req *DeliverContractRequest) Path() string {
	return fmt.Sprintf("/my/contracts/%s/deliver", req.Id)
}

func (req *DeliverContractRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*DeliverContractRequest) AuthRequired() bool {
	return true
}

func (*DeliverContractRequest) AuthAllowed() bool {
	return true
}
