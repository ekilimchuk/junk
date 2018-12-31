package main

import (
	"../util"
	"fmt"
	"log"
	"os"
)

const (
	FileName = "./configs/secret-cli.config.json.example"
)

func showUsage() {
	fmt.Println("Usage: ./secret-cli <list|add|status|remove|fingers> -h")
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		os.Exit(1)
	}
	c := util.CliConfig{}
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
	case "add":
		AddAction(c.Server)
	case "remove":
		RemoveAction(c.Server)
	case "status":
		StatusAction()
	case "fingers":
		FingersAction(c.Server)
	default:
		showUsage()
		os.Exit(1)
	}
}
