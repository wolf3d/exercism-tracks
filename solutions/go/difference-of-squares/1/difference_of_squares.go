package diffsquares

func SquareOfSum(n int) int {
	var op int = n * ( 1 + n ) / 2
    return op * op
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2 * n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
