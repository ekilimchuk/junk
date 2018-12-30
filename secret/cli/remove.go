package main

import (
    "log"
    "os"
	"flag"
)


func RemoveAction() {
	dst := flag.String("dst", "", "is a remote dir name.")
	flag.CommandLine.Parse(os.Args[2:])
	if *dst == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	log.Printf("%s\n", *dst)
}
