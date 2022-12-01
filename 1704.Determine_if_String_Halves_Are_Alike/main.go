package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("book"))
}
func halvesAreAlike(s string) bool {
	a := s[:len(s)/2]
	b := s[len(s)/2:]

	resA := 0
	resB := 0
	for _, r := range a {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			resA++
		}
	}

	for _, r := range b {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			resB++
		}
	}

	if resA == resB {
		return true
	}
	return false
}
