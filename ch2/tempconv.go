// Package tempconv converts temperatures between Celsius and Fahrenheit.

package ch2

// package tempconv

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
)

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

const (
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
