package main

import (
	"fmt"
	"go-starter-course/basics/structs/types"
)

func main() {
	vec := types.Vector3D{"Home", 123, 15, 124}
	sphericalLat, sphericalLong := vec.ToSphericalCoordinates()
	fmt.Println(sphericalLat, sphericalLong)
	vec.ChangeLabel("Home in Radians").
		ChangeCoordinates(types.Coordinate(sphericalLat), types.Coordinate(sphericalLong))
	fmt.Printf("%+v", vec)
}
