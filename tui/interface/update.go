package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			line := m.input.Value() + "\n"
			return m, writePTY([]byte(line), m) //TODO: Send input to the server
		}
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	case initPTYMsg:
		m.ptmx = msg.f
		return m, readPTY(m)
	case ptyOutMsg:
		m.content += string(msg.data)
		m.vp.SetContent(m.content)
		return m, readPTY(m)
	case errMsg:
		m.err = msg.err
	}
	return m, nil
}
