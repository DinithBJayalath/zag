package utils

import (
	"io"

	"github.com/fyne-io/terminal"

	pb "terminal/grpcService"
)

type Router struct {
	dst io.WriteCloser
	term *terminal.Terminal
	prefix string
	NLRequest func(prompt, cwd string) (LLMResponse pb.LLMResponse, err error)
}