package requests

import (
	"fmt"
	"io"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type ListContractsRequest struct {
	Page int
}

func NewListContractsRequest(page int) *ListContractsRequest {
	return &ListContractsRequest{
		Page: page,
	}
}

func (*ListContractsRequest) Method() string {
	return "GET"
}

func (req *ListContractsRequest) Path() string {
	return fmt.Sprintf("/my/contracts?limit=%d&page=%d", models.MAX_LIMIT, req.Page)
}

func (*ListContractsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*ListContractsRequest) AuthRequired() bool {
	return true
}

func (*ListContractsRequest) AuthAllowed() bool {
	return true
}
