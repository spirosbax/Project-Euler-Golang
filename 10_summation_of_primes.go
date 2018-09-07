package main

import (
	"fmt"
	"helper"
)

func main() {
	limit := 2000000
	var sum = 0

	// for all numbers smaller than the limit
	for i := 2; i < limit; i++ {
		// check if it is prime
		if helper.IsPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
