package main

import "fmt"

func incrementer(initVal int, inc int) func() int {
	sum := initVal
	return func() int {
		sum += inc
		return sum
	}
}

func main() {
	evenInc := incrementer(0, 2)
	oddInc := incrementer(1, 2)
	fmt.Println(evenInc())
	fmt.Println(oddInc())
	fmt.Println(evenInc())
	fmt.Println(evenInc())
	fmt.Println(evenInc())
	fmt.Println(oddInc())
}
