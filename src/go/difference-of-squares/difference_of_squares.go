package diffsquares

import "math"

func SquareOfSum(n int) int {
	// S(i) = n(n+1)/2
	// => S(i)^2 = n^2(n+1)^2/4
	return int((math.Pow(float64(n), 2)) * (math.Pow(float64(n)+1, 2)) / 4)
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
