//go:build ignore

package main

import "fmt"

// Snippet 7: Struct-Embedding mit Generics
// Testet: Eingebettete Structs, Konstruktorfunktion, direkter Feldzugriff

type Named struct {
	Name string
}

type Box[T any] struct {
	Named
	Value T
}

func NewBox[T any](name string, v T) Box[T] {
	return Box[T]{Named: Named{Name: name}, Value: v}
}

func main() {
	b := NewBox("answer", 42)
	fmt.Println(b.Name, b.Value) // answer 42
}