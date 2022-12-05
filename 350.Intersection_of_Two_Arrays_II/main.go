package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var res []int
	for _, i := range nums1 {
		if _, ok := m[i]; ok {
			m[i]++
		} else {
			m[i] = 1
		}

	}

	for _, j := range nums2 {
		if v, ok := m[j]; v > 0 && ok {
			res = append(res, j)
			m[j]--
		}

	}
	return res
}
