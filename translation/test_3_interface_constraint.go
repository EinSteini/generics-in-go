//go:build ignore

package main

import "fmt"

// Snippet 3: Interface-Embedding als Constraint
// Testet: Interface-Embedding, Methoden-Constraint, Generische Funktion mit Struct-Constraint

type Stringer interface {
	String() string
}

type Validator interface {
	IsValid() bool
}

type Entry interface {
	Stringer
	Validator
}

func PrintValid[T Entry](items []T) {
	for _, item := range items {
		if item.IsValid() {
			fmt.Println("✓", item.String())
		} else {
			fmt.Println("✗", item.String())
		}
	}
}

// --- Konkreter Typ ---

type Email struct {
	Address string
}

func (e Email) String() string  { return e.Address }
func (e Email) IsValid() bool   { return len(e.Address) > 3 }

func main() {
	emails := []Email{
		{Address: "alice@example.com"},
		{Address: "ab"},
		{Address: "bob@test.org"},
	}
	PrintValid(emails)
}