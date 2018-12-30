package main

import (
	"fmt"
	"../api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ClientConn struct {
	*grpc.ClientConn
}

func NewClient(addrs string, port string) (*ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", addrs, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &ClientConn{conn}, nil
}

func (s *ClientConn) List(path string) (*api.ListResult, error) {
    c := api.NewSecretClient(s.ClientConn)
    return c.List(context.Background(), &api.ListMessage{Path: path})
}

func (s *ClientConn) Add(path string) (*api.AddMessage, error) {
    c := api.NewSecretClient(s.ClientConn)
    return c.Add(context.Background(), &api.AddMessage{Aeskey: path})
}

func (s *ClientConn) Remove(path string) (*api.RemoveMessage, error) {
    c := api.NewSecretClient(s.ClientConn)
    return c.Remove(context.Background(), &api.RemoveMessage{Path: path})
}

func (s *ClientConn) Status(path string) (*api.StatusMessage, error) {
    c := api.NewSecretClient(s.ClientConn)
    return c.Status(context.Background(), &api.StatusMessage{Path: path})
}

func (s *ClientConn) Close() {
	if s.ClientConn != nil {
		s.ClientConn.Close()
	}
}
