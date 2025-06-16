package ch1

// package main

import (
	"fmt"
	"log"
	"net/http"
)

func server1demo() {
	// func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("server1demo started on localhost:8000")
}

// Handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
