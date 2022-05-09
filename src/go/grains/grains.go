package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Only accepts between 1 and 64")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() (total uint64) {
	// If I remember my maths right, the optimized version would be "Square(65)-1" (with a change
	// to the Square function to permit >64 as arg), but here I'm optimizing for readability)
	for i := 1; i < 65; i++ {
		// Man I am really not loving GoLang's error-handling. Shouldn't we throw an exception (or,
		// idiomatically, return an error) if Square returns one?
		val, _ := Square(i)
		total += val
	}
	return
}
