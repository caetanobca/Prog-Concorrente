package main

import (
	"fmt"
	"time"
)

// Função que retorna um canal de Itens a serem leiloados
func itemsStream() chan int {
	itensChan := make(chan int)
	return itensChan
}

/* A função bid, internamente, executa uma requisição rest para o item
 * recebido como parâmetro. A resposta da requisição rest é retornada com o
 * lance obtido (bid) para o item
 */

// // Com timeout
// func bid(item int) <-chan Bid {
// 	time.Sleep(3 * time.Second)
// 	chBid := make(chan Bid)
// 	chBid <- Bid{item, 3, false}
// 	return chBid
// }

// Sem timeout
func bid(item int) Bid {
	time.Sleep(3 * time.Second)
	return Bid{item, 3, false}
}

type Bid struct {
	item      int
	bidValue  int
	bidFailed bool
}

func main() {
	idChan := make(chan int)
	go generateInput(idChan)

	nServers := 5
	// timeout := 6
	bidChan := handle(nServers, idChan)

	for bid := range bidChan {
		fmt.Println("Bid received! ", bid)
	}
}

// // Com timeout (incompleta)
// func handle(nServers int, itensChan <-chan int, timeout int) <-chan Bid {

// 	bidChan := make(chan Bid, nServers)
// 	joinCh := make(chan int, nServers)

// 	for i := 0; i < nServers; i++ {
// 		go func() {
// 			for item := range itensChan {
// 				timer := time.Tick(time.Second * 1)
// 				var result Bid
// 				chBid := bid(item)
// 				select {
// 				case result = <-chBid:
// 					fmt.Print("oi")
// 				case <-timer:
// 					// result = Bid{item, -1, true}
// 					fmt.Print("opa")
// 				default:
// 					fmt.Print("ola")
// 				}
// 				bidChan <- result
// 			}
// 			joinCh <- 1
// 		}()

// 	}

// 	go func() {
// 		for i := 0; i < nServers; i++ {
// 			<-joinCh
// 		}
// 		close(bidChan)
// 	}()

// 	return bidChan
// }

// Sem timeout
func handle(nServers int, itensChan <-chan int) <-chan Bid {

	bidChan := make(chan Bid, nServers)
	joinCh := make(chan int, nServers)

	for i := 0; i < nServers; i++ {
		go func() {
			for item := range itensChan {
				bidChan <- bid(item)
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
