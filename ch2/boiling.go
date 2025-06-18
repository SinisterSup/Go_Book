// Boiling prints the boiling point of water.
package ch2

// package main

import "fmt"

const boilingF = 212.0 // boiling point of water in Fahrenheit

func boilingPoint() {
	// func main() {
	f := boilingF
	c := (f - 32) * 5 / 9 // fahrenheit to Celsius conversion
	fmt.Printf("boiling point of water = %g F or %g C \n", f, c)
	// Output:
	// boiling point of water = 212 F or 100 C
}
