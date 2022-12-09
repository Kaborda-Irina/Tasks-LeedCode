package main

import "fmt"

func main() {
	fmt.Println(isLongPressedName("saeed", "ssaaed"))
	fmt.Println(isLongPressedName("alex", "aaleex"))
}
func isLongPressedName(name string, typed string) bool {

	i := 0
	for j, _ := range typed {
		if i < len(name) && name[i] == typed[j] {
			i++
		} else if j == 0 || typed[j] != typed[j-1] {
			return false
		}
	}
	return i == len(name)
}
