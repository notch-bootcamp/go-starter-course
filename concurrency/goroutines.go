package main

import (
	"fmt"
	"time"
)

func print(x interface{}) {
	fmt.Println(x)
	time.Sleep(500 * time.Millisecond)
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		print(i)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		print(string(char))
	}
}

func main() {
	go printNumbers()
	printLetters()
}
