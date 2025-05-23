package main

import (
	"fmt"
)

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	a := []string{"Go", "Lang"}
	b := []string{"Go", "Lang"}
	c := []string{"Golang"}

	fmt.Println(equal(a, b))
	fmt.Println(equal(a, c))
}
