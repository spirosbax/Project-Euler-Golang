// simple sieve of eratosthenis to calculate Nth prime
package main

import (
	"fmt"
	"math"
)

const n = 110000

func main() {
	primes := make(map[int]bool)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if primes[i] {
			k := 0
			j := int(math.Pow(float64(i), 2.0)) + k*i
			for j <= n {
				primes[j] = false
				k++
				j = int(math.Pow(float64(i), 2.0)) + k*i
			}
		}
	}

	realprimes := []int{}
	for i := 2; i <= n; i++ {
		if primes[i] {
			realprimes = append(realprimes, i)
		}
	}

	fmt.Println(realprimes[10000]) // slices are 0 indexed
}
