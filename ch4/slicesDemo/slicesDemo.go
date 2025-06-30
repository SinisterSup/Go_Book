package slicesdemo

// package main

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

	// Slice beyond capacity i.e., cap(slice) causes a panic, but slicing beyong len(slice)
	// extends the slice without panic.
	fmt.Println("Slicing beyond capacity: ", summer[:20]) // panic: out of range
	endlessSummer := summer[:5]                           // extends the slice without panic (within capacity)
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
}
