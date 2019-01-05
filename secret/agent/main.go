package main

import (
	"../util"
	"fmt"
	"log"
	"os"
)

const (
	FileName = "./configs/secret-agent.config.json.example"
)

func showUsage() {
	fmt.Println("Usage: ./agent <list|sync> -h")
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		os.Exit(1)
	}
	c := util.AgentConfig{}
	if err := util.LoadConfig(&c, FileName); err != nil {
		log.Fatalf("%v", err)
	}
	k := Keys{}
	if err := k.LoadKeys(c); err != nil {
		log.Fatalf("%v", err)
	}
	switch os.Args[1] {
	case "list":
		ListAction(c.Server)
	case "sync":
		SyncAction(c.Server)
	default:
		showUsage()
		os.Exit(1)
	}
}
