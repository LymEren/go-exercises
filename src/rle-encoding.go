package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(s string) string {
	var result string
	for len(s) > 0 {
		letter := s[0]
		stringlen := len(s)
		s = strings.TrimLeft(s, string(letter))

		rep := stringlen - len(s)
		if rep > 1 {
			// Adding repeat number
			result += strconv.Itoa(rep)
		}
		// Adding repeated letter
		result += string(letter)
	}
	return result
}

func RunLengthDecode(s string) string {
	var result string
	for len(s) > 0 {
		i := strings.IndexFunc(s, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		n := 1
		if i != 0 {
			n, _ = strconv.Atoi(s[:i])
		}
		result += strings.Repeat(string(s[i]), n)
		s = s[i+1:]
	}
	return result
}

func main() {

	fmt.Println(RunLengthEncode("AABCCCDEEEE"))
	fmt.Println(RunLengthDecode("2AB3CD4E"))

}
