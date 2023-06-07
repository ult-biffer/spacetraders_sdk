package requests

import (
	"fmt"
	"io"
)

type ScanWaypointsRequest struct {
	Symbol string
}

func NewScanWaypointsRequest(symbol string) *ScanWaypointsRequest {
	return &ScanWaypointsRequest{
		Symbol: symbol,
	}
}

func (*ScanWaypointsRequest) Method() string {
	return "POST"
}

func (req *ScanWaypointsRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/scan/waypoints", req.Symbol)
}

func (*ScanWaypointsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*ScanWaypointsRequest) AuthRequired() bool {
	return true
}

func (*ScanWaypointsRequest) AuthAllowed() bool {
	return true
}
