package main

import (
	"fmt"
)

func main() {
	var sumOfSquares int = 0
	var squareOfSum int = 0
	for i := 1; i <= 100; i++ {
		sumOfSquares = sumOfSquares + i*i
		squareOfSum = squareOfSum + i
	}
	squareOfSum = squareOfSum * squareOfSum
	fmt.Println("Square of sum - sumOfSquares = ", squareOfSum, " - ", sumOfSquares, " = ", squareOfSum-sumOfSquares)
}
