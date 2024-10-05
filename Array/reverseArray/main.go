package main

import (
	"fmt"
)

func main(){
	arr := []int{1,2,3,4,5}
	fmt.Println(reverseArray(arr))
}


func reverseArray(arr []int) []int {
	start , end := 0 ,len(arr)-1

	for start < end {
		arr[start] ,arr[end] = arr[end], arr[start]
		start ++
		end --
	}
	return arr
}