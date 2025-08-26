package ui

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	// TODO: implement view logic
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.vp.View(),
		m.path+"\n"+m.input.View(),
	)
}
