package main

import "fmt"

func main() {
	fmt.Println(romanToInt("LVIII"))
}

func romanToInt(s string) int {
	m := make(map[rune]int)
	{
		m['I'] = 1
		m['V'] = 5
		m['X'] = 10
		m['L'] = 50
		m['C'] = 100
		m['D'] = 500
		m['M'] = 1000
	}
	max := 0
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		num := m[rune(s[i])]

		if num >= max {
			max = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}
