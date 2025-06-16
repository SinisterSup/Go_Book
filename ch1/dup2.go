package ch1

// package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup2Main() {
	// func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if containsDuplicateLines(arg) {
				fmt.Printf("File %s contains duplicate lines.\n", arg)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s \n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := strings.TrimSpace(input.Text())
		// fmt.Printf("Processing line: %s\t with length %d\n", line, len(line))
		if line == "" {
			continue
		}
		counts[line]++
	}
	// Note: ignoring potential errors from input.Err()
}

// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
func containsDuplicateLines(fileName string) bool {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}
	lineCounts := make(map[string]int)
	hasDuplicates := false
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lineCounts[line]++
		if lineCounts[line] > 1 {
			hasDuplicates = true
			break
		}
	}
	return hasDuplicates
}
