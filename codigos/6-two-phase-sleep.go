/*
 * two-phase sleep (3.0) - Crie um programa que recebe um número inteiro n como argumento e cria n goroutines.
 * Cada uma dessas goroutines deve dormir por um tempo aleatório de no máximo 5 segundos. Depois que  acordar,
 * cada thread deve sortear um outro número aleatório s (entre 0 e 10).  Somente depois de todas as n goroutines
 * terminarem suas escolhas (ou seja, ao fim da primeira fase), começamos a segunda fase. Nesta segunda fase,
 * a n-ésima goroutine criada deve dormir pelo tempo s escolhido pela goroutine n - 1 (faça a contagem de maneira modular,
 * ou seja, a primeira goroutine dorme conforme o número sorteado pela última).
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)

	joinChan := make(chan int)
	defer close(joinChan)

	var wg = sync.WaitGroup{}

	channels := make([]chan int, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan int, 1)
	}

	for i := 0; i < n; i++ {
		in_channel := ((i + n) % n)
		out_channel := ((i + 1) % n)
		wg.Add(1)
		go func_random(joinChan, channels[in_channel], channels[out_channel], &wg)
	}

	for i := 0; i < n; i++ {
		<-joinChan
	}

	fmt.Printf("No total foram executadas %d goroutines\n", n)
}

func func_random(joinChan chan int, inChan <-chan int, outChan chan<- int, wg *sync.WaitGroup) {

	st := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(st))
	out := rand.Intn(10)
	outChan <- out

	wg.Done()
	wg.Wait()

	in := <-inChan

	fmt.Printf("Esperando por %d segundos\n", in)
	time.Sleep(time.Second * time.Duration(in))

	joinChan <- 0
}
