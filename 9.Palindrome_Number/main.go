package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}
func isPalindrome(x int) bool {
	res := 0
	num := x
	for num > 0 {
		rem := num % 10
		res = (res * 10) + rem
		num /= 10
	}
	if x == res {
		return true
	} else {
		return false
	}
}
