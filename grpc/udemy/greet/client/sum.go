package main

import (
	"context"
	"log"

	pb "github.com/udemy/greet/proto"
)

func doSum(c pb.GreetServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 5,
	})

	if err != nil {
		log.Fatalf("Fail in requesting to server %v\n", err)
	}

	log.Printf("result %v", res.Result)
}
