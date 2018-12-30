package main

import (
	"../api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("%v", err)
	}
	s := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterSecretServer(grpcServer, &s)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("%v", err)
	}
}
