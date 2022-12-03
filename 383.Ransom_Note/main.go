package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "ab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, r := range magazine {
		m[r]++
	}
	for _, r := range ransomNote {
		if val, ok := m[r]; !ok || val == 0 {
			return false
		}
		m[r]--
	}
	return true
}
