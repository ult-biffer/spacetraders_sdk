package requests

import (
	"fmt"
	"io"
)

type GetSystemRequest struct {
	SystemSymbol string
}

func NewGetSystemRequest(system string) *GetSystemRequest {
	return &GetSystemRequest{
		SystemSymbol: system,
	}
}

func (*GetSystemRequest) Method() string {
	return "GET"
}

func (req *GetSystemRequest) Path() string {
	return fmt.Sprintf("/systems/%s", req.SystemSymbol)
}

func (*GetSystemRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetSystemRequest) AuthRequired() bool {
	return false
}

func (*GetSystemRequest) AuthAllowed() bool {
	return true
}
