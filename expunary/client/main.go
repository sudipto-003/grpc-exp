package main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/sudipto-003/grpc-exp/expunary/proto"
)

var addr string = "0.0.0.0:5000"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewASMDServiceClient(conn)

	callAdd(c)
}