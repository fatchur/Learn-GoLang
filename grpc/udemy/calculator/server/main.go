package main

import (
	"log"
	"net"

	pb "github.com/udemy/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error in listenign tcp addr %v\n", err)
	}

	log.Printf("Success in listening tcp addres %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error in serving grpc server %v\n", err)
	}

}
