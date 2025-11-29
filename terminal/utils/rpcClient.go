package utils

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "terminal/grpcService"
)

func RPCClient() pb.NLAgentClient {
	var opts []grpc.DialOption
	conn, err := grpc.NewClient("http://localhost:50051", opts...)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil
	}
	defer conn.Close()
	client := pb.NewNLAgentClient(conn)
	return client
}

func RPCRequest(ctx context.Context,prompt string, cwd string, client pb.NLAgentClient) *pb.LLMResponse {
	request := &pb.PromptRequest{Prompt: prompt, Temperature: 1.0}
	response, err := client.SendPrompt(ctx, request)
	if err != nil {
		log.Printf("Error calling the server: %s", err)
		return nil
	}
	return response
}