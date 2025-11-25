package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"google.golang.org/grpc"

	"terminal/ui"
	pb "terminal/grpcService"
)

func main() {
	termApp := app.New()
	termWindow := termApp.NewWindow("Zag")
	term := ui.AttachTerminal(termApp)
	termWindow.SetContent(container.NewStack(term))
	termWindow.Resize(fyne.NewSize(900, 600))
	termWindow.ShowAndRun()
}


func RPCClient() {
	var opts []grpc.DialOption
	conn, err := grpc.NewClient("http://localhost:50051", opts...)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	defer conn.Close()
	client := pb.NewNLAgentClient(conn)
	// response, err := client.SendPrompt()
}