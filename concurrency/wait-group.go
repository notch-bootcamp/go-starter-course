package main

import (
	"fmt"
	"sync"
)

var (
	wg               sync.WaitGroup
	waitGroupCounter int
	waitGroupMutex   sync.Mutex
)

func simpleIncrement() {
	defer wg.Done()
	waitGroupMutex.Lock()
	defer waitGroupMutex.Unlock()
	fmt.Println("Acquired the mutex for write")
	waitGroupCounter++
	fmt.Printf("Counter value was incremented: %d\n", waitGroupCounter)
}

func simpleAccess() {
	defer wg.Done()
	waitGroupMutex.Lock()
	defer waitGroupMutex.Unlock()
	fmt.Println("Acquired the mutex for read")
	fmt.Printf("Counter value: %d\n", waitGroupCounter)
}

func main() {
	wg.Add(2)
	go simpleIncrement()
	go simpleAccess()
	wg.Wait()
}
