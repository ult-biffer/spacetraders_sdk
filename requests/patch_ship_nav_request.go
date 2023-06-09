package requests

import (
	"fmt"
	"io"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type PatchShipNavRequest struct {
	Symbol     string                   `json:"-"`
	FlightMode models.ShipNavFlightMode `json:"flightMode"`
}

func NewPatchShipNavRequest(symbol string, mode models.ShipNavFlightMode) *PatchShipNavRequest {
	return &PatchShipNavRequest{
		Symbol:     symbol,
		FlightMode: mode,
	}
}

func (*PatchShipNavRequest) Method() string {
	return "PATCH"
}

func (req *PatchShipNavRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/nav", req.Symbol)
}

func (req *PatchShipNavRequest) Body() (io.Reader, error) {
	return marshal(req)
}

func (*PatchShipNavRequest) AuthRequired() bool {
	return true
}

func (*PatchShipNavRequest) AuthAllowed() bool {
	return true
}
