package main

import (
	"fmt"
	"time"
)

func numbers(num chan int) {
	for i := 1; i <= 5; i++ {
		num <- i
	}
	close(num)
}

func main() {
	strat := time.Now()
	ch := make(chan int)
	go numbers(ch)
	for num := range ch {
		fmt.Println(num)
	}
	end := time.Since(strat)
	fmt.Println(end)
}
