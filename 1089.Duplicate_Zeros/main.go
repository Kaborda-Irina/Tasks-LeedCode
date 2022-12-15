package main

import "fmt"

func main() {
	fmt.Println(duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0}))
}
func duplicateZeros(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i+1 < len(arr) {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}

	}
	return arr
}
