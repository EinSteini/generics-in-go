//go:build ignore

package main

import "fmt"

// WITHOUT generics: separate types, duplicated code

// IntStack only works for int.
type IntStack struct{ items []int }

func (s *IntStack) Push(v int)        { s.items = append(s.items, v) }
func (s *IntStack) Len() int          { return len(s.items) }
func (s *IntStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// StringStack is an identical copy — just with string instead of int.
type StringStack struct{ items []string }

func (s *StringStack) Push(v string)        { s.items = append(s.items, v) }
func (s *StringStack) Len() int             { return len(s.items) }
func (s *StringStack) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// AnyStack avoids the duplication but loses type safety:
// compiler cannot catch s.Push(42) on a "string stack" — only a runtime panic can.
type AnyStack struct{ items []any }

func (s *AnyStack) Push(v any)       { s.items = append(s.items, v) }
func (s *AnyStack) Len() int         { return len(s.items) }
func (s *AnyStack) Pop() (any, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}


// WITH GENERICS: one implementation, full type safety

type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Len() int { return len(s.items) }
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func main() {
	fmt.Println("=== Without generics ===")

	var intStack IntStack
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	if v, ok := intStack.Pop(); ok {
		fmt.Println("IntStack popped:", v)
	}

	var stringStack StringStack
	stringStack.Push("go")
	stringStack.Push("generics")
	if v, ok := stringStack.Pop(); ok {
		fmt.Println("StringStack popped:", v)
	}

	// AnyStack: both pushes compile fine, error only surfaces at runtime
	var anyStack AnyStack
	anyStack.Push(42)
	anyStack.Push("oops")
	if v, ok := anyStack.Pop(); ok {
		_ = v.(string) // runtime type assertion, panics if wrong
		fmt.Println("AnyStack top was a string: ", v)

		// This would panic at runtime:
		// _ = v.(int)
	}

	fmt.Println("\n=== With generics ===")

	var genericIntStack Stack[int]
	genericIntStack.Push(10)
	genericIntStack.Push(20)
	genericIntStack.Push(30)

	// This would not compile: cannot use string as int
	// genericIntStack.Push("hello")

	if v, ok := genericIntStack.Pop(); ok {
		fmt.Println("Stack[int] popped:", v)
	}

	var genericStringStack Stack[string]
	genericStringStack.Push("go")
	genericStringStack.Push("generics")

	// This would not compile: cannot use int as string
	// genericStringStack.Push(123)
	
	if v, ok := genericStringStack.Pop(); ok {
		fmt.Println("Stack[string] popped:", v)
	}
}
