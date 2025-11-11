package ui

import (
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"github.com/fyne-io/terminal"
)

func AttachTerminal(termApp fyne.App) fyne.CanvasObject {
	//Setting up the initial path
	os.Setenv("SHELL", "/bin/zsh")
	tmp, _ := os.MkdirTemp("", "zag-zdotdir-")
	_ = os.WriteFile(filepath.Join(tmp, ".zshrc"), []byte(`
	setopt PROMPT_SUBST
	PROMPT='%F{#089DDA}%2~ ❯%f '
	RPROMPT=''
	`), 0600)
	os.Setenv("ZDOTDIR", tmp)
	// Setting up the terminal widget
	term := terminal.New()
	go func() {
		_ = term.RunLocalShell()
		log.Printf("Shell exited: %d", term.ExitCode())
		termApp.Quit()
	}()
	return term
}
