package main

import "fmt"

const (
	_ = 1 << (10 * iota) // ignore first value by assigning to blank identifier
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func untypedConstsDemo() {
	// the choice of literal may affect the result of a constant division expression (or any expression that involves untyped constants).
	var f float64 = 212
	fmt.Println("Printing Celsius conversion of 212 Fahrenheit:")
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64 expression
	fmt.Println(5 / 9 * (f - 32))     // "0"; 5 / 9 is an integer expression, so the result is truncated to an integer
	fmt.Println(5.0 / 9 * (f - 32))   // "100"; 5.0 / 9 is an untyped float64 expression, so the result is a float64
	fmt.Println(5 / 9.0 * (f - 32))   // "100"; 5 / 9.0 is also an untyped float64, so the result is a float64
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0 / 9.0 is an untyped float64 constant, so the result is a float64

	fmt.Println()
	// only constants can be untyped.
	// when an untyped constant is assigned to a variable,
	// or appears on the right-hand side of a variable declaration with an explicit type,
	// the constant is implicitly converted to the type of the variable (if possible) or the type specified in the declaration.
	fmt.Println("Printing untyped constants:")
	var r float64 = 3 + 0i      // untyped complex -> float64
	fmt.Printf("%T %[1]v\n", r) // "3"
	r = 2                       // untyped integer -> float64
	fmt.Printf("%T %[1]v\n", r) // "2"
	r = 1e123                   // untyped floating-point -> float64
	fmt.Printf("%T %[1]v\n", r) // "1e+123"
	r = 'a'                     // untyped rune -> float64
	fmt.Printf("%T %[1]v\n", r) // "97"

	// // the statements above are thus equivalent to:
	// var r float64 = float64(3 + 0i)
	// r = float64(2)
	// r = float64(1e123)
	// r = float64('a')

	fmt.Println()
	// For literal values, the Go compiler uses the type of the literal to determine the type of the constant.
	// i.e., Syntax determines flavor of the constant.
	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32"
	fmt.Printf("%T\n", 'a')    // "int32" (rune is an alias for int32)
}

// func main() {
// 	untypedConstsDemo()
// }
