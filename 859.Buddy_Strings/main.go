package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("abcd", "cbad"))
	fmt.Println(buddyStrings("abbc", "baac"))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == goal[j] && s[j] == goal[i] {
				sl := []byte(s)
				sl[i], sl[j] = sl[j], sl[i]
				if string(sl) == goal {
					return true
				}
			}
		}
	}

	return false
}
