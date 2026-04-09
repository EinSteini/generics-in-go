//go:build ignore

package main

// Methods are not allowed to introduce type parameters

import (
	"fmt"
	"strings"
)

// Container is a generic wrapper holding a slice of T.
type Container[T any] struct {
	items []T
}

// Methods on generic types CANNOT introduce new type parameters.
// We'd like a Map method that transforms T -> U, but this is illegal in Go:
//
// func (c Container[T]) Map[U any](f func(T) U) Container[U] {
// 	result := Container[U]{items: make([]U, len(c.items))}
// 	for i, v := range c.items {
// 		result.items[i] = f(v)
// 	}
// 	return result
// }

// WORKAROUND: use a top-level function that introduces both type parameters.
func Map[T any, U any](c Container[T], f func(T) U) Container[U] {
	result := Container[U]{items: make([]U, len(c.items))}
	for i, v := range c.items {
		result.items[i] = f(v)
	}
	return result
}

func main() {
	nums := Container[int]{items: []int{1, 2, 3, 4}}

	// nums.Map(func(n int) string { ... })  // Cannot do this — methods cannot have type parameters!

	strs := Map(nums, func(n int) string {
		return strings.Repeat("*", n)
	})
	fmt.Println(strs.items) // Output: [* ** *** ****]
}