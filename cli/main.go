package main

import (
	"context"
	"log"
	"time"

	hservice "cli/hservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := hservice.NewHookServiceClient(conn)

	callSend(client)
}

func callSend(client hservice.HookServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Send(ctx, &hservice.MessageRequest{Value: "Hello"}) 
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Ok)
}
