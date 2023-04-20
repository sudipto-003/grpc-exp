package main

import (
	"log"
	"net"

	pb "github.com/sudipto-003/grpc-exp/expunary/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:5000"

type Server struct {
	pb.ASMDServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterASMDServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Panicf("Failed to serve %v\n", err)
	}

}