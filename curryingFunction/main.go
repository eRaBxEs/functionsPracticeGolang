package main

import (
	"fmt"
)

func multiply(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

// currying a feature from functional programming allows function with multiple arguments to be transformed into series of functions, each taking a single argument
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func main() {
	squarefunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squarefunc(5))
	fmt.Println(doubleFunc(5))
}
