package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3}
	fmt.Println(arr)

	firstFibs := [5]int{1, 1, 2, 3, 5}
	fmt.Println(firstFibs)
	var fibsSlice []int = firstFibs[1:5]
	fmt.Println(fibsSlice)
	fibsSlice[1] = 10
	fmt.Println(firstFibs)
}
