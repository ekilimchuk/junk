package series

// All returns all the contiguous substrings.
func All(n int, s string) (ret []string) {
	for i := range s[:len(s)+1-n] {
		ret = append(ret, s[i:i+n])
	}
	return ret
}

// UnsafeFirst returns a first the contiguous substring.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns a first the contiguous substring with a safe mode.
func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}
