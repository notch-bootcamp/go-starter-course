package first

import "fmt"

const firstHiddenVar = "hidden"

var FirstVar = firstHiddenVar

func SomeFunction() {
	fmt.Println(firstHiddenVar)
}
