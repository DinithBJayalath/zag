package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	termApp := app.New()
	termWindow := termApp.NewWindow("Hello World")
	termWindow.SetContent(widget.NewLabel("Hello World"))
	termWindow.ShowAndRun()
}
