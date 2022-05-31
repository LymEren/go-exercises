package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsIsogram(str string) bool {

	isogram := make(map[rune]bool)

	for _, val := range strings.ToLower(str) {
		if isogram[val] && unicode.IsLetter(val) {
			return false
		}
		isogram[val] = true
	}
	return true
}

func main() {
	fmt.Println(IsIsogram("eyyub"))
}
