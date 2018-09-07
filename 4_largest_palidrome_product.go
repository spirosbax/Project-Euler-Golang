package main

import (
	"fmt"
	// "math"
)

const largest_3digit = 998001

func main() {
	var largest_palindrome int64 = 0
	for i := 0; i <= largest_3digit; i++ {
		if isPalindrome(int64(i)) {
			largest_palindrome = Max(largest_palindrome, int64(i))
		}
	}
	fmt.Println(largest_palindrome)
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func isPalindrome(num int64) bool {
	tempNum := num
	var numReversed int64 = 0
	for tempNum > 0 {
		numReversed = numReversed*10 + tempNum%10
		tempNum = tempNum / 10
	}

	if num-numReversed == 0 {
		return true
	}
	return false
}
