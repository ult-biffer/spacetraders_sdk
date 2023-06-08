package requests

import "io"

type GetMyAgentRequest struct{}

func (*GetMyAgentRequest) Method() string {
	return "GET"
}

func (*GetMyAgentRequest) Path() string {
	return "/my/agent"
}

func (*GetMyAgentRequest) Body() (io.Reader, error) {
	return nil, nil
}

func (*GetMyAgentRequest) AuthRequired() bool {
	return true
}

func (*GetMyAgentRequest) AuthAllowed() bool {
	return true
}
