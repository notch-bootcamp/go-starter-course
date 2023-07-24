package main

import (
	"fmt"
	"math"
)

var i, j int

func main() {
	i = 42
	j = 7
	fmt.Println(i, j)

	var p = 10
	fmt.Println(p)
	r := 11
	fmt.Println(r)
	const isTrue = true
	str := "always"
	fmt.Printf("%d is %s %t\n", i, str, isTrue)

	a, b := 4, 3
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Printf("Hypotenuse is: %f\n", c)
}
