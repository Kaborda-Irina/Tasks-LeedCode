package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTheDifference("b", "bb"))
}

func findTheDifference(s string, t string) byte {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for i, v := range t {
		if _, ok := m[v]; !ok {
			return t[i]
		} else {
			m[v]--
			if m[v] < 0 {
				return t[i]
			}
		}
	}
	return '0'
}
