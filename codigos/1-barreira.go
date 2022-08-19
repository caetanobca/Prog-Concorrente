package main

//Barreira
type Barrier struct {
	size  int
	s0    chan int
	s1    chan int
	mutex chan int
}
