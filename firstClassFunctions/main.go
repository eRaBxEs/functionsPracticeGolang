package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// find a first class functions below ahich is defined as a parameter for another function
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	firstResult := arithmetic(a, b)
	return arithmetic(firstResult, c)

}
func main() {
	totalsum := aggregate(2, 3, 4, add)
	fmt.Println("total sum is", totalsum)

	totalProd := aggregate(2, 3, 4, mul)
	fmt.Println("total product is", totalProd)
}
