/* FizzBuzz will print out all of the numbers from 1 to N replacing any divisible by 3
with "Fizz", and divisible by 5 with "Buzz", and any divisible by both with "Fizz Buzz". */

package main

import "fmt"

func FizzBuzz(n int) {

	for i := 1; i < n; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(" ", i)
		}

		if i%3 == 0 && i%5 != 0 {
			fmt.Print(" Fizz")
		}
		if i%5 == 0 && i%3 != 0 {
			fmt.Print(" Buzz")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(" Fizz Buzz")
		}
	}

}

func main() {
	FizzBuzz(17)
}
