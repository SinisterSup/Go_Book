// Ftoc prints two Farenheit temperatures and their Celsius conversions.
package ch2

// package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("Freezing temp is %g F = %g C \n", freezingF, fToC(freezingF))
	fmt.Printf("Boiling temp is %g F = %g C \n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
