package main

import (
	"log"
	"net"

	pb "github.com/udemy/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error in listen: %v\n", err)
	}

	log.Printf("Listening to: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error in serving grpc %v/n", err)
	}
}
