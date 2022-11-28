package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(10))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}
