package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{6, 6, 6, 6}))
}
func distributeCandies(candyType []int) int {
	n := len(candyType) / 2
	m := make(map[int]int)

	for _, v := range candyType {
		m[v]++
	}
	countType := len(m)
	if countType < n {
		return countType
	}

	return n
}
