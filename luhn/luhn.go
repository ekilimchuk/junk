package luhn

import (
	"errors"
	"strconv"
)

// Get and check a valid card number.
func getValidCode(s string) ([]int, error) {
	numbers := make([]int, 0, len(s))
	for _, v := range s {
		c, err := strconv.Atoi(string(v))
		if err != nil {
			// Skip a valid rune.
			if v == ' ' {
				continue
			}
			// Invalid string.
			return numbers, errors.New("String is not valid")
		}
		numbers = append(numbers, c)
	}
	// Invalid size.
	if len(numbers) < 2 {
		return numbers, errors.New("Single digit string in not valid.")
	}
	return numbers, nil
}

// Check all numbers on zero
func isAllZeros(numbers []int) bool {
	for _, v := range numbers {
		if v != 0 {
			return false
		}
	}
	return true
}

// Sum all numbers
func sumNumbers(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// Valid determines whether or not string numbers is valid per the Luhn formula.
func Valid(s string) bool {
	numbers, err := getValidCode(s)
	if err != nil {
		return false
	}
	if isAllZeros(numbers) {
		return len(numbers) > 1
	}
	// Starting from the right and double every second digit.
	for i := len(numbers) - 2; i >= 0; i -= 2 {
		n := numbers[i] * 2
		if n > 9 {
			numbers[i] = n - 9
			continue
		}
		numbers[i] = n
	}
	return sumNumbers(numbers)%10 == 0
}
