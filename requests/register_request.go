package requests

import (
	"bytes"
	"encoding/json"
	"io"
)

type RegisterRequest struct {
	Faction string  `json:"faction"`
	Symbol  string  `json:"symbol"`
	Email   *string `json:"email,omitempty"`
}

func NewRegisterRequest(symbol string, faction string, email *string) *RegisterRequest {
	return &RegisterRequest{
		Faction: faction,
		Symbol:  symbol,
		Email:   email,
	}
}

func (rr *RegisterRequest) Method() string {
	return "POST"
}

func (rr *RegisterRequest) Path() string {
	return "/register"
}

func (rr *RegisterRequest) Body() (io.Reader, error) {
	bts, err := json.Marshal(rr)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bts), nil
}

func (rr *RegisterRequest) AuthRequired() bool {
	return false
}
