package pty

import (
	"context"
	"os"
	"os/exec"
	"fmt"
	"github.com/creack/pty"
)

type PTYSession struct {
	cmd *exec.Cmd
	ptmx *os.File
	ctx context.Context
}

func startPty(ctx context.Context) (*PTYSession, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "bin/bash"
	}
	cmd := exec.Command(shell)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, fmt.Errorf("Faild to start pty: %s", err)
	}
	session := &PTYSession{cmd: cmd, ptmx: ptmx, ctx: ctx}
	//TODO: add functions for I/O
	return session,nil
}