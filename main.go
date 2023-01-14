package main

import (
	"context"
	"fmt"
	hservice "grpctest/hservice"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type Tile38Server struct {
	hservice.HookServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hservice.RegisterHookServiceServer(s, &Tile38Server{})
	fmt.Println("Server started")
	s.Serve(lis)
}

func (s *Tile38Server) Send(ctx context.Context, in *hservice.MessageRequest) (*hservice.MessageReply, error) {
	fmt.Println("Received: ", in)
	return &hservice.MessageReply{Ok:true}, nil
}