package requests

import (
	"fmt"
	"io"
)

type GetFactionRequest struct {
	Symbol string
}

func NewGetFactionRequest(symbol string) *GetFactionRequest {
	return &GetFactionRequest{
		Symbol: symbol,
	}
}

func (*GetFactionRequest) Method() string {
	return "GET"
}

func (req *GetFactionRequest) Path() string {
	return fmt.Sprintf("/factions/%s", req.Symbol)
}

func (*GetFactionRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetFactionRequest) AuthRequired() bool {
	return false
}

func (*GetFactionRequest) AuthAllowed() bool {
	return true
}
