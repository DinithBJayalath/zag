package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.vp.Width = msg.Width
		m.vp.Height = msg.Height - 3
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			line := m.input.Value() + "\n"
			m.input.SetValue("")
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
		m.vp.GotoBottom()
		return m, readPTY(m)
	case errMsg:
		m.err = msg.err
	}
	return m, nil
}
