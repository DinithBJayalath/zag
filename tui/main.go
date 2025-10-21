package main

import (
	"context"
	"embed"
	"fmt"
	"tui/internal/pty"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type App struct {
	ctx context.Context
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	fmt.Println("Wails startup: context captured")
}

func (a *App) domReady(ctx context.Context) {
	fmt.Println("DOM ready: starting PTY session")
	if _, err := pty.StartPty(ctx); err != nil {
		fmt.Println("Failed to start PTY:", err)
	}
}

func main() {
	// Create an instance of the app structure
	app := &App{}
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Zag",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
