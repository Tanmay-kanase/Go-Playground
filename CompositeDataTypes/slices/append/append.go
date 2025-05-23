package main

import (
	"fmt"
)

func main() {
	var runes []rune
	for _, r := range "Hello, BF" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // Output: ['H' 'e' 'l' 'l' 'o' ',' ' ' 'B' 'F']
}
