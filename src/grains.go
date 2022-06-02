/* Calculate the number of grains of wheat on a chessboard
given that the number on each square doubles. */

package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {

	var totaly uint64 = 0
	if number < 1 || number > 64 {
		return 0, errors.New("n cannot be less than 1 or greater than 64.")
	} else {
		for i := 0; i < number; i++ {
			totaly = uint64(math.Pow(2, float64(i)))
		}
	}
	return totaly, nil
}

func Total() uint64 {
	var totaly uint64 = 0
	for i := 0; i < 64; i++ {
		totaly += uint64(math.Pow(2, float64(i)))
	}
	return totaly
}
