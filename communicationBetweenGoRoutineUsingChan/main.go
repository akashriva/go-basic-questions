package main

import (
	"fmt"
	"sync"
)

// Goroutine A: Sends initial data
func goroutineA(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i // Send data to channel
	}
	close(ch) // Close channel when done sending
}

// Goroutine B: Receives data from A, processes it, and sends to C
func goroutineB(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fmt.Println("in go routine B------", num)
		processed := num * 2 // Example processing
		out <- processed
	}
	close(out) // Close channel when done sending
}

// Goroutine C: Receives data from B, processes it, and sends to D
func goroutineC(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in { 
		processed := num + 10 // Example processing
		out <- processed
	}
	close(out) // Close channel when done sending
}

// Goroutine D: Receives data from C and performs the final operation
func goroutineD(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fmt.Println("Final output:", num)
	}
}

func main() {
	var wg sync.WaitGroup

	// Channels for communication
	chAtoB := make(chan int)
	chBtoC := make(chan int)
	chCtoD := make(chan int)

	// Add goroutines to WaitGroup
	wg.Add(4)

	// Start goroutines
	go goroutineA(chAtoB, &wg)
	go goroutineB(chAtoB, chBtoC, &wg)
	go goroutineC(chBtoC, chCtoD, &wg)
	go goroutineD(chCtoD, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}
