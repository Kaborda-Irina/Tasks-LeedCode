package main

import "fmt"

func main() {
	fmt.Println(isAnagram("aacc", "ccac"))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[string]int{}

	for _, r := range s {
		m[string(r)]++
	}

	for _, r := range t {
		if _, ok := m[string(r)]; !ok {
			return false
		}
		if m[string(r)] == 0 {
			return false
		}
		m[string(r)]--
	}

	return true
}
