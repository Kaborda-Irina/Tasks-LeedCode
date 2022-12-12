package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}

func removeOuterParentheses(s string) string {
	res := ""
	open := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if open > 0 {
				res += "("
			}
			open++
		} else if s[i] == ')' {
			if open > 1 {
				res += ")"
			}
			open--
		}
	}
	return res
}
