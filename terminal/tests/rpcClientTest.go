package tests

import (
	"context"
	"fmt"
	"strconv"
	pb "terminal/grpcService"
	"terminal/utils"
)

func TestRPCClient() {
	conn := utils.RPCConn()
	client := pb.NewNLAgentClient(conn)
	for i := range 3 {
		prompt := "make a new folder named test-" + strconv.Itoa(i)
		cwd := ""
		response := utils.RPCRequest(context.Background(), prompt, cwd, client)
		fmt.Print(response.Command)
	}
}