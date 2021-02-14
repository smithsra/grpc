package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smithsra/grpc/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello im a calculator client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connet: %v", err)
	}

	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c calcpb.CalcServiceClient) {
	fmt.Println("Starting to do a sum Unary rpc")
	req := &calcpb.SumRequest{
		Input: &calcpb.NumsToAdd{
			FirstNum:  10,
			SecondNum: 3,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling sum rpc: %v", err)
	}
	log.Printf("response from Sum: %v", res.Total)
}
