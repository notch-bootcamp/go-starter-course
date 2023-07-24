package main

import (
	"fmt"
	"go-starter-course/basics/structs/types"
)

func main() {
	a := types.Vector3D{"A", 0, 1, 2}
	b := types.Vector3D{Label: "A", Longitude: 1, Elevation: 2}
	fmt.Printf("Are they equal? %t", a == b)
}
