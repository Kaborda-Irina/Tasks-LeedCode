package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	var num int
	if n == 0 || n == 1 {
		return 1
	}
	tmp := []int{1, 2}
	for i := 2; i < n; i++ {
		tmp = append(tmp, tmp[i-1]+tmp[i-2])
	}
	return tmp[n-1]

	return num
}
