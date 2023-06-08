package requests

import (
	"fmt"
	"io"
)

type FulfillContractRequest struct {
	Id string
}

func NewFulfillContractRequest(id string) *FulfillContractRequest {
	return &FulfillContractRequest{
		Id: id,
	}
}

func (*FulfillContractRequest) Method() string {
	return "POST"
}

func (req *FulfillContractRequest) Path() string {
	return fmt.Sprintf("/my/contracts/%s/fulfill", req.Id)
}

func (*FulfillContractRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*FulfillContractRequest) AuthRequired() bool {
	return true
}

func (*FulfillContractRequest) AuthAllowed() bool {
	return true
}
