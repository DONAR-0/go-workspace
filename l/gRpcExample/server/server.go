package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/donar-0/go-workspace/l/gRpcExample/helloworld"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context,
	in *helloworld.HelloRequest) (*helloworld.HelloReply,
	error) {
	log.Printf("Received: %v", in.GetName())

	return &helloworld.HelloReply{
		Message: "Hello" + in.GetName(),
	}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: "Hello Again" + in.GetName(),
	}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
