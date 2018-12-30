package main

import (
    "log"
	"fmt"
	"os"
)

func showListUsage() {
	fmt.Println("Usage: ./secret-cli list")
	fmt.Println("	shows all remote dirs.")
}

func ListAction() {
	if len(os.Args) > 2 {
		showListUsage()
		os.Exit(1)
	}
	c, err := NewClient("", "8888");
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
