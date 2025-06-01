package main

import (
	"log"
	"zag/interface"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
