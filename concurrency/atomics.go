package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	atomicWg      sync.WaitGroup
	atomicCounter int64
)

func atomicIncrement() {
	defer atomicWg.Done()
	atomic.AddInt64(&atomicCounter, 1)
}

func atomicAccess() {
	defer atomicWg.Done()
	currentValue := atomic.LoadInt64(&atomicCounter)
	fmt.Printf("Current atomic value: %d\n", currentValue)
}

func main() {
	atomicWg.Add(2)
	go atomicIncrement()
	go atomicAccess()
	atomicWg.Wait()
}
