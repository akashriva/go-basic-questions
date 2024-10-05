package main

import (
	"fmt"
	"sync"
)

func printHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

func printWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("World")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go printHello(&wg)
	go printWorld(&wg)

	wg.Wait()
}
