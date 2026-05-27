package main

import (
	"fmt"
)

// Type Constraint für numerische Typen
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

// Generische Funktion Min mit Constraint
func Min[T Number](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

// Generische Map-Funktion
func Map[T, U any](items []T, fn func(item T) U) []U {
	result := make([]U, len(items))
	for i := 0; i < len(items); i++ {
		result[i] = fn(items[i])
	}
	return result
}

// Hauptprogramm
func main() {
	fmt.Println(Min(3, 7))
	fmt.Println(Min(2.5, 1.8))

	words := []string{"hello", "world"}
	lengths := Map(words, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)
}