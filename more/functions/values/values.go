package main

import (
	"fmt"
	"math"
)

func doMath(doesMath func(float64, float64) float64, x, y float64) float64 {
	if x == 0 || y == 0 {
		return 0
	}
	return doesMath(x, y)
}

func main() {

	square := func(r float64) float64 {
		return r * r
	}

	hypotenus := func(a, b float64) float64 {
		return math.Sqrt(square(a) + square(b))
	}

	fmt.Println(hypotenus(5, 12))
	fmt.Println(doMath(hypotenus, 5, 12))
	fmt.Println(doMath(hypotenus, 0, 12))
}
