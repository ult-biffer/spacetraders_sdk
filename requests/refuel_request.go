package requests

import (
	"fmt"
	"io"
)

type RefuelRequest struct {
	Symbol string
}

func NewRefuelRequest(symbol string) *RefuelRequest {
	return &RefuelRequest{
		Symbol: symbol,
	}
}

func (*RefuelRequest) Method() string {
	return "POST"
}

func (req *RefuelRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/refuel", req.Symbol)
}

func (*RefuelRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*RefuelRequest) AuthRequired() bool {
	return true
}

func (*RefuelRequest) AuthAllowed() bool {
	return true
}
