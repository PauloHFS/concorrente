package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produtor(nums chan<- int) {
	for {
		nums <- rand.Intn(10)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func consumidor(nums <-chan int) {
	for {
		select {
		case v := <-nums:
			if v%2 == 0 {
				fmt.Printf("Par %d\n", v)
			} else {
				fmt.Printf("Impar %d\n", v)
			}
		}
	}
}

func main() {
	nums := make(chan int)

	go produtor(nums)
	go consumidor(nums)

	time.Sleep(60 * time.Second)
}
