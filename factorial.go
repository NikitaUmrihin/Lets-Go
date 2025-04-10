package main

import (
	"fmt"
)

// A recursive function to calculate the factorial of a number
func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)

}

func main() {
	input := 8
	answer := Factorial(8)
	fmt.Printf("%d! = %d", input, answer)
}
