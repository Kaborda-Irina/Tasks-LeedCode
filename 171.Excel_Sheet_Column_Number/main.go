package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("AB"))
}

func titleToNumber(columnTitle string) int {

	res := 0
	f := 0.0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += (int(columnTitle[i]) - int('A') + 1) * int(math.Pow(26, f))
		f++
	}

	return res
}
