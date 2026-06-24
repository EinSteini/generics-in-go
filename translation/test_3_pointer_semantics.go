//go:build ignore

package main

import "fmt"

// Snippet 6: Call-by-value vs. Call-by-pointer
// Testet: Go kopiert Structs by-value, TS teilt Objekte by-reference

type Pair[T any] struct {
	First, Second T
}

// Call-by-value: mutiert nur die lokale Kopie, gibt neues Pair zurück
func SwapCopy[T any](p Pair[T]) Pair[T] {
	p.First, p.Second = p.Second, p.First
	return p
}

// Call-by-pointer: mutiert das Original direkt
func SwapInPlace[T any](p *Pair[T]) {
	p.First, p.Second = p.Second, p.First
}

func main() {
	p := Pair[int]{First: 10, Second: 20}

	q := SwapCopy(p)
	fmt.Println(p.First, p.Second) // 10 20 (unverändert!)
	fmt.Println(q.First, q.Second) // 20 10 (neue Kopie)

	SwapInPlace(&p)
	fmt.Println(p.First, p.Second) // 20 10 (mutiert)
}