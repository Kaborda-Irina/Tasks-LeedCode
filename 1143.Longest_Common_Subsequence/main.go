package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	m := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		m[i] = make([]int, n2+1)
	}
	for i1 := 1; i1 <= n1; i1++ {
		for i2 := 1; i2 <= n2; i2++ {
			if text1[i1-1] == text2[i2-1] {
				m[i1][i2] = m[i1-1][i2-1] + 1
			} else {
				m[i1][i2] = max(m[i1][i2-1], m[i1-1][i2])
			}
		}
	}
	return m[n1][n2]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
