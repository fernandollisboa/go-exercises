package main

import (
	"fmt"
	"math/rand"
	"time"
)

// produtor cria 5 numero inteiros n aleatorios entre 0 e 6 e depois dorme por n segundos

func producer(ch chan<- int) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		num := rand.Intn(7)
		fmt.Println(num)
		ch <- num

	}
	close(ch)
}

func doSomething(num int) Tuple {
	time.Sleep(time.Duration(num) * time.Second)
	return Tuple{num, num + 1}
}

type Tuple struct {
	original int
	newVal   int
}

// consumidor recebe de produtor
// faz duas tentativas doSomething de 2.5 segundos
// se acertar responde com {num, num + 1}
// em caso de timout responde com {num,-1}
func consumer(ch <-chan int, join chan int) chan Tuple {
	returnedTupleCh := make(chan Tuple)

	go func() {
		for num := range ch {
			auxCh := make(chan Tuple)

			go func() {
				auxCh <- doSomething(num)
			}()

			tick := time.Tick(time.Duration(5/2) * time.Second)
			var returnedTuple Tuple

		loop:
			for try := 0; try < 3; try++ {
				select {
				case returnedTuple = <-auxCh:
					break loop
				case <-tick:
					if try == 2 {
						returnedTuple = Tuple{num, -1}
					}
				}
			}
			returnedTupleCh <- returnedTuple
		}
		close(returnedTupleCh)
		join <- 0
	}()

	return returnedTupleCh
}

func main() {
	numbersCh := make(chan int, 5)
	join := make(chan int)

	go producer(numbersCh)
	outCh := consumer(numbersCh, join)

	for item := range outCh {
		fmt.Println(item)
	}

	fmt.Println("Cabou a execução")
	<-join
}
