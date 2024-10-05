package main

import (
	"fmt"
	"sync"
)

func writeSomeThing(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("write.....")
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func ReadUsingChannel(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Read....", num)
	}
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(2)

	go writeSomeThing(ch, &wg)
	go ReadUsingChannel(ch, &wg)

	wg.Wait()

}
