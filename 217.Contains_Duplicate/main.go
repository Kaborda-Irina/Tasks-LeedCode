package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	fmt.Println(m)
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}
