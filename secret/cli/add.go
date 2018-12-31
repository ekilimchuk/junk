package main

import (
	"flag"
	"log"
	"os"
)

func AddAction(server string) {
	src := flag.String("src", "", "is a source dir path.")
	dst := flag.String("dst", "", "is a remote dir name.")
	flag.CommandLine.Parse(os.Args[2:])
	if *src == "" || *dst == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	log.Printf("%s %s %s\n", *src, *dst, server)
}
