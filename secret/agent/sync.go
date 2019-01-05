package main

import (
	"fmt"
	"os"
)

func showSyncUsage() {
	fmt.Println("Usage: ./agent sync")
	fmt.Println("\tsync all remote dirs.")
}

func SyncAction(server string) {
	if len(os.Args) < 2 {
		showSyncUsage()
		os.Exit(1)
	}
	fmt.Printf("SYNC %s\n", server)
}
