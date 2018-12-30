package main

import (
    "log"
	"fmt"
	"os"
)

func showFingersUsage() {
	fmt.Println("Usage: ./secret-cli fingers")
	fmt.Println("\tshows all remote rsa pub fingers.")
}

func FingersAction() {
	if len(os.Args) > 2 {
		showListUsage()
		os.Exit(1)
	}
	c, err := NewClient("", "8888");
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
