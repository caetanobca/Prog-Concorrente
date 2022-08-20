package main

import (
	"fmt"
	"time"
)

func main() {
	nServers := 5
	timeout := 1
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
	chBid := make(chan Bid)

	go func() {
		defer close(chBid)
		time.Sleep(3 * time.Second)
		chBid <- Bid{item, 3, false}
	}()

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

				timer := time.Tick(time.Second * time.Duration(timeout))
				var result Bid

				chBid := bid(item)
				select {
				case result = <-chBid:
				case <-timer:
					result = Bid{item, -1, true}
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
