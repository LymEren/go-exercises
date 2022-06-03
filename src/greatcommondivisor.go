package main

import "fmt"

func GCD(a, b int) int {
	greatnum := 1
	if b > a {
		b, a = a, b
	}
	if a > b {
		if a%b == 0 {
			return b
		} else {
			for i := 1; i <= b; {
				if a%i == 0 && b%i == 0 {
					greatnum *= i
					a /= i
					b /= i
				}
				i++
			}
		}
	}
	return greatnum
}

func main() {
	fmt.Println(GCD(30, 100))
}
