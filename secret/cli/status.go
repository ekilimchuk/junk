package main

import (
	"flag"
	"log"
	"os"
)

func StatusAction() {
	dst := flag.String("dst", "", "is a remote dir name.")
	flag.CommandLine.Parse(os.Args[2:])
	if *dst == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	log.Printf("%s\n", *dst)
}
