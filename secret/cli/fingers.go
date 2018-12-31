package main

import (
	"fmt"
	"log"
	"os"
)

func showFingersUsage() {
	fmt.Println("Usage: ./secret-cli fingers")
	fmt.Println("\tshows all remote rsa pub fingers.")
}

func FingersAction(server string) {
	if len(os.Args) > 2 {
		showListUsage()
		os.Exit(1)
	}
	c, err := NewClient(server)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer c.Close()
	response, err := c.Fingers("./")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("./Fingers list:")
	for _, finger := range response.Fingers {
		fmt.Printf("\t%s\n", finger)
	}
}
