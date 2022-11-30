package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
}

func isPowerOfTwo(n int) bool {

	if n <= 0 {
		return false
	}
	for n >= 2 {
		if n%2 == 0 {
			n /= 2
		} else {

			return false
		}
	}
	return true
}
