package main

import (
	"context"
	"log"

	pb "github.com/udemy/calculator/proto"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient) {
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: -7})
	if err != nil {
		err, ok := status.FromError(err)

		if ok {
			log.Printf("Error status code: %v\n", err.Code())
			log.Printf("error message: %v\n", err.Message())
			log.Printf("%v", err.Details())
		} else {
			log.Printf("Error in parsing error from grpc return")
		}

		return
	}

	log.Printf("result: %v\n", res.Result)
	return
}
