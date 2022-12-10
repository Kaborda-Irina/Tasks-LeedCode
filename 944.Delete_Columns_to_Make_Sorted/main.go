package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}

func minDeletionSize(strs []string) int {
	res := 0
	for i := 0; i < len(strs[0]); i++ {
		for i2 := 0; i2 < len(strs)-1; i2++ {
			if strs[i2][i] > strs[i2+1][i] {
				res++
				break
			}
		}
	}
	return res

}
