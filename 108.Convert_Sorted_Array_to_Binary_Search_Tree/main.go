package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	mid := (n - 1) >> 1
	res := &TreeNode{nums[mid],
		sortedArrayToBST(nums[:mid]),
		sortedArrayToBST(nums[mid+1:])}
	return res
}
