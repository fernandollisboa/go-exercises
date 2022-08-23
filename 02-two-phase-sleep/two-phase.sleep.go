package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)

	routinesCh := make(chan int)
	roundTwoCh := make(chan int, n-1)
	roundThreeCh := make(chan int, n-1)

	for i := 0; i < n; i++ {
		go func(i int) {
			// seconds := rand.Intn(6)
			// time.Sleep(time.Duration(seconds) * time.Second)
			seconds := rand.Intn(6)
			time.Sleep(time.Duration(seconds) * time.Second)
			routinesCh <- seconds
			fmt.Println("Rotina", i, "dormiu", seconds, "segundos")
		}(i)
	}

	for i := 0; i < n; i++ {
		x := <-routinesCh
		fmt.Println("Segunda rodada: ")
		roundTwoCh <- x
	}

	lastTime := <-routinesCh

	fmt.Println("Segunda rodada: ")

	go func(seconds int) {
		time.Sleep(time.Duration(seconds) * time.Second)
		roundThreeCh <- 1
		fmt.Println("Rotina  0 dormiu", seconds, "segundos")
	}(lastTime)

	for i := 1; i < n; i++ {
		nextTime := <-roundTwoCh
		go func(seconds int, i int) {
			time.Sleep(time.Duration(seconds) * time.Second)
			roundThreeCh <- 1
			fmt.Println("Rotina", i, "dormiu", seconds, "segundos")
		}(nextTime, i)
	}

	for i := 0; i < n; i++ {
		<-roundThreeCh
	}
}
