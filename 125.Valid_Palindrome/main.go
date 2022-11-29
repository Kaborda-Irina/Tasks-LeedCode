package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	newS := strings.TrimSpace(strings.ToLower(s))
	alpha := "abcdefghijklmnopqrstuvwxyz0123456789"
	resStr := ""
	palStr := ""
	for _, r := range newS {
		letter := string(r)
		if strings.Contains(alpha, letter) {
			resStr += letter
		}

	}
	for i := len(resStr) - 1; i < len(resStr) && i >= 0; i-- {
		palStr += string(resStr[i])
	}
	if resStr == palStr {
		return true
	}
	return false
}
