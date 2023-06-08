package requests

import (
	"fmt"
	"io"
)

type AcceptContractRequest struct {
	Id string
}

func NewAcceptContractRequest(id string) *AcceptContractRequest {
	return &AcceptContractRequest{
		Id: id,
	}
}

func (*AcceptContractRequest) Method() string {
	return "POST"
}

func (req *AcceptContractRequest) Path() string {
	return fmt.Sprintf("/my/contracts/%s/accept", req.Id)
}

func (*AcceptContractRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*AcceptContractRequest) AuthRequired() bool {
	return true
}

func (*AcceptContractRequest) AuthAllowed() bool {
	return true
}
