package utils

import (
	"context"
	"io"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/fyne-io/terminal"

	pb "terminal/grpcService"
)

type Router struct {
	dst    io.WriteCloser
	term   *terminal.Terminal
	prefix string
	buf    []rune
}

type ResultCommand struct {
	Command     string
	Explanation string
	IsDangerous bool
}

func NewRouter(dst io.WriteCloser, term *terminal.Terminal) *Router {
	return &Router{dst: dst, term: term, prefix: "nl"}
}

func (r *Router) Write(text []byte) (int, error) {
	for _, b := range text {
		switch b {
		case '\r', '\n':
			line := string(r.buf)
			r.buf = r.buf[:0]
			if r.IsNL(line) {
				ch := r.SendPrompt(context.Background(), line)
				fyne.Do(func() {
					go func() {
						response := <-ch
						result := &ResultCommand{Command: response.Command, Explanation: response.Explanation, IsDangerous: response.IsDangerous}
						_, _ = r.dst.Write([]byte(result.Command))
						_, _ = r.dst.Write([]byte{'\n'})
					}()
				})
				continue
			}
			if _, err := r.dst.Write([]byte{b}); err != nil {
				return 0, err
			}
		case 0x7f:
			if len(r.buf) > 0 {
				r.buf = r.buf[:len(r.buf)-1]
			}
			if _, err := r.dst.Write([]byte{b}); err != nil {
				return 0, err
			}
		default:
			r.buf = append(r.buf, rune(b))
			if _, err := r.dst.Write([]byte{b}); err != nil {
				return 0, err
			}
		}
	}
	return len(text), nil
}

func (r *Router) Close() error { return r.dst.Close() }

func (r *Router) IsNL(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, r.prefix)
}

func (r *Router) SendPrompt(ctx context.Context, line string) <-chan *pb.LLMResponse {
	ch := make(chan *pb.LLMResponse)
	conn := RPCConn()
	client := pb.NewNLAgentClient(conn)
	defer conn.Close()
	go func() {
		defer close(ch)
		prompt := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), r.prefix))
		cwd := ""
		c, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		ch <- RPCRequest(c, prompt, cwd, client)
	}()
	return ch
}
