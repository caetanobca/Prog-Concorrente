package main

import (
	"fmt"
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
	get_biggest_aux(&elements, 0, ARRAY_SIZE, &biggest, 0)
	return biggest[0]
}

func get_biggest_aux(elements *[ARRAY_SIZE]int, begin int, end int, biggest *[2]int, position int) {
	if end-begin < 100 {
		biggest_local := begin
		for i := begin + 1; i < end; i++ {
			if elements[i] > elements[biggest_local] {
				biggest_local = i
			}
		}
		biggest[position] = biggest_local
	} else {
		var biggest_local [2]int

		middle := (begin + end) / 2

		get_biggest_aux(elements, begin, middle, &biggest_local, 0)
		get_biggest_aux(elements, middle+1, end, &biggest_local, 1)

		if elements[biggest_local[0]] > elements[biggest_local[1]] {
			biggest[position] = biggest_local[0]
		} else {
			biggest[position] = biggest_local[1]
		}
	}
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
