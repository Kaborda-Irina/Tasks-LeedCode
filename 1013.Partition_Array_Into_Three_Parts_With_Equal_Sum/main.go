package main

import "fmt"

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
}
func canThreePartsEqualSum(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	separate := 0

	for _, v := range arr {
		separate += v
	}

	if separate%3 != 0 {
		return false
	}

	sum := 0
	for _, v := range arr {
		sum += v
		if sum == separate/3 {
			sum = 0
		}
	}

	return sum == 0
}
