package main

// package slicesdemo

import (
	"fmt"
	"unicode/utf8"
)

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

// gopl.io/ch4/nonempty
// NonEmpty returns a slice holding only the non-empty strings from the input slice.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	// The zero value of a slice is nil, which is different from an empty slice.
	// A nil slice has no underlying array, so it cannot be modified.
	// An empty slice has an underlying array of length 0, so it can be modified.
	// Hence, we can use the zero value of a slice to check if the input slice is nil or empty.
	if len(strings) == 0 {
		return strings // return the empty slice
	}

	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s // keep the non-empty string
			i++
		}
	}
	return strings[:i] // return the modified slice with non-empty strings
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice with the same underlying array
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// To remove an element from the middle of a slice,
// preserving the other of the remaining elements,
// use copy to slide the elements to the left.
func remove(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice // return the original slice if index is out of range
	}
	copy(slice[i:], slice[i+1:]) // shift elements to the left
	return slice[:len(slice)-1]  // return the modified slice
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

	stringData := []string{"Hello", "", "World", "", "Go", "Programming", ""}
	fmt.Printf("%q\n", nonempty(stringData)) // ["Hello" "World" "Go" "Programming"]
	stringDataCopy := []string{"Hello", "", "World", "", "Go", "Programming", ""}
	fmt.Printf("%q\n", nonempty2(stringDataCopy)) // ["Hello" "World" "Go" "Programming"]
	fmt.Printf("%q\n", stringData)                // `["Hello" "World" "Go" "Programming" "Go" "Programming" ""]`

	// Stack implementation using slices
	stack := []int{}
	stack = append(stack, 1, 2, 3)
	v := 4
	stack = append(stack, v)     // push v onto the stack
	top := stack[len(stack)-1]   // top element of the stack
	stack = stack[:len(stack)-1] // pop the top element
	fmt.Printf("Stack after push and pop: %v, top: %d\n", stack, top)

	slice := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(slice, 2)) // Remove element at index 2, output: [5 6 8 9]

	arrayPntr := &[...]int{1, 3, 5, 7, 9}
	fmt.Println("Original array pointer:", arrayPntr)
	reverseArr(arrayPntr) // Reverse the array using a pointer
	fmt.Println("Reversed array pointer:", arrayPntr)

	slice = []int{2, 4, 6, 8, 10, 12}
	fmt.Println("Original slice to be rotated:", slice)
	slice = rotateSlice(slice, 2) // Rotate the slice in a single pass
	fmt.Println("Rotated slice in single pass:", slice)

	intSlice := []int{4, 5, 6, 6, 7, 7, 7, 3, 3, 9, 1, 1, 1, 1, 2, 2, 4, 8}
	fmt.Println("Original intSlice with adjacent duplicates ", intSlice)
	intSlice = discardAdjDuplicates(intSlice)
	fmt.Println("Discarded adjacent Duplicates in-place:", intSlice)

	testStr := "Hello, 世界"
	fmt.Println("Original string:", testStr)
	reversedStr := reverseString(testStr)
	fmt.Println("Reversed string:", reversedStr)

	helloBytes := []byte("Hello, 世界") // []byte{72, 101, 108, 108, 111, 44, 32, 228, 184, 150, 229, 155, 189}
	fmt.Println("Original byte slice:", helloBytes)
	// reverseBytes(helloBytes)
	// fmt.Println("Reversed byte slice:", helloBytes)
	reverseUTF8ByteSlice(helloBytes) // Reverse the byte slice in place
	fmt.Println("Reversed byte slice:", helloBytes)
	fmt.Println("Reversed byte slice as string:", string(helloBytes))
}

/* Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice. */
func reverseArr(arrP *[5]int) {
	for l, r := 0, len(arrP)-1; l < r; l, r = l+1, r-1 {
		(*arrP)[l], (*arrP)[r] = (*arrP)[r], (*arrP)[l]
	}
}

/* Exercise 4.4: Write a version of rotate that operates in a single pass. */
func rotateSlice(slice []int, k int) []int {
	if len(slice) == 0 || k <= 0 || k >= len(slice) {
		return slice // return the original slice if empty or i is out of range
	}
	k %= len(slice)
	// slice = append(slice[k:], slice[:k]...) // rotate the slice in a single pass
	// return slice
	return append(slice[k:], slice[:k]...)
}

/* Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice. */
func discardAdjDuplicates(intSlice []int) []int {
	i := 1
	for _, val := range intSlice[1:] {
		if val != intSlice[i-1] {
			intSlice[i] = val
			i++
		}
	}
	return intSlice[:i]
}

/* Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces
* (see unicode.IsSpace) in a UTF-8 encoded []byte slice into a single ASCII space. */

/* Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a
* UTF-8 encoded string, in place. Can you do it without allocating new memory? */
func reverseString(str string) string {
	runes := []rune(str)
	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}

func reverseBytes(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}

func reverseUTF8ByteSlice(s []byte) {
	reverseBytes(s) // reverse the byte slice
	i := 0
	for i < len(s) {
		_, size := utf8.DecodeRune(s[i:]) // decode the rune at index i
		// Reverse the bytes within this single character.
		reverseBytes(s[i : i+size]) // reverse the bytes of the rune
		i += size                   // move to the next rune
	}
}
