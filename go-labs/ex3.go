// Construa um pipeline em que uma goroutine gere strings aleatórias,
//enquanto uma segunda filtre as strings que contém somente valores alpha,
//e uma terceira escreva os valores filtrados na saída padrão;

package main

import (
	"fmt"
	"math/rand"
	"unicode"
)

func random_string() string {
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func validate_alpha_string(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func main() {
	random_strings := make(chan string)
	alpha_strings := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			random_strings <- random_string()
		}
		close(random_strings)
	}()

	go func() {
		for s := range random_strings {
			if validate_alpha_string(s) {
				alpha_strings <- s
			}
		}
		close(alpha_strings)
	}()

	for s := range alpha_strings {
		fmt.Println(s)
	}
}
