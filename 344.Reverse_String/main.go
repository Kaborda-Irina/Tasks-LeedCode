package main

import "fmt"

func main() {
	fmt.Println(reverseString([]byte{'h', 'e', 'l', 'l', 'o'}))
}
func reverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}
