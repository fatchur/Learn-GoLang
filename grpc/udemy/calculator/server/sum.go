package main

import (
	"context"

	pb "github.com/udemy/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	result := in.FirstNumber + in.SecondNumber
	return &pb.SumResponse{
		Result: result,
	}, nil
}
