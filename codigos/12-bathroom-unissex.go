package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LightSwitch struct {
	count *int
	mutex chan int
}

func LightSwitch_init() LightSwitch {
	var c int
	c = 0

	return LightSwitch{count: &c, mutex: make(chan int, 1)}
}

func (ls LightSwitch) lock(s chan int) {
	ls.mutex <- 0
	*ls.count = *ls.count + 1

	if *ls.count == 1 {
		s <- 0
	}

	<-ls.mutex
}

func (ls LightSwitch) unlock(s chan int) {
	ls.mutex <- 0
	*ls.count--

	if *ls.count == 0 {
		<-s
	}

	<-ls.mutex
}
func useBathroom(ls LightSwitch, multiplex chan int, s chan int, n int, join chan int) {
	fmt.Printf("entrou\n")
	ls.lock(s)
	multiplex <- 1
	fmt.Printf("usando o banheiro %d \n", n)
	number := rand.Intn(5)
	time.Sleep(time.Duration(number) * time.Second)
	<-multiplex
	ls.unlock(s)
	join <- 1
}

func main() {
	var n int
	fmt.Scan(&n)

	s := make(chan int, 1)
	join := make(chan int)

	maleSwitch := LightSwitch_init()
	femaleSwitch := LightSwitch_init()

	maleMultiplex := make(chan int, 3)
	femaleMultiplex := make(chan int, 3)

	defer close(s)
	defer close(maleMultiplex)
	defer close(femaleMultiplex)
	defer close(join)
	defer close(maleSwitch.mutex)
	defer close(femaleSwitch.mutex)

	for i := 0; i < n; i++ {
		number := rand.Intn(1000)
		if number%2 == 0 {
			fmt.Printf(" sorteado %d 0\n", number)
			go useBathroom(maleSwitch, maleMultiplex, s, number, join)
		} else {
			fmt.Printf(" sorteado %d 1\n", number)
			go useBathroom(femaleSwitch, femaleMultiplex, s, number, join)
		}
	}

	for i := 0; i < n; i++ {
		<-join
	}

}
