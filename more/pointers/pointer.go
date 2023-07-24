package main

import (
	"fmt"
	"go-starter-course/basics/structs/types"
)

func main() {
	var addressToMeaning *int
	meaningOfEverything := 42
	addressToMeaning = &meaningOfEverything
	fmt.Println("This is pointer value", addressToMeaning)
	fmt.Println("This is a value to what we point", *addressToMeaning)
	*addressToMeaning = 50
	fmt.Println("This is a changed value to what we point", *addressToMeaning)

	home := types.Vector3D{"Home", 123, 15, 124}
	fmt.Println(home)
	pointerToHome := &home
	pointerToHome.Label = "Away from home"
	pointerToHome.Longitude = -15
	fmt.Println(home)
}
