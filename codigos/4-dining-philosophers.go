package main

// Como evitar stravation?
func main() {
	var forks [5]chan int
	for i := range forks {
		forks[i] = make(chan int, 1)
	}
}

func getFork(philosopher int, forks []chan int) {
	lFork := (philosopher + 1) % 5
	rFork := philosopher

	forks[lFork] <- 0
	forks[rFork] <- 0
}

func putFork(philosopher int, forks []chan int) {
	lFork := (philosopher + 1) % 5
	rFork := philosopher

	<-forks[lFork]
	<-forks[rFork]
}
