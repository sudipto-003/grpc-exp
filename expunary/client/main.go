package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/sudipto-003/grpc-exp/expunary/proto"
)

var grpc_addr string = "35.247.2.220:80"

func main() {
	conn, err := grpc.Dial(grpc_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewASMDServiceClient(conn)

	callAdd(c)
	callSub(c)
	callMul(c)
	callDiv(c)
}