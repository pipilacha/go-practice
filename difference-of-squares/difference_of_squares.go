package diffsquares

func SquareOfSum(n int) int {
	s := (n * (n + 1)) / 2
	return s * s
}

func SumOfSquares(n int) int {
	s := (n * (n + 1) * (2*n + 1)) / 6
	return s
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
