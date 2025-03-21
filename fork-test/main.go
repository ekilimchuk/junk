package main

import (
	"log"
	"math"
	"os/exec"
)

func main() {
	for i := uint64(0); i <= math.MaxUint64; i++ {
		if i > math.MaxInt {
			log.Printf("Int overflow!\n")
		}
		log.Printf("Iter %v\n", i)
		cmd := exec.Command("./script.sh")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
