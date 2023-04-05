package main

import (
	"log"

	pb "github.com/udemy/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	opts := []grpc.DialOption{}
	tls := false
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("error in creating credentials %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Error in connectiong to addr %v\n", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	doSum(c)
	doGreetWithDeadline(c)
}
