package requests

import (
	"fmt"
	"io"
)

type GetContractRequest struct {
	Id string
}

func NewGetContractRequest(id string) *GetContractRequest {
	return &GetContractRequest{
		Id: id,
	}
}

func (*GetContractRequest) Method() string {
	return "GET"
}

func (req *GetContractRequest) Path() string {
	return fmt.Sprintf("/my/contracts/%s", req.Id)
}

func (*GetContractRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetContractRequest) AuthRequired() bool {
	return true
}

func (*GetContractRequest) AuthAllowed() bool {
	return true
}
