package main

import (
	"log"
	"net"

	"../api"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("%v", err)
	}
	s := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterPingServer(grpcServer, &s)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("%v", err)
	}
}
