package main

import (
	"fmt"
)

func main() {
	for i := 1; i >= 0; i++ {
		if isDivisibleByAll(int64(i)) == true {
			fmt.Println("I is:", i)
			break
		}
	}
}

func isDivisibleByAll(num int64) bool {
	for j := 1; j <= 20; j++ {
		// fmt.Println(num, " % ", j, " = ", num%int64(j))
		if !(num%int64(j) == 0) {
			return false
		}
	}
	return true
}
