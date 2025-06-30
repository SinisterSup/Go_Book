package main

// package slicesdemo

import "fmt"

// gopl.io/ch4/rev
// function reverse reverses a slices of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, i int) string {
	reverse(s[:i]) // reverse first i elements
	reverse(s[i:]) // reverse remaining elements
	reverse(s)     // reverse the entire slice
	return fmt.Sprintf("Rotated slice: %v", s)
}

func sliceEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// gopl.io/ch4/append
func appendInt(x []int, newElement int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) { // if there is capacity
		z = x[:zlen] // extend the slice length
	} else {
		// There is insufficient space. Allocate a new array slice.
		// Grow by doubling, for amortized linear complexity.
		zcap := max(zlen, 2*len(x))
		z = make([]int, zlen, zcap) // create a new slice with the new capacity
		copy(z, x)                  // a built-in function to copy elements from x to z
	}
	z[len(x)] = newElement // append the new element
	return z
}

// func appendInt(x []int, newElement int) []int {
//   if len(x) < cap(x) { // if there is capacity, append the new element
//     x = x[:len(x)+1] // extend the slice length
//     x[len(x)-1] = newElement
//     return x
//   }
//   // if no capacity, create a new slice with double the capacity and copy elements
//   newSlice := make([]int, len(x), 2*cap(x)+1)
//   copy(newSlice, x)
//   newSlice[len(x)] = newElement
//   return newSlice
// }

func main() {
	months := [...]string{
		1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December",
	}

	q2 := months[4:7]     // April, May, June
	summer := months[3:6] // April, May, June
	rainy := months[6:9]  // June, July, August

	fmt.Println("Q2:", q2)
	fmt.Println("Summer:", summer)
	fmt.Println("Rainy:", rainy)

	for _, s := range summer {
		for _, q := range q2 {
			if s == q {
				fmt.Printf("%s appears in both Summer and Q2\n", s)
			}
		}
	}

	// // Slice beyond capacity i.e., cap(slice) causes a panic,
	// fmt.Println("Slicing beyond capacity: ", summer[:20]) // panic: out of range
	//
	// but slicing beyond len(slice) extends the slice without panic.
	endlessSummer := summer[:5] // extends the slice without panic (within capacity)
	fmt.Println("Endless Summer: ", endlessSummer)

	// a slice contains a pointer to an underlying array,
	// hence passing a slice to a function allows the function to modify the original underlying array.
	nums := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(nums[:]) // reverse the entire array
	fmt.Println("Reversed nums:", nums)

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice1)
	fmt.Println("Rotated slice:", rotate(slice1, 2)) // Rotate the slice by 2 positions

	// slice literal implicitly creates an array variable of the right size and yields a slice that points to it.
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice 1 equals Slice 2: ", sliceEqual(slice1, slice2)) // Check if slices are equal
	// slices are not comparable, so we cannot use `==` operator directly.

	// The zero value of a slice is nil, which is different from an empty slice.
	// slice zero values demo
	var sli []int                          // len(sli) == 0, sli == nil
	sli = nil                              // sli is nil, len(sli) == 0
	sli = []int(nil)                       // sli is nil, len(sli) == 0
	sli = []int{}                          // len(sli) == 0, sli != nil
	fmt.Println(sli, len(sli), sli == nil) // Output: [] 0 true

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r) // appending runes to the slice
	}
	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']

	// using to achieve the same, with build-in conversion function `[]rune()` of rune slices
	runes = []rune("Hello, 世界") // creating a rune slice from a string
	fmt.Printf("%q\n", runes)   // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']

	var x, y []int
	for i := range 10 {
		y = appendInt(x, i) // appending integers to the slice
		fmt.Printf("%d cap=%d\t %v\n", i, cap(y), y)
		x = y
	}
}
