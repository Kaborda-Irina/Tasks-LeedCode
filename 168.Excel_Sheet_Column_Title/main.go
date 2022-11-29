package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1))
}

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		columnNumber--
		res = string(byte(columnNumber%26)+'A') + res
		columnNumber /= 26
	}
	return res
}
