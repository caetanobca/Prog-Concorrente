package main

import (
	"fmt"
	"time"
)

func main() {
	nServers := 5
	timeout := 6
	bidChan := handle(nServers, timeout)

	for bid := range bidChan {
		fmt.Println("Bid received! ", bid)
	}
}

// Função que retorna um canal de Itens a serem leiloados
func itemsStream() chan int {
	itensChan := make(chan int)
	go generateInput(itensChan)
	return itensChan
}

/* A função bid, internamente, executa uma requisição rest para o item
 * recebido como parâmetro. A resposta da requisição rest é retornada com o
 * lance obtido (bid) para o item
 */
func bid(item int) <-chan Bid {
	time.Sleep(3 * time.Second)
	chBid := make(chan Bid)
	chBid <- Bid{item, 3, false}
	return chBid
}

type Bid struct {
	item      int
	bidValue  int
	bidFailed bool
}

func handle(nServers int, timeout int) <-chan Bid {

	bidChan := make(chan Bid, nServers)
	joinCh := make(chan int, nServers)
	itensChan := itemsStream()

	for i := 0; i < nServers; i++ {
		go func() {
			for item := range itensChan {
				timer := time.Tick(time.Second * 1)
				var result Bid
				chBid := bid(item)
				select {
				case result = <-chBid:
					fmt.Print("oi")
				case <-timer:
					// result = Bid{item, -1, true}
					fmt.Print("opa")
				default:
					fmt.Print("ola")
				}
				bidChan <- result
			}
			joinCh <- 1
		}()

	}

	go func() {
		for i := 0; i < nServers; i++ {
			<-joinCh
		}
		close(bidChan)
	}()

	return bidChan
}

func generateInput(chanInt chan<- int) {
	for i := 0; i < 15; i++ {
		chanInt <- i
	}
	close(chanInt)
}
