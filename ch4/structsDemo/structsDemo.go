// package structsdemo

package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	// The zero value of a struct is composed of the zero values of each of its fields.
	var dilbert Employee
	// the dot notation is used to access fields of a struct
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	position := &dilbert.Position
	*position = "Software Engineer"   // dilbert writes code all day now!
	*position = "Senior " + *position // promoted, for writing more lines of code

	// the dot notation is also works with a pointer to a struct
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Name = "Dilbert"
	// (*employeeOfTheMonth).Position += " (proactive team player)"
	// the implicit dereferencing of the pointer, can be simplified
	// because Go does automatic dereferencing, hence we can write:
	employeeOfTheMonth.Position += " (proactive team player)"
	dilbert.ID = 1         // assigned an employee ID
	dilbert.ManagerID = 42 // assigned a manager ID

	id := dilbert.ID
	fmt.Printf("Employee ID: %d\n", id)
	fmt.Printf("dilbert: %+v\n", dilbert)

	// Struct LIterals
	// A value of a struct type can be created using a struct literal.
	// that specifies the values of its fields.
	//
	// There are two forms of struct literals:
	// 1. requires a value be specified for every field, in the order they are defined
	// The first form is more concise, but requires you to specify a value for every field
	point1 := Point{1, 2}

	// 2. a struct literal that specifies the values of some or all of its fields by name
	//    & their corresponding values. If a field is ommitted, it will be set to its zero value.
	// The second form is more readable and allows you to specify only the fields you want to
	point2 := Point{Y: 3} // X will be set to its zero value, which is 0

	fmt.Printf("Point 1 Scaled by 2: %+v\n", ScalePoint(point1, 2))
	fmt.Printf("Point 2 Scaled by 3: %+v\n", ScalePoint(point2, 3))
}

type Point struct{ X, Y int }

func ScalePoint(p Point, factor int) Point {
	// ScalePoint takes a Point and a factor, and returns a new Point
	// that is the result of scaling the original Point by the factor.
	return Point{X: p.X * factor, Y: p.Y * factor}
}

// For efficiency, larger structs are often passed to or returned from functions
// indirectly using a pointer to the struct.
// when the struct is large, it avoids copying the entire struct
//
// this is required if the function must modify its argument, since in a call-by-value
// language like Go, the function receives a copy of the argument, not the original.
func employeeBonus(emp *Employee, percent int) int {
	// employeeBonus takes a pointer to an Employee and a percentage,
	// and returns the bonus amount based on the employee's salary.
	if emp == nil {
		return 0
	}
	return emp.Salary * percent / 100
}

func AwardAnnualRaise(emp *Employee) {
	emp.Salary = emp.Salary * 110 / 100 // 10% raise
}
