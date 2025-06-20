// package utf8demo
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println("String:", s)
	fmt.Println(len(s), "bytes")
	fmt.Println(utf8.RuneCountInString(s), " no. of runes")

	fmt.Println()

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d \t %c \t %v-byteSize \n", i, r, size)
		i += size
	}

	// Go's range loop, when applied to a string, performs UTF-8 decoding implicitly.
	for i, r := range s {
		fmt.Printf("%d \t %q \t %d \n", i, r, r)
	}
	// for i, r := range s {
	// 	fmt.Printf("%d \t %q \t %d \n", i, r, r)
	// }

	// n := 0
	// for range s {
	// 	n++
	// }

	// omitting range variables if need to calculate length of string
	n := 0
	for range s {
		n++
	}

	// or we can just call utf8.RuneCountInString(s)

	// "program" in Japanese katakana
	japanese := "プログラム"
	fmt.Printf("% x\n", japanese) // e3 83 97 e3 83 a9 e3 82 b0 e3 83 a9 e3 83 a0

	r := []rune(japanese) // convert string to rune slice
	fmt.Printf("%x\n", r) // [30d7 30ed 30b0 30e9 30e0]

	fmt.Println(string(r)) // "プログラム"

	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京", not "0x4eac"

	fmt.Println(string(1234567)) // block?
}
