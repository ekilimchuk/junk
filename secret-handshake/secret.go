package secret

// Handshake decodes a binary code to phrases.
func Handshake(n uint) (ret []string) {
	if n&1 == 1 {
		ret = append(ret, "wink")
	}
	if n&2 == 2 {
		ret = append(ret, "double blink")
	}
	if n&4 == 4 {
		ret = append(ret, "close your eyes")
	}
	if n&8 == 8 {
		ret = append(ret, "jump")
	}
	if n&16 == 16 {
		for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
			ret[i], ret[j] = ret[j], ret[i]
		}
	}
	return ret
}
