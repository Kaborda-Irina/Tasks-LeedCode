package main

import "fmt"

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
}

func moveZeroes(nums []int) []int {
	k := 0
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
			nums = append(nums[:k], nums[k+1:]...)
			i--
			continue

		}
		k++
	}
	for i := 0; i < zero; i++ {
		nums = append(nums, 0)
	}
	return nums
}
