package main

import (
	"fmt"
)

// look at this first class function defined as a parameter for the conversion function below
func conversion(converted func(int) int, x, y, z int) (int, int, int) {
	convertedX := converted(x)
	convertedY := converted(y)
	convertedZ := converted(z)

	return convertedX, convertedY, convertedZ
}

// you can either specify a function to satisfy the first class function defined above with this
func double(a int) int {
	return a + a
}

func main() {

	// and use it as a parameter to pass into the function as below
	newX, newY, newZ := conversion(double, 1, 2, 3)
	fmt.Println(newX, newY, newZ)

	// or simply create an anonymous function and use it quickly rather than having needed to have created the function named double above
	newX, newY, newZ = conversion(func(a int) int {
		return a + a
	}, 1, 2, 3)

	fmt.Println(newX, newY, newZ)

}
