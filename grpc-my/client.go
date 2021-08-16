package main

import (
	"context"
	"grpc-my/chat"
	"log"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connection: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from Client!"})
	if err != nil {
		log.Fatalf("error while calling Say Hello: %s", err)
	}
	log.Printf("response from server: %s", response.Body)
}
