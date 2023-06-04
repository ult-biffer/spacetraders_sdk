package requests

import (
	"io"
	"net/http"
)

type GenericHttpError struct {
	Body []byte `json:"body"`
}

func NewGenericHttpError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		body = []byte{}
	}

	return &GenericHttpError{
		Body: body,
	}
}

func (err *GenericHttpError) Error() string {
	return string(err.Body)
}
