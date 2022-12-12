package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID"))
}
func diStringMatch(s string) []int {
	l := len(s)
	ret := make([]int, l+1)
	indexI, indexD := 0, len(s)
	for i := 0; i <= l; i++ {
		if i == l {
			ret[i] = indexI
			break
		}
		if s[i] == 'D' {
			ret[i] = indexD
			indexD--
		}
		if s[i] == 'I' {
			ret[i] = indexI
			indexI++
		}
	}
	return ret
}
