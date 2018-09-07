package main

import (
	"fmt"
)

func main() {
	var first_fib int = 1
	var second_fib int = 2
	var last_fib int = second_fib
	var sum int = 0
	for last_fib <= 4000000 {
		if last_fib%2 == 0 {
			sum = sum + last_fib
		}
		last_fib = first_fib + second_fib
		first_fib = second_fib
		second_fib = last_fib
	}
	fmt.Println(sum)
}
