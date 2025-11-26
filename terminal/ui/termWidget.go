package ui

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"fyne.io/fyne/v2"
	"github.com/creack/pty"
	"github.com/fyne-io/terminal"
)

func AttachTerminal(termApp fyne.App) fyne.CanvasObject {
	//Setting up the initial path
	cmd := exec.Command("bin/zsh")
	tmp, _ := os.MkdirTemp("", "zag-zdotdir-")
	_ = os.WriteFile(filepath.Join(tmp, ".zshrc"), []byte(`
	setopt PROMPT_SUBST
	PROMPT='%F{#089DDA}%2~ ❯%f '
	RPROMPT=''
	`), 0600)
	os.Setenv("ZDOTDIR", tmp)
	// Setting up the PTY
	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
	defer ptmx.Close()
	// Setting up the terminal widget
	term := terminal.New()
	go func() {
		_ = term.RunLocalShell()
		log.Printf("Shell exited: %d", term.ExitCode())
		termApp.Quit()
	}()
	return term
}