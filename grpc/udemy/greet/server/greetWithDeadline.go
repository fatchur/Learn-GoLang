package main

import (
	"context"
	"log"
	"time"

	pb "github.com/udemy/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet with deal in server called")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, status.Error(codes.Canceled, "Processing time exceed")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello - " + in.FirstName,
	}, nil
}
