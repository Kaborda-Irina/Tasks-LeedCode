package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
func lengthOfLastWord(s string) int {

	new := strings.Split(strings.TrimSpace(s), " ")
	return len(new[len(new)-1])
}
