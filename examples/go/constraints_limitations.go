//go:build ignore

package main

// Generic methods are not allowed to introduce type bounds

import "fmt"

// Summable is a constraint requiring an Add method.
type Summable interface {
	Add(other Summable) Summable
}

// Stack is a generic stack with no constraint on T.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// ERROR: cannot use a more specific constraint in the receiver.
// We'd like to restrict T to Summable only for this method, but Go forbids it:
//
// func (s Stack[T Summable]) SumAll() T {
// 	var acc T
// 	for _, v := range s.items {
// 		acc = acc.Add(v).(T)
// 	}
// 	return acc
// }

// WORKAROUND: use a top-level function that introduces the constraint
func SumAll[T Summable](s Stack[T]) T {
	var acc T
	for _, v := range s.items {
		acc = acc.Add(v).(T)
	}
	return acc
}

type Money struct{ cents int }

func (m Money) Add(other Summable) Summable {
	return Money{cents: m.cents + other.(Money).cents}
}

func (m Money) String() string {
	return fmt.Sprintf("$%d.%02d", m.cents/100, m.cents%100)
}

func main() {
	var s Stack[Money]
	s.Push(Money{cents: 350})
	s.Push(Money{cents: 1275})
	s.Push(Money{cents: 99})

	// s.SumAll() // Cannot do this — constraint not allowed in receiver!
	total := SumAll(s)
	fmt.Println(total) // Output: $17.24
}