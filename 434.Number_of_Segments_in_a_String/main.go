package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSegments("      Hello,          my name is John"))
}
func countSegments(s string) int {
	segments := false
	count := 0
	for _, v := range s {
		if v == ' ' && segments {
			segments = false
			count += 1
		} else if v != ' ' {
			segments = true
		}
	}
	if segments {
		count++
	}
	return count
}
