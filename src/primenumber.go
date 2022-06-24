/* Given a number n, determine what the nth prime is.

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13. */

package main

import (
	"errors"
	"math"
)

//IsPrime check if a number is prime
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}
	is := true
	max := math.Sqrt(float64(number))
	for i := 2; i <= int(max); i++ {
		if number%i == 0 {
			is = false
			break
		}
	}
	return is
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	counter := 0
	i := 0
	if n <= 0 {
		return n, errors.New("'n' is equal or less than zero")
	}
	for counter < n {
		if IsPrime(i) {
			counter++
		}
		i++
	}
	nthPrime := i - 1
	return nthPrime, nil
}
