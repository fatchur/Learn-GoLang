package main

import (
	"context"
	"log"

	pb "github.com/udemy/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		log.Fatalf("Fail in requesting to server %v\n", err)
	}

	log.Printf("result %v", res.Result)
}
