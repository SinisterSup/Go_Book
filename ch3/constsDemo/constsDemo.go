// package constsDemo
package main

import (
	"fmt"
	"time"
)

// Since const values are known to compiler,
// constant expressions may appear in types,
// specifically as length of an array type.
const IPv4Len = 4

// parseIPv4 parses an IPv4 address (d.d.d.d) from a byte slice.
func parseIPv4(s string) {
	var p [IPv4Len]byte
	// ...
	fmt.Println("Parsed IPv4 address:", p)
}

func main() {
	// type is inferred from the expression on the right-hand side.
	// notice how Type of timeout and noDelay is time.Duration
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T  %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T  %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T  %[1]v\n", time.Minute) // "time.Duration 1m0s"

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // 1 1 2 2

	// Demo const generator iota
	// In const declarations, the iota identifier represents successive untyped integer constants.
	// the value of iota starts at 0 and increments by 1 for each const declaration in the same block.
	//
	// type weekday here of this kind are often called enumerated types., or enums for short.
	type Weekday int
	const (
		Sunday Weekday = iota // 0
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	// netflagsDemo
	var v Flags = FlagUp | FlagMulticast
	fmt.Printf("%b  %t\n", v, IsUp(v)) // 10001 true
	TurnDown(&v)
	fmt.Printf("%b  %t\n", v, IsUp(v)) // 10000 false
	SetBroadcast(&v)
	fmt.Printf("%b  %t\n", v, IsUp(v))   // 10010 false
	fmt.Printf("%b  %t\n", v, IsCast(v)) // 10010 true
}

type Flags uint

const (
	FlagUp           Flags = 1 << iota // 1 << 0 = 1 - is up
	FlagBroadcast                      // 1 << 1 = 2 - supports broadcast access capability
	FlagLoopback                       // 1 << 2 = 4 - is a loopback interface
	FlagPointToPoint                   // 1 << 3 = 8 - belongs to a point-to-point link
	FlagMulticast                      // 1 << 4 = 16 - supports multicast access capability
)

// gopl.io/ch3/netflag
func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }
