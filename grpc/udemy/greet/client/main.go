package main

import (
	"log"

	pb "github.com/udemy/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error in connectiong to addr %v\n", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	doSum(c)
	doGreetWithDeadline(c)
}
