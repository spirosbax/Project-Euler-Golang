package helper

import (
	"math"
)

func IsPrime(n int) bool {
	// implement primality test
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
