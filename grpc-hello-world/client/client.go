package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-hello-world/proto"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../certs/server.pem", "dev")
	if err != nil {
		log.Println("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creds))
	defer conn.Close()
	if err != nil {
		log.Println(err)
	}
	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer : "Grpc",
	}
	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}
	log.Println(r.Message)
}