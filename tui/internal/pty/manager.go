package pty

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type PTYSession struct {
	cmd  *exec.Cmd
	ptmx *os.File
	ctx  context.Context
}

func startPty(ctx context.Context) (*PTYSession, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "bin/bash"
	}
	cmd := exec.Command(shell)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, fmt.Errorf("faild to start pty: %s", err)
	}
	session := &PTYSession{cmd: cmd, ptmx: ptmx, ctx: ctx}
	go session.ReadOutput()
	session.ListenForInput()
	return session, nil
}

func (s *PTYSession) ReadOutput() {
	reader := bufio.NewReader(s.ptmx)
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil {
			break
		}
		runtime.EventsEmit(s.ctx, "pty-output", buf[:n])
	}
}

func (s *PTYSession) ListenForInput() {
	runtime.EventsOn(s.ctx, "pty-input", func(optionalData ...interface{}) {
		if len(optionalData) > 0 {
			if data, ok := optionalData[0].(string); ok {
				s.ptmx.Write([]byte(data))
			}
		}
	})
}

func (s *PTYSession) Close() {
	s.ptmx.Close()
	s.cmd.Process.Kill()
}
