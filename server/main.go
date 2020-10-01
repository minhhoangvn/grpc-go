package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/minhhoangvn/grpc-go/helloworld"

	"google.golang.org/grpc"
)


type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) sayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello Server Reply: " + in.GetName()}, nil
}

func main(port int) {
	log.Printf("gRPC Server Start Port: %d", port)
	lis, err := net.Listen("tcp",  fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	greaterServer := server{}

	pb.RegisterGreeterService(grpcServer, &pb.GreeterService{SayHello: greaterServer.sayHello})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Run(port int){
	main(port)
}