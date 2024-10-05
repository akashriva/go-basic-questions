package main

import (
	"fmt"
)

func main() {

	arr := [7]int{1, 2, 3, 4, 6, 7, 8}
	fmt.Println("missing number in given array", findMissingNo(arr))

}

func findMissingNo(arr [7]int) int {
	lastEle := arr[len(arr)-1]

	total := lastEle * (lastEle + 1) / 2

	calGivenArr := 0
	for i := 0; i < len(arr); i++ {
		calGivenArr += arr[i]
	}
	missingEle := total - calGivenArr

	return missingEle
}
