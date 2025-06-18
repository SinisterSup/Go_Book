// Package tempconv performs Celsius and Fahrenheit temperature conversions.

package ch2

// import "fmt"

// // package tempconv
//
// type (
// 	Celsius    float64
// 	Fahrenheit float64
// )
//
// const (
// 	AbsoluteZeroC Celsius = -273.15
// 	FreezingC     Celsius = 0
// 	BoilingC      Celsius = 100
// )
//
// const (
// 	AbsoluteZeroF Fahrenheit = -459.67
// 	FreezingF     Fahrenheit = 32
// 	BoilingF      Fahrenheit = 212
// )
//
// func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
//
// func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
//
// func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// // Testing Celsius type's String method
// func main() {
// 	c := FToC(BoilingF)
// 	fmt.Println(c)
// 	// fmt.Println(c.String())
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%s\n", c)
// 	fmt.Printf("%g\n", c)
// 	fmt.Println(float64(c)) // convert to float64
// }
