// package mapsDemo

package main

import "fmt"

// `equal` function checks whether two maps with string keys and int values are equal.
// i.e 2 maps contain the same keys and the same associated values.
func equal(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k1, v1 := range map1 {
		if v2, ok := map2[k1]; !ok || v2 != v1 {
			return false
		}
	}
	return true
}

func main() {
	// built-in function `make` to create a map
	// ages := make(map[string]int) // mapping from strings to ints

	// map literal to create a map, populated with some initial key/value pairs
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	// // this is equivalent to
	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 34

	// An alternative expression for a new empty map is:
	// map[string]int{}
	fmt.Println("Understanding map basics: ")
	fmt.Println("A map is a reference to a hash table, which is a data structure that maps keys to values.")

	ages["alice"] = 32                                          // update value for key "alice"
	fmt.Println("Updated key `alice`'s age is:", ages["alice"]) // prints 32

	delete(ages, "charlie") // delete key "charlie"

	// A map lookup using a key that isn't present returns the zero value for the map's value type.
	ages["bob"]++
	fmt.Printf("New key `bob`'s age is %d\n", ages["bob"]) // prints 1

	// Map element is not a variable, so we cannot take its address
	// _ = &ages["bob"]  // compile error: cannot take address of map element
	// growing a map might cause rehashing of existing elements into new storage locations,
	// so the address of a map element may change, thus potentially invalidating any pointers to it.
	//
	var agesNew map[string]int
	fmt.Println("nil value of map check - agesNew == nil?:", agesNew == nil)
	fmt.Println("len of empty map check - len(agesNew) == 0?:", len(agesNew) == 0)
	// Most operations on mapps, including lookup, delete, len, and range loops,
	// are safe to perform on a nil map refernece, since it behanves like an empty map.
	// But storing to a nil map causes a runtime panic.
	// ages["carol"] = 21 // panic: assignment to entry in nil map

	// maps cannot be compared with each other directly using `==` operator,
	// the only legal comparison is to compare the map to nil.
	agesNew = map[string]int{
		"alice": 32,
		"bob":   1,
	}

	fmt.Println("Comparing two maps using `equal` function - equal(ages, agesNew):", equal(ages, agesNew))
	ages["bob"]++
	fmt.Println("Comparing updated 'ages' map - equal(ages, agesNew):", equal(ages, agesNew))
}

// Sometimes we need a map or set whose keys are slices
// but, map's keys must be comparable,
// Hece for slices, we need to use a workaround.
// First, we convert the slice to a string, using a helper function like k
// maps each key to string representation of the slice.
// Then we can use the string(applying the helper function) before using it as a key in the map.
var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("eq %v", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

// The value type of a map can itself be a composite type, such as a map, slice, or struct.
// gopl.io/ch4/graph
var graph = make(map[string]map[string]bool)

func AddEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func HasEdge(from, to string) bool {
	return graph[from][to]
}

// This HasEdge shows how the zero value of a missing map entry can be often put to work:
// 1. if `from` is not a key in `graph`, then `graph[from]` returns nil,
// and the expression `graph[from][to]` evaluates to false without causing a panic.
// 2. if `from` is a key in `graph`, but `to` is not a key in the map
// then zero value of the map's value type(false) is returned.

/* Exercise 4.8: Modify charcount to count the number of occurrences of each Unicode category
* count letters, digits, punctuation, and so on in their Unicode categories., using the unicode package. */

/* Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text file.
* Call input.Split(bufio.ScanWords) before the first call to Scan to break the input
* into words instead of lines. */
