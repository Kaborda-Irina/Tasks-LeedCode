package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(1))
}
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for n >= 4 {
		if n%4 == 0 {
			n /= 4
		} else {
			return false
		}
	}

	return n == 1
}
