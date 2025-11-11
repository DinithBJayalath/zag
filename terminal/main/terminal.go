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
	term := ui.AttachTerminal(termApp)
	termWindow.SetContent(container.NewStack(term))
	termWindow.Resize(fyne.NewSize(900, 600))
	termWindow.ShowAndRun()
}
