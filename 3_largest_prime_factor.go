package main

import (
	"fmt"
	"math"
)

const num = 600851475143

func main() {
	var largest_prime int = 1

	for i := int(math.Sqrt(float64(num))); i >= 1; i-- {
		fmt.Println(i)
		if isPrime(i) {
			if num%i == 0 && i > largest_prime {
				largest_prime = i
				fmt.Println("set largest_prime to ", largest_prime)
			}
		}
	}
	fmt.Println(largest_prime)
}

func isPrime(n int) bool {
	// implement primality test
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
