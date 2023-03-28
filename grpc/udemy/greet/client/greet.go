package main

import (
	"context"
	"log"

	pb "github.com/udemy/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "kkhdhdhdh",
	})

	if err != nil {
		log.Fatalf("Fail in requesting to server %v\n", err)
	}

	log.Printf("result %s", res.Result)
}
