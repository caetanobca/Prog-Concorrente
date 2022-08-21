package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Request struct {
	id   int
	size int
}

func limitCap_wait(req Request, bucket chan int, mutex chan int) {
	size := req.size
	for {
		mutex <- 1

		if size <= len(bucket) {
			for i := 0; i < size; i++ {
				token := <-bucket
				fmt.Printf("Req {%d} Token removido: %d\n", req.id, token)
			}
			<-mutex
			break
		}
		<-mutex
	}
	//if size <= len(bucket) {
	//	for i := 0; i < size; i++ {
	//		token := <-bucket
	//		fmt.Printf("Req {%d} Token removido: %d\n", req.id, token)
	//	}
	//} else {
	//	//time.Sleep(10)
	//	limitCap_wait(req, bucket)
	//}

}

func run(req Request, bucket chan int, mutex chan int) {
	limitCap_wait(req, bucket, mutex)
}

func fill(bucket chan int, freq int) {
	for {
		time.Sleep(time.Duration(freq) * time.Second)
		if len(bucket) < cap(bucket) {
			token := rand.Intn(9900) + 100
			bucket <- token
		}
	}
}

func main() {
	var bs int
	var freq_s int
	fmt.Println("Insira tamanho do bucket")
	fmt.Scan(&bs)
	fmt.Println("Insira frequência de criação dos tokens por segundo")
	fmt.Scan(&freq_s)

	joinch := make(chan int)
	mutex := make(chan int, 1)

	freq := 1 / freq_s
	bucket := make(chan int, bs)

	go fill(bucket, freq)

	for i := 0; i < 10; i++ {

		go func(i int) {

			id := (i + 1) * 100
			for {
				size := rand.Intn(10)
				id++
				req := Request{id: id, size: size}

				run(req, bucket, mutex)

			}
		}(i)
	}

	<-joinch

}
