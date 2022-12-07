package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("deeee"))
	fmt.Println(validPalindrome("abcba"))
}
func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return isValidPalindrome(s, i+1, j) || isValidPalindrome(s, i, j-1)
		}
		i++
		j--
	}
	return true
}

func isValidPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true

}
