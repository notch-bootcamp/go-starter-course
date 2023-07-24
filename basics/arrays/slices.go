package main

import "fmt"

func main() {
	s := []int{1, 1, 2, 3, 5}
	fmt.Printf("len=%d c=%d %v\n", len(s), cap(s), s)
	s = s[:0]
	fmt.Printf("len=%d c=%d %v\n", len(s), cap(s), s)
	s = s[:4]
	fmt.Printf("len=%d c=%d %v\n", len(s), cap(s), s)
	s = s[2:]
	fmt.Printf("len=%d c=%d %v\n", len(s), cap(s), s)
}
