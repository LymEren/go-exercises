package scrabble

import "unicode"

// SCRABBLE EYYUB EREN - This is a different solution :)
//
func Score(word string) int {
	counter := 0
	// i used rune datatype for the for-range loop
	onePoint := []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}
	twoPoint := []rune{'D', 'G'}
	threePoint := []rune{'B', 'C', 'M', 'P'}
	fourPoint := []rune{'F', 'H', 'V', 'W', 'Y'}
	fivePoint := []rune{'K'}
	eightPoint := []rune{'J', 'X'}
	tenPoint := []rune{'Q', 'Z'}

	for _, metter := range word {
		letter := unicode.ToUpper(metter)
		for j := 0; j <= 9; j++ {
			if letter == (onePoint[j]) {
				counter += 1
			}
			if j < 2 && letter == twoPoint[j] {
				counter += 2
			}
			if j < 4 && letter == threePoint[j] {
				counter += 3
			}
			if j < 5 && letter == fourPoint[j] {
				counter += 4
			}
			if j < 1 && letter == fivePoint[j] {
				counter += 5
			}
			if j < 2 && letter == eightPoint[j] {
				counter += 8
			}
			if j < 2 && letter == tenPoint[j] {
				counter += 10
			}
		}
	}
	return counter
}
