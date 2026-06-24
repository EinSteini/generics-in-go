//go:build ignore

package main

import "fmt"

// Snippet 2: Generischer Struct mit Methoden
// Testet: Generischer Typ, Pointer-Receiver, Zero Value, Comma-Ok-Idiom

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func main() {
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("Len:", s.Len())
	if v, ok := s.Pop(); ok {
		fmt.Println("Popped:", v)
	}
}