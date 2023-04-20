package main

import (
	"context"
	pb "github.com/sudipto-003/grpc-exp/expunary/proto"

)

func (s *Server) Add(ctx context.Context, in *pb.CalcRequest) (*pb.ASMResponse, error) {
	res := in.Var1 + in.Var2
	return &pb.ASMResponse{
		Val: res,
	}, nil
}

func (s *Server) Sub(ctx context.Context, in *pb.CalcRequest) (*pb.ASMResponse, error) {
	res := in.Var1 - in.Var2
	return &pb.ASMResponse{
		Val: res,
	}, nil
}

func (s *Server) Mul(ctx context.Context, in *pb.CalcRequest) (*pb.ASMResponse, error) {
	res := in.Var1 * in.Var2
	return &pb.ASMResponse{
		Val: res,
	}, nil
}

func (s *Server) Div(ctx context.Context, in *pb.CalcRequest) (*pb.DivResponse, error) {
	if in.Var2 == 0 {
		return &pb.DivResponse{
			Val1: 0,
			Val2: 0,
		}, nil
	}
	div := in.Var1 / in.Var2
	mod := in.Var1 % in.Var2
	return &pb.DivResponse{
		Val1: div,
		Val2: mod,
	}, nil
}