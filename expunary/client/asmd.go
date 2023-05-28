package main

import (
	"context"
	"log"

	pb "github.com/sudipto-003/grpc-exp/expunary/proto"
)

func callAdd(c pb.ASMDServiceClient) {
	req := &pb.CalcRequest{
		Var1: 10,
		Var2: 5,
	}
	add, err := c.Add(context.Background(), req)

	if err != nil {
		log.Panicf("Could not call Add: %v\n", err)
	}

	log.Printf("Add: %d\n", add.Val)
}

func callSub(c pb.ASMDServiceClient) {
	req := &pb.CalcRequest{
		Var1: 10,
		Var2: 5,
	}

	sub, err := c.Sub(context.Background(), req)
	if err != nil {
		log.Panicf("Could not call Sub: %v\n", err)
	}
	log.Printf("Sub: %d\n", sub.Val)
}

func callMul(c pb.ASMDServiceClient) {
	req := &pb.CalcRequest{
		Var1: 10,
		Var2: 5,
	}
	mul, err := c.Mul(context.Background(), req)

	if err != nil {
		log.Panicf("Could not call Mul: %v\n", err)
	}
	log.Printf("Mul: %d\n", mul.Val)
}

func callDiv(c pb.ASMDServiceClient) {
	req := &pb.CalcRequest{
		Var1: 10,
		Var2: 5,
	}

	div, err := c.Div(context.Background(), req)
	if err != nil {
		log.Panicf("Could not call Div: %v\n", err)
	}
	log.Printf("Divident: %d, Reminder: %d\n", div.Val1, div.Val2)
}