package main

import (
	"errors"
	"fmt"
)

type Histogram map[rune]int
type DNA string

func (d DNA) Counts() (Histogram, error) {

	var h = map[rune]int{
		'A': 0,
		'T': 0,
		'G': 0,
		'C': 0,
	}

	for _, value := range d {

		switch value {
		case 'A', 'T', 'G', 'C':
			h[value]++
		default:
			return h, errors.New("Wrong DNA Sequence!")
		}
	}

	return h, nil
}

func main() {
	fmt.Println(DNA.Counts("ATCC"))
}
