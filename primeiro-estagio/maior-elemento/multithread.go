package main

import (
	"fmt"
	"sync"
	"time"
)

const ARRAY_SIZE = 9999999

func main() {
	elements := get_elements(105)

	start := time.Now()
	biggest := get_biggest(elements)
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)

	fmt.Print("Position: ")
	fmt.Println(int(biggest))
	fmt.Print("Value: ")
	fmt.Println(elements[biggest])
}

func get_biggest(elements [ARRAY_SIZE]int) int {
	var biggest [2]int
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	get_biggest_aux(&elements, 0, ARRAY_SIZE, &biggest, 0, wg)

	return biggest[0]
}

func get_biggest_aux(elements *[ARRAY_SIZE]int, begin int, end int, biggest *[2]int, position int, wg *sync.WaitGroup) {
	if end-begin < 100 {
		biggest_local := begin
		for i := begin + 1; i < end; i++ {
			if elements[i] > elements[biggest_local] {
				biggest_local = i
			}
		}
		biggest[position] = biggest_local
	} else {
		var new_wg = &sync.WaitGroup{}

		var biggest_local [2]int

		middle := (begin + end) / 2

		new_wg.Add(2)
		go get_biggest_aux(elements, begin, middle, &biggest_local, 0, new_wg)
		go get_biggest_aux(elements, middle+1, end, &biggest_local, 1, new_wg)
		new_wg.Wait()

		if elements[biggest_local[0]] > elements[biggest_local[1]] {
			biggest[position] = biggest_local[0]
		} else {
			biggest[position] = biggest_local[1]
		}
	}
	wg.Done()
}

func get_elements(biggest int) [ARRAY_SIZE]int {
	var elements [ARRAY_SIZE]int

	for i := 0; i < ARRAY_SIZE; i++ {
		if i == biggest {
			elements[i] = 100
		} else {
			elements[i] = i % 11
		}
	}
	return elements
}
