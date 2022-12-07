package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}
func reverseStr(s string, k int) string {
	sl := []byte(s)
	for l := 0; ; l += 2 * k {
		r := l + k - 1
		if r >= len(sl) {
			r = len(sl) - 1
		}
		for i := l; i < r; i, r = i+1, r-1 {
			sl[i], sl[r] = sl[r], sl[i]
		}
		if (l + 2*k) >= len(sl) {
			break
		}
	}
	return string(sl)
}
