package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{5, 7, 2, 4}))
}
func sortArrayByParityII(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	result := make([]int, len(nums))
	one, two := 0, 1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			result[one] = nums[i]
			one += 2
		} else {
			result[two] = nums[i]
			two += 2
		}
	}

	return result
}
