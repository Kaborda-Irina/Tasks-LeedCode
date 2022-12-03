package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
func wordPattern(pattern string, s string) bool {
	mP := make(map[rune]string)
	mS := make(map[string]rune)

	sl := strings.Split(s, " ")
	if len(pattern) != len(sl) {
		return false
	}

	for i, r := range pattern {
		if val, ok := mP[r]; ok {
			if sl[i] != val {
				return false
			}
		} else if val, ok := mS[sl[i]]; ok {
			if val != r {
				return false
			}
		} else {
			mP[r] = sl[i]
			mS[sl[i]] = r
		}
	}
	return true
}
