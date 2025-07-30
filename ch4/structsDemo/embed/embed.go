// package embed.go

package main

import "fmt"

// Go's unusual struct embedding mechanism lets us use
// one named struct type as an anonymous field in another struct type.
type Circle0 struct {
	X, Y, Radius int
}

type Wheel0 struct {
	X, Y, Radius, Spokes int
}

// factoring out common parts
type Point struct {
	X, Y int
}

type Circle1 struct {
	Center Point
	Radius int
}
type Wheel1 struct {
	Circle Circle1
	Spokes int
}

// Go lets us declare a field with a type but no name,
// such fields are called anonymous fields.
// The type of the field is used as the name of the field.
// The type of the field must be a named type, or pointer to a named type.
type Circle2 struct {
	Point
	Radius int
}
type Wheel2 struct {
	Circle2
	Spokes int
}

func main() {
	var w0 Wheel0
	w0.X = 6
	w0.Y = 7
	w0.Radius = 5
	w0.Spokes = 15

	var w1 Wheel1
	w1.Circle.Center.X = 6
	w1.Circle.Center.Y = 7
	w1.Circle.Radius = 5
	w1.Spokes = 15

	var w2 Wheel2
	w2.X = 6      // equivalent to w2.Circle2.Point.X
	w2.Y = 7      // equivalent to w2.Circle2.Point.Y
	w2.Radius = 5 // equivalent to w2.Circle2.Radius
	w2.Spokes = 15

	w := Wheel2{
		Circle2: Circle2{
			Point:  Point{X: 8, Y: 9},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Println("Wheel0:", w0)
	fmt.Println("Wheel1:", w1)
	fmt.Println("Wheel2:", w2)
	fmt.Printf("%#v\n", w)
	fmt.Printf("%+v\n", w)
}
