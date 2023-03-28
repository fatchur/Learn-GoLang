package main

import (
	"context"
	"fmt"
	"math"

	pb "github.com/udemy/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	var result float32
	if in.Number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Input argument is negative number"),
		)
	}
	result = float32(math.Sqrt(float64(in.Number)))
	return &pb.SqrtResponse{Result: result}, nil
}
