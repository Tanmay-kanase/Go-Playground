package main

import "fmt"

func main(){
	const freezing , boiling = 32.0 , 212.0
	fmt.Printf("%gF = %gC \n" , freezing , fToC(freezing))
	fmt.Printf("%gF = %gC \n" , boiling , fToC(boiling))


}

func fToC(f float64) float64{
	return (f-32) * 5/9
}