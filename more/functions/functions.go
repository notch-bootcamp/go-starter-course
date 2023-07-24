package main

import (
	"fmt"
	"go-starter-course/more/functions/function"
	"go-starter-course/more/functions/returns"
	"go-starter-course/more/functions/variadic"
)

func main() {
	lengthGuess := 42
	word := "meaning of everything"
	resultOfGuess := function.GuessLength(lengthGuess, word)
	fmt.Printf("You are guessin that length of '%s' is %d: %t\n", word, lengthGuess, resultOfGuess)

	firstHalf, secondHalf := returns.SplitInHalf(word)
	fmt.Printf("A half of '%s' is '%s' and '%s'\n", word, firstHalf, secondHalf)

	fmt.Printf("Sum of three consecutive numbers is: %d\n", variadic.SumItAll(1, 2, 3))

	fibNums := []int{1, 1, 2, 3, 5}
	fmt.Printf("Sum of first five Fib numbers: %d\n", variadic.SumItAll(fibNums...))
}
