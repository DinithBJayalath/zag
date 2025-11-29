package utils

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "terminal/grpcService"
)

func RPCClient(ctx context.Context,prompt string, cwd string) *pb.LLMResponse {
	var opts []grpc.DialOption
	conn, grpcErr := grpc.NewClient("http://localhost:50051", opts...)
	if grpcErr != nil {
		log.Printf("Error: %s", grpcErr.Error())
		return nil
	}
	defer conn.Close()
	client := pb.NewNLAgentClient(conn)
	request := &pb.PromptRequest{Prompt: prompt, Temperature: 1.0}
	response, clientErr := client.SendPrompt(ctx, request)
	if clientErr != nil {
		log.Printf("Error calling the server: %s", clientErr)
		return nil
	}
	return response
}
