package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}

func isPerfectSquare(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}

	l := 0
	r := num / 2
	for l <= r {
		m := (l + r) / 2
		v := m * m
		if v == num {
			return true
		} else if v > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
