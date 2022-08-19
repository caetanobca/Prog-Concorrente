/*
 * Fork-sleep-join (3.0) - Crie um programa que recebe um número inteiro n como argumento e cria n goroutines.
 * Cada uma dessas goroutines deve dormir por um tempo aleatório de no máximo 5 segundos. A main-goroutine deve
 * esperar todas as goroutines filhas terminarem de executar para em seguida escrever na saída padrão o valor de n.
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)

	joinChan := make(chan int)
	defer close(joinChan)
	for i := 0; i < n; i++ {
		go func() {
			st := rand.Intn(5)
			time.Sleep(time.Second * time.Duration(st))
			joinChan <- 0
		}()
	}
	for i := 0; i < n; i++ {
		<-joinChan
	}

	fmt.Println(n)
}
