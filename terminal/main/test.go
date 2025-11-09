package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/fyne-io/terminal"
)

func main() {
	termApp := app.New()
	termWindow := termApp.NewWindow("Zag")
	term := terminal.New()

	go func() {
		_ = term.RunLocalShell()
		fmt.Printf("Shell exited: %d", term.ExitCode())
		termApp.Quit()
	}()

	termWindow.SetContent(container.NewStack(term))
	termWindow.Resize(fyne.NewSize(900, 600))
	termWindow.ShowAndRun()
}
