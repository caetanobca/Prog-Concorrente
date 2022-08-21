package main

import "math/rand"

var n_chann chan int
var mutex chan int
var mutex_0 chan int
var mutex_1 chan int
var count_0 int
var count_1 int

func main() {
	n := 5

	n_chann = make(chan int, n)
	mutex = make(chan int, 1)
	mutex_0 = make(chan int, 1)
	count_0 = 0
	mutex_1 = make(chan int, 1)
	count_1 = 0

	req_chan := req_maker()

	for req := range req_chan {
		go handle(req)
	}
}
func handle(req Request) {
	if req.req_type == 0 {
		handle_aux0(req)
	} else {
		handle_aux1(req)
	}
	<-n_chann
}
func handle_aux0(req Request) {
	mutex <- 0
	mutex_1 <- 0
	if count_1 == 0 {
		mutex_0 <- 0
		count_0++
		n_chann <- 0
		<-mutex_0
	}
	<-mutex_1
	<-mutex
	exec(req)
}

func handle_aux1(req Request) {
	mutex <- 0
	mutex_0 <- 0
	if count_0 == 0 {
		mutex_1 <- 0
		count_1++
		n_chann <- 0
		<-mutex_1
	}
	<-mutex_0
	<-mutex
	exec(req)
}

type Request struct {
	req_type int
}

func exec(req Request) {}

func req_maker() <-chan Request {
	req_chan := make(chan Request)
	go func() {
		defer close(req_chan)
		for {
			n := rand.Intn(1)
			req_chan <- Request{req_type: n}
		}
	}()
	return req_chan
}
