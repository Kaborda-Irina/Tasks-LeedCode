package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(6))
}
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n >= 3 {
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}
	return n == 1
}
