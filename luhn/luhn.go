package luhn

import (
	"strconv"
)

// Valid determines whether or not string numbers is valid per the Luhn formula.
func Valid(s string) bool {
	numbers, ok := getValidCode(s)
	if !ok {
		return false
	}
	sum := 0
	// Starting from the right and double every second digit.
	for i := len(numbers) - 1; i >= 0; i-- {
		if (i+1)%2 == 0 {
			if numbers[i] *= 2; numbers[i] > 9 {
				numbers[i] -= 9
			}
		}
		sum += numbers[i]
	}
	return sum%10 == 0
}

// Get and check a valid card number.
func getValidCode(s string) ([]int, bool) {
	numbers := make([]int, 0, len(s))
	for _, v := range s {
		c, err := strconv.Atoi(string(v))
		if err != nil {
			// Skip a valid rune.
			if v == ' ' {
				continue
			}
			// Invalid string.
			return numbers, false
		}
		numbers = append(numbers, c)
	}
	return numbers, len(numbers) > 2
}
