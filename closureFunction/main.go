package main

import (
	"fmt"
)

// A closure is a function that references variable outside it's own body,
// here a closure is the anonymous function defined as a return value to concatter because it has access to the variable doc
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word
		return doc
	}
}

func main() {
	wordAggregator := concatter()
	wordAggregator("That ")
	wordAggregator("the God ")
	wordAggregator("of our Lord Jesus ")
	wordAggregator("may give unto you ")
	wordAggregator("the spirit of wisdom ")
	wordAggregator("and revelation ")
	wordAggregator("in the knowledge of ")

	fmt.Println(wordAggregator("Him. "))

}
