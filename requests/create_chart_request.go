package requests

import (
	"fmt"
	"io"
)

type CreateChartRequest struct {
	Symbol string
}

func NewCreateChartRequest(symbol string) *CreateChartRequest {
	return &CreateChartRequest{
		Symbol: symbol,
	}
}

func (req *CreateChartRequest) Method() string {
	return "POST"
}

func (req *CreateChartRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/chart", req.Symbol)
}

func (req *CreateChartRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (req *CreateChartRequest) AuthRequired() bool {
	return true
}

func (req *CreateChartRequest) AuthAllowed() bool {
	return true
}
