//go:build ignore

package main

import "fmt"

// Snippet 1: Generische Funktion mit Constraint
// Testet: Type Constraint, Typparameter, Multiple Return Values

type Number interface {
	~int | ~float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Map[T any, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, v := range items {
		result[i] = fn(v)
	}
	return result
}

func main() {
	fmt.Println(Min(3, 7))
	fmt.Println(Min(2.5, 1.8))

	words := []string{"hello", "world"}
	lengths := Map(words, func(s string) int { return len(s) })
	fmt.Println(lengths)
}