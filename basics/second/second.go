package main

import (
	"fmt"
	"go-starter-course/basics/first"
)

func someFunction() {
	first.SomeFunction()
	fmt.Println("Exported var: ", first.FirstVar)
}
