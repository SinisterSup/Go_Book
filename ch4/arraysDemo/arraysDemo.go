// package arraysdemo
package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	// By default, the elements of a new variable are initially set to zero value.
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)
	fmt.Println(r[2]) // "0"

	arr := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr)

	arr = [4]int{4, 5, 6, 7}
	fmt.Println(arr)
	// arr = [3]int{2, 3, 4}
	// // compile error: cannot use [3]int{...} (value of type [3]int) as [4]int value in assignment

	// it is also possible to specify a list of index and value pairs
	symbol := [...]string{USD: "dollar", EUR: "euro", GBP: "pound", RMB: "yen"}
	fmt.Println(RMB, symbol[RMB]) // "3 yen"
}

/* Go passes arguments (i.e paramter variable values) by value,
this behavior is different from languages (c/c++) that implicitly pass arguments(arrays) by reference.
Go treats arrays like any other type. (passing them by value)
this includes arrays/slices, hence passing large arrays can be inefficient */

// Instead we can explicitly pass a pointer to an array so that any modifications
// the function makes to array elements will be visible to the caller.
func zero1(arrPtr *[32]byte) {
	for i := range arrPtr {
		arrPtr[i] = 0
	}
}

func zero2(arrPtr *[32]byte) {
	*arrPtr = [32]byte{}
}

/* Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
* (See PopCount from ch2/popcount) */

/* Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but
* supports a command-line flag to print the SHA384 or SHA512 hash instead. */
