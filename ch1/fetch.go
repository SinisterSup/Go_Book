// Fetch prints the content found at each of the URLs provided as command-line arguments.
package ch1

// package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchDemo() {
	// func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

/* Exercise 1.7: The function call to io.Copy(dst, src) reads from src and writes to dst.
* Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
* buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy. */

/* Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing.
* You might want to use strings.HasPrefix for this. If the URL has a http:// prefix, */

/* Exercise 1.9: Modify fetch to print the HTTP response status code, found in resp.Status. */
