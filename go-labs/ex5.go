package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func goroutineFunc(id int, done chan struct{}) {
	rand.Seed(time.Now().UnixNano())
	sleepTime := time.Duration(rand.Intn(5000)) * time.Millisecond
	fmt.Printf("Goroutine %d: Dormindo por %v\n", id, sleepTime)
	time.Sleep(sleepTime)
	fmt.Printf("Goroutine %d: Acordou!\n", id)
	done <- struct{}{}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: ./programa <n>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Erro ao converter o argumento para um número inteiro:", err)
		os.Exit(1)
	}

	done := make(chan struct{})

	for i := 1; i <= n; i++ {
		go goroutineFunc(i, done)
	}

	for i := 1; i <= n; i++ {
		<-done
	}

	fmt.Printf("Todas as %d goroutines terminaram.\n", n)
}
