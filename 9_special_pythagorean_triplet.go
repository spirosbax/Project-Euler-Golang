package main

import (
	"fmt"
)

func main() {
	for a := 0; a <= 332; a++ {
		for b := a + 1; b <= 499; b++ {
			for c := b + 1; c <= 1000; c++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					fmt.Println(a, b, c, " = ", a*b*c)
				}
			}
		}
	}
}
