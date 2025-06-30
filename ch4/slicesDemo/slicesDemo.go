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

	nums := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(nums[:]) // reverse the entire array
	fmt.Println("Reversed nums:", nums)
}
