package main

import (
	"fmt"
	"sync"
)

type Driver struct {
	name     string
	assigned bool
}

var (
	drivers = []Driver{
		{name: "sachin"},
		{name: "Ramesh"},
		{name: "yashpal"},
	}

	mu sync.Mutex
)

func assignedDriver(customer string, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	for i := range drivers {
		if !drivers[i].assigned {
			drivers[i].assigned = true
			fmt.Printf("%s is assigned to customer %s \n", drivers[i].name, customer)
			return
		}
	}
	fmt.Printf("No drivers available for customer %s\n ", customer)
}

func main() {
	var wg sync.WaitGroup

	customers := []string{"Rahul", "Raju", "Ratan", "Ram"}

	wg.Add(len(customers))

	for _, customer := range customers {
		go assignedDriver(customer, &wg)
	}
	wg.Wait()
	fmt.Println("assignment Completed !")
}
