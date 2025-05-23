package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	b := make([]int, 3, 5)

	fmt.Println(a)             
	fmt.Println(b, cap(b))   
	b = append(b, 1, 2)        
	fmt.Println(b)             						
}
