package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/minhhoangvn/grpc-go/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:8884"
	defaultName = "world"
)

func main(port int, name string) {
	// Set up a connection to the server.
	log.Printf("gRPC Client Start Port: %d", port)
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func Run(port int, name string) {
	main(port, name)
}
