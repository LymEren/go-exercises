/* Use the Sieve of Eratosthenes to find all the primes from 2 up to a given number.

The Sieve of Eratosthenes is a simple, ancient algorithm for finding all prime numbers up to any given limit.
It does so by iteratively marking as composite (i.e. not prime) the multiples of each prime,
starting with the multiples of 2. It does not use any division or remainder operation. */

package main

import "fmt"

func Sieve(n int) []int {
	// n+1 allows us to start the loop at 2
	// without a subtracting 2 from each iteration
	// (this just ignores indexes 0, 1)
	marks := make([]bool, n+1)
	//preallocate the array; upper bound of half (no even numbers > 2)
	primes := make([]int, n/2)
	pidx := 0
	for i := 2; i <= n; i++ {
		if marks[i] {
			continue
		}
		primes[pidx] = i
		pidx++
		// walk the array by increments of the current value
		// eliminating all multiples of i
		for m := i; m <= n; m += i {
			marks[m] = true
		}
	}
	return primes[:pidx]
}

func main() {
	fmt.Println(Sieve(15))
}
