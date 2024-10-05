package main

import (
	"fmt"
	"sort"
)

func maxPrduct(arr []int) int {
	sort.Ints(arr)

	n := len(arr)

	product1 := arr[n-1] * arr[n-2]
	product2 := arr[0] * arr[1]

	if product1 > product2 {
		return product1
	}

	return product2
}

func main() {
	arr := []int{-200, -100, 5, 7, 1, 56, 34}
	result := maxPrduct(arr)
	fmt.Println(result)
}
