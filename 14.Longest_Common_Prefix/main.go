package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	elem := strs[0]

	for _, str := range strs {
		for !strings.HasPrefix(str, elem) {
			elem = elem[0 : len(elem)-1]
			if len(elem) == 0 {
				elem = ""
			}
		}
	}

	return elem
}
