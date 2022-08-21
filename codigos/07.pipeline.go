/*
 * Pipeline (4.0) - Crie um programa organizado como um pipeline de goroutines. Esse programa deve receber como
 * argumento um caminho absoluto para um diretório. Uma goroutine deve navegar na árvore que tem como raiz o
 * diretório passado como argumento. Essa goroutine deve passar para uma próxima goroutine do pipeline o nome
 * dos arquivos encontrados na busca dos diretórios, ou seja, ignore os diretórios. Esta segunda goroutine deve
 * ler o primeiro byte de conteúdo de cada um desses arquivos e escrever na saída padrão o nome dos arquivos que
 * tem esse valor do primeiro byte sendo par.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func find(root string, filesChan chan string, cont int) {
	if cont == 0 {
		defer close(filesChan)
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		addr := root + "/" + f.Name()
		if f.IsDir() {
			find(addr, filesChan, cont+1)
		} else {
			filesChan <- addr
		}
	}
}

func readFiles(filesChan chan string, join chan int) {
	for file := range filesChan {
		fmt.Println(file)
	}
	join <- 1
}

func main() {
	var root string
	fmt.Scan(&root)

	filesChan := make(chan string)
	joinchan := make(chan int)

	go find(root, filesChan, 0)
	go readFiles(filesChan, joinchan)

	<-joinchan
}
