package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPerfectNumber(28))
}
func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sums := 1
	sk := int(math.Sqrt(float64(num))) + 1

	for i := 2; i < sk; i++ {
		if num%i != 0 {
			continue
		}
		sums += (num / i) + i
	}

	return sums == num
}
