package ui

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	path    string //TODO: use this to show current path in prompt
	vp      viewport.Model
	input   textinput.Model
	ptmx    *os.File
	content string
	err     error
}

func InitialModel() Model {
	initialInput := textinput.New()
    initialInput.Focus()
    initialInput.CharLimit = 256
    initialInput.Width = 50
	vp := viewport.New(80, 20)
	path := "~"
	return Model{
		path:   path,
		vp:     vp,
		input:  initialInput,
	}
}

func (m Model) Init() tea.Cmd {
	return startShell()
}
