package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 3, 1, 3}))
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := 0
	for i := 0; i < len(nums); i++ {

		res ^= nums[i]
	}
	return res
}
