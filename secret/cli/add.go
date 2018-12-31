package main

import (
	"../util"
	"flag"
	"fmt"
	"io/ioutil"
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
	c, err := NewClient(server)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer c.Close()
	file, err := ioutil.ReadFile(fmt.Sprintf("./%s.tar", *dst))
	if err != nil {
		log.Fatalf("%v", err)
	}
	response, err := c.Add(*dst, "qwerty12", file)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%s\n", response.Result)
}
