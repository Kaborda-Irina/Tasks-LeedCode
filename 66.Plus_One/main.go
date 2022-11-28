package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}
func plusOne(digits []int) []int {
	result := []int{1}
	if len(digits) == 0 {
		return []int{}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	result = append(result, digits...)
	return result
}
