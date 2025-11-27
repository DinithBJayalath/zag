package utils

import (
	"io"
	"strings"

	"github.com/fyne-io/terminal"

	pb "terminal/grpcService"
)

type Router struct {
	dst io.WriteCloser
	term *terminal.Terminal
	prefix string = "nl"
	NLRequest func(prompt, cwd string) (pb.LLMResponse, error)
	buf []rune
}

func NewRouter(dst io.WriteCloser, term *terminal.Terminal, NLRequest func(string, string) (pb.LLMResponse, error)) *Router {
	return &Router{dst: dst, term: term, prefix: "nl", NLRequest: NLRequest}
}

func (r *Router) Write(text []byte) (int, error) {
	for _, b := range text {
		switch b {
		case '\r', '\n':
			line := string(r.buf)
			r.buf = r.buf[:0]
			if r.IsNL(line) {
				//Create the function to handle the NL request first
			}
		}
	}
}

func (r *Router) IsNL(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, r.prefix)
}