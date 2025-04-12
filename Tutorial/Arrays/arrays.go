package main

import (
	"fmt"
)

// Function to demonstrate passing array as value (copy)
func modifyArray(arr [3]int) {
	arr[0] = 100
	fmt.Println("Inside modifyArray (copy):", arr)
}

// Function to demonstrate passing array as pointer (reference)
func modifyArrayByPointer(arr *[3]int) {
	arr[0] = 200
	fmt.Println("Inside modifyArrayByPointer (reference):", *arr)
}

func main() {
	// -------------------------------
	// ✅ Basic Declaration
	// -------------------------------
	var arr1 [3]int // default initialization (all elements 0)
	fmt.Println("arr1:", arr1)

	// -------------------------------
	// ✅ Declaration with Initialization
	// -------------------------------
	arr2 := [3]int{1, 2, 3}
	fmt.Println("arr2:", arr2)

	// -------------------------------
	// ✅ Partial Initialization
	// Remaining elements are zeroed
	// -------------------------------
	arr3 := [5]int{1, 2}
	fmt.Println("arr3:", arr3)

	// -------------------------------
	// ✅ Compiler Inferred Length
	// -------------------------------
	arr4 := [...]string{"Go", "Java", "Python"}
	fmt.Println("arr4:", arr4)

	// -------------------------------
	// ✅ Index-based Initialization
	// -------------------------------
	arr5 := [5]int{1: 10, 3: 20}
	fmt.Println("arr5:", arr5) // [0 10 0 20 0]

	// -------------------------------
	// ✅ Iterating with for loop
	// -------------------------------
	fmt.Print("Iterating arr2: ")
	for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i], " ")
	}
	fmt.Println()

	// -------------------------------
	// ✅ Iterating with for-range
	// -------------------------------
	fmt.Print("Iterating arr4 with range: ")
	for index, value := range arr4 {
		fmt.Printf("[%d] = %s ", index, value)
	}
	fmt.Println()

	// -------------------------------
	// ✅ Multidimensional Array
	// -------------------------------
	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D matrix:", matrix)

	// -------------------------------
	// ✅ Access & Modify Elements
	// -------------------------------
	arr2[1] = 100
	fmt.Println("Modified arr2:", arr2)

	// -------------------------------
	// ✅ Array Comparison (Same length & type)
	// -------------------------------
	a := [2]int{1, 2}
	b := [2]int{1, 2}
	fmt.Println("a == b:", a == b)

	// -------------------------------
	// ✅ Arrays are Value Types (copy on assignment)
	// -------------------------------
	original := [3]int{1, 2, 3}
	copyArr := original
	copyArr[0] = 999
	fmt.Println("original:", original)
	fmt.Println("copyArr:", copyArr)

	// -------------------------------
	// ✅ Passing Arrays to Functions
	// -------------------------------
	arrFunc := [3]int{10, 20, 30}
	modifyArray(arrFunc)           // passed by value
	fmt.Println("After modifyArray:", arrFunc)

	modifyArrayByPointer(&arrFunc) // passed by reference
	fmt.Println("After modifyArrayByPointer:", arrFunc)

	// -------------------------------
	// ✅ Length and Capacity
	// -------------------------------
	fmt.Printf("Length of arr3: %d\n", len(arr3))
	// (arrays have no separate capacity, unlike slices)

	// -------------------------------
	// ✅ Array of Structs
	// -------------------------------
	type Person struct {
		name string
		age  int
	}
	people := [2]Person{
		{name: "Alice", age: 21},
		{name: "Bob", age: 25},
	}
	fmt.Println("Array of structs:", people)
}
