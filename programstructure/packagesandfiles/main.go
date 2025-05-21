package main

import (
	"go-playground/programstructure/packagesandfiles/tempconv" 
	"fmt"
)

func main() {
	fmt.Println(tempconv.BoilingC)                    // 100°C
	fmt.Println(tempconv.CToF(tempconv.BoilingC))     // 212°F
	fmt.Println(tempconv.FToC(212))                   // 100°C
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // -273.15°C
}
