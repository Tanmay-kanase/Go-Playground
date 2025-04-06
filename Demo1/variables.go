package main

import (
	"fmt"
)

func main() {
	// -------------------------------
	// ✅ Basic Data Types
	// -------------------------------
	var aInt int = 42                     // default int
	var aInt8 int8 = -8                   // 8-bit signed integer
	var aInt16 int16 = -16000             // 16-bit signed integer
	var aInt32 int32 = -32000             // 32-bit signed integer
	var aInt64 int64 = -64000000000       // 64-bit signed integer

	var aUint uint = 42                   // unsigned int
	var aUint8 uint8 = 255                // 8-bit unsigned integer
	var aUint16 uint16 = 65535            // 16-bit unsigned integer
	var aUint32 uint32 = 4294967295       // 32-bit unsigned integer
	var aUint64 uint64 = 18446744073709551615 // 64-bit unsigned integer
	var aByte byte = 'A'                  // alias for uint8
	var aRune rune = '🦊'                 // alias for int32 (Unicode code point)

	var aFloat32 float32 = 3.14           // 32-bit float
	var aFloat64 float64 = 3.1415926535   // 64-bit float

	var aComplex64 complex64 = 1 + 2i     // complex with float32 parts
	var aComplex128 complex128 = 2 + 3i   // complex with float64 parts

	var aBool bool = true                 // boolean
	var aString string = "Hello, Go!"     // string

	// -------------------------------
	// ✅ Composite Types
	// -------------------------------
	var anArray [5]int = [5]int{1, 2, 3, 4, 5}           // array of 5 ints
	var aSlice []string = []string{"go", "is", "fun"}   // dynamic array
	var aMap map[string]int = map[string]int{"Alice": 25, "Bob": 30} // key-value store
	var aStruct = struct {
		id   int
		name string
	}{id: 1, name: "Structy"}                            // anonymous struct
	var aChan chan int = make(chan int)                  // unbuffered channel

	// -------------------------------
	// ✅ Reference / Interface / Function Types
	// -------------------------------
	var aPointer *int = &aInt                            // pointer to int
	var anInterface interface{} = "I'm anything!"        // empty interface
	var aFunc func(int) string = func(x int) string {
		return fmt.Sprintf("Got %d", x)
	}

	// -------------------------------
	// ✅ Constants & Short Declaration
	// -------------------------------
	const pi float64 = 3.14159                           // constant
	short := "GoLang"                                    // short variable declaration

	// -------------------------------
	// ✅ Type Alias & Custom Types
	// -------------------------------
	type myInt int
	var x myInt = 100                                    // custom type

	type Person struct {
		name string
		age  int
	}
	var p Person = Person{name: "John", age: 30}         // named struct

	type ScoreMap = map[string]int
	var sm ScoreMap = map[string]int{"Math": 90, "CS": 95} // alias type

	fmt.Printf("aInt = %d\n", aInt)
	fmt.Printf("aInt8 = %d\n", aInt8)
	fmt.Printf("aInt16 = %d\n", aInt16)
	fmt.Printf("aInt32 = %d\n", aInt32)
	fmt.Printf("aInt64 = %d\n", aInt64)

	fmt.Printf("aUint = %d\n", aUint)
	fmt.Printf("aUint8 = %d\n", aUint8)
	fmt.Printf("aUint16 = %d\n", aUint16)
	fmt.Printf("aUint32 = %d\n", aUint32)
	fmt.Printf("aUint64 = %d\n", aUint64)

	fmt.Printf("aByte = %c\n", aByte)
	fmt.Printf("aRune = %c\n", aRune)

	fmt.Printf("aFloat32 = %.2f\n", aFloat32)
	fmt.Printf("aFloat64 = %.10f\n", aFloat64)

	fmt.Printf("aComplex64 = %v\n", aComplex64)
	fmt.Printf("aComplex128 = %v\n", aComplex128)

	fmt.Printf("aBool = %t\n", aBool)
	fmt.Printf("aString = %s\n", aString)

	fmt.Printf("anArray = %v\n", anArray)
	fmt.Printf("aSlice = %v\n", aSlice)
	fmt.Printf("aMap = %v\n", aMap)
	fmt.Printf("aStruct = %+v\n", aStruct)
	fmt.Printf("aChan = %v\n", aChan)

	fmt.Printf("aPointer = %v (value = %d)\n", aPointer, *aPointer)
	fmt.Printf("anInterface = %v\n", anInterface)
	fmt.Printf("aFunc(7) = %s\n", aFunc(7))

	fmt.Printf("pi = %.5f\n", pi)
	fmt.Printf("short = %s\n", short)
	fmt.Printf("x (myInt) = %d\n", x)
	fmt.Printf("p (Person) = %+v\n", p)
	fmt.Printf("sm (ScoreMap) = %v\n", sm)

}
