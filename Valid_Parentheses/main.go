package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	sl := []rune{}
	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, bracket := range s {
		if val, ok := brackets[bracket]; ok {
			sl = append(sl, val)
			continue
		}

		if len(sl)-1 < 0 || bracket != sl[len(sl)-1] {
			return false
		}

		sl = sl[0 : len(sl)-1]
	}

	return len(sl) == 0
}
