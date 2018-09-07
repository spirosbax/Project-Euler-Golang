package main

import (
	"fmt"
	"helper"
	"math"
)

const num = 600851475143

func main() {
	var largest_prime int = 1

	for i := int(math.Sqrt(float64(num))); i >= 1; i-- {
		fmt.Println(i)
		if helper.IsPrime(i) {
			if num%i == 0 && i > largest_prime {
				largest_prime = i
				fmt.Println("set largest_prime to ", largest_prime)
			}
		}
	}
	fmt.Println(largest_prime)
}
