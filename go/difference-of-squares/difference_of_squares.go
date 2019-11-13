package diffsquares

//SquareOfSum calculates the square of the sum of the n numbers
func SquareOfSum(n int) int {
    sum := (n) * (n + 1) / 2
    return sum * sum
}

//SumOfSquares calculates the sum of squares of n numbers
func SumOfSquares(n int) int {
    sumOfSq := n * (n + 1) * (2*n + 1) / 6

    return sumOfSq
}

//Difference calculates the difference between square of sum and sum of squares
func Difference(n int) int {
    return SquareOfSum(n) - SumOfSquares(n)
}
