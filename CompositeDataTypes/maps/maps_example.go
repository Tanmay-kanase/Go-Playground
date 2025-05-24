package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 1. Creating a map
func creatingMaps() {
	fmt.Println("1. Creating maps")
	ages1 := make(map[string]int)
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println("ages1:", ages1)
	fmt.Println("ages2:", ages2)
}

// 2. Access, update, delete map elements
func mapOperations() {
	fmt.Println("\n2. Map operations")
	ages := map[string]int{
		"alice": 31,
	}
	ages["alice"] = 32
	fmt.Println("Updated alice:", ages["alice"])
	delete(ages, "alice")
	fmt.Println("After delete:", ages)
	ages["bob"] = ages["bob"] + 1
	fmt.Println("Increment bob:", ages["bob"])
}

// 3. Iterating over a map
func iterateMap() {
	fmt.Println("\n3. Iterating map")
	ages := map[string]int{
		"alice":   31,
		"bob":     28,
		"charlie": 34,
	}
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

// 4. Sorted iteration
func sortedIteration() {
	fmt.Println("\n4. Sorted iteration")
	ages := map[string]int{
		"alice":   31,
		"bob":     28,
		"charlie": 34,
	}
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

// 5. Nil maps
func nilMaps() {
	fmt.Println("\n5. Nil maps")
	var ages map[string]int
	fmt.Println("Is nil:", ages == nil)
	fmt.Println("Length 0:", len(ages) == 0)
	// Uncommenting below will panic
	// ages["carol"] = 21
}

// 6. Check if key exists
func keyCheck() {
	fmt.Println("\n6. Key check")
	ages := map[string]int{"bob": 25}
	if age, ok := ages["bob"]; ok {
		fmt.Println("Found bob:", age)
	} else {
		fmt.Println("bob not found")
	}
}

// 7. Comparing two maps
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func compareMaps() {
	fmt.Println("\n7. Compare maps")
	a := map[string]int{"A": 0}
	b := map[string]int{"A": 0}
	c := map[string]int{"B": 42}
	fmt.Println("a == b:", equal(a, b))
	fmt.Println("a == c:", equal(a, c))
}

// 8. Dedup: Using map as set
func dedup() {
	fmt.Println("\n8. Dedup from stdin (type and press Ctrl+D):")
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

// 9. Using slices as map keys via string conversion
var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func sliceAsMapKey() {
	fmt.Println("\n9. Slices as map keys")
	Add([]string{"go", "lang"})
	Add([]string{"go", "lang"})
	Add([]string{"java", "lang"})
	fmt.Println("Count go, lang:", Count([]string{"go", "lang"}))
	fmt.Println("Count java, lang:", Count([]string{"java", "lang"}))
}

func main() {
	creatingMaps()
	mapOperations()
	iterateMap()
	sortedIteration()
	nilMaps()
	keyCheck()
	compareMaps()
	sliceAsMapKey()

}
