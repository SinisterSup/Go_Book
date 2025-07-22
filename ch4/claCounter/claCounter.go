package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var (
	c = flag.Bool("c", false, "count Characterwise frequency")
	w = flag.Bool("w", false, "count Wordwise frequency")
	l = flag.Bool("l", false, "count Linewise frequency")
)

func countCharacters(input *os.File, lettersCounter, digitsCounter, specialsCounter, othersCounter map[rune]int) {
	text := bufio.NewReader(input)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	for {
		r, n, err := text.ReadRune() // returns rune, number of bytes read, and error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "countCharacters: Error reading rune: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		utflen[n]++
		switch {
		case unicode.IsLetter(r):
			lettersCounter[r]++
		case unicode.IsDigit(r):
			digitsCounter[r]++
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			specialsCounter[r]++
		default:
			othersCounter[r]++
		}

		// fmt.Println("Displaying utf8 rune lengths & counts:")
		// fmt.Print("len\tcount\n")
		// for i, n := range utflen {
		// 	if i > 0 {
		// 		fmt.Printf("%d\t%d\n", i, n)
		// 	}
		// }
		if invalid > 0 {
			fmt.Fprintf(os.Stderr, "countCharacters: %d invalid UTF-8 characters\n", invalid)
		}
	}
}

func countWords(file *os.File, wordsCounter map[string]int) {
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordsCounter[word]++
	}
}

// func countWords(file *os.File, wordsCounter map[string]int) {
// 	input := bufio.NewScanner(file)
// 	for input.Scan() {
// 		line := input.Text()
// 		words := strings.Fields(line)
// 		for _, word := range words {
// 			wordsCounter[word]++
// 		}
// 	}
// }

func countLines(file *os.File, linesCounter map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		linesCounter[input.Text()]++
	}
}

func main() {
	flag.Parse()
	if !*c && !*w && !*l {
		files := flag.Args()
		if len(files) != 0 {
			fmt.Println("Defaulting to Counting Characterwise freq from Standard Input")
			fmt.Println("No command line arguments provided!")
			fmt.Println("user -h or -help flag to get help on available flags")
		}

		// TODO: Implement default behavior
		lettersCounter := make(map[rune]int)
		digitsCounter := make(map[rune]int)
		specialsCounter := make(map[rune]int)
		othersCounter := make(map[rune]int)
		countCharacters(os.Stdin, lettersCounter, digitsCounter, specialsCounter, othersCounter)

		fmt.Println("Displaying Letters, Digits, Specials, and Others frequency maps:")
		fmt.Println("Letter\tCount")
		for letter, count := range lettersCounter {
			fmt.Printf("%q\t%d\n", letter, count)
		}
		fmt.Println()
		fmt.Println("Digit\tCount")
		for digit, count := range digitsCounter {
			fmt.Printf("%c\t%d\n", digit, count)
		}
		fmt.Println()
		fmt.Println("Special\tCount")
		for special, count := range specialsCounter {
			fmt.Printf("%q\t%d\n", special, count)
		}
		fmt.Println()
		fmt.Println("Other\tCount")
		for other, count := range othersCounter {
			fmt.Printf("%q\t%d\n", other, count)
		}
		fmt.Println()
	}
	if *c {
		files := flag.Args()
		lettersCounter := make(map[rune]int)
		digitsCounter := make(map[rune]int)
		specialsCounter := make(map[rune]int)
		othersCounter := make(map[rune]int)
		for _, fileArg := range files {
			f, err := os.Open(fileArg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "CharacterCounter: Error opening file %s: %v\n", fileArg, err)
				continue
			}
			countCharacters(f, lettersCounter, digitsCounter, specialsCounter, othersCounter)
			f.Close()
		}

		fmt.Println("Displaying for all input Arg Files - Letters, Digits, Specials, and Others frequency maps:")
		fmt.Println("Letter\tCount")
		for letter, count := range lettersCounter {
			fmt.Printf("%q\t%d\n", letter, count)
		}
		fmt.Println()
		fmt.Println("Digit\tCount")
		for digit, count := range digitsCounter {
			fmt.Printf("%c\t%d\n", digit, count)
		}
		fmt.Println()
		fmt.Println("Special\tCount")
		for special, count := range specialsCounter {
			fmt.Printf("%q\t%d\n", special, count)
		}
		fmt.Println()
		fmt.Println("Other\tCount")
		for other, count := range othersCounter {
			fmt.Printf("%q\t%d\n", other, count)
		}
		fmt.Println()

	} else if *w {
		files := flag.Args()
		wordsCounter := make(map[string]int)
		for _, fileArg := range files {
			f, err := os.Open(fileArg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "WordCounter: Error opening file %s: %v\n", fileArg, err)
				continue
			}
			countWords(f, wordsCounter)
			f.Close()
		}

		fmt.Println("Displaying for all Arg Text Files - Words frequency map:")
		fmt.Println("Word\tCount")
		for word, count := range wordsCounter {
			fmt.Printf("%s\t%d\n", word, count)
		}
		fmt.Println()

	} else if *l {
		files := flag.Args()
		linesCounter := make(map[string]int)
		for _, fileArg := range files {
			f, err := os.Open(fileArg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "LineCounter: Error opening file %s: %v\n", fileArg, err)
				continue
			}
			countLines(f, linesCounter)
			f.Close()
		}

		fmt.Println("Displaying for all Arg Text Files - Line frequency map:")
		fmt.Println("Line\tCount")
		for line, count := range linesCounter {
			fmt.Printf("%s\t%d\n", line, count)
		}
		fmt.Println()
	}
}
