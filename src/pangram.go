// Panagram A pangram or holoalphabetic sentence is a sentence using every letter of a given alphabet at least once.

package main

import (
	"fmt"
	"unicode"
)

func Pangram(str string) bool {
	var cont int

	for _, value := range str {
		// ASCII CODE
		i := uint(unicode.ToLower(value) - 97)
		// if i smaller than a, in other words its not in the alphabet
		if i >= 0 && i < 26 {
			// Byte logical operators
			cont |= 1 << i

		}
	}
	return cont == (1<<26)-1
}

func main() {
	fmt.Println(Pangram("The quick brown fox jumps over the lazy dog"))
}
