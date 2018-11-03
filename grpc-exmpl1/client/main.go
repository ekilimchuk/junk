package main

import (
	"log"

	"../api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
//	var conn *grps.ClientConn
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()
	c := api.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "ping"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Response: %s", response.Greeting)
}
