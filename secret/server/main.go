package main

import (
	"../api"
	"../util"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	FileName = "./configs/secret-server.config.json.example"
)

func main() {
	c := util.ServerConfig{}
	if err := util.LoadConfig(&c, FileName); err != nil {
		log.Fatalf("%v", err)
	}
	k := Keys{}
	if err := k.LoadKeys(c); err != nil {
		log.Fatalf("%v", err)
	}
	cipher, err := util.EncryptOAEP([]byte("send reinforcements, we're going to advance"), k.ServerPublicKey)
	if err != nil {
		log.Fatalf("%v", err)
	}
	text, err := util.DecryptOAEP(cipher, k.ServerPrivateKey)
	list, err := net.Listen("tcp", c.ServerListen)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("QQQQQ: %s\n", text)
	s := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterSecretServer(grpcServer, &s)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("%v", err)
	}
}
