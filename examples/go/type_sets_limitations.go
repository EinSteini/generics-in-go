//go:build ignore

package main

// The use of type sets is strictly limited to built-in primitives

import (
	"fmt"
	"slices"
)

// Works: type set of built-in primitives
type Sortable interface {
	~int | ~float64 | ~string
}

func Sort[T Sortable](vals []T) []T {
	sorted := make([]T, len(vals))
	copy(sorted, vals)
	slices.SortFunc(sorted, func(a, b T) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	return sorted
}

// ERROR: cannot use struct types in a type set
// Go forbids using struct literals or composite types in union constraints.
//
// type Coordinate interface {
// 	struct{ X, Y int } | struct{ X, Y, Z int }   // compile error
// }

// ERROR: type sets with non-basic named types cannot be used for operations
// Even if you define a union of custom types, the compiler cannot infer
// which fields or methods are available — even shared ones:
//
// type Point struct{ X, Y int }
// type Segment struct{ Start, End Point }
//
// type Geometry interface {
// 	Point | Segment
// }
//
// func DescribeGeometry[T Geometry](g T) string {
// 	return fmt.Sprintf("%v", g)  // only fmt works via any
// 	// g.X                // ERROR: g.X undefined (type T has no field or method X)
// 	// return g.Describe()   // ERROR: g.Describe undefined (type T has no field or method Describe)
// }

// WORKAROUND for custom types: use a method-based interface instead.
type Describable interface {
	Describe() string
}

type Point struct{ X, Y int }

func (p Point) Describe() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

type Segment struct {
	Start Point
	End   Point
}

func (s Segment) Describe() string {
	return fmt.Sprintf("Segment(%s -> %s)", s.Start.Describe(), s.End.Describe())
}

func PrintAll[T Describable](items []T) {
	for _, item := range items {
		fmt.Println(item.Describe())
	}
}

func main() {
	// Type set with primitives works fine
	ints := []int{30, 10, 20}
	fmt.Println("Sorted ints:", Sort(ints))

	floats := []float64{3.0, 1.5, 2.5}
	fmt.Println("Sorted floats:", Sort(floats))

	words := []string{"cherry", "apple", "banana"}
	fmt.Println("Sorted strings:", Sort(words))

	fmt.Println()

	// For custom types, we must fall back to method-based interfaces
	points := []Point{{1, 2}, {3, 4}}
	PrintAll(points)

	segments := []Segment{{Start: Point{0, 0}, End: Point{1, 1}}, {Start: Point{2, 2}, End: Point{4, 5}}}
	PrintAll(segments)
	// Output:
	// Point(1, 2)
	// Point(3, 4)
	// Segment(Point(0, 0) -> Point(1, 1))
	// Segment(Point(2, 2) -> Point(4, 5))
}
