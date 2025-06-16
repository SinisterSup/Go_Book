// Server2 is a minimal "echo" and counter server.
package ch1

// package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func server2demo() {
	// func main() {
	http.HandleFunc("/", pathHandler)  // each request to "/" calls pathHandler
	http.HandleFunc("/count", counter) // each request to "/count" calls counter
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var (
	mu    sync.Mutex
	count int
)

// pathHandler echoes the Path component of the request Url r.
func pathHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // Lock to protect shared count variable
	count++
	mu.Unlock() // Unlock after incrementing count
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls to the pathHandler.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}
