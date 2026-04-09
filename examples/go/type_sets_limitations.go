//go:build ignore

package main

// The use of type sets is strictly limited to built-in primitives

import "fmt"


// Works: type set of built-in primitives
type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Numeric](vals []T) T {
	var total T
	for _, v := range vals {
		total += v
	}
	return total
}

// ERROR: cannot use struct types in a type set
// Go forbids using struct literals or composite types in union constraints.
//
// type Coordinate interface {
// 	struct{ X, Y int } | struct{ X, Y, Z int }   // compile error
// }

// ERROR: type sets with non-basic named types cannot be used for operations
// Even if you define a union of custom types, the compiler cannot infer
// which operators are available:
//
// type Point struct{ X, Y int }
// type Segment struct{ Start, End Point }
//
// type Geometry interface {
// 	Point | Segment
// }
//
// func Describe[T Geometry](g T) string {
// 	return fmt.Sprintf("%v", g)  // only fmt works via any
// 	// g.X                       // error: no field X on type T
// }

// WORKAROUND for custom types: use a method-based interface instead.
type Describable interface {
	Describe() string
}

type Point struct{ X, Y int }

func (p Point) Describe() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

type Circle struct {
	Center Point
	Radius float64
}

func (c Circle) Describe() string {
	return fmt.Sprintf("Circle(center=%s, r=%.1f)", c.Center.Describe(), c.Radius)
}

func PrintAll[T Describable](items []T) {
	for _, item := range items {
		fmt.Println(item.Describe())
	}
}

func main() {
	// Type set with primitives works fine
	ints := []int{10, 20, 30}
	fmt.Println("Sum of ints:", Sum(ints)) // Output: Sum of ints: 60

	floats := []float64{1.5, 2.5, 3.0}
	fmt.Println("Sum of floats:", Sum(floats)) // Output: Sum of floats: 7

	fmt.Println()

	// For custom types, we must fall back to method-based interfaces
	points := []Point{{1, 2}, {3, 4}}
	PrintAll(points)
	// Output:
	// Point(1, 2)
	// Point(3, 4)
}