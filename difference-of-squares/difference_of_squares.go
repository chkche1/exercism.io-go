package diffsquares

/*
SquareOfSum returns the square of the sum of the first n natural numbers
*/
func SquareOfSum(n int) int {
	squareSum := (1 + n) * n / 2
	return squareSum * squareSum
}

/*
SumOfSquares returns the sum of the squares of the first n natural numbers
*/
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

/*
Difference returns the difference of the square of the sum of the first n natural numbers and the sum of the squares of the first n natural numbers
*/
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
