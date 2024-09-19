package main

import (
	"fmt"
	"sync"
)

type Auction struct {
	highestBid int
	mu         sync.Mutex
}

func (a *Auction) PlaceBid(bid int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.mu.Lock()
	defer a.mu.Unlock()

	if bid > a.highestBid {
		a.highestBid = bid
		fmt.Printf("New highest bid: %d\n", bid)
	} else {
		fmt.Printf("Bid of %d is lower than the current highest bid of %d\n", bid, a.highestBid)
	}
}

func main() {
	auction := &Auction{highestBid: 0}
	var wg sync.WaitGroup

	bids := []int{100, 150, 200, 180}
	wg.Add(len(bids))

	for _, bid := range bids {
		go auction.PlaceBid(bid, &wg)
	}

	wg.Wait()
	fmt.Println("Auction complete. Highest bid:", auction.highestBid)
}
