// CF converts its numeric argument to Celsius and Fahrenheit.

package ch2

// package main

// "github.com/adonovan/gopl.io/ch2/tempconv"

// func cf() {
// 	// func main() {
// 	for _, arg := range os.Args[1:] {
// 		t, err := strconv.ParseFloat(arg, 64)
// 		if err != nil {
// 			fmt.FPrintf(os.Stderr, "cf: %v\n", err)
// 			os.Exit(1)
// 		}
// 		f := tempconv.Fahrenheit(t)
// 		c := tempconv.Celsius(t)
// 		fmt.Printf("%s = %s, %s = %s\n",
// 			f, tempconv.FToC(f), c, tempconv.CToF(c))
// 	}
// }
