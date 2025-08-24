package ui

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/creack/pty"
)

type ptyOutMsg struct{ data []byte }
type initPTYMsg struct{ f *os.File }
type errMsg struct{ err error }

func startShell() tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("/bin/zsh", "-l")
		ptmx, err := pty.Start(cmd)
		if err != nil {
			return errMsg{err: err}
		}
		return initPTYMsg{f: ptmx}
	}
}

func readPTY(m Model) tea.Cmd {
	return func() tea.Msg {
		buf := make([]byte, 1024)
		n, err := m.ptmx.Read(buf)
		if err != nil {
			return errMsg{err: err}
		}
		return ptyOutMsg{data: buf[:n]}
	}
}

func writePTY(buf []byte, m Model) tea.Cmd {
	return func() tea.Msg {
		_, err := m.ptmx.Write(buf)
		if err != nil {
			return errMsg{err: err}
		}
		return nil
	}
}
