package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Acquired the mutex for write")
		counter++
		fmt.Printf("Counter value was incremented: %d\n", counter)
		time.Sleep(10 * time.Millisecond)
		mutex.Unlock()
	}
}

func access() {
	time.Sleep(10 * time.Millisecond)
	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Acquired the mutex for read")
		fmt.Printf("Counter value: %d\n", counter)
		time.Sleep(10 * time.Millisecond)
		mutex.Unlock()
	}
}

func main() {
	go increment()
	go access()
	time.Sleep(1 * time.Second)
}
