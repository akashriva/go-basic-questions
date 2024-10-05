package main

import (
	"fmt"
	"sync"
)

func goRoutineA(chanA, chanB chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {

		<-chanA // Wait for signal to start
		fmt.Print("hello\n")
		chanB <- true // Signal goRoutineB to proceed
	}
}

func goRoutineB(chanB, chanC chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {

		<-chanB // Wait for goRoutineA
		fmt.Print("sachin\n")
		chanC <- true // Signal goRoutineC to proceed
	}
}

func goRoutineC(chanC, chanA chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {

		<-chanC // Wait for goRoutineB
		fmt.Print("bhai\n")
		// chanA <- true
		// Final routine; no need to signal chanA

	}
}

func main() {
	var wg sync.WaitGroup

	chanA := make(chan bool)
	chanB := make(chan bool)
	chanC := make(chan bool)

	wg.Add(3)

	go goRoutineA(chanA, chanB, &wg)
	go goRoutineB(chanB, chanC, &wg)
	go goRoutineC(chanC, chanA, &wg)

	// Start the sequence by sending a signal to chanA
	chanA <- true

	// Wait for all goroutines to complete
	wg.Wait()

	// No need to close channels explicitly
	close(chanA)
	close(chanB)
	close(chanC)
}
