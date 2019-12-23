package main

import (
	"context"
	"log"
	"net"

	pb "github.com/alexmspina/portapp/internal/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	creds, _ := credentials.NewServerTLSFromFile("./test/certs/localhost.cer.pem", "./test/certs/localhost.key.pem")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
