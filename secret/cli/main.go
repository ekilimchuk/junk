package main

import (
	"fmt"
	"os"
)

func showUsage() {
	fmt.Println("Usage: ./secret-cli <list|add|status|remove|fingers> -h")
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "list":
		ListAction()
	case "add":
		AddAction()
	case "remove":
		RemoveAction()
	case "status":
		StatusAction()
	case "fingers":
		FingersAction()
	default:
		showUsage()
		os.Exit(1)
	}
}