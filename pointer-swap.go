package main

import (
	"fmt"
)

// A function that swaps the values of two integers using pointers
func BadPointerSwapper(n1, n2 *int) {
	//  Risk of integer overflow: If *n1 + *n2 exceeds the maximum integer limit, it could cause undefined behavior
	*n2 = *n1 + *n2
	*n1 = *n2 - *n1
	*n2 = *n2 - *n1
}

// Direct swapping using pointers
func GoodPointerSwapper(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}

func main() {
	n1 := 15
	n2 := 10

	fmt.Println("\nSwap the values of two integers using pointers:")
	fmt.Println("\tn1 =", n1, "-> ( address :", &n1, ")")
	fmt.Println("\tn2 =", n2, "-> ( address :", &n2, ")")

	GoodPointerSwapper(&n1, &n2)

	fmt.Println("AFTER SWAP")
	fmt.Println("\tn1 =", n1, "-> ( address :", &n1, ")")
	fmt.Println("\tn2 =", n2, "-> ( address :", &n2, ")")

}
