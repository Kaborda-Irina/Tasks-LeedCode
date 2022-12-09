package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(isMonotonic([]int{1, 3, 2}))
}
func isMonotonic(nums []int) bool {
	up, down := true, true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			down = false
		}
		if nums[i] < nums[i-1] {
			up = false
		}
	}
	return up || down
}
