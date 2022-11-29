package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	max := 0
	num := 0
	for k, i := range m {
		if i > max {
			max = i
			num = k
		}
	}
	return num
}
