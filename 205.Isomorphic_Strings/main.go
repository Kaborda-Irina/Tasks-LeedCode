package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}

func isIsomorphic(s string, t string) bool {
	first := make(map[byte]int)
	second := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := first[s[i]]; !ok {
			first[s[i]] = i
		}
		if _, ok := second[t[i]]; !ok {
			second[t[i]] = i
		}
		if first[s[i]] != second[t[i]] {
			return false
		}
	}

	return true
}
