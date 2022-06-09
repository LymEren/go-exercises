/* Determine if a triangle is equilateral, isosceles, or scalene.

An equilateral triangle has all three sides the same length.

An isosceles triangle has at least two sides the same length.
(It is sometimes specified as having exactly two sides the same length,
	but for the purposes of this exercise we'll say at least two.)

A scalene triangle has all sides of different lengths. */

package main

import (
	"fmt"
	"math"
)

type Kind string

// Iota is an identifier which is used with constant and
// which can simplify constant definitions that use auto increment numbers.

const (
	NaT Kind = "Not a Triangle"
	Equ      = "Equilateral"
	Iso      = "Isosceles"
	Sca      = "Scalene"
)

func KindFromSides(a, b, c float64) Kind {

	if math.IsNaN(a+b+c) || a+b <= c || a+c <= b || b+c <= a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

func main() {
	fmt.Println("Triangle Type is: ", KindFromSides(10, 12, 14))
	fmt.Println("Triangle Type is: ", KindFromSides(5, 12, 13))
	fmt.Println("Triangle Type is: ", KindFromSides(6, 6, 6))
	fmt.Println("Triangle Type is: ", KindFromSides(6, 6, 8))
	fmt.Printf("%v, %T ", KindFromSides(6, 6, 8), KindFromSides(6, 6, 8))
}
