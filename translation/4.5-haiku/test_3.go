package main

import (
	"fmt"
)

// Stringer definiert die String-Methode
type Stringer interface {
	String() string
}

// Validator definiert die IsValid-Methode
type Validator interface {
	IsValid() bool
}

// Entry erweitert sowohl Stringer als auch Validator
type Entry interface {
	Stringer
	Validator
}

// PrintValid gibt alle Items aus, die das Entry-Interface erfüllen
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

// Email implementiert das Entry-Interface
type Email struct {
	Address string
}

// NewEmail ist der Konstruktor für Email
func NewEmail(address string) *Email {
	return &Email{
		Address: address,
	}
}

// String implementiert das Stringer-Interface
func (e *Email) String() string {
	return e.Address
}

// IsValid implementiert das Validator-Interface
func (e *Email) IsValid() bool {
	return len(e.Address) > 3
}

func main() {
	emails := []*Email{
		NewEmail("alice@example.com"),
		NewEmail("ab"),
		NewEmail("bob@test.org"),
	}

	PrintValid(emails)
}