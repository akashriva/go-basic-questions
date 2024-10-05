package main

import (
	"fmt"
	"sync"
	"time"
)

func numbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	start := time.Now()
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go numbers(ch, &wg)

	for num := range ch {
		fmt.Println(num)
	}
	wg.Wait()
	end := time.Since(start)
	fmt.Println(end)
}
