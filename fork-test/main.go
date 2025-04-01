package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
)

func worker(n uint32, results chan<- string) {
	pwd := os.Getenv("PWD")
	for i := uint64(0); i <= math.MaxUint64; i++ {
		if i > math.MaxInt {
			results <- fmt.Sprintf("Worker %v int overflow!")
		}
		if i%1000 == 0 {
			results <- fmt.Sprintf("%v %v", n, i)
		}
		cmd := exec.Command(pwd + "/script.sh")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	workersCount := uint32(64)
	results := make(chan string)
	for n := uint32(0); n <= workersCount; n++ {
		go worker(n, results)
	}
	for {
		log.Printf("%v\n", <-results)
	}
}
