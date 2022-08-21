package main

import (
	"fmt"
	//"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)

	routinesCh := make(chan int)
	for i := 0; i < n; i++ {
		go func() {
			//seconds := rand.Intn(5)
			time.Sleep(time.Duration(1) * time.Second)
			routinesCh <- 1
		}()
	}

	for j := 0; j < n; j++ {
		<-routinesCh
	}
	close(routinesCh)
	fmt.Println(n)

}
