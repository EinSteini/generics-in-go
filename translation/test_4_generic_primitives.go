//go:build ignore

package main

import "fmt"

// Snippet 4: Einfache generische Funktion mit primitiven Typen
// Testet: Grundlegende Generics, Typinferenz, Slice-Erzeugung

func Repeat[T any](v T, n int) []T {
	result := make([]T, n)
	for i := range result {
		result[i] = v
	}
	return result
}

func main() {
	fmt.Println(Repeat("go", 3)) // [go go go]
	fmt.Println(Repeat(42, 2))   // [42 42]
}