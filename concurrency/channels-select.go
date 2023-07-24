package main

import (
	"fmt"
	"time"
)

func printNumbersWithChanel(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(750 * time.Millisecond)
	}
	close(ch)
}

func fibGenerator(ch chan int) {
	a, b := 0, 1
	for i := 0; i < 5; i++ {
		ch <- b
		a, b = b, a+b
		time.Sleep(750 * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	fib := make(chan int)
	go printNumbersWithChanel(ch)
	go fibGenerator(fib)
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				ch = nil
			} else {
				fmt.Println(num)
			}
		case f, ok := <-fib:
			if !ok {
				fib = nil
			} else {
				fmt.Println("Fib: ", f)
			}
		default:
			if ch == nil && fib == nil {
				fmt.Println("All channels are done...")
				return
			}
			//fmt.Println("Doing something default - some other tasks maybe...")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
