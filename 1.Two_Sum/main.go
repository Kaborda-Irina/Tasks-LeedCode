package main

import "fmt"

func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 5, 5, 2, 15}, 10))
}

func twoSum(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		for next := i + 1; next < len(nums); next++ {
			if nums[i]+nums[next] == target {
				result = append(result, i, next)
			}
		}
	}

	return result
}
