package requests

import (
	"fmt"
	"io"
)

type CreateSurveyRequest struct {
	Symbol string
}

func NewCreateSurveyRequest(symbol string) *CreateSurveyRequest {
	return &CreateSurveyRequest{
		Symbol: symbol,
	}
}

func (req *CreateSurveyRequest) Method() string {
	return "POST"
}

func (req *CreateSurveyRequest) Path() string {
	return fmt.Sprintf("/my/ships/%s/survey", req.Symbol)
}

func (req *CreateSurveyRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (req *CreateSurveyRequest) AuthRequired() bool {
	return true
}

func (req *CreateSurveyRequest) AuthAllowed() bool {
	return true
}
