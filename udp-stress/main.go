package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	h = flag.String("h", "", "a hosts list")
	t = flag.Int("t", 1, "a parallel operation count")
	n = flag.Bool("6", false, "use udp6")
)

func getHosts(s string) (result []string) {
	for _, v := range strings.Split(s, ",") {
		result = append(result, strings.TrimSpace(v))
	}
	return result
}

func run(count int, hosts []string, proto string) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				for _, host := range hosts {
					port := 1024
					conn, err := net.Dial(proto, fmt.Sprintf("%s:%d", host, port))
					if err != nil {
						fmt.Printf("%s:%d - %v\n", host, port, err)
						break
					}
					fmt.Fprintf(conn, "blah")
					conn.Close()
				}
			}
		}()
	}
	wg.Wait()
}

func main() {
	flag.Parse()
	if *h == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	proto := "udp4"
	if *n {
		proto = "udp6"
	}
	hosts := getHosts(*h)
	fmt.Printf("Host(s): %v\n", hosts)
	fmt.Printf("Parallel operation count: %v\n", *t)
	run(*t, hosts, proto)
}
