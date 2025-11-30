package utils

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "terminal/grpcService"
)

func RPCConn() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil
	}
	return conn
}

func RPCRequest(ctx context.Context, prompt string, cwd string, client pb.NLAgentClient) *pb.LLMResponse {
	request := &pb.PromptRequest{Prompt: prompt, Temperature: 1.0}
	response, err := client.SendPrompt(ctx, request)
	if err != nil {
		log.Printf("Error calling the server: %s", err)
		return nil
	}
	return response
}
