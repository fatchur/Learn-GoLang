package main

import (
	"context"
	"log"
	"time"

	pb "github.com/udemy/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "fatchur",
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Processing time exceed time limit, cancelled")
			} else {
				log.Fatalf("Unexpected grpc error: %v\n", e.Message())
			}
		} else {
			log.Fatalf("Non grpc error: %v\n", err)
		}
		return
	}

	log.Printf("success %s", res.Result)
}
