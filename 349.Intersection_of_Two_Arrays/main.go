package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	m := make(map[int]bool)
	for _, i := range nums1 {
		m[i] = true
	}
	for _, j := range nums2 {
		if _, ok := m[j]; ok {
			result = append(result, j)
			delete(m, j)
		}
	}
	return result
}
