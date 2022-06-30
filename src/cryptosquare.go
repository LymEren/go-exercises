/*
Implement the classic method for composing secret messages called a square code.

Given an English text, output the encoded version of that text.

First, the input is normalized: the spaces and punctuation are removed from the English text and the message is down-cased.

Then, the normalized characters are broken into rows. These rows can be regarded as forming a rectangle when printed with intervening newlines. */

package main

import (
	"fmt"
	"strings"
)

func Encode(in string) string {
	if in == "" {
		return ""
	}

	var norm string
	for i := 0; i < len(in); i++ {
		l := strings.ToLower(string(in[i]))
		if strings.ContainsAny(l, "qwertyuioplkjhgfdsazxcvbnm1234567890") {
			norm += l
		}
	}
	r := 1
	c := 1
	for r*c < len(norm) {
		if r == c {
			c++
		} else {
			r++
		}
	}
	var out string
	for j := 0; j < c; j++ {
		for k := 0; k < r; k++ {
			if j+c*k < len(norm) {
				out += string(norm[j+c*k])
			} else {
				out += " "
			}
		}
		if j != c-1 {
			out += " "
		}
	}
	return out
}

func main() {

	// RESULT
	fmt.Println(Encode("If man was meant to stay on the ground, god would have given us roots."))

}
