package main

import (
	"fmt"
)

func main() {
	fmt.Println(modifyString("????v??"))
}
func modifyString(s string) string {
	sl := []rune(s)
	for i := 0; i < len(sl); i++ {
		if sl[i] == '?' {
			for _, r := range "abcdeifghjklmnopqustwzy" {
				if (i == 0 || sl[i-1] != r) && (i+1 == len(sl) || sl[i+1] != r) {
					sl[i] = r
					break
				}
			}
		}
	}
	return string(sl)
}
