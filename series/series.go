package series

func All(n int, s string) (ret []string) {
	for i, _ := range s[:len(s)+1-n] {
		ret = append(ret, s[i:i+n])
	}
	return ret
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}
