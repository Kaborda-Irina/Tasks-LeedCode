package main

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}

func findErrorNums(nums []int) []int {
	m := make(map[int]int)
	sum := 0
	r := make([]int, 2)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			r[0] = v
		} else {
			m[v] = 1
			sum += v
		}
	}
	l := len(nums)
	r[1] = l*(1+l)/2 - sum
	return r
}
