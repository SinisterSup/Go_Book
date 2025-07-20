package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)       // a set of strings
	input := bufio.NewScanner(os.Stdin) // read from standard input
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line) // print the line if it has not been seen before
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Deduplication complete. Unique lines printed above.")
	fmt.Println("Total unique lines:", len(seen))
	fmt.Println("Unique lines are store in map:", seen)
}
