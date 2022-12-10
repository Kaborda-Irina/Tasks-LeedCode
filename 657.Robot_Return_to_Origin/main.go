package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
}

func judgeCircle(moves string) bool {
	b := []byte(moves)
	u, d, l, r := 0, 0, 0, 0
	for _, v := range b {
		if v == 'U' {
			u++
		}
		if v == 'D' {
			d++
		}
		if v == 'L' {
			l++
		}
		if v == 'R' {
			r++
		}
	}
	return u == d && r == l
}
