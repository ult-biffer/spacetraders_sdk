package requests

import (
	"fmt"
	"io"

	"github.com/ult-biffer/spacetraders_sdk/models"
)

type ListFactionsRequest struct {
	Page int
}

func NewListFactionsRequest(page int) *ListFactionsRequest {
	return &ListFactionsRequest{
		Page: page,
	}
}

func (*ListFactionsRequest) Method() string {
	return "GET"
}

func (req *ListFactionsRequest) Path() string {
	return fmt.Sprintf("/factions?limit=%d&page=%d", models.MAX_LIMIT, req.Page)
}

func (*ListFactionsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*ListFactionsRequest) AuthRequired() bool {
	return false
}

func (*ListFactionsRequest) AuthAllowed() bool {
	return true
}
