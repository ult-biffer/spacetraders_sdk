package requests

import (
	"fmt"
	"io"
)

type ScanShipsRequest struct {
	Symbol string
}

func NewScanShipsRequest(symbol string) *ScanShipsRequest {
	return &ScanShipsRequest{
		Symbol: symbol,
	}
}

func (*ScanShipsRequest) Method() string {
	return "POST"
}

func (req *ScanShipsRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/scan/ships", req.Symbol)
}

func (*ScanShipsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*ScanShipsRequest) AuthRequired() bool {
	return true
}

func (*ScanShipsRequest) AuthAllowed() bool {
	return true
}
