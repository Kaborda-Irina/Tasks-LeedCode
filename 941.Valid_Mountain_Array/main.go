package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 1}))
	fmt.Println(validMountainArray([]int{0, 2, 3, 5, 2, 1}))
}
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	i := 0
	l := len(arr)
	for i+1 < l && arr[i] < arr[i+1] {
		i++
	}
	if i == 0 || i == l-1 {
		return false
	}
	for i+1 < l && arr[i] > arr[i+1] {
		i++
	}
	return i == l-1
}
