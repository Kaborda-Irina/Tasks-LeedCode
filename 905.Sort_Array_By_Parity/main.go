package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
func sortArrayByParity(nums []int) []int {
	result := make([]int, len(nums))
	i := 0
	j := len(nums) - 1
	for _, num := range nums {
		if num%2 == 0 {
			result[i] = num
			i++
		} else {
			result[j] = num
			j--
		}
	}
	return result
}
