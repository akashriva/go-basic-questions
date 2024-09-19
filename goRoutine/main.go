package main

import (
	"fmt"
	"time"
)

func number(nums *[]int) {
	for i := 1; i <= 5; i++ {
		*nums = append(*nums, i) // Append i to the shared slice
	}
}

func main() {
	start := time.Now()
	fmt.Println(start)
	nums := []int{}
	go number(&nums)
	time.Sleep(1 * time.Second)
	for _, value := range nums {
		fmt.Println(value)
	}
	end := time.Since(start)

	fmt.Println("comleted time :--", end)
}
