package api

import (
	"log"
	"golang.org/x/net/context"
	"io/ioutil"
	"bufio"
	"os"
)

type Server struct {}

func (s *Server) List(ctx context.Context, in *ListMessage) (*ListResult, error) {
	log.Printf("Receive: list %s", in.Path)
	files, err := ioutil.ReadDir(in.Path)
	if err != nil {
        log.Println(err)
		return nil, err
    }
	list := make([]string, 0)
	for _, f := range files {
		list = append(list, f.Name())
	}
	return &ListResult{Dirs: list}, nil
}

func (s *Server) Remove(ctx context.Context, in *RemoveMessage) (*RemoveMessage, error) {
	log.Printf("Receive: %s", in.Path)
	return &RemoveMessage{Path: in.Path}, nil
}

func (s *Server) Add(ctx context.Context, in *AddMessage) (*AddMessage, error) {
	log.Printf("Receive: %s", in.Aeskey)
	return &AddMessage{Aeskey: in.Aeskey}, nil
}

func (s *Server) Status(ctx context.Context, in *StatusMessage) (*StatusMessage, error) {
	log.Printf("Receive: %s", in.Path)
	return &StatusMessage{Path: in.Path}, nil
}

func (s *Server) Fingers(ctx context.Context, in *FingersMessage) (*FingersResult, error) {
	log.Printf("Receive: fingers - list rsa pub fingers.")
	file, err := os.Open("./known_hosts")
	if err != nil {
        log.Println(err)
		return nil, err
    }
	defer file.Close()
	scanner := bufio.NewScanner(file)
	list := make([]string, 0)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return &FingersResult{Fingers: list}, nil
}
