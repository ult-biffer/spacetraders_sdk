package requests

import (
	"fmt"
	"io"
	"net/http"
	"spacetraders_sdk/responses"
)

type GenericHttpError struct {
	Response *responses.ErrorResponse
}

func NewGenericHttpError(resp *http.Response) error {
	response, err := responses.NewErrorResponse(resp)

	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		msg := fmt.Sprintf("error response did not match schema: %s", string(body))
		response = responses.NewFakeErrorResponse(msg)
	}

	return &GenericHttpError{
		Response: response,
	}
}

func (err *GenericHttpError) Error() string {
	return err.Response.Error.Message
}
