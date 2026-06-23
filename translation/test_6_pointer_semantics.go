//go:build ignore

package main

import "fmt"

// Snippet 6: Call-by-value und Zeiger-Semantik
// Testet: Generische Zeiger-Parameter, In-Place-Mutation vs. Value-Semantik

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func Double[T ~int | ~float64](v T) T {
	return v * 2 // Original bleibt unverändert
}

func main() {
	x, y := 10, 20
	Swap(&x, &y)
	fmt.Println(x, y) // 20 10

	n := 5
	fmt.Println(Double(n), n) // 10 5
}