package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}

func searchInsert(nums []int, target int) int {
	result := -1
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			result = i
			break
		}
	}

	return result
}
