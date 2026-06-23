package main

import "fmt"

// Named represents an entity with a name.
// This struct corresponds to the TypeScript 'Named' class.
type Named struct {
	Name string
}

// In Go, a struct's zero value handles the default constructor parameter
// (e.g., Named{} results in Name: "", matching constructor(name: string = "")).
// For the purpose of 'Box' embedding 'Named', we directly initialize 'Named' within 'NewBox'.

// Box holds a value of any type T and is named.
// It corresponds to the TypeScript 'Box<T>' class.
// 'Named' is embedded, which is Go's way of achieving composition/inheritance,
// promoting the 'Name' field directly to 'Box'.
type Box[T any] struct {
	Named // Embedded struct: provides the Name field and effectively "extends" Named.
	Value T
}

// NewBox creates a new Box instance with a given name and value.
// This generic function acts as the factory/constructor for Box[T],
// mirroring both the TypeScript 'Box' constructor and the 'NewBox' factory function.
func NewBox[T any](name string, value T) Box[T] {
	// Initialize the embedded 'Named' struct.
	// This corresponds to the 'super(name)' call in the TypeScript constructor.
	return Box[T]{
		Named: Named{Name: name}, // Initialize the embedded Named struct with the provided name.
		Value: value,
	}
}

// main function to demonstrate the usage.
// Corresponds to the TypeScript 'main' function.
func main() {
	// Create a new Box instance with a string name and an integer value.
	b := NewBox("answer", 42)
	// Access the 'Name' field (promoted from the embedded 'Named' struct)
	// and the 'Value' field.
	fmt.Println(b.Name, b.Value) // Outputs: answer 42
}