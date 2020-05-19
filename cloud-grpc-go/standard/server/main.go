package main

import (
	"log"
	"net"

	pbf "cloud_grpc_go/pdf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8488"
)

type helloServer struct{}

func (s *helloServer) SayHello(ctx context.Context, in *pbf.HelloRequest) (*pbf.HelloResponse, error) {
	return &pbf.HelloResponse{Message: "Hello " + in.Name+"，I'm Go grpc server"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbf.RegisterHelloServer(s, &helloServer{})

	reflection.Register(s)

	log.Printf("Server started at port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}