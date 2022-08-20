package main

type LightSwitch struct {
	count int
	mutex chan int
}

func LightSwitch_init() LightSwitch {
	return LightSwitch{count: 0, mutex: make(chan int, 1)}
}

func (ls LightSwitch) lock(s chan int) {
	ls.mutex <- 0
	ls.count += 1
	if ls.count == 1 {
		s <- 0
	}
	<-ls.mutex
}

func (ls LightSwitch) unlock(s chan int) {
	ls.mutex <- 0
	ls.count -= 1
	if ls.count == 0 {
		<-s
	}
	<-ls.mutex
}

func main() {

	a := make(chan int, 1)
	b := make(chan int, 1)
	ls := LightSwitch_init()
	go ls.lock(a)
	go ls.lock(b)
}
