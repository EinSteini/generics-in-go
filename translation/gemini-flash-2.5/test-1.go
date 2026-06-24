package main

import "fmt"

// Repeat creates a slice containing the value 'v' repeated 'n' times.
//
// Equivalent to TypeScript:
// export function Repeat<T>(v: T, n: number): T[] { ... }
func Repeat[T any](v T, n int) []T {
	// In Go, slices are typically initialized with make.
	// make([]T, n) creates a slice of length 'n', with all elements
	// initialized to their zero value. These will be immediately
	// overwritten by the loop, making it functionally equivalent to
	// TypeScript's `new Array(n)` for this purpose.
	result := make([]T, n)

	for i := 0; i < n; i++ {
		result[i] = v
	}
	return result
}

func main() {
	fmt.Println(Repeat("go", 3)) // Erwartete Ausgabe: [go go go]
	fmt.Println(Repeat(42, 2))   // Erwartete Ausgabe: [42 42]
}