package main

import "fmt"

/* Factor takes in a list of primes and a number and factors that number with the provided primes.
The returned numbers can be in any order; tests will sort them in increasing
order to make checking easier.

Examples with remainders:

   Factor([2,5], 30) // []int{2,5,3}
   Factor([3,3,5], 720) // []int{3,3,5,16}
   Factor([], 4) // []int{4}
*/

func Factor(primes []int, number int) []int {

	for i, val := range primes {

		number /= val
		if i+1 == len(primes) {
			primes = append(primes, number)
		}
	}
	return primes
}

func main() {
	fmt.Println(Factor([]int{2, 2, 5}, 420))
	return
}
