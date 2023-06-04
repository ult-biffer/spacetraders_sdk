package requests

import (
	"fmt"
	"io"
	"spacetraders_sdk/models"
)

type ListShipsRequest struct {
	Page int
}

func NewListShipsRequest(page int) *ListShipsRequest {
	return &ListShipsRequest{
		Page: page,
	}
}

func (req *ListShipsRequest) Method() string {
	return "GET"
}

func (req *ListShipsRequest) Path() string {
	return fmt.Sprintf("/my/ships?limit=%d&page=%d", models.MAX_LIMIT, req.Page)
}

func (req *ListShipsRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (req *ListShipsRequest) AuthRequired() bool {
	return true
}

func (req *ListShipsRequest) AuthAllowed() bool {
	return true
}
