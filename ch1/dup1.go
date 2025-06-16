package ch1

// package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
func dup1Main() {
	input := bufio.NewScanner(os.Stdin)
	lineCounts := make(map[string]int)
	for input.Scan() {
		lineCounts[input.Text()]++
	}
	// Note: ignoring potential errors from input.Err()
	for line, count := range lineCounts {
		if count > 1 {
			fmt.Printf("Count: %d \t Line: %s\n", count, line)
		}
	}
}
