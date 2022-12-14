package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
}

func reverseOnlyLetters(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		for i < j && !(unicode.IsLetter(runes[i])) {
			i++
		}
		for i < j && !(unicode.IsLetter(runes[j])) {
			j--
		}

		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)

}
