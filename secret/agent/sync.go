package main

import (
	"fmt"
	"os"
	//	"../util"
	"flag"
	"log"
)

func SyncAction(server string) {
	dst := flag.String("dst", "", "is a remote dir name.")
	flag.CommandLine.Parse(os.Args[2:])
	if *dst == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	c, err := NewClient(server)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer c.Close()
	response, err := c.Sync(*dst)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%s\n", string(response.Aeskey))
}
