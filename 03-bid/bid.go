package main

import (
	"fmt"
	"time"
)

func bid(item int) Bid {
	time.Sleep(3 * time.Second)
	return Bid{item, 3, false}
}

type Bid struct {
	item      int
	bidValue  int
	bidFailed bool
}

func generateInput(chanInt chan<- int) {
	for i := 0; i < 15; i++ {
		chanInt <- i
	}
	close(chanInt)
}

func itemsStream() chan int {
	itensChan := make(chan int)
	go generateInput(itensChan)
	return itensChan
}

func handle(nServers int) chan Bid {
	itemsCh := itemsStream()
	bidsCh := make(chan Bid)
	joinCh := make(chan int)

	for i := 0; i < nServers; i++ {
		go func() {
			for item := range itemsCh {
				bidsCh <- bid(item)
			}
			joinCh <- 1
		}()
	}

	go func() {
		for i := 0; i < nServers; i++ {
			<-joinCh
		}
		close(bidsCh)
		close(joinCh)
	}()

	return bidsCh
}

func main() {
	var nServers int
	fmt.Scan(&nServers)
	bidCh := handle(nServers)

	for bid := range bidCh {
		fmt.Println(bid)
	}
}
