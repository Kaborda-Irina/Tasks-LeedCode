package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	k := 0
	num := 0
	result := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			k = i
		} else {
			if num == k {
				result = append(result, fmt.Sprintf("%d", nums[num]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[num], nums[k]))
			}
			k = i
			num = i
		}
	}
	if k == num {
		result = append(result, fmt.Sprintf("%v", nums[k]))
	} else {
		result = append(result, fmt.Sprintf("%v->%v", nums[num], nums[k]))
	}
	return result
}
