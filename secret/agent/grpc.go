package main

import (
	"../api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ClientConn struct {
	*grpc.ClientConn
}

func NewClient(server string) (*ClientConn, error) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &ClientConn{conn}, nil
}

func (s *ClientConn) List(path string) (*api.ListResult, error) {
	c := api.NewSecretClient(s.ClientConn)
	return c.List(context.Background(), &api.ListMessage{Path: path})
}

func (s *ClientConn) Sync(dir string) (*api.SyncResult, error) {
	c := api.NewSecretClient(s.ClientConn)
	return c.Sync(context.Background(), &api.SyncMessage{Dirname: dir})
}

func (s *ClientConn) Close() {
	if s.ClientConn != nil {
		s.ClientConn.Close()
	}
}
