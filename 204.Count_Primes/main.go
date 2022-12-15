package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	nums := make([]bool, n+1)
	for i := 2; i < n; i++ {
		if nums[i] == false {
			for j := i + i; j < n; j += i {
				nums[j] = true
			}
		}
	}
	cnt := 0
	for i := 2; i < n; i++ {
		if nums[i] == false {
			cnt++
		}
	}
	return cnt
}
