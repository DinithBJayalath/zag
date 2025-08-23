package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	path  string
	input string
}

func InitialModel() Model {
	path := "~"
	return Model{
		path: path,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
