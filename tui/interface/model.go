package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
)

type Model struct {
	input       textinput.Model
	prompt      string
	command     string
	explanation string
	stage       string // "input", "thinking", "confirmation", "executing"
	errMsg      string
}

func InitialModel() Model {
	input := textinput.New()
	input.Placeholder = "ZAG> "
	input.Focus()
	return Model{
		input: input,
		stage: "input",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
