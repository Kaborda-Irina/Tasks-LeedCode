package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:k], nums[k+1:]...)
			//len(nums)-1
			i--
			continue

		}
		k++
	}
	return k
}
