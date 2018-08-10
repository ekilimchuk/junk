package main

import (
	"fmt"
	"strconv"
)

// Decode decodes a input string
func Decode(s string) string {
	var res, count string
	for _, v := range s {
		switch v {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			count += string(v)
		default:
			c := 1
			if count != "" {
				c, _ = strconv.Atoi(count)
			}
			for i := 0; i < c; i++ {
				res += fmt.Sprintf("%s", string(v))
			}
			count = ""
		}
	}
	return res
}

func getStr(n int, r rune) string {
	if n == 1 {
		return string(r)
	}
	return fmt.Sprintf("%d%s", n, string(r))
}

// Encode encodes a input string
func Encode(s string) string {
	if len(s) == 0 {
		return ""
	}
	var res string
	var tmpChar rune = rune(s[0])
	var countChar int
	for _, v := range s {
		if tmpChar == v {
			countChar++
		} else {
			res += getStr(countChar, tmpChar)
			countChar = 1
		}
		tmpChar = v
	}
	res += getStr(countChar, tmpChar)
	return res
}

func main() {
	// Decode
	s := "aacbbmvvv"
	fmt.Printf("%s =  %s\n", s, Encode(s))
	s = "aagghhyyyyyyyyyyyyycbbfghjmvvv"
	fmt.Printf("%s =  %s\n", s, Encode(s))
	s = "a"
	fmt.Printf("%s =  %s\n", s, Encode(s))
	s = ""
	fmt.Printf("%s =  %s\n", s, Encode(s))
	s = "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW"
	fmt.Printf("%s =  %s\n", s, Encode(s))
	// Encode
	s = "2ac2bm3v"
	fmt.Printf("%s =  %s\n", s, Decode(s))
	s = "2a2g2h13yc2bfghjm3v"
	fmt.Printf("%s =  %s\n", s, Decode(s))
	s = "a"
	fmt.Printf("%s =  %s\n", s, Decode(s))
	s = ""
	fmt.Printf("%s =  %s\n", s, Decode(s))
	s = "127A127A2A"
	fmt.Printf("%s =  %s\n", s, Decode(s))

}
