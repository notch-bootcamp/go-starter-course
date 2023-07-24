package main

import "fmt"

/*
Multiline comment
Goes here
*/

const firstCont = "first"

var firstLine = firstCont

func main() {
	// Single line comment: following line writes to stdout
	fmt.Printf("This is %s line of my code\n", firstLine)
}
