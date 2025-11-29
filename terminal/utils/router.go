package utils

import (
	"context"
	"io"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/fyne-io/terminal"

	pb "terminal/grpcService"
)

type Router struct {
	dst    io.WriteCloser
	term   *terminal.Terminal
	prefix string
	buf    []rune
	client pb.NLAgentClient
}

func NewRouter(dst io.WriteCloser, term *terminal.Terminal, client pb.NLAgentClient) *Router {
	return &Router{dst: dst, term: term, prefix: "nl", client: client}
}

func (r *Router) Write(text []byte) (int, error) {
	for _, b := range text {
		switch b {
		case '\r', '\n':
			line := string(r.buf)
			r.buf = r.buf[:0]
			if r.IsNL(line) {
				response := r.SendPrompt(line)
				fyne.Do(func(){
					go func() {
						_, _ = r.dst.Write([]byte(response.Response))
                		_, _ = r.dst.Write([]byte{'\n'})
					}()
				})
				continue
			}
			if _, err := r.dst.Write([]byte{b}); err != nil { return 0, err }
		case 0x7f:
			if len(r.buf) > 0 { r.buf = r.buf[:len(r.buf)-1]}
			if _, err := r.dst.Write([]byte{b}); err != nil { return 0, err }
		default:
			r.buf = append(r.buf, rune(b))
			if _, err := r.dst.Write([]byte{b}); err != nil { return 0, err }
		}
	}
	return len(text), nil
}

func (r *Router) Close() error { return r.dst.Close() }

func (r *Router) IsNL(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, r.prefix)
}

func (r *Router) SendPrompt(line string) *pb.LLMResponse {
	var response *pb.LLMResponse
	go func() {
		prompt := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), r.prefix))
		cwd := ""
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		response = RPCRequest(ctx, prompt, cwd, r.client)
	}()
	return response
}
