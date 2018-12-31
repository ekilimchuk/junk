package main

import (
	"../util"
	"flag"
	"fmt"
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
	if err := util.Tar(*src, fmt.Sprintf("./%s.tar", *dst)); err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%s %s %s\n", *src, *dst, server)
}
