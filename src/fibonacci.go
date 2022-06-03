package main

import "fmt"

// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377

func Fibonacci(n int) int {
	series := make(map[int]int)
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			series[i] = i
		}
		if i > 1 {
			series[i] = series[i-1] + series[i-2]
		}
	}
	return series[n]
}

func main() {
	fmt.Println(Fibonacci(7))
}
