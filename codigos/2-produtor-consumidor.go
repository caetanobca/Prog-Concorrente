package main

import (
	"fmt"
	"math/rand"
)

//chan<- send only channel
func producer(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- rand.Intn(100)
	}
	close(ch)
}

//<-chan recieve only channel
func consumer(ch <-chan int, join chan<- int) {
	for i := range ch {
		if i%2 == 1 {
			fmt.Printf("Numero impar: %d \n", i)
		}
	}
	join <- 0
}

func main() {

	ch := make(chan int)
	join := make(chan int)

	go producer(ch)
	go consumer(ch, join)

	<-join
}
