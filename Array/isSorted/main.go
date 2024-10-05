package main

import (
	"fmt"
)

func main() {
	// arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{2, 4, 1, 5, 7}

	result := isSorted(arr2)

	if result == true {
		fmt.Println("Array is sorted")
	} else {
		fmt.Println("Array is unsorted")
	}

}

func isSorted(arr []int) bool {
	for i := 0; i > len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
