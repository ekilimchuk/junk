package main

import (
	"fmt"
	"log"
	"os"
)

func showListUsage() {
	fmt.Println("Usage: ./agent list")
	fmt.Println("\tshows all remote dirs.")
}

func ListAction(server string) {
	if len(os.Args) > 2 {
		showListUsage()
		os.Exit(1)
	}
	c, err := NewClient(server)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer c.Close()
	response, err := c.List("./")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("./")
	for _, name := range response.Dirs {
		fmt.Printf("\t%s\n", name)
	}
}
