package main

import "fmt"

// Stack ist ein generischer Stack mit Elementen vom Typ T.
type Stack[T any] struct {
	items []T
}

// NewStack erstellt einen neuen, leeren Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: []T{},
	}
}

// Push fügt ein Element v oben auf den Stack.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop entfernt und gibt das oberste Element vom Stack zurück.
// Das zweite Return-Value gibt an, ob das Pop erfolgreich war
// (false, wenn der Stack leer ist).
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Len gibt die Anzahl der Elemente im Stack zurück.
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// main
func main() {
	s := NewStack[int]()
	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("Len:", s.Len())
	v, ok := s.Pop()
	if ok {
		fmt.Println("Popped:", v)
	}
}