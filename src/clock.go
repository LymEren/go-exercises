/* Implement a clock that handles times without dates.

You should be able to add and subtract minutes to it.

Two clocks that represent the same time should be equal to each other. */

package main

import "fmt"

const dayInMin = 1440

// Clock base structure.
type Clock struct {
	minutes int
}

// New create new Clock{}.
func New(h, m int) Clock {
	return Clock{(h*60 + m) % dayInMin}
}

// String pretty print Clock{}.
func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add minutes to existing Clock{}.
func (c Clock) Add(newMin int) Clock {
	return New(0, c.minutes+newMin)
}

// Subtract minutes from existing Clock{}.
func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}

func main() {

	var London Clock = Clock{
		minutes: 480,
	}

	fmt.Println(New(8, 0))
	fmt.Println(London.String())
	fmt.Println(London.minutes)
	London = London.Add(22)
	fmt.Println(London.minutes)
	London = London.Subtract(10)
	fmt.Println(London.minutes)

}
