package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type ListSystemsRequest struct {
	Page int
}

func NewListSystemsRequest(page int) *ListSystemsRequest {
	return &ListSystemsRequest{
		Page: page,
	}
}

func (*ListSystemsRequest) Method() string {
	return "GET"
}

func (req *ListSystemsRequest) Path() string {
	return fmt.Sprintf("/systems?limit=%d&page=%d", models.MAX_LIMIT, req.Page)
}

func (*ListSystemsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*ListSystemsRequest) AuthRequired() bool {
	return false
}

func (*ListSystemsRequest) AuthAllowed() bool {
	return true
}
