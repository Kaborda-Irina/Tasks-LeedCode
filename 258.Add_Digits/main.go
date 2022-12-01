package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
}
func addDigits(num int) int {
	if num >= 0 && num <= 9 {
		return num
	}
	result := 0
	for num > 0 {

		result = result + num%10
		num /= 10
	}
	return addDigits(result)
}
