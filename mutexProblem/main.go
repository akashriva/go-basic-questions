package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (a *BankAccount) deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
	fmt.Printf("Deposited: %d, Balance: %d\n", amount, a.balance)

}
func (a *BankAccount) withdraw(amount int, wg *sync.WaitGroup) {

	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Println(a.balance >= amount)
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrawal amount %d , balance left %d\n", amount, a.balance)
	} else {
		fmt.Println("Insufficient balance")
	}
}

func main() {
	account := BankAccount{balance: 0}
	var wg sync.WaitGroup

	wg.Add(3)
	go account.deposit(10000, &wg)
	go account.withdraw(2000, &wg)
	go account.withdraw(10000, &wg)

	wg.Wait()
	fmt.Println("Final Balance:", account.balance)
}
