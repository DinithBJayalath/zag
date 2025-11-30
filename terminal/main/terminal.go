package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"terminal/ui"
)

func main() {
	termApp := app.New()
	termWindow := termApp.NewWindow("Zag")
	term, ptmx, conn := ui.AttachTerminal(termApp)
	termWindow.SetContent(container.NewStack(term))
	termWindow.Resize(fyne.NewSize(900, 600))
	defer ptmx.Close()
	defer conn.Close()
	termWindow.ShowAndRun()
}
