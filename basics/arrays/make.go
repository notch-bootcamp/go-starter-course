package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Printf("l=%d, c=%d, %v\n", len(s), cap(s), s)
	s = make([]int, 5, 100)
	fmt.Printf("l=%d, c=%d, %v\n", len(s), cap(s), s)

	firstFibSlice := []int{1, 1, 2, 3, 5}
	fmt.Printf("l=%d, c=%d, %v\n", len(firstFibSlice), cap(firstFibSlice), firstFibSlice)
	firstFibSlice = append(firstFibSlice, 8, 13, 21)
	fmt.Printf("l=%d, c=%d, %v\n", len(firstFibSlice), cap(firstFibSlice), firstFibSlice)
}
