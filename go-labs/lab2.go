package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type gerador func() int

func producer(canal chan<- int, gen gerador, wg *sync.WaitGroup) {
	defer wg.Done()

	for true {
		canal <- gen()
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func consumer(p <-chan int, i <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	tick := time.Tick(1 * time.Second)
	for true {
		select {
		case v := <-p:
			fmt.Printf("consumiu %d\n", v)
		case v := <-i:
			fmt.Printf("consumiu %d\n", v)
		case <-tick:
			fmt.Println("tick")
		}
	}
}

func main() {
	pares := make(chan int)
	impares := make(chan int)
	var wg sync.WaitGroup

	wg.Add(3)

	go producer(pares, func() int {
		num := rand.Intn(10)
		if num%2 != 0 {
			num += 1
		}
		fmt.Printf("Par produziu: %d\n", num)
		return num
	}, &wg)

	go producer(impares, func() int {
		num := rand.Intn(10)
		if num%2 == 0 {
			num += 1
		}
		fmt.Printf("Impar produziu: %d\n", num)
		return num
	}, &wg)

	go consumer(pares, impares, &wg)

	wg.Wait()
}
