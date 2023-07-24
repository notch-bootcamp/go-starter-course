package main

import (
	"fmt"
	"go-starter-course/basics/structs/types"
	c "go-starter-course/consts"
)

func main() {
	locations := []types.Vector3D{
		{c.Zero, 0, 0, 0},
		{c.Home, 45, 15, 124},
		{c.Google, 37, -122, 8},
		{},
	}
	for i, location := range locations {
		switch location.Label {
		case c.Zero:
			fmt.Printf("This is zero location at index %d\n", i)
		case c.Home:
			fmt.Printf("We are %s at index %d\n", location.Label, i)
		case c.Google:
			fmt.Printf("We are in %s at index %d\n", location.Label, i)
		default:
			fmt.Printf("We don't know about location '%s' at index %d", location.Label, i)
		}
	}
}
