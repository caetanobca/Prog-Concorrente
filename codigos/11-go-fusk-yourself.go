package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func find(root string, filesChan chan string, join chan int) {
	find_aux(root, filesChan)
	close(filesChan)
	join <- 1

}

func find_aux(root string, filesChan chan string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		addr := root + "/" + f.Name()
		if f.IsDir() {
			find_aux(addr, filesChan)
		} else {
			filesChan <- addr
		}
	}
}

func readFiles(filesChan chan string) {
	for file := range filesChan {
		fmt.Println(file)
	}
}

func main() {
	var root string
	fmt.Scan(&root)

	filesChan := make(chan string)
	joinchan := make(chan int)

	find(root, filesChan, joinchan)

	<-joinchan
}
