package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type TransferCargoRequest struct {
	SourceSymbol      string             `json:"-"`
	DestinationSymbol string             `json:"shipSymbol"`
	TradeSymbol       models.TradeSymbol `json:"tradeSymbol"`
	Units             int                `json:"units"`
}

func NewTransferCargoRequest(source string, dest string, item models.TradeSymbol, units int) *TransferCargoRequest {
	return &TransferCargoRequest{
		SourceSymbol:      source,
		DestinationSymbol: dest,
		TradeSymbol:       item,
		Units:             units,
	}
}

func (*TransferCargoRequest) Method() string {
	return "POST"
}

func (req *TransferCargoRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/transfer", req.SourceSymbol)
}

func (req *TransferCargoRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*TransferCargoRequest) AuthRequired() bool {
	return true
}

func (*TransferCargoRequest) AuthAllowed() bool {
	return true
}
