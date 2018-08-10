package main

import (
	"fmt"
	"strconv"
)

// CheckLuhn determines whether or not string numbers is valid per the Luhn formula.
func CheckLuhn(s string) bool {
	var i, sum int
	for _, v := range s {
		c, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		}
		if i%2 == 0 {
			d := c * 2
			if d < 9 {
				sum += d
			} else {
				sum += d - 9
			}
		} else {
			sum += c
		}
		i++
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

func main() {
	s := "4539 1488 0343 6467"
	fmt.Printf("%s - %v\n", s, Luhn(s))
	s = "8273 1232 7352 0569"
	fmt.Printf("%s - %v\n", s, Luhn(s))
}
