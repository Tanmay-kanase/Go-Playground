package main

import (
	"fmt"
)

func main() {
	var s1 []int             
	s2 := []int(nil)          
	s3 := []int{}            
	s4 := make([]int, 3)[3:] 

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil) 
	fmt.Println(s4 == nil) 

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s3), cap(s3))
	fmt.Println(len(s4), cap(s4))
}
