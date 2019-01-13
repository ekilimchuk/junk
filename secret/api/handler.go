package api

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
)

type Server struct{}

func (s *Server) List(ctx context.Context, in *ListMessage) (*ListResult, error) {
	log.Printf("Receive: list %s", "./tmp")
	files, err := ioutil.ReadDir("./tmp")
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

func (s *Server) Add(ctx context.Context, in *AddMessage) (*AddResult, error) {
	log.Printf("Receive: %s %s", in.Aeskey, in.Dirname)
	if err := os.Mkdir(fmt.Sprintf("./tmp/%s", in.Dirname), 0755); err != nil {
		return &AddResult{Result: "Error: Add"}, err
	}
	if err := ioutil.WriteFile(fmt.Sprintf("./tmp/%s/key", in.Dirname), []byte(in.Aeskey), 0600); err != nil {
		return &AddResult{Result: "Error: Add"}, err
	}
	if err := ioutil.WriteFile(fmt.Sprintf("./tmp/%s/blob", in.Dirname), in.Blob, 0600); err != nil {
		return &AddResult{Result: "Error: Add"}, err
	}
	return &AddResult{Result: "Ok"}, nil
}

func (s *Server) Sync(ctx context.Context, in *SyncMessage) (*SyncResult, error) {
	log.Printf("Receive: %s", in.Dirname)
	key, err := ioutil.ReadFile(fmt.Sprintf("./tmp/%s/key", in.Dirname))
	if err != nil {
		log.Printf("Warrning: %v", err)
		return &SyncResult{Aeskey: []byte(""), Blob: []byte("")}, err
	}
	blob, err := ioutil.ReadFile(fmt.Sprintf("./tmp/%s/blob", in.Dirname))
	if err != nil {
		log.Printf("Warrning: %v", err)
		return &SyncResult{Aeskey: []byte(""), Blob: []byte("")}, err
	}
	return &SyncResult{Aeskey: key, Blob: blob}, nil
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
