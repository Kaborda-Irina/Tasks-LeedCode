package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; count < n && i < len(flowerbed); i++ {
		if flowerbed[i] != 0 {
			continue
		}
		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}
		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			continue
		}

		flowerbed[i] = 1
		count++
		if i < len(flowerbed)-1 {
			flowerbed[i+1] = -1
		}
	}
	return count >= n
}
