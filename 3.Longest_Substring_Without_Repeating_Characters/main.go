package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	res := 0
	last_pos := -1
	for i := 0; i < len(s); i++ {
		if pos, ok := m[s[i]]; ok && last_pos < pos {
			last_pos = pos
		}
		res = max(res, i-last_pos)
		m[s[i]] = i
	}
	return res
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
