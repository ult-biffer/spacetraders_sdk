package requests

import (
	"fmt"
	"io"
)

type ScanSystemsRequest struct {
	Symbol string
}

func NewScanSystemsRequest(symbol string) *ScanSystemsRequest {
	return &ScanSystemsRequest{
		Symbol: symbol,
	}
}

func (*ScanSystemsRequest) Method() string {
	return "POST"
}

func (req *ScanSystemsRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/scan/systems", req.Symbol)
}

func (*ScanSystemsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*ScanSystemsRequest) AuthRequired() bool {
	return true
}

func (*ScanSystemsRequest) AuthAllowed() bool {
	return true
}
