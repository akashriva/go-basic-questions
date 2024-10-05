package main

import (
	"fmt"
)

func firstRepeatElement(arr []int) int {
	elementMap := make(map[int]bool)

	for _, value := range arr {
		if elementMap[value] {
			return value
		}
		elementMap[value] = true
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 3, 5, 3}
	result := firstRepeatElement(arr)
	if result != -1 {
		fmt.Println("First Repeated element is :", result)
	} else {
		fmt.Println("num not repeated")
	}
}
