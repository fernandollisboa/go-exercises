package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bid(item int) Bid {
	secs := rand.Intn(10)
	time.Sleep(time.Duration(secs) * time.Second)
	fmt.Println("O item", item, "vai demorar", secs, "segundos")
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

func handle(nServers int, timeoutSeconds int) chan Bid {
	itemsCh := itemsStream()
	bidsCh := make(chan Bid)
	joinCh := make(chan int)

	for i := 0; i < nServers; i++ {
		go func() {
			for item := range itemsCh {
				timeCh := time.Tick(time.Duration(timeoutSeconds) * time.Second)

				auxChan := make(chan Bid)
				go func() {
					auxChan <- bid(item)
				}()

				var _bid Bid
				select {
				case <-timeCh:
					_bid = Bid{item, -1, true}
					fmt.Println("O item", item, "FALHOU")
				case _bid = <-auxChan:
					fmt.Println("O item", item, "processou")
				}
				bidsCh <- _bid
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
	nServers := 3
	timeoutSeconds := 3
	bidCh := handle(nServers, timeoutSeconds)

	for bid := range bidCh {
		fmt.Println(bid)
	}
}
