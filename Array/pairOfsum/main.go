package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 5, 7, -1, 5}

	target := 6

	pairOfSum(arr, target)

}

func pairOfSum(arr []int, target int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				fmt.Printf("Pair of sum: (%v, %v)\n", arr[i], arr[j])
			}
		}
	}

}
