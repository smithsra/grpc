package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/smithsra/grpc/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calc(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstNum := req.GetInput().GetFirstNum()
	secondNum := req.GetInput().GetSecondNum()
	total := firstNum + secondNum
	res := &calcpb.SumResponse{
		Total: total,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
