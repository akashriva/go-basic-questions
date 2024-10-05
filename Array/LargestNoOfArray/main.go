package main

import (
	"fmt"
)

func main() {
	arr := [6]int{9, 5, 10, 67, 8, 1}
	fmt.Println("Largest Number in given Array", largestNumber(arr))

}

func largestNumber(arr [6]int) int {
	if len(arr) < 1 {
		return -1
	}
	largest := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	return largest
}
