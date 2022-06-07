/* Given a number determine whether or not it is valid per the Luhn formula.
The Luhn algorithm is a simple checksum formula used to validate a variety of identification numbers,
such as credit card numbers and Canadian Social Insurance Numbers.
The task is to check if a given string is valid. */

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Valid(input string) bool {
	var sum int
	cardNumber := strings.Join(strings.Fields(input), "")
	// fmt.Println(cardNumber)

	if len(cardNumber) < 2 {
		return false
	}
	b := 1
	for i := len(cardNumber) - 1; i >= 0; i-- {
		var char = rune(cardNumber[i])
		if unicode.IsDigit(char) == false {
			return false
		}
		num := int(char - 48)
		// For the ASCII codes

		if b%2 == 0 {
			if num*2 > 9 {
				num = num*2 - 9
			} else {
				num = num * 2
			}
			sum += num
		} else {
			sum += num
		}
		b++
	}
	return sum%10 == 0
}

func main() {
	fmt.Println(Valid("059"))

	return
}
