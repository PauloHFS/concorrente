package main

import (
	"fmt"
	"time"
)

const MAX = 10

func produtor(canal chan int) {
	i := 1
	for {
		canal <- i
		fmt.Println("Enviou:", i)
		i++
	}
}

func main() {
	canal := make(chan int, MAX) // Canal com capacidade 10

	go produtor(canal)
	for {
		valor := <-canal
		fmt.Println("Recebeu:", valor)
		time.Sleep(1 * time.Second)
	}
}
