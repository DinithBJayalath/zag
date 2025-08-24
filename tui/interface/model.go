package ui

import (
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	path  string
	vp    viewport.Model
	input string
	ptmx  *os.File
}

func InitialModel() Model {
	path := "~"
	return Model{
		path: path,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(startShell(), readPTY(m))
}
