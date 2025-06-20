// package stringsdemo

package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// basename removes directory components and a '.' suffix.
// i.e removes any prefix of s that looks like a file system path
// with components separated by slashes '/' and
// it removes any suffix that looks like a file type:
//
// gopl.io/ch3/basename1
func basename1(s string) string {
	// Discard last '/' and everything before it
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// gopl.io/ch3/basename2
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if '/' not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// gopl.io/ch3/comma
// comma inserts commas in a non-negative decimal integer string.
func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma1(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	commaSepString := ""
	for i := n - 3; i >= 0; i -= 3 {
		commaSepString = s[i:i+3] + "," + commaSepString
		if i-3 < 0 {
			commaSepString = s[0:i] + "," + commaSepString
		}
	}
	return commaSepString[:len(commaSepString)-1] // Remove the last comma
}

// gopl.io/ch3/printints
// intsToString is like fmt.Sprintf(values) but adds commas
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()

	// gh copilot suggestion!
	//
	//	var sb strings.Builder
	//	for i, v := range values {
	//	  if i > 0 {
	//	    sb.WriteString(", ")
	//	  }
	//	  sb.WriteString(fmt.Sprintf("%d", v))
	//	}
	//	return sb.String()
}

func main() {
	fmt.Println(basename1("/usr/local/bin/test.go")) // Output: test
	fmt.Println(basename2("/usr/local/bin/test.go")) // Output: test
	fmt.Println(comma1("1234567890"))                // Output: 1,234,567,890
	fmt.Println(comma2("1234567890"))                // Output: 1,234,567,890

	s := "hey,hello,ciao,hola,hallo,bonjour"
	b := []byte(s)
	r := []rune(s)
	splitStrings := strings.Split(string(b), ",")
	fmt.Println(splitStrings, reflect.TypeOf(splitStrings), reflect.TypeOf(splitStrings[0]))
	fmt.Println(r)
	for _, runeval := range r {
		fmt.Printf("%q \t %c \t %d \t %T\n", runeval, runeval, runeval, runeval)
	}
	fmt.Println(b)
	for _, byteVal := range b {
		fmt.Printf("%q \t %c \t %d \t %T\n", byteVal, byteVal, byteVal, byteVal)
	}
	s2 := string(b)
	fmt.Println(s2)

	fmt.Println(intsToString([]int{1, 2, 3, 4, 5})) // Output: [1, 2, 3, 4, 5]
}

/* Exercise 3.10: Write a non-recursive version of comman,
* using bytes.Buffer instead of string concatentation. */

/* Exercise 3.11: Enhance 'comma' func so that it deals with floating-point numbers,
* and an optional sign. */

/* Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,
* that is, they contain the same letters in a different order. */
